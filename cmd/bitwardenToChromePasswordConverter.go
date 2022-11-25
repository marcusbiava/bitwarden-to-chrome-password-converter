/*
Copyright © 2022 Marcus Venicius Nunes Biava
*/
package cmd

import (
	"github.com/marcusbiava/bitwarden-to-chrome-password-converter/app"
	"github.com/spf13/cobra"
)

// bitwardenToChromePasswordConverterCmd represents the bitwardenToChromePasswordConverter command
var bitwardenToChromePasswordConverterCmd = &cobra.Command{
	Use:   "bitwardenToChromePasswordConverter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		csvPath := args[0]
		
		app.Run(csvPath)
	},
}

func init() {
	rootCmd.AddCommand(bitwardenToChromePasswordConverterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bitwardenToChromePasswordConverterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bitwardenToChromePasswordConverterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
