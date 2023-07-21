package roaring_demo

import (
	"bytes"
	"fmt"
	"github.com/RoaringBitmap/roaring"
	"testing"
)

/*
https://darjun.github.io/2022/07/17/godailylib/roaring/
https://www.jb51.net/article/276679.htm
roaring
	地址 github.com/RoaringBitmap/roaring
	集合是软件中的基本抽象。
	实现集合的方法有很多，例如 hash set、tree等。要实现一个整数集合，
	位图（bitmap，也称为 bitset 位集合，bitvector 位向量）是个不错的方法。

roaring对的 bitset 进行了压缩，让他对memory的需求更小

	bitMap结构体中最主要的是维护可一个  []container这个一个接口，
		有三种container: arraycontainer,bitmapcontainer,runcontainer
			arraycontainer:最大值是4096，当容量超过这个值的时候会将当前 Container 替换为BitmapContainer。
				你可以在源码中看到这种代码bcRet := ac.toBitmapContainer()
			bitmapcontainer:存的是uint64的一个切片
				BitmapContainer 中无论存储了 1 个还是存储了 65536 个数据，其占用的空间都是同样的 8 kb （4096）
			RunContainer 又称行程长度压缩算法(Run Length Encoding)，在连续数据上压缩效果显著。
				RunContainer 原理在连续出现的数字，只会记录其初始数字和后续数量，举个例子：
				数列 22，它会压缩为 22,0；
				数列 22,23,24 它会压缩为 22,3；
				数列 22,23,24,32,33，它会压缩为 22,3,32,1；
				其中，short[] valueslength中存储的就是压缩后的数据。
				可以看出，这种压缩算法在性能和数据的连续性（紧凑性）关系极为密切，
				在连续的 100 个 short，可以将 200 字节压缩成 4 个 kb。
				对于不连续的 100 个 short，编码完之后会从 200 字节变为 400 kb。
				如果要分析RunContainer的容量，我们可以做下面两种极端的假设：
				最优情况，只存在一个数据或者一串连续数字，存储 2 个 short 会占用 4 kb。
				最差情况，0~65535 的范围内填充所有的不连续数字，(全部奇数位或全部偶数位)，需要存储 65536 个short 占用 128 kb。
*/

func Test_Roating(t *testing.T) {
	//创建位图并添加集合元素
	bitmap1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	fmt.Println(bitmap1.String())         //打印集合
	fmt.Println(bitmap1.GetCardinality()) //返回集合大小
	fmt.Println(bitmap1.Contains(5))      //检查输入的元素是否在集合中

	//创建空位图
	bitmap2 := roaring.New()
	bitmap2.Add(30)
	bitmap2.Add(31)
	bitmap2.Add(32)
	fmt.Println(bitmap2.String())

	//bitmap1 执行并集
	bitmap1.Or(bitmap2)
	fmt.Println(bitmap1)
	fmt.Println(bitmap2)

	//bitmap1 执行交集
	bitmap1.And(bitmap2)
	fmt.Println(bitmap1)
	fmt.Println(bitmap2)
}

// Test_Iterator roaring支持迭代
func Test_Iterator(t *testing.T) {
	bitmap := roaring.BitmapOf(2, 3, 4, 1, 5, 6)
	iterator := bitmap.Iterator() //获取迭代器
	for iterator.HasNext() {      //我们发现迭代输出的结果是有序的，说明他在内部已经将数据进行排序存储了
		fmt.Println(iterator.Next()) //1 2 3 4 5 6
	}
}

// Test_Goroutine roaring 支持位图集合运算的并行执行。可以指定使用多少个 goroutine 对集合执行交集、并集等
func Test_Goroutine(t *testing.T) {
	bitmap1 := roaring.BitmapOf(1, 2, 3)
	bitmap2 := roaring.BitmapOf(4, 5, 6)

	//取并集
	//第一个参数 parallelism 决定要使用多少worker
	and := roaring.ParAnd(4, bitmap1, bitmap2)
	fmt.Println(and)

	//交集
	or := roaring.ParOr(1, bitmap1, bitmap2)
	fmt.Println(or)

}

// Test_ReadAndWrite 对位图的I/O读写
func Test_ReadAndWrite(t *testing.T) {
	bm := roaring.BitmapOf(1, 3, 5, 7, 100, 300, 500, 700)

	buf := &bytes.Buffer{}
	bm.WriteTo(buf)

	newBm := roaring.New()
	newBm.ReadFrom(buf)
	if bm.Equals(newBm) {
		fmt.Println("write and read back ok.")
	}
}
