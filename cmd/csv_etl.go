/*
Copyright © 2023 MYM
*/
package cmd

import (
	"github.com/mymhimself/simple-csv-reader/apps/csvetl"
	"github.com/spf13/cobra"
)

// scriptsCmd represents the scripts command
var scriptsCmd = &cobra.Command{
	Use:   "csv-etl",
	Short: "A Simple CSV Extractor-Transformer-Loader",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		run, err := csvetl.NewApp(cmd)
		if err != nil {
			panic(err)
		}

		// start the application
		return run()
	},
}

func init() {
	rootCmd.AddCommand(scriptsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scriptsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scriptsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
