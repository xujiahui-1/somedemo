package main

import (
	"fmt"
	"testing"
)

//Adapter Pattern 适配器模式
//它的功能是将一个类的接口变成客户端所期望的另一种接口，从而使原本因接口不匹配而导致无法在一起工作的两个类能够一起工作，属于结构型设计模式

// 适配器与原对象（被适配对象）实现不同的接口，适配器的特点在于兼容，客户端通过适配器的接口完成跟自己不兼容的原对象的访问交互。
// 代理与原对象（被代理对象）实现相同的接口，代理模式的特点在于隔离和控制，代理直接转发原对象的返回给客户端，
// 但是可以在调用原始对象接口的前后做一些额外的辅助工作，AOP编程的实现也是利用这个原理。

// Adaptee 我们像用这个接口，但是与客户期待的不一致，所以需要一个适配器
type Adaptee interface {
	SpecificRequest() string
}

// Adaptee接口实现类
type adapteeImpl struct{}

// NewAdaptee  初始化函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// adapteeImpl 类实现  Adaptee接口
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// -------------------------------
// 适配器接口
type Target interface {
	Request() string
}

// Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee Adaptee
}

// 这里是精华，在适配器结构体实现适配器接口的方法里，调用 Adaptee 接口中的方法
func (a *adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

// NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}

// ------------------ 例子2---------------------
// 我们的接口（新接口）——音乐播放
type MusicPlayer interface {
	play(fileType string, fileName string)
}

// 在网上找的已实现好的库 音乐播放
// ( 旧接口）
type ExistPlayer struct {
}

func (*ExistPlayer) playMp3(fileName string) {
	fmt.Println("play mp3 :", fileName)
}
func (*ExistPlayer) playWma(fileName string) {
	fmt.Println("play wma :", fileName)
}

// 适配器
type PlayerAdaptor struct {
	// 持有一个旧接口
	existPlayer ExistPlayer
}

// 实现新接口
func (player *PlayerAdaptor) play(fileType string, fileName string) {
	switch fileType {
	case "mp3":
		player.existPlayer.playMp3(fileName)
	case "wma":
		player.existPlayer.playWma(fileName)
	default:
		fmt.Println("暂时不支持此类型文件播放")
	}
}
