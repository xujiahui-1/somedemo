package main

import (
	"fmt"
	"github.com/carlmjohnson/requests"
	"golang.org/x/net/context"
)

// 本程序体验requests： 一个更高效的HTTP请求包
// 地址 https://github.com/llitfkitfk/go-best-practice

type placeholder struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var s string
	err := requests.
		URL("http://example.com").
		ToString(&s).Fetch(context.Background())

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
		fmt.Println("成功")
	}

	var post placeholder
	errr := requests.
		URL("https://jsonplaceholder.typicode.com").
		Pathf("/posts/%d", 1).
		ToJSON(&post).
		Fetch(context.Background())

	if errr != nil {
		fmt.Println(errr)
	} else {
		fmt.Println(post)
		fmt.Println("成功")
	}
}
