package main

import (
	"fmt"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

// Go调用Lua
func main1() {
	lu := lua.NewState()
	defer lu.Close()

	err := lu.DoFile("D:\\GoCode\\GoIntroductionToMaster\\GoWithLua_gopher-lua\\main.lua")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = lu.CallByParam(lua.P{
		Fn:      lu.GetGlobal("add"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(1), lua.LNumber(2))

	if err != nil {
		fmt.Println(err.Error())
	}
	ret := lu.Get(-1)
	lu.Pop(1)
	//两个返回值
	// ret := lu.Get(2)
	//	lu.Pop(2)
	res, ok := ret.(lua.LNumber)
	if ok {
		fmt.Println(res)
	}
}

// Lua调用Go方法
func main2() {
	lu := lua.NewState()
	defer lu.Close()

	lu.SetGlobal("add", lu.NewFunction(add))

	//go
	err := lu.DoFile("D:\\GoCode\\GoIntroductionToMaster\\GoWithLua_gopher-lua\\main2.lua")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
func add(L *lua.LState) int {
	arg1 := L.ToInt(1)
	arg2 := L.ToInt(2)
	ret := arg1 + arg2
	L.Push(lua.LNumber(ret))
	return 1
}

// LuaTable转为GoStruct

type Role struct {
	Name string
}
type Person struct {
	Name      string
	Age       string
	WorkPlace string
	Role      []*Role
}

func main() {
	L := lua.NewState()
	if err := L.DoString(`
		person = {
		  name = "Michel",
		  age  = "31", -- weakly input
		  work_place = "San Jose",
		  role = {
			{
			  name = "Administrator"
			},
			{
			  name = "Operator"
			}
		  }
		}
`); err != nil {
		panic(err)
	}
	var person Person
	if err := gluamapper.Map(L.GetGlobal("person").(*lua.LTable), &person); err != nil {
		panic(err)
	}
	fmt.Printf("%s %s", person.Name, person.Age)
}
