/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	_ "inlaid-cli/pkg/viperloader"
	"inlaid-cli/internal/reporting"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// reportingCmd represents the reporting command
var reportingCmd = &cobra.Command{
	Use:   "reporting",
	Run: func(cmd *cobra.Command, args []string) {
		Reporting()
	},
}

func init() {
	rootCmd.AddCommand(reportingCmd)
}

func Reporting() {
	reply, _ := reporting.Daily(viper.GetString("path"))
	fmt.Println(reply)
}
