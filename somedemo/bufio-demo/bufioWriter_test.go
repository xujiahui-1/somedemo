package bufio_demo

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"sync"
	"testing"
)

func Test_NewWriter(t *testing.T) {
	b1 := bytes.NewBuffer(make([]byte, 0))

	writer := bufio.NewWriter(b1)
	writer.WriteString("12345") //写入字符串
	writer.Flush()              //把所有缓冲区数据写入最底层writer。不刷新写不进去奥
	fmt.Println(b1)
	b2 := bytes.NewBuffer(make([]byte, 0))
	writer.Reset(b2) //Reset从新给writer赋予缓冲区
	writer.WriteString("abcdsassssssssssss")
	writer.Flush()
	fmt.Println(b2)

}
func Test_Bufferedd(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 20))
	writer := bufio.NewWriter(buffer)
	fmt.Println("当前可用字节数", writer.Available()) //4096
	fmt.Println("以使用字节数", writer.Buffered())   //0

	writer.WriteString("111111111111")
	fmt.Println("当前可用字节数", writer.Available()) //4084
	fmt.Println("以使用字节数", writer.Buffered())   //12
}

func Test_ReadFrom(t *testing.T) { //从reader读取
	//创建reader
	reader := strings.NewReader("abcdedasfsafsa")

	//创建writer
	buffer := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(buffer)

	n, err := writer.ReadFrom(reader) //返回读取的字节数和err ,并且ReadFrom无需刷新
	fmt.Println(n, err)
	fmt.Println(buffer)
}
func Test_Write(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	p := []byte("abcde")

	writer := bufio.NewWriter(buffer)
	writer.Write(p) //将p中的数据写入到缓冲
	writer.Flush()

	fmt.Println(buffer)
}

func Test_ReadWrite(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(buffer) //创建writer
	reader := strings.NewReader("abcd")
	r := bufio.NewReader(reader) //创建reader

	//创建一个 ReadWrite
	rw := bufio.NewReadWriter(r, w)
	readString, err := rw.ReadString('n')
	fmt.Println(readString, err)
}
func Test_qqq(t *testing.T) {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-done)
	}
}
