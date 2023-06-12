package struct_demo

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
空结构体特点：零内存占用，地址相同，无状态
1.为什么他零内存占用呢？
	在golang中，存在一个变量叫zerobase，这是一个 uintptr 全局变量，占用 8 个字节。
	当你定义了 内存 size 为 0 的内存分配，那么golang就会把这个&zerobase这个地址付给他
	所以无论你有多少个空结构体，他底层都是这个zerobase
*/

// Test_size 测试空结构体内存大小
func Test_size(t *testing.T) {
	//声明, 匿名空结构体
	var e struct{}
	//命名空结构体
	type EmptyStruct struct{}
	var b EmptyStruct

	var str string
	//测试一下是否真不占用内存
	fmt.Println(unsafe.Sizeof(e))   //0
	fmt.Println(unsafe.Sizeof(b))   //0
	fmt.Println(unsafe.Sizeof(str)) //16

}

// Test_address 测试两个空结构体的内存地址是否一样
func Test_address(t *testing.T) {
	a := struct{}{}
	c := struct{}{}
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &c)

}

// Test_diff 测试内置字段时地址如何分配
func Test_diff(t *testing.T) {
	type Object1 struct {
		s struct{}
		b byte
	}

	// Object2 类型变量占用 8 个字节
	type Object2 struct {
		s struct{}
		n int64
	}
	//struct {} 在最前面
	o1 := Object1{}
	o2 := Object2{}
	//发现和之后的变量内存大小是一致的
	fmt.Println(unsafe.Sizeof(o1)) //1
	fmt.Println(unsafe.Sizeof(o2)) //8
}

// Test_diff2 struct {} 在中间
func Test_diff2(t *testing.T) {
	type Object1 struct {
		b  byte
		s  struct{}
		b1 int64
	}

	o1 := Object1{}
	//因为要满足8个字节的对其规则，所有是16，但是编译器不会对 struct { } 做任何字节填充。
	fmt.Println(unsafe.Sizeof(o1)) //16
}

// Test_diff3 struct {} 在最后
func Test_diff3(t *testing.T) {
	type Object1 struct {
		b byte
		s struct{}
	}

	type Object2 struct {
		n int64
		s struct{}
	}

	type Object3 struct {
		n int16
		m int16
		s struct{}
	}

	type Object4 struct {
		n int16
		m int64
		s struct{}
	}

	o1 := Object1{}
	o2 := Object2{}
	o3 := Object3{}
	o4 := Object4{}
	//编译器在遇到这种 struct {} 在最后一个字段的场景，
	//会进行特殊填充，struct { } 作为最后一个字段，会被填充对齐到前一个字段的大小，地址偏移对齐规则不变
	fmt.Println(unsafe.Sizeof(o1)) //2
	fmt.Println(unsafe.Sizeof(o2)) //16
	fmt.Println(unsafe.Sizeof(o3)) //6
	fmt.Println(unsafe.Sizeof(o4)) //24

	type Part1 struct {
		a bool  //1
		b int32 //4
		c int8  //1
		d int64 //8
		e byte  //1
	}
	o5 := Part1{}
	fmt.Println(unsafe.Sizeof(o5))
}

// Test_offset 查看变量的对齐值
func Test_offset(t *testing.T) {
	type Demo struct {
		a int32
		b bool
		c int16
	}
	demo := Demo{}
	alignOfDemo := unsafe.Alignof(demo)
	fmt.Printf("Demo 对齐值：%d\n", alignOfDemo)

	a := int32(0)
	alignOfA := unsafe.Alignof(a)
	fmt.Printf("a 对齐值：%d\n", alignOfA)

	b := true
	alignOfB := unsafe.Alignof(b)
	fmt.Printf("b 对齐值：%d\n", alignOfB)
}
