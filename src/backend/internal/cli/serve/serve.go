package serve

import (
	"fmt"

	"github.com/Roongjin/ChatApplication/src/backend/internal/config"
	"github.com/Roongjin/ChatApplication/src/backend/internal/controller"
	"github.com/Roongjin/ChatApplication/src/backend/internal/controller/chat"
	"github.com/Roongjin/ChatApplication/src/backend/internal/third-party/databases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve [FLAGS]...",
	Short: "Serve application",
	RunE: func(cmd *cobra.Command, args []string) error {
		pathCfgFiles, err := cmd.Flags().GetStringSlice("config-file")
		if err != nil {
			return err
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		appCfg := config.MustReadMultipleAppConfigFiles(pathCfgFiles)
		if debug {
			printAppConfig(appCfg)
		}

		ipAddr := appCfg.Network.En0.IpAddr
		lanOrigin := fmt.Sprintf("http://%s:5173", ipAddr)

		// DBs and controllers
		db := databases.ConnectSQLDB(appCfg.Database.Postgres.DSN)
		handler := controller.NewHandler(db)
		redisClient := databases.ConnectRedis(appCfg.Database.Redis.DSN)

		// Chat system
		chatEntity := chat.NewChat(db, redisClient, &handler.Chat)
		defer chatEntity.Close()

		r := gin.Default()

		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173", lanOrigin},
			AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))

		r.POST("/authen/:name", handler.User.Authentication)

		chat := r.Group("/chat")
		{
			chat.GET("/ws/:userId", chatEntity.ServeWS)
			chat.GET("/online-users", handler.User.GetOnlineUsers)
		}

		r.Run()

		return nil
	},
}
