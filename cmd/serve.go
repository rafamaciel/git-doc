package cmd

import (
	"github.com/rafamaciel/git-doc/internal/server"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "provide documentation via http",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag(DefaultPortOption.Name).Value.String()
		host := cmd.Flag(DefaultHostOption.Name).Value.String()
		source := cmd.Flag(DefaultSourceOption.Name).Value.String()

		s := server.Server{
			Port:   port,
			Host:   host,
			Source: source,
		}

		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String(
		DefaultPortOption.Name,
		DefaultPortOption.Default,
		DefaultPortOption.Description,
	)

	serveCmd.PersistentFlags().String(
		DefaultHostOption.Name,
		DefaultHostOption.Default,
		DefaultHostOption.Description,
	)

	serveCmd.PersistentFlags().String(
		DefaultSourceOption.Name,
		DefaultSourceOption.Default,
		DefaultSourceOption.Description,
	)
}
