package strconv_demo

import (
	"fmt"
	"strconv"
	"testing"
)

// strconv string类型转换包、 将string转换成其他基本类型

func Test_conver(t *testing.T) {

	// int -> string
	strconv.Itoa(21)
	//string -> int
	strconv.Atoi("1234")

	// Parse  string -> value
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(b, f, i, u)

	//Format value -> string
	s2 := strconv.FormatBool(true)
	s3 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s4 := strconv.FormatInt(-42, 16)
	s5 := strconv.FormatUint(42, 16)
	fmt.Println(s2, s3, s4, s5)
}

func Test_Append(t *testing.T) {
	//AppendBool 向byte切片添加true or false
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

	//AppendFloat  向byte切片添加 float
	b32 := []byte("float32:")
	//dst 原列表
	//f 需要append到列表的浮点数
	//fmt 转换格式 'b' 'e' 'E' 'f' 'g'或'G'
	//prec 浮点数精度
	//bitSize 32或64，32对应float32，64对应float64
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', 10, 32)
	fmt.Println(string(b32))

	//将字符串s转换为双引号引起来的字符串，并将结果追加到dst的尾部，返回追加后的[]byte。其中的特殊字符将被转换为转义字符
	c := []byte("quote:")
	c = strconv.AppendQuote(c, `"你好"`)
	fmt.Println(string(c))

}
