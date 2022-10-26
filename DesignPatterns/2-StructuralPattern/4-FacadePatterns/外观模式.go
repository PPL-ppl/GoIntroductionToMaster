package main

import "fmt"

// 电视机

type Tv struct {
}

func (t *Tv) On() {

}
func (t *Tv) Off() {

}

// VoiceBox音响

type VoiceBox struct {
}

func (t *VoiceBox) On() {

}
func (t *VoiceBox) Off() {

}

// Light灯光

type Light struct {
}

func (t *Light) On() {

}
func (t *Light) Off() {

}

// XBox

type XBox struct {
}

func (t *XBox) On() {

}
func (t *XBox) Off() {

}

// MicroPhone

type MicroPhone struct {
}

func (t *MicroPhone) On() {

}
func (t *MicroPhone) Off() {

}

// MicroPhone

type Projector struct {
}

func (t *Projector) On() {

}
func (t *Projector) Off() {

}

type HomePlayerFacade struct {
	tv         Tv
	voiceBox   VoiceBox
	light      Light
	xBox       XBox
	microPhone MicroPhone
	projector  Projector
}

// 进入Ktv模式

func (hm *HomePlayerFacade) DoKtv() {
	fmt.Println("家庭影院进入KTV模式")
	hm.tv.On()
	hm.projector.On()
	hm.microPhone.On()
}

// 进入游戏模式

func (hm *HomePlayerFacade) DoGame() {
	fmt.Println("家庭影院进入Game模式")
	hm.xBox.On()
	hm.projector.On()
	hm.microPhone.On()
}

func main() {
	homePlayerFacade := new(HomePlayerFacade)
	homePlayerFacade.DoKtv()
}
