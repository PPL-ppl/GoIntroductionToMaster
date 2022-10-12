package main

import "fmt"

func main() {
	s := Student{
		"name",
	}
	name := s.GetName()
	fmt.Print(name)
	t := T{}
	t.M3()

}

type Student struct {
	name string
}

func (s Student) GetName() string {
	return s.name
}

type T struct {
}

func (T) M1() {}
func (T) M2() {}
func (T) M3() {
	fmt.Println("222")
}

type E interface {
	M1()
	M2()
}
type F interface {
	M3()
}
type I interface {
	E
	F
	M1()
	M2()
	M3()
}
type E1 interface {
	M1()
	M2()
	M3()
}
type E2 interface {
	M1()
	M2()
	M4()
}
type E3 struct {
	E1
	E2
}
