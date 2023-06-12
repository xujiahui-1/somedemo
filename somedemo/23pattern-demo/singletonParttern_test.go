package main

import "sync"

// 单例模式 singleton parttern
/*
 单例模式:顾名思义空置实例数量， 确保一个类型只有一个实例
	饿汉模式：程序初始化时就创建实例，不管他用不用 Eager Initialization
		推荐使用init()函数
	懒汉模式：就是延迟加载，只有用的时候才创建类型实例 Lazy Initialization

*/
type Instance struct {
	Name string
}

// EagerInstance 饿
func EagerInstance() *Instance {
	return MyIns
}

var MyIns *Instance

func init() {
	MyIns = &Instance{Name: "徐家汇"}
}

//	懒:我们必须使用原子性操作来防止并发情况下我们的实例被多次创建
//
// 我们可以借助其sync库中自带的并发同步原语Once来实现
var instance *Instance
var once sync.Once

func GetLazyInstance() *Instance {
	once.Do(func() {
		instance = &Instance{Name: "xxx"}
	})
	return instance
}
