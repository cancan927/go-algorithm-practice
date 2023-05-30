package main

import "fmt"

/**
使用goroutine按序依次打印一百次dog，cat，fish，使用channel控制主携程等待
*/

func main() {
	dogCh := make(chan struct{})
	catCh := make(chan struct{})
	fishCh := make(chan struct{})

	done := make(chan struct{})

	go DogPrint(dogCh, catCh)
	go CatPrint(catCh, fishCh)
	go FishPrint(fishCh, dogCh, done)
	dogCh <- struct{}{}

	<-done
}

func DogPrint(in <-chan struct{}, out chan<- struct{}) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Printf("第%d次Dog\n", i)
		out <- struct{}{}
	}
}
func CatPrint(in <-chan struct{}, out chan<- struct{}) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Printf("第%d次Cat\n", i)
		out <- struct{}{}
	}
}

func FishPrint(in <-chan struct{}, out chan<- struct{}, end chan<- struct{}) {
	for i := 0; i < 100; i++ {
		<-in
		fmt.Printf("第%d次Fish\n", i)
		if i < 99 {
			out <- struct{}{}
		} else {
			end <- struct {
			}{}
		}
	}
}
