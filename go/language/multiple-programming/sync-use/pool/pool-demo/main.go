package main

import (
	"sync"
	"fmt"
)

func main() {
	// 定义对象池中的类型
	type instance struct {
		body []byte
	}

	var pool sync.Pool

	pool.New = func() any {
		return &instance{
			body: []byte("pool"),
		}
	}

	for i := 0 ; i < 10000 ; i ++ {
		inst, _ := pool.Get().(*instance)
		fmt.Printf("body: %s\n", inst.body)
		pool.Put(inst)
	}

}