package serve

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/config"
	"github.com/Roongjin/ChatApplication/src/backend/internal/controller"
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

		db := databases.ConnectSQLDB(appCfg.Database.Postgres.DSN)
		handler := controller.NewHandler(db)
		// redisClient := databases.ConnectRedis(appCfg.Database.Redis.DSN)
		// chatEntity := chat.NewChat(db, redisClient, &handler.Chat)

		r := gin.Default()

		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))

		r.POST("/", handler.User.Authentication)

		r.Run()

		return nil
	},
}
