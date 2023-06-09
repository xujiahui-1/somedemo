package main

import (
	"fmt"
	"github.com/h2non/filetype"
	"os"
)

// filetype 一个文件校验用的包
// git地址 https://github.com/h2non/filetype
// API 地址 https://pkg.go.dev/github.com/h2non/filetype#section-readme
func main2() {

	//校验文件type
	file, err := os.ReadFile("filetype-demo/test.png")
	if err != nil {
		fmt.Println("error:", err)
	}
	//匹配推断给定缓冲区的文件类型，检查其幻数签名
	kind, err := filetype.Match(file)
	if err != nil {
		fmt.Println("error:", err)
	}
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return
	}
	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)

	//Check type class 检验文件是不是图片
	if filetype.IsImage(file) {
		fmt.Println("File is an image")
	} else {
		fmt.Println("Not an image")
	}
	//检测扩展名是否支持该文件
	if filetype.IsSupported("jpg") {
		fmt.Println("Extension supported")
	} else {
		fmt.Println("Extension not supported")
	}

	//通过文件头判断
	file2, err := os.Open("filetype-demo/test2.pdf")
	if err != nil {
		fmt.Println("error:", err)
	}
	head := make([]byte, 261)
	//Read方法从文件中读取字节并存储在byte中
	file2.Read(head)

	if filetype.IsImage(head) {
		fmt.Println("文件是一个图片")
	} else {
		fmt.Println("文件不是一个图片")
	}
}
