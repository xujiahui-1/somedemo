package crypto_rand_demo

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"testing"
)

/*
rand包实现了用于加解密的更安全的随机数生成器。
	包里有个var Reader io.Reader
		Reader是一个全局、共享的密码用强随机数生成器。在Unix类型系统中，会从/dev/urandom读取；而Windows中会调用CryptGenRandom API。



*/

// Test_Reader 使用reader
func Test_Reader(t *testing.T) {
	b := make([]byte, 32)
	//ReadFull从rand.Reader中精确的读取len(b)个字节填充进b
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		fmt.Println(err)
	}
	str := base64.URLEncoding.EncodeToString(b)
	fmt.Println(b)
	fmt.Println(str) //转码成了string
}

// Test_Int  func Int() : 返回一个在[0, max)区间服从均匀分布的随机值，如果max<=0则会panic。
func Test_Int(t *testing.T) {
	//从128开始，这样就能够将(max.BitLen() % 8) == 0的情况包含在里面
	for n := 128; n < 140; n++ {
		b := new(big.Int).SetInt64(int64(n)) //将new(big.Int)设为int64(n)并返回new(big.Int)
		fmt.Printf("max Int is : %v\n", b)
		i, err := rand.Int(rand.Reader, b)
		if err != nil {
			fmt.Printf("Can't generate random value: %v, %v", i, err)
		}
		fmt.Printf("rand Int is : %v\n", i)
	}
}
