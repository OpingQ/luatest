package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"testing"
	"time"
	"usedllfile/luacall"
)

func TestSetReadOnly(t *testing.T) {
	L := lua.NewState()
	L2 := lua.NewState()
	defer L.Close()
	defer L2.Close()
	s := time.Now()
	if lfunc,err := L.LoadString(`a = 1 + 1`); err != nil{
		fmt.Println(err)
		return
	}else{
		for i:= 0; i < 1000; i++{
			L2.Push(lfunc)
			L2.PCall(0,lua.MultRet,nil)
		}
	}
	fmt.Println(time.Since(s))
}

func Test_loadModule(t *testing.T) {
	L := lua.NewState()
	L2 := lua.NewState()
	defer L.Close()
	defer L2.Close()
	s := time.Now()
	proto,_ := luacall.CompileString(`a = 1 + 1`)
	lfunc := L.NewFunctionFromProto(proto)
	for i:= 0; i < 1000; i++{
		L2.Push(lfunc)
		L2.PCall(0,lua.MultRet,nil)
	}

	fmt.Println(time.Since(s))
}

func Test_mainCallingLuaFromGo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_mainCompileCall(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_mainDoAndLoad(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_moduleCall2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_moduleCall201(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_moduleCall202(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}



//
//func loadDllSum(){
//	dll32 := syscall.NewLazyDLL(`test.dll`)
//	g := dll32.NewProc("Sum")
//	arg1 := uintptr(5)
//	arg2 := uintptr(159)
//	r, _, _ := g.Call(arg1,arg2)
//	fmt.Println(r)
//}
