package cmd

import (
	"log"

	"github.com/Roongjin/ChatApplication/src/backend/internal/migrate"
)

func main() {
	if err := migrate.MigrateCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
