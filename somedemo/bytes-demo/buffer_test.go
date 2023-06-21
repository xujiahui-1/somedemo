package bytes_demo

import (
	"bytes"
	"fmt"
	"testing"
)

//  test for bytes/buffer

func Test_Buffer(t *testing.T) {

	buffer1 := bytes.NewBuffer([]byte("hello"))   //create  a new buffer
	buffer2 := bytes.NewBufferString("helloqwer") //create a new buffer from string
	fmt.Println(buffer1)
	fmt.Println(buffer2)
	buffer1.Write([]byte(" xujiahui"))          //add string to buffer
	buffer2.WriteByte('c')                      //add byte to buffer
	fmt.Println(buffer2.Cap())                  //Cap return the capacity of the buffer's
	buffer2.Grow(8)                             //增加容量
	fmt.Println(buffer2.Len())                  //bytes的长度
	fmt.Printf("%s\n", string(buffer2.Next(5))) //切割出来前几位，
	fmt.Println(buffer2.Cap())
	fmt.Println(buffer2)

	str := make([]byte, 5)
	n, err := buffer2.Read(str) //从buffer中读取数据到str中，返回读取的数据长度和error
	fmt.Println(n, err)
	fmt.Println(str)

}
