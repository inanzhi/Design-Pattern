package singleton_pattern

import "sync"

//import "sync"
//
//var mu sync.Mutex
//
//// Singleton 懒汉模式 非并发安全
//type Singleton struct {
//}
//
//var ins *Singleton
//
//func GetInsOr() *Singleton {
//	if ins == nil {
//		mu.Lock()
//		if ins == nil {
//			ins = &Singleton{}
//		}
//		mu.Unlock()
//	}
//
//	return ins
//}

var once sync.Once

// Singleton 懒汉模式 非并发安全
type Singleton struct {
}

var ins *Singleton

func GetInsOr() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
	})
	return ins
}
