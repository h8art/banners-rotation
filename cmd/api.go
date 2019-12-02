package cmd

import (
	"banner-otus/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run api",
	Long:  `Run api server`,
	Run: func(cmd *cobra.Command, args []string) {
		api.RunApi()
	},
}
