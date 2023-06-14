package main

import (
	"fmt"
	"testing"
	"time"
)

// 这个是一个计数程序
func count(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
	close(c) //关闭channel
}

func Test_channel(t *testing.T) {
	c := make(chan string) //创建一个channel

	go count(5, "🐑", c)
	for {
		massage, open := <-c //每次从channel接受消息的时候都可以额外获得一个布尔值
		fmt.Println(massage)
		if !open {
			break
		}
	}
}

// 当我们要同时接受多个channel的消息时，可以使用select来接受最新的消息，防止阻塞
func Test_select(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "🐑"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "🐮"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	//无限循环读取管道中的消息
	for {
		select {
		case massage := <-c1:
			fmt.Println(massage)
		case massage := <-c2:
			fmt.Println(massage)
		}
	}
}
