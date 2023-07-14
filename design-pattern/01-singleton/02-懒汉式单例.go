package _1_singleton

import (
	"sync"
	"sync/atomic"
)

type singleton2 struct {
}

//定义一个锁
var lock sync.Mutex

//标记位
var initialized uint32

var instance2 *singleton2

func GetInstance2() *singleton2 {
	//只使用锁的话性能很差
	if atomic.LoadUint32(&initialized) == 1 { //标记为1，已经被初始化了
		return instance2
	}
	lock.Lock()
	defer lock.Unlock()
	//只有首次调用才生成这个单例对象
	if instance2 == nil { //并发下可能会出现多个协程进入这个分支，因此需要锁
		instance2 = new(singleton2)
		//设置标记位
		atomic.StoreUint32(&initialized, 1)
		return instance2
	}
	return instance2
}
