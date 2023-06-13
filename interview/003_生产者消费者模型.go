package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
基于channel实现一个生产者消费者模型
- 生产者每秒生产一个随机数
- 消费者每秒读取并打印生产出来的数
*/
func main() {

	queue := make(chan int, 10)
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go producer(queue)
	}
	for i := 0; i < 10; i++ {
		go consumer(queue)
	}
	<-done

}

func producer(queue chan<- int) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		queue <- rand.Intn(100)
	}
}

func consumer(queue <-chan int) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		fmt.Println(<-queue)
	}
}
