/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"inlaid-cli/pkg/resolve"
	"inlaid-cli/pkg/record"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// inlaidCmd represents the inlaid command
var inlaidCmd = &cobra.Command{
	Use:   "inlaid",
	Run: func(cmd *cobra.Command, args []string) {
		Inlaid()
	},
}

func init() {
	rootCmd.AddCommand(inlaidCmd)
}

func Inlaid() {
	context := resolve.ResolveGitCommitContext(viper.GetString("entity"))
	record.Write(viper.GetString("project"), context)
}
