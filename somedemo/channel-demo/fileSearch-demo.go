package main

import (
	"fmt"
	"os"
	"time"
)

// 使用携程来完成文件夹扫描
var workerCount = 0                   //当前工人计数器
var maxWorkerCount = 32               //最多允许有32个工人
var searchRequest = make(chan string) //包工头给工人指派活
var workerDone = make(chan bool)      //工人告诉包工头活干完了
var foundMatch = make(chan bool)      //传输找到搜索结果的消息
var query = "linked-list"             //文件夹名
var matches = 0                       //文件计数器

func main() {
	startTime := time.Now()
	workerCount = 1
	go Search("/Users/jiahui.xu.ex/GolandProjects/", true)
	waitForWorkers()
	fmt.Println("matches", matches)
	fmt.Println(time.Since(startTime))
}
func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest: //有请求消息了，新建携程分配任务
			workerCount++
			go Search(path, true)
		case <-workerDone: //有人干完了
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch: //有人返回结果了
			matches++
		}
	}
}
func Search(path string, master bool) { //master 来告诉主函数运行完毕了
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				//喊话，说我找到了，
				foundMatch <- true
			}
			if file.IsDir() {
				//看看有没有空闲的工人
				if workerCount < maxWorkerCount {
					//有，往searchRequest里传输路径，叫他派工人
					searchRequest <- path + name + "/"
				} else {
					Search(path+name+"/", false)
				}
			}
		}
		if master {
			workerDone <- true
		}

	} else {
		fmt.Println(err)
		panic(err)
	}
}
