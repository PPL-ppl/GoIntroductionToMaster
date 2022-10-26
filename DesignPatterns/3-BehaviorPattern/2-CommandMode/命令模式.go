package main

import "fmt"

// 医生

type Doctor struct {
}

func (doctor *Doctor) treatEye() {
	fmt.Println("治疗眼睛")
}
func (doctor *Doctor) treatNose() {
	fmt.Println("治疗鼻子")
}

// 治疗眼睛订单

type CommandTreatEye struct {
	docter *Doctor
}

func (c *CommandTreatEye) Treat() {
	c.docter.treatEye()
}

// 治疗鼻子订单

type CommandTreatNose struct {
	docter *Doctor
}

func (c *CommandTreatNose) Treat() {
	c.docter.treatNose()
}

//抽象的病单 -命令

type Command interface {
	Treat()
}

// 护士 分配病单

type Nurse struct {
	CmdList []Command //命令集合
}

// 发送病单

func (n Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat() //多态 调用具体的命令
	}
}
func main() {
	//医生
	d := new(Doctor)
	//将这个订单给这个医生
	eye := CommandTreatEye{d}
	eye.Treat()

	// 引入护士
	cmdEye := CommandTreatEye{d}
	cmdNose := CommandTreatNose{d}
	n := new(Nurse)
	n.CmdList = append(n.CmdList, &cmdEye, &cmdNose)
	n.Notify() //多态调用
}
