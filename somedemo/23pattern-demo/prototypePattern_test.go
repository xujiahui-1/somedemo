package main

import (
	"fmt"
	"testing"
)

/*
原型模式
	如果一个类的有非常多的属性，层级还很深。每次构造起来，不管是直接构造还是用建造者模式，都要对太多属性进行复制，那么推荐使用原型模式
	通过复制、拷贝或者叫克隆已有对象的方式来创建新对象的设计模式叫做原型模式，被拷贝的对象也被称作原型对象。

	所有的原型对象都要暴露一个Clone方法，这个方法给外部调用者一个机会，来从自己这里0成本克隆新对象
*/

// 原型接口:那么动物必须实现这个接口
type Cloneable interface {
	Clone() Cloneable
}

// 克隆实验室 我们有一个动物实验室，可以克隆各种动物，
type CloneLab struct {
	animals map[string]Cloneable
}

// 创建克隆实验室
func NewCloneLab() *CloneLab {
	return &CloneLab{animals: make(map[string]Cloneable)}
}

// 获取一只克隆羊
func (p *CloneLab) Get(name string) Cloneable {
	return p.animals[name]
}

// set动物当前属性，//这里值是Cloneable对象，而所有动物都实现了这个接口，所以调用这个方法的时候，传入任何动物都可以
func (p *CloneLab) Set(name string, prototype Cloneable) {
	p.animals[name] = prototype
}

// 测试
var lab *Cloneable

// 羊
type Sheep struct {
	name   string
	weight int
}

// 羊实现Cloneable接口
func (s *Sheep) Clone() Cloneable {
	tc := *s
	return &tc
}

// Test_clone 使用动物直接克隆
func Test_clone(t *testing.T) {
	sheep1 := &Sheep{
		name:   "sheep",
		weight: 10,
	}

	sheep2 := sheep1.Clone()
	fmt.Println(&sheep2)
	fmt.Println(&sheep1)
	if sheep2 == sheep1 {
		t.Fail()
	}
}

// Test_clone2 通过克隆实验室
func Test_clone2(t *testing.T) {
	cloneLab := NewCloneLab()

	cloneLab.Set("sheep", &Sheep{name: "sheep1", weight: 100}) //把动物放入克隆实验室
	c := cloneLab.Get("sheep").Clone()                         //获取
	cw := c.(*Sheep)
	c2 := cloneLab.Get("sheep").Clone() //获取
	cw2 := c2.(*Sheep)
	fmt.Println(&cw)
	fmt.Println(&cw2)
}
