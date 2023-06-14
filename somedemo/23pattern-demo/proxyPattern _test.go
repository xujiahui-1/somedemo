package main

import (
	"fmt"
	"testing"
)

/*
代理模式
	代理模式是一种结构型设计模式。 其中代理控制着对于原对象的访问， 并允许在将请求提交给原对象的前后进行一些处理，从而增强原对象的逻辑处理。
	代理模式能够在不改变原始类（或叫被代理类）代码的情况下，通过引入代理类来给原始类附加功能。 一般代理类和被代理类有同一个父类。

*/

// 小汽车
type Car struct {
}

// 汽车行为接口
type Vehicle interface {
	Drive() //驾驶方法
}

// 小汽车实现Driver()方法
func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

//如果想判断开车人是否成年，
//给Car结构体加一个Age字段显然是不合理的，因为我们要表示的驾驶员的年龄而不是车的车龄。
//同理驾驶员年龄的判断我们也不应该加在 Car 实现的 Drive() 方法里，
//这样会导致每个实现 Vehicle 接口的类型都要在自己的 Drive() 方法里加上类似的判断

// 驾驶员
type Driver struct {
	Age int
}

// 使用代理把车和驾驶员包装起来
type CarProxy struct {
	Vehicle Vehicle
	Driver  *Driver
}

// 创建代理
func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

// 这样的话我们接可以通过，用包装类型代理vehicle属性的 Drive() 行为时，给它加上驾驶员的年龄限制。
func (c *CarProxy) Drive() {
	if c.Driver.Age >= 16 {
		c.Vehicle.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}
func Test_(t *testing.T) {
	car := NewCarProxy(&Driver{12})
	car.Drive() // 输出 Driver too young!
	car2 := NewCarProxy(&Driver{22})
	car2.Drive() // 输出 Car is being driven
}
