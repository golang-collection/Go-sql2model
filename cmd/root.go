package cmd

/**
* @Author: super
* @Date: 2020-09-15 08:28
* @Description:
**/

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(sqlCmd)
}