package bufio_demo

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func Test_ReadString(t *testing.T) {

	//通过strings.NewReader获取两个reader。
	str1 := strings.NewReader("A\nBCDE")
	str2 := strings.NewReader("12345")

	bufReader := bufio.NewReader(str1)
	readString, _ := bufReader.ReadString('\n')
	fmt.Println(readString)

	bufReader.Reset(str2) //重制
	readString2, _ := bufReader.ReadString('\n')
	fmt.Println(readString2)

}
func Test_NewReaderSize(t *testing.T) {

	//通过strings.NewReader获取两个reader。
	IOReader1 := strings.NewReader("A\nBCDE")
	buf2 := bufio.NewReaderSize(IOReader1, 4069)
	s, _ := buf2.ReadString('\n') //直到读取\n为止
	fmt.Println(s)

}
func Test_Reader(t *testing.T) {
	reader := strings.NewReader("ABCDEEEEEEEEEEEEEEEEEEEEEEEE")
	newReader := bufio.NewReader(reader)
	p := make([]byte, 5)
	for {
		n, err := newReader.Read(p)
		if err == io.EOF {
			fmt.Println("读取结束")
			break
		} else {
			fmt.Println(string(p[0:n]))
		}
	}

}
func Test_Peek(t *testing.T) {
	reader := strings.NewReader("ABCDeeeeeeeeeeeeeeeee")
	newReader := bufio.NewReader(reader)
	peek, err := newReader.Peek(5)
	if err != nil {
		if err == io.EOF {
			fmt.Println("读取结束")
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(string(peek))
}
func Test_ReadByte(t *testing.T) {
	reader := strings.NewReader("ABCDeeeeeeeeeeeeeeeee")
	newReader := bufio.NewReader(reader)
	readByte, _ := newReader.ReadByte()
	readByte2, _ := newReader.ReadByte()
	fmt.Println(string(readByte))
	fmt.Println(string(readByte2))

	//如果没有数据的情况下呢

	r := strings.NewReader("")
	r2 := bufio.NewReader(r)
	b, er := r2.ReadByte()
	fmt.Println(b)
	fmt.Println("错误", er)
}

// 向buffer中放入最近一次读取的最后一个字节
func Test_UnReadByte(t *testing.T) {
	reader := strings.NewReader("ABCDeeeeeeeeeeeeeeeee")
	newReader := bufio.NewReader(reader)
	p := make([]byte, 3)
	//读取了ABC
	n, err := newReader.Read(p)
	fmt.Println(n, err)
	//只放回去了C
	err2 := newReader.UnreadByte()
	if err2 != nil {
		fmt.Println(err2)
	}
	//测试重复读取 发现报错bufio: invalid use of UnreadByte
	//ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")这个错
	err3 := newReader.UnreadByte()
	if err3 != nil {
		fmt.Println(err3)
	}
	//再次读取，发现是C
	readByte, err := newReader.ReadByte()
	fmt.Println(string(readByte))
	readByte2, err := newReader.ReadByte()
	fmt.Println(string(readByte2))
}

func Test_ReadRune_UnReadRune(t *testing.T) {
	reader := strings.NewReader("序abcde")
	newReader := bufio.NewReader(reader)
	r, size, err := newReader.ReadRune()
	fmt.Println(string(r), size, err) //序

	//UnreadRune后再读取
	err = newReader.UnreadRune()
	fmt.Println(err)
	q, sizew, erre := newReader.ReadRune()
	fmt.Println(string(q), sizew, erre) //还是序
}

func Test_ReadLine(t *testing.T) {
	reader := strings.NewReader("序assssssssssssssssssss\r\ncde\nsdsafsafsadfsad")
	newReader := bufio.NewReader(reader)
	line, prefix, err := newReader.ReadLine()
	fmt.Println(string(line), prefix, err)
}

func Test_Buffered(t *testing.T) {
	reader := strings.NewReader("123456")
	newReader := bufio.NewReader(reader)
	p := make([]byte, 3)
	n, err := newReader.Read(p)
	fmt.Println(n, err)
	buffered := newReader.Buffered()
	fmt.Println(buffered)
}

func Test_Discard(t *testing.T) {
	reader := strings.NewReader("徐家汇想潇洒的撒")
	newReader := bufio.NewReader(reader)
	discarded, err := newReader.Discard(4)
	fmt.Println(discarded, err)

	readString, err := newReader.ReadString('q')
	fmt.Println(readString, err)
}

func Test_Size(t *testing.T) {
	reader := strings.NewReader("aaaaa")

	readerSize := bufio.NewReaderSize(reader, 3)
	size := readerSize.Size()
	fmt.Println(size)

}
