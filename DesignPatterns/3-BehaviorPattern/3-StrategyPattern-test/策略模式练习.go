package main

import "fmt"

type planA struct {
}

func (p *planA) Go() {
	fmt.Println("打八折")
}

type planB struct {
}

func (p *planB) Go() {
	fmt.Println("打六折")
}

type plan interface {
	Go()
}

type Company struct {
	m plan
}

func (cpm *Company) SetPlan(plan2 plan) {
	cpm.m = plan2
}
func main() {
	cm := new(Company)
	a := new(planA)
	cm.SetPlan(a)
	cm.m.Go()
	b := new(planB)
	cm.SetPlan(b)
	cm.m.Go()
}
