package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"usedllfile/luacall"
)

func mainCompileCall() {
	L := lua.NewState()
	L2 := lua.NewState()
	defer L.Close()
	proto, err := luacall.CompileFile(`luafolder/hello.lua`)
	if err != nil {
		fmt.Println(err)
		return
	}
	luacall.DoCompiledFile(L, proto)
	luacall.DoCompiledFile(L2, proto)

}
