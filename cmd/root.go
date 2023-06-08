/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	_ "inlaid-cli/pkg/viperloader"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "inlaid-cli",
	Run: func(cmd *cobra.Command, args []string) {
		Inlaid()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 绑定 flags
	rootCmd.Flags().String("path", "", "")
	viper.BindPFlag("path", rootCmd.Flags().Lookup("path"))

	rootCmd.Flags().String("entity", "", "")
	viper.BindPFlag("entity", rootCmd.Flags().Lookup("entity"))

	rootCmd.Flags().String("project", "", "")
	viper.BindPFlag("project", rootCmd.Flags().Lookup("project"))
}
