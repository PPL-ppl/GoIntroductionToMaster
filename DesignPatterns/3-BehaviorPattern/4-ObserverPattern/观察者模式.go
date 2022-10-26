package main

import "fmt"

// 抽象观察者   不要让Listen触发事件不然会循环奔溃

type Listener interface {
	ToDo()
}

// 被观察者
type Notifier interface {
	AddListener(listener Listener)
	Remove(listener Listener)
	Notify()
}
type StudentZhangSan struct {
	BadThing string
}

func (s *StudentZhangSan) ToDo() {
	fmt.Println("吃啥1")
}
func (s *StudentZhangSan) ToBad() {
	fmt.Println(s.BadThing)
}

type StudentLisi struct {
	BadThing string
}

func (s *StudentLisi) ToDo() {
	fmt.Println("吃啥2")
}
func (s *StudentLisi) ToBad() {
	fmt.Println(s.BadThing)
}

type StudentNotifier struct {
	listener []Listener
}

func (s *StudentNotifier) AddListener(listener Listener) {
	s.listener = append(s.listener, listener)
}

func (s *StudentNotifier) Remove(listener Listener) {

	for i, l := range s.listener {
		if l == listener {
			//拼接slice
			s.listener = append(s.listener[:i], s.listener[i+1:]...)
		}
	}

}

func (s *StudentNotifier) Notify() {
	for _, l := range s.listener {
		l.ToDo()
	}
}

func main() {
	zhangSan := new(StudentZhangSan)
	zhangSan.BadThing = "hello1"
	zhangSan.ToBad()
	lisi := new(StudentLisi)
	lisi.BadThing = "hello2"
	lisi.ToBad()
	notifier := new(StudentNotifier)
	notifier.AddListener(zhangSan)
	notifier.AddListener(lisi)
	notifier.Notify()
}
