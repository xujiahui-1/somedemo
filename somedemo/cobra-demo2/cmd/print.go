package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 本地标志
var printFlag string

func NewPrintCmd() *cobra.Command {
	var printCmd = &cobra.Command{
		Use:   "print [OPTIONS] [COMMANDS]",
		Short: "print XXXXXXXXXXXXXXXXXXX",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run print...")
			fmt.Printf("printFlag: %v\n", printFlag)
			fmt.Printf("verbor: %v\n", Verbose)
		},
	}
	printCmd.Flags().StringVarP(&printFlag, "flag", "f", "", "print flag for local")
	return printCmd
}
