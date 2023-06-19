package sync_demo

import (
	"fmt"
	"sync"
	"testing"
)

// Test_Once 函数，无论你调用多少次，他都只执行一次
// 如果你的Once.Do函数内再次调用Do。那么会导致死锁
func Test_Once(t *testing.T) {
	one := sync.Once{}

	var a = func() {
		fmt.Println("测试一次")
		one.Do(func() {
		})
	}
	one.Do(a)
	one.Do(a)
	one.Do(a)
}

/*
https://juejin.cn/post/7189966589042556989
sync.Map 是一个读写分离的概念
它里面包括read map和dirty map，
read是只读的 dirty map是可读写的

基本就是先找read map，如果没找到，在加锁找dirty map
map 结构体中维护了一个misses，记录read map中没有找到的次数，到达一定次数，会将dirty map的数据同步到 read map
还有就是在使用Range遍历的时候，也会同步，同步后dirty map会值为nil，然后misses会初始化为0


amended： read map中维护了amended，这个标识来判断我们的read map和 dirty mop是否完全一致
因为我们写入新key时，会先写入dirty mop，不会写入read map，所以我们查找的时候，就要加锁查dirty map
但是每次都加锁，太浪费资源了，所以如果read map和 dirty mop是否完全一致，我们就不需要加锁，
不一致则需要加锁， amended就是来判断一不一致的bool

*/
// Test_Map sync.Map
// 当给定密钥的条目只写入一次但读取多次时
// 当多个 goroutine 对不相交的键集读、写和覆盖条目时,
// 这两种情况推荐使用 sync.Map, 他可以极大的减少Mutex or RWMutex.的锁竞争
// Map 在第一次使用后不能被复制
func Test_Map(t *testing.T) {
	mymap := sync.Map{}
	//写入
	mymap.Store("key1", "value1")
	mymap.Store("key2", "value2")
	//读取
	value, ok := mymap.Load("key1")
	fmt.Println(value, ok)
	mymap.Load("key2")
}

func Test_Range(t *testing.T) {
	a := sync.Map{}
	a.Store("1", 1)
	a.Store("2", 2)
	a.Store("3", 3)

	a.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

// Test_DeleteAndLoadAndDelete Delete 底层就是调用了LoadAndDelete
// LoadAndDelete会返回该key的值和该key是否存在，不存在，则返回  <nil> false
// sync.Map的删除操作并不是真正删除，而是把底层的entry的状态设置成nil，这样就可以不加锁，实现删除
// 例外的情况是，它在 read map 中找不到，然后就需要加锁，然后做 double checking，然后再去 dirty map 中查找了。
// 当key存在于 read map 中，则直接删除。（设置 entry 指针为 nil，但是不会删除 read map 中的 key）
// 当key只存在于 dirty map 中，则直接删除。这种情况下，会删除 dirty map 中的 key。
func Test_DeleteAndLoadAndDelete(t *testing.T) {
	a := sync.Map{}
	a.Store("1", "一")
	a.Store("2", "二")
	a.Store("3", "三")

	value, loaded := a.LoadAndDelete("4")
	fmt.Println(value, loaded)
	a.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

// Test_LoadOrStore 存在则加载不存在则存入,
// key 在 read map 中找到，尝试在 read map 中 Load 或 Store，操作成功则返回。找不到则加锁，然后二次检查（double checking）。
// 在 read map 中依然找不到，但是 key 在 dirty map 中找到，尝试在 dirty map 中 Load 或 Store，操作成功则返回。（missLocked）
// key 不存在，往 dirty map 中写入 key 和 value。（如果 dirty map 为 nil，则先进行初始化），然后read map 的 amended 修改为 true。
func Test_LoadOrStore(t *testing.T) {
	a := sync.Map{}
	a.Store("1", "一")
	a.Store("2", "二")
	a.Store("3", "三")
	actual, loaded := a.LoadOrStore("1", "111")
	fmt.Println(actual, loaded)
	a.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

// Test_Swap Store 实际上是对 Swap 方法的调用 ,交换值
// 如果有，则修改，没有，则存进去
func Test_Swap(t *testing.T) {
	a := sync.Map{}
	a.Store("1", "一")
	a.Store("2", "二")
	a.Store("3", "三")
	previous, loaded := a.Swap("0", "二")
	fmt.Println(previous, loaded)
	a.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
