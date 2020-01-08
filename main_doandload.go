package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"usedllfile/luacall"
)

func mainDoAndLoad(){
	// 創建時候都會消耗資源
	// 創建虛擬機
	l := lua.NewState()
	// 關閉虛擬機
	defer l.Close()
	
	luacall.DoStringOnLua(l,`print("Hello World")`)
	luacall.DoFileOnLua(l,`luafolder/hello.lua`)

	fmt.Println("-------------------------------")
	luacall.LoadStringOnLua(l,`print("Hello World")`)
	luacall.LoadFileOnLua(l,`luafolder/hello.lua`)
	//luacall.LoadOnLua(l,`luafolder/hello.lua`,"<string>")
}