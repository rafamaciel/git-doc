package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var workingDir string

var rootCmd = &cobra.Command{
	Use:   "git-doc",
	Short: "",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func GetDefaultWorkingDir() string {
	workingDir, _ = os.Getwd()
	return workingDir
}
func init() {

	rootCmd.PersistentFlags().StringVar(&workingDir, "workdir", GetDefaultWorkingDir(), "current directory path")

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
