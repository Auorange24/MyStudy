package main

import (
	"sync"
)

func main() {
	type instance struct {
		body []byte
	}

	var pool sync.Pool

	pool.New = func() any {
		return &instance{
			body: []byte{
				"key": "value",
			},
		}
	}

	for i := 0 ; i < 10000 ; i ++ {
		inst, _ := pool.Get().(*instance)
	}

}