package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

type host struct {
	value string
}

func (h *host) String() string {
	return h.value
}

func (h *host) Set(v string) error {
	h.value = v
	return nil
}

func (h *host) Type() string {
	return "host"
}

// pflag包，配置命令行参数的
func main() {
	var ip *int = pflag.Int("ip", 1234, "help message for ip")
	var port int
	pflag.IntVar(&port, "port", 8080, "help message for port")
	var h host
	pflag.Var(&h, "host", "help message for host")

	//解析命令行参数
	pflag.Parse()
	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("port: %d\n", port)
	fmt.Printf("host: %+v\n", h)

	fmt.Printf("NFlag: %v\n", pflag.NFlag()) // 返回已设置的命令行标志个数
	fmt.Printf("NArg: %v\n", pflag.NArg())   // 返回处理完标志后剩余的参数个数
	fmt.Printf("Args: %v\n", pflag.Args())   // 返回处理完标志后剩余的参数列表
	fmt.Printf("Arg(1): %v\n", pflag.Arg(1)) // 返回处理完标志后剩余的参数列表中第 i 项
}
