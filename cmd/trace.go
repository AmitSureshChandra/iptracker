/*
Copyright © 2022 NAME HERE <akumar00029@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide IP to trace")
			return
		}
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
