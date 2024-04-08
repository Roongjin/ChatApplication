package main

import (
	"log"

	"github.com/Roongjin/ChatApplication/src/backend/internal/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
