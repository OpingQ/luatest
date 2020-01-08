package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"usedllfile/luacallmodule"
)
var command01 = `
local m = require("mathmodule")
local result = m.add("c01a",5,8)
print(result)
`

var command02 = `
local m = require("mathmodule")
local result = m.add("c01a",5,8)
print(result)
result = m.muti("c01m",5,8)
print(result)
print(m.gargs)
print(glargs)
-- m.gargs = "7797sdsd" 變成只讀後該變數不能被賦值
`

func moduleCall(){
	fmt.Println("01 Moudle------------")
	moduleCall201()
	fmt.Println("02 Moudle------------")
	moduleCall202()
}

func moduleCall201(){
	l := lua.NewState()
	defer l.Close()
	l.RegisterModule("mathmodule",exports01)
	if err := l.DoString(command01); err != nil {
		panic(err)
	}
}

func moduleCall202(){
	l := lua.NewState()
	defer l.Close()
	l.PreloadModule("mathmodule",loadModule)
	if err := l.DoString(command02); err != nil {
		panic(err)
	}
}

// 可供組合 包含有exports01 及 exports02 的集成
func loadModule(L *lua.LState) int{
	mod := L.NewTable()
	L.SetFuncs(mod, exports01)
	L.SetFuncs(mod, exoports02)

	// 把欄位設定在table中
	// 在 LTable 中設定值 key: name,  value: "gomodule789"
	L.SetField(mod,"gargs",lua.LString("modlocalArgs"))
	// 該LState裡的全域
	L.SetGlobal("glargs",lua.LString("GobalArgs"))

	// 可以把mod變成只讀
	L.Push(SetReadOnly(L,mod))

	return 1
}

// SetReadOnly 把該Table設定為ReadOnly
func SetReadOnly(l *lua.LState, table *lua.LTable) *lua.LUserData {
	ud := l.NewUserData()
	mt := l.NewTable()
	// 设置表中域的指向为 table
	l.SetField(mt, "__index", table)
	// 限制对表的更新操作
	l.SetField(mt, "__newindex", l.NewFunction(func(state *lua.LState) int {
		state.RaiseError("not allow to modify table")
		return 0
	}))
	ud.Metatable = mt
	return ud
}


var exports01 = map[string]lua.LGFunction{
	"add": luacallmodule.Add,
}

var exoports02 = map[string]lua.LGFunction{
	"muti": luacallmodule.Muti,
}