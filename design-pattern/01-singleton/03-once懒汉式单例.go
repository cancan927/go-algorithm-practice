package _1_singleton

import "sync"

var once sync.Once

type singleton3 struct {
}

var instance3 *singleton3

func GetInstance3() *singleton3 {
	once.Do(func() {
		instance3 = new(singleton3)
	})
	return instance3
}
