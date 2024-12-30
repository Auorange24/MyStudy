package skiplist

import (
	"math/rand"
	"fmt"
)

type Node struct {
	Nexts []*Node
	Key, Val int
}

func NewNode(key, value, level int) *Node {
	node := &Node{}
	node.Nexts = make([]*Node, level)
	node.Key = key
	node.Val = value
	return node
}

type Skiplist struct {
    Head *Node
}

func NewSkiplist() *Skiplist {
	sl := &Skiplist{}
	sl.Head = NewNode(-1, -1, 0)
	return sl
}


/*
将一组key-value写入到有序表中，时间复杂度为O(log(N))
*/
func (sl *Skiplist) Put(key, value int) {
	head := sl.Head
	node := NewNode(key, value, calLevel(1))

	// 如果新节点的层数大于当前head的层数 需要进行更新
	for i := len(head.Nexts) ; i < len(node.Nexts) ; i ++ {
		// 先对Nexts slice 进行扩容
		head.Nexts = append(head.Nexts, node)
		node.Nexts[i] = nil
	}

	// 逐层往下遍历，找到要插入的位置的前一个节点
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		for head.Nexts[i] != nil && head.Nexts[i].Val < value {
			head = head.Nexts[i]
		}
		node.Nexts[i] = head.Nexts[i]
		head.Nexts[i] = node
	}
}

func calLevel(level int) int {
	for true {
		a := rand.Intn(10)
		fmt.Println("a = ", a)
		if a > 5 { 
			level ++
		} else {
			break
		}
	}
	return level
}