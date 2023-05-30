package main

import (
	"fmt"
	"sync"
)

/**
使用goroutine按序依次打印一百次dog，cat，fish，使用waitGroup控制主携程等待
*/

func main() {
	dogCh := make(chan struct{})
	catCh := make(chan struct{})
	fishCh := make(chan struct{})
	var wg = &sync.WaitGroup{}
	wg.Add(3)

	go PrintDog(dogCh, catCh, wg)
	go PrintCat(catCh, fishCh, wg)
	go PrintFish(fishCh, dogCh, wg)

	dogCh <- struct{}{}
	wg.Wait()
}

func PrintDog(in <-chan struct{}, out chan<- struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Printf("这是第%d次Dog\n", i)
		out <- struct{}{}
	}
	wg.Done()
}

func PrintCat(in <-chan struct{}, out chan<- struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Printf("这是第%d次Cat\n", i)
		out <- struct{}{}
	}
	wg.Done()
}

func PrintFish(in <-chan struct{}, out chan<- struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Printf("这是第%d次Fish\n", i)
		if i < 99 {
			out <- struct{}{}
		}
	}
	wg.Done()
}
