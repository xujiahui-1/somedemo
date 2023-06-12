package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 一个或者多个操作在 CPU 执行的过程中不被中断的特性，称为原子性（atomicity） 。
// 这些操作对外表现成一个不可分割的整体，他们要么都执行，要么都不执行，外界不会看到他们只执行到一半的状态。
/*
	那么互斥锁和原子性操作的区别在哪里
		互斥锁用来保护一段逻辑，由操作系统的调度器实现
		原子操作保护一个变量的更新，由底层硬件指令直接提供支持
	使用sync/atomic 包 执行原子性操作， 基本包括，增减 ADD ，载入 Load ，存储  Store，比较并交换 CAS，交换

*/

// mutexAdd  使用互斥锁实现的并发计数器
func Test_mutexAdd(t *testing.T) {
	var a int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a += 1
			mu.Unlock()
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use mutex a is %d, spend time: %v\n", a, timeSpends)
	//use mutex a is 100000000, spend time: 286155862
}

// Test_atomicity 使用原则性操作
func Test_atomicity(t *testing.T) {
	var a int32 = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use atomic a is %d, spend time: %v\n", atomic.LoadInt32(&a), timeSpends)
	//use mutex a is 1000000, spend time: 267207300
}
