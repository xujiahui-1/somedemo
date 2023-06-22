package bytes_demo

import (
	"bytes"
	"fmt"
	"testing"
	"unicode"
)

/*  function demo for  bytes package  */

func Test_MappeFunc(t *testing.T) {
	b := []byte("Hello ")

	fmt.Printf("%s\n", bytes.ToUpper(b)) // 大写
	fmt.Printf("%s\n", bytes.ToLower(b)) //小写
	fmt.Printf("%s\n", bytes.ToTitle(b)) //标题（UTF-8-encoded）

	str := []byte("AHOJ VÝVOJÁRİ golang")
	fmt.Printf("%s\n", bytes.ToLowerSpecial(unicode.AzeriCase, str)) //用于特殊的大小写规则
	fmt.Printf("%s\n", bytes.ToUpperSpecial(unicode.AzeriCase, str))
	fmt.Printf("%s\n", bytes.ToTitleSpecial(unicode.AzeriCase, str))

}

// Test_Compare test for two  bytes[] is same or not
func Test_Compare(t *testing.T) {
	a := []byte("b")
	b := []byte("B")

	compare := bytes.Compare(a, b) // Compare 字典顺序比，the result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	fmt.Println(compare)

	equal := bytes.Equal(a, b) // if a==b retrun ture else false
	fmt.Println(equal)

	fold := bytes.EqualFold(a, b) // case-folding  Equal
	fmt.Println(fold)
}

// Test_Trim test for clean a bytes[]
func Test_Trim(t *testing.T) {
	str := []byte("!!!!abcde!!!!!")

	fmt.Printf("%s\n", bytes.Trim(str, "!!"))      //off all leading and trailing
	fmt.Printf("%s\n", bytes.TrimLeft(str, "!!"))  //off all leading
	fmt.Printf("%s\n", bytes.TrimRight(str, "!!")) //off all trailing

	str2 := []byte("     QQ       ")
	fmt.Printf("%s \n", str2)
	fmt.Printf("%s\n", bytes.TrimSpace(str2))                //off all leading and trailing white space
	fmt.Printf("%s\n", bytes.TrimSuffix(str, []byte("!!!"))) // off suffix
	fmt.Printf("%s\n", bytes.TrimPrefix(str, []byte("!!!"))) //off Prefix
}

// Test_Split for split a bytes[]
func Test_Split(t *testing.T) {
	b := []byte("a,b,c")
	fmt.Printf("%s \n", bytes.Split(b, []byte(",")))          //以 ", "切割切片
	fmt.Printf("%s \n", bytes.SplitN(b, []byte(","), 2))      //以 ", "切割切片 并指定子切片数
	fmt.Printf("%s \n", bytes.SplitAfter(b, []byte(",")))     //结果包含分隔符
	fmt.Printf("%s \n", bytes.SplitAfterN(b, []byte(","), 2)) //从后往前切两次

	str := []byte("a     b     c")

	fmt.Printf("%s \n", bytes.Fields(str)) //以空白切割

	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", "))) //以,链接多个切片

	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2)) //将 na 重复两次后返回
}

// Test_SubSlice test for SubSlice
func Test_SubSlice(t *testing.T) {
	b := []byte("test_xujiahui.com")

	fmt.Println(bytes.HasPrefix(b, []byte("test")))                 //do slice have a prefix?
	fmt.Println(bytes.HasSuffix(b, []byte(".com")))                 //do this have a suffix?
	fmt.Println(bytes.Contains(b, []byte("xu")))                    //do this contains like xu?
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'f')) //UTF-8
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'ö'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大'))
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))      //return the index if is present ,or -1
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('k')))      //一个字节的时候
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))            //UTF-8
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go"))) //从后往前查找

}

func Test_Other(t *testing.T) {
	b := []byte("test_xujiahui.com")
	clone := bytes.Clone(b) //克隆，内存地址是不同的
	fmt.Println(&b == &clone)

	fmt.Println(bytes.Count([]byte("cheese"), []byte("e"))) //Count，查看出现了几次该元素

}
