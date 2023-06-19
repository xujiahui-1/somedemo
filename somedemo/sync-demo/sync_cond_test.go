package sync_demo

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Test_Cond 用来想要共享资源的那些携程，但是一些简单场景下还是推荐使用channel
func Test_Cond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{}) //创建cond实例
	for i := 0; i < 10; i++ {
		go listen(c) //十个携程调用
	}
	time.Sleep(time.Second)
	go broadcast(c) //唤醒

	ck := make(chan os.Signal, 1)
	signal.Notify(ck, os.Interrupt)
	<-ck
}

var status uint32

func listen(c *sync.Cond) {
	c.L.Lock() //加锁

	for atomic.LoadUint32(&status) != 1 { //原子读取
		c.Wait() //开始等待
	}
	fmt.Println("等待结束")
	c.L.Unlock()
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreUint32(&status, 1)
	c.Broadcast() //唤醒所有等待
	c.L.Unlock()
}
