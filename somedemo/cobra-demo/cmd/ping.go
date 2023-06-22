/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command

func NewServerCmd() *cobra.Command {
	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help() // Ignore error
		},
	}
	return pingCmd
}
