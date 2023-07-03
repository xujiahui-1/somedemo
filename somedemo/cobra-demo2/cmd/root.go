package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

/*
https://xie.infoq.cn/article/915006cf3760c99ad0028d895
命令（COMMAND）：命令表示要执行的操作。
	在执行 Run 函数前后，我么可以执行一些钩子函数，其作用和执行顺序如下：
		PersistentPreRun：在 PreRun 函数执行之前执行，对此命令的子命令同样生效。
		PreRun：在 Run 函数执行之前执行。
		Run：执行命令时调用的函数，用来编写命令的业务逻辑。
		PostRun：在 Run 函数执行之后执行。
		PersistentPostRun：在 PostRun 函数执行之后执行，对此命令的子命令同样生效
*/

/*
参数（ARG）：是命令的参数，一般用来表示操作的对象。
	Args 属性类型为一个函数：func(cmd *Command, args []string) error，可以用来验证参数。
	Cobra 内置了以下验证函数：
	NoArgs：如果存在任何命令参数，该命令将报错。
	ArbitraryArgs：该命令将接受任意参数。
	OnlyValidArgs：如果有任何命令参数不在 Command 的 ValidArgs 字段中，该命令将报错。
	MinimumNArgs(int)：如果没有至少 N 个命令参数，该命令将报错。
	MaximumNArgs(int)：如果有超过 N 个命令参数，该命令将报错。
	ExactArgs(int)：如果命令参数个数不为 N，该命令将报错。
	ExactValidArgs(int)：如果命令参数个数不为 N，或者有任何命令参数不在 Command 的 ValidArgs 字段中，该命令将报错。
	RangeArgs(min, max)：如果命令参数的数量不在预期的最小数量 min 和最大数量 max 之间，该命令将报错

当然我们可以自定义
*/

/*
标志（FLAG）：是命令的修饰，可以调整操作的行为。

	对于全局标志，可以定义在根命令 rootCmd 上。

	除了将命令行标志的值绑定到变量，我们也可以将标志绑定到 Viper，这样就可以使用 viper.Get() 来获取标志的值了。
*/

var Region string

// 持久标志FLAG
var Verbose bool

// viper绑定用变量
var author string

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hugo", //命令名称
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io`,
		Run: func(cmd *cobra.Command, args []string) { //Run属性是一个函数，当执行命令时会调用此函数
			fmt.Println("run hugo...")
			fmt.Printf("Verbose: %v\n", Verbose)
			fmt.Printf("Region: %v\n", Region)
			fmt.Printf("Author: %v\n", viper.Get("author"))
			fmt.Printf("Config: %v\n", viper.AllSettings())
		},
	}
	//添加子命令
	rootCmd.AddCommand(NewVersionCmd())
	rootCmd.AddCommand(NewPrintCmd())
	//持久标志FLAG 绑定到变量了
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	//必选标志 绑定到变量了
	rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	rootCmd.MarkFlagRequired("region") //MarkFlagRequired必须填写！！！！

	//添加标志，并绑定viper
	rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))

	//cobra对配置的初始化
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
	return rootCmd
}
func Execute() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil { //Execute 命令的入口
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgFile string

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		config := viper.New()
		config.AddConfigPath(".")
		config.SetConfigName("config")
		config.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
