package sync_demo

import (
	"fmt"
	"sync"
	"testing"
)

var a int

// Test_Mutex 互斥锁 mutual exclusion lock
// 使用互斥锁完成多线程下多 变量的增加
func Test_Mutex(t *testing.T) {
	var wg sync.WaitGroup
	var mu sync.Mutex
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
	fmt.Println(a)
}
