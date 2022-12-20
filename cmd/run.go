/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/amrom66/eth-exporter/pkg"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a daemon backend",
	Long:  `run a daemon backend`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Capture()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
