package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 锁
var lock sync.Mutex

// 数量
var count int

func main() {
	// 开始时间
	startTime := time.Now().Unix()
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				TestRequest()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	endTime := time.Now().Unix()
	fmt.Printf("请求时间: %vs", endTime-startTime)
	fmt.Printf("请求次数: %v", count)
}

func TestRequest() {
	response, _ := http.Get("http://localhost:8080/user/list")
	all, _ := io.ReadAll(response.Body)
	// 加锁
	lock.Lock()
	// 数量加一
	count++
	// 解锁
	lock.Unlock()
	go func() {
		fmt.Println(string(all))
	}()
}
