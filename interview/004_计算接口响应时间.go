package main

import (
	"fmt"
	"strconv"
	"time"
)

func timeSpent(handler func(a, b int) string) func(a, b int) string {
	return func(a, b int) string {
		start := time.Now()
		ret := handler(a, b)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func main() {
	timeFn := timeSpent(func(a, b int) string {
		time.Sleep(time.Second)
		return strconv.Itoa(a + b)
	})

	timeFn(3, 5)

}
