package cmd

import (
	"fmt"

	"github.com/rafamaciel/git-doc/internal"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize the documentation in the repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.GenerateSkeleton()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {

	rootCmd.AddCommand(initCmd)

}
