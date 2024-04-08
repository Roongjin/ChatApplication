package cli

import (
	"github.com/spf13/cobra"

	"github.com/Roongjin/ChatApplication/src/backend/internal/cli/serve"
)

func init() {
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "Run in debug mode")
	RootCmd.PersistentFlags().StringSliceP("config-file", "c", []string{}, "Path to configuration file")
	RootCmd.AddCommand(serve.ServeCmd)
}

var RootCmd = &cobra.Command{
	Use:   "Chat-system backend",
	Short: "Backend service of Chat Application",
}
