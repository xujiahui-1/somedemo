package viper_demo

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
	"time"
)

func Test_DemoViper(t *testing.T) {
	// create a viper instance
	config := viper.New()
	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	config.SetDefault("redis.port", 6381)

	if err := config.ReadInConfig(); err != nil { //Find and read the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}
	set := config.IsSet("q") //确认该key是否存在
	if set {
		fmt.Println("q is set")
	} else {
		fmt.Println("q is  not set")
	}
	fmt.Println(config.Get("server.port")) //get the value from config file
	fmt.Println(config.Get("server.tomcat.uri-encoding"))

}

// Test_DemoViper2 deomo Get function
func Test_DemoViper2(t *testing.T) {
	config := viper.New()
	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	set := config.IsSet("server.servlet.context-path")
	fmt.Println(set)
	//读取字符串类型切片
	slice := config.GetStringSlice("server.servlet.context-path")
	fmt.Println(slice[1])

	//读取int类型切片
	intSlice := config.GetIntSlice("server.port")
	fmt.Println(intSlice)

	//获取时间？
	duration := config.GetDuration("server.tomcat.uri-encoding")
	fmt.Println(duration)

	//获取字符串
	getString := config.GetString("server.string")
	fmt.Println(getString)

	//获取一整个域中的所有的键值对
	stringMap := config.GetStringMap("server.tomcat")
	fmt.Println(stringMap)
	//获取一整个域中的所有的键值对
	settings := config.AllSettings()
	fmt.Println(settings)
}

// Test_DemoViper3 demo Set function
func Test_DemoViper3(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	//写配置,如果文件已经存在则不行了
	viper.Set("app_name", "awesome web")
	viper.Set("log_level", "DEBUG")
	viper.Set("mysql.ip", "127.0.0.1")
	viper.Set("mysql.port", 3306)
	viper.Set("mysql.user", "root")
	viper.Set("mysql.password", "123456")
	viper.Set("mysql.database", "awesome")

	viper.Set("redis.ip", "127.0.0.1")
	viper.Set("redis.port", 6381)

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}

}

// Test_DemoViper4 监听配置文件，viper 可以监听文件修改，热加载配置
func Test_DemoViper4(t *testing.T) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	//监听配置文件是否被修改了
	viper.WatchConfig()
	fmt.Println("redis port before sleep: ", viper.Get("server.string"))
	time.Sleep(time.Second * 10)
	fmt.Println("redis port after sleep: ", viper.Get("server.string"))

}
