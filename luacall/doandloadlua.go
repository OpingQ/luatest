package luacall

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)


func DoStringOnLua(l *lua.LState,source string){
	if err := l.DoString(source);err != nil {
		panic(err)
	}
}

func DoFileOnLua(l *lua.LState,path string){
	if err := l.DoFile(path);err != nil {
		panic(err)
	}
}

func LoadStringOnLua(ls *lua.LState,source string){
	if fn,err := ls.LoadString(source); err != nil{
		fmt.Println(err)
		return
	} else {
		// 1.Push
		// 2.Call
		ls.Push(fn)
		err = ls.PCall(0, lua.MultRet, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func LoadFileOnLua(ls *lua.LState,path string){
	if fn,err := ls.LoadFile(path); err != nil{
		fmt.Println(err)
		return
	}else {

		ls.Push(fn)
		err = ls.PCall(0, lua.MultRet, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
