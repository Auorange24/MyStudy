package skiplist

import (
	"math/rand/v2"
	"testing"
	"fmt"
)

func TestPut(t *testing.T) {
	sl := NewSkiplist()
	for i := 0 ; i < 20 ; i ++ {
		sl.Put(rand.IntN(50), rand.IntN(100))
	}

	for i := len(sl.Head.Nexts) - 1; i >= 0 ; i -- {
		fmt.Println("level : ", i)
		for sl.Head.Nexts[i] != nil {
			fmt.Println("value = ", sl.Head.Nexts[i].Val)
		}
	}
}