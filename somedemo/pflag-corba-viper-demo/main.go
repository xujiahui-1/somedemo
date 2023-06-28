/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"LeetCode/pflag-corba-viper-demo/cmd"
	"fmt"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
