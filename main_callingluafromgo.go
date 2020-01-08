package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

const luaCode = `
function concat(a, b)
	return a .. " + " .. b
end
`

func mainCallingLuaFromGo(){
	L := lua.NewState()
	defer L.Close()
	// 在 L 裡運行該 luaCode 編譯並執行
	if err := L.DoString(luaCode);err != nil {
		panic(err)
	}

	// 編譯執行之後, 使用CallByParam來使用luaCode的方法
	// 傳入參數分2類, 1.要呼叫的方法資訊, 2.往後放入要傳入的參數列
	if err := L.CallByParam(lua.P{
		Fn:L.GetGlobal("concat"),// Lua的方法名
		NRet:    1,                     // 返回參數的數量
		Protect: true,                  // 是否會返回error或是panic
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}


	// 從stack裡取值並把值轉型成lua.LString
	// -1為stack的最上一面一層, 依序為-2, -3 以此類推
	if str, ok := L.Get(-1).(lua.LString); ok {
		fmt.Println(str)
	}

	// 從stack裡Pop多少數量出來
	// 若沒Pop出來再使用一次L.Get(-1)會取得跟上面一樣的值
	// 若有Pop出來取不到值, 前提是stack上僅有一個回傳值
	L.Pop(1)
}
