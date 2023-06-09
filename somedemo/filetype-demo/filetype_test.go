package main

import (
	"github.com/h2non/filetype"
	"net/http"
	"os"
	"testing"
)

// 测试一下这个包的速度
// 测试命令，  go test -bench=. -benchmem
func readFile(filename string) []byte {
	buf, _ := os.ReadFile(filename)
	return buf

}

// 测试文件检测类型,对比一下他检测的块还是标准包kuai
func BenchmarkMatch(b *testing.B) {
	//读取文件
	buf := readFile("filetype-demo/test.png")
	//重制计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filetype.Match(buf)
	}
}
func BenchmarkDetect(b *testing.B) {
	//读取文件
	buf := readFile("filetype-demo/test.png")
	//重制计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//检查给定数据的内容类型Content-Type,最多检测512byte数据,如果有效的话,该函数返回一个MIME类型,否则的话,返回一个"application/octet-stream"
		http.DetectContentType(buf)
	}
}
