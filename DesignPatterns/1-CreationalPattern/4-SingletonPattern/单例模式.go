package main

// 使用场景 只能有一个对象的地方  1 日志

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 保证一个类永远只能有一个对象，且该对象功能依然能被其他模块使用

// 保证类非公有化，外界不能通过这个类直接创建对象
type singleton struct{}

// 私有指针，不能被外部模块访问
var instance *singleton = new(singleton) //饿汉式
func GetInstance() *singleton {
	return instance
}

var instance2 *singleton //懒汉式   Lock和原子同时使用 线程安全的

var initialized uint32
var lock sync.Locker

func GetInstance2() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance2
	}
	lock.Lock()
	defer lock.Unlock()
	if instance2 == nil {
		instance2 = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance2
}

var instance3 *singleton //Go语言专属 Once 线程安全的

var once sync.Once

func GetInstance3() *singleton {
	once.Do(func() {
		instance3 = new(singleton)
	})
	return instance3
}
func main() {
	getInstance1 := GetInstance()
	getInstance2 := GetInstance()
	if getInstance1 == getInstance2 {
		fmt.Println("对的")
	}
}
