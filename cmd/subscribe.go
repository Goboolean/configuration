/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subscribeCmd represents the subscrible command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe {stockID}-{Location}",
	Short: "Subscribe stock data",
	Long: `{stockID} is the unique code of each stock.
	{Location} is a country code defined in ISO 3166-1.
	For example, country code of korea is "ko" and the united state is "us".
	{Location} must be lower case.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subscribe called")
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
