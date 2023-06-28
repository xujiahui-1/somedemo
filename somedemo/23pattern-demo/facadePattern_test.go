package main

import (
	"fmt"
	"testing"
)

//外观模式
/*
医院设立的接待员的角色就是我们今天要介绍的外观模式，
系统通过引入外观模式让需要调用多个子系统各自部分的功能接口以完成的需求，
变为调用方只需要跟外观提供的统一功能进行交互即可
*/

//例子：有台电脑，我们如果需要启动的化，需要启动cpu，内存，硬盘，所以我们提供一个开机键，一键启动所有

const (
	BOOT_ADDRESS = 0
	BOOT_SECTOR  = 0
	SECTOR_SIZE  = 0
)

// CPU
type CPU struct {
}

func (c *CPU) Freeze() {
	fmt.Println("Freeze")
}
func (c *CPU) Jump(position int) {
	fmt.Println("CPU.Jump()")
}

func (c *CPU) Execute() {
	fmt.Println("CPU.Execute()")
}

// Memory
type Memory struct{}

func (m *Memory) Load(position int, data []byte) {
	fmt.Println("Memory.Load()")
}

// HardDrive
type HardDrive struct{}

func (hd *HardDrive) Read(lba int, size int) []byte {
	fmt.Println("HardDrive.Read()")
	return make([]byte, 0)
}

// 门面结构体
type ComputerFacade struct {
	processor *CPU
	ram       *Memory
	hd        *HardDrive
}

// 创建一个门面
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{new(CPU), new(Memory), new(HardDrive)}
}

// start 封装所有调用
func (c *ComputerFacade) start() {
	c.processor.Freeze()
	c.ram.Load(BOOT_ADDRESS, c.hd.Read(BOOT_SECTOR, SECTOR_SIZE))
	c.processor.Jump(BOOT_ADDRESS)
	c.processor.Execute()
}

func Test_Facde(t *testing.T) {
	//既可以调用start，又可以直接调用内部的下级接口
	com := NewComputerFacade()
	com.start()
	com.processor.Execute()
}
