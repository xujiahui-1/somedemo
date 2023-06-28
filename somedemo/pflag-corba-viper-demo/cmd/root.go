/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ENV = "DB_NAME"

// rootCmd represents the base command when called without any subcommands
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "test",                                                                //命令名
		Short: "A brief description of your application",                             //描述
		Long:  `A longer description that spans multiple lines and  likely contains`, //描述
		Run: func(cmd *cobra.Command, args []string) { //执行调用的函数
			// 读取env变量
			viper.AutomaticEnv()
			env := viper.Get(ENV)
			if env == nil {
				fmt.Println("您未配置环境变量，将使用命令行传入参数")
			} else {
				fmt.Println("您配置了环境变量，", env)
			}
			_ = cmd.Help() // Ignore error
		},
	}
	// 通过Pflag处理命令行参数
	pflag.StringP("name", "n", "", "name")
	pflag.Parse()
	//绑定命令行参数
	viper.BindPFlags(pflag.CommandLine)
	getString := viper.GetString("name")
	fmt.Println("您通过命令行传入了参数", getString)
	rootCmd.AddCommand(NewEnvCmd())
	return rootCmd
}
