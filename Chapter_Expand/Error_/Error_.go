package main

import (
	"errors"
	"fmt"
)

func main1() {
	err := fun1(2)
	fmt.Println(err)
	fun2()
	fmt.Println(3)
}

func fun1(int) error {
	err := errors.New("your first demo error")
	//err = nil
	return err
}
func fun2() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	fmt.Println(1)
	panic("hello panic")
	fmt.Println(2)
}
