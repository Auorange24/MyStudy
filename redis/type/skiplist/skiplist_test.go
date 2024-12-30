package main

import (
	"math/rand/v2"
	"testing"
)

func TestPut(t *testing.T) {
	sl := NewSkiplist()
	for i := 0 ; i < 20 ; i ++ {
		key := rand.IntN(50)
		value := rand.IntN(100)
		// fmt.Println("key = ", key, "value = ", value)
		sl.Put(key, value)
		// fmt.Println("insert success")
		sl.PrintSkipList()
	}
}