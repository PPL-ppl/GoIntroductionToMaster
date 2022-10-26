package main

import "fmt"

// 厨师

type Cook struct {
}

func (c *Cook) CookFish() {
	fmt.Println("CookFish")
}
func (c *Cook) CookCake() {
	fmt.Println("CookFish")
}
func (c *Cook) CookBlue() {
	fmt.Println("CookFish")
}

type CookFish struct {
	cook *Cook
}

func (c *CookFish) CookThing() {
	c.cook.CookFish()
}

type CookCake struct {
	cook *Cook
}

func (c *CookCake) CookThing() {
	c.cook.CookCake()
}

type CookBlue struct {
	cook *Cook
}

func (c *CookBlue) CookThing() {
	c.cook.CookBlue()
}

type CommandCoke interface {
	CookThing()
}
type Waiter struct {
	CaseList []CommandCoke
}

func (w Waiter) Notify() {
	if w.CaseList == nil {
		return
	}
	for _, case2 := range w.CaseList {
		case2.CookThing()
	}
}

func main() {
	c := new(Cook)
	cookFish := CookFish{c}
	cookBlue := CookBlue{c}
	cookCake := CookCake{c}
	w := new(Waiter)

	w.CaseList = append(w.CaseList, &cookCake, &cookFish, &cookBlue)
	w.Notify()
}
