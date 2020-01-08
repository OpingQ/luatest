package luacallmodule

import (
	lua "github.com/yuin/gopher-lua"
	"strconv"
)

func Add(L *lua.LState) int{
	// Lua 的參數從 1 開始, 依序取變數
	a1 := L.ToString(1)
	a2 := L.ToInt(2)
	a3 := L.ToInt(3)

	L.Push(lua.LString(a1+" : " + strconv.Itoa(a2+a3)))
	// 返為參數為1個
	return 1
}

func Muti(L *lua.LState) int{
	// Lua 的參數從 1 開始
	a1 := L.ToString(1)
	a2 := L.ToInt(2)
	a3 := L.ToInt(3)

	L.Push(lua.LString(a1+" : " + strconv.Itoa(a2*a3)))
	// 返為參數為1個
	return 1
}