package cmd

import (
	"log"

	"github.com/Roongjin/ChatApplication/src/backend/internal/cli/migrate"
)

func main() {
	if err := migrate.MigrateCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
