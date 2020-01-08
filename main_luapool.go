package main

import (
	"usedllfile/luacall"
)

func mainLuaPool(){
	pool := luacall.LuaPool
	// 初始化池子的大小
	for i := 0; i < 10; i++ {
		L := pool.New()
		pool.Put(L)
	}
	// 取出來作使用
	s := pool.Get()
	s.DoString(`a = 1 + 1`)
	// 放回池子中
	pool.Put(s)
}

