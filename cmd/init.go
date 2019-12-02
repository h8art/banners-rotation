package cmd

import (
	"banner-otus/pkg/banners"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize postgresql",
	Long:  `Init all tables in db`,
	Run: func(cmd *cobra.Command, args []string) {
		err := banners.InitDB()
		if err != nil {
			fmt.Println(err)
		}
	},
}
