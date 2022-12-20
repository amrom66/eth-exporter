/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pkg "github.com/amrom66/eth-exporter/pkg"
)

// config for program
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eth-exporter",
	Short: "capture all your connection ip",
	Long:  `eth-exporter is a tool to capture all connection and flush to database.`,
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .eth-exporter.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(pwd)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".eth-exporter")
	}
	viper.SetDefault("app.name", "eth-exporter")
	viper.SetDefault("app.log", "/var/log/eth-exporter")

	if err := viper.ReadInConfig(); err == nil {
		logPath := viper.GetString("app.log")
		_, err := os.Stat(logPath)
		if err != nil {
			err := os.MkdirAll(logPath, 0766)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		defer glog.Flush()
		flag.Set("log_dir", logPath)
		flag.Set("alsologtostderr", "true")
		flag.Parse()
		glog.Infoln("Using config file:", viper.ConfigFileUsed())
	}

	pkg.Device = viper.GetString("app.instance.device")
	pkg.Token = viper.GetString("app.db.token")
	pkg.Url = viper.GetString("app.db.url")
	pkg.Bucket = viper.GetString("app.db.bucket")
	pkg.Org = viper.GetString("app.db.org")
	pkg.Instance = viper.GetString("app.db.instance")
}
