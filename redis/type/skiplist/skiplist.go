package main

import (
	"errors"
	"fmt"
	"math/rand"
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
	Size int
}

func NewSkiplist() *Skiplist {
	sl := &Skiplist{}
	sl.Head = NewNode(-1, -1, 1)
	return sl
}


/*
将一组key-value写入到有序表中，时间复杂度为O(log(N))
*/
func (sl *Skiplist) Put(key, value int) {
	// fmt.Println("skipList : ", sl.Head)
	head := sl.Head
	node := NewNode(key, value, calLevel(1))

	// fmt.Println("新生成的节点的高度为", len(node.Nexts), node)

	// 逐层往下遍历，找到要插入的位置的前一个节点
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		if i >= len(node.Nexts) {
			// fmt.Println("执行 continue")
			continue
		}
		for head.Nexts[i] != nil && head.Nexts[i].Val < value {
			head = head.Nexts[i]
		}
		node.Nexts[i] = head.Nexts[i]
		head.Nexts[i] = node
		// fmt.Println("head > node : head", head.Nexts)
		// fmt.Println("head > node : node", node.Nexts)
	}

	// 如果新节点的层数大于当前head的层数 需要进行更新
	for i := len(head.Nexts) ; i < len(node.Nexts) ; i ++ {
		// 先对Nexts slice 进行扩容
		head.Nexts = append(head.Nexts, node)
		// fmt.Println("node > head : head", head.Nexts)
		// fmt.Println("node > head : node", node.Nexts)
	}

}

func calLevel(level int) int {
	for {
		a := rand.Intn(10)
		// fmt.Println("a = ", a)
		if a > 5 { 
			level ++
		} else {
			break
		}
	}
	return level
}

func (sl *Skiplist) PrintSkipList() {
	// fmt.Println("PrintSkipList")
	front := sl.Head
	for i := len(sl.Head.Nexts) - 1; i >= 0 ; i -- {
		fmt.Print("level = ", i, " :")
		head := front
		for head.Nexts[i] != nil {
			// fmt.Println("print : ", head.Nexts)
			fmt.Print(head.Nexts[i].Val, " ")
			head = head.Nexts[i] 
		}
		fmt.Printf("\n")
	}
}

func (sl *Skiplist) Get(key int) int {
	head := sl.Head 
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		for head.Nexts[i] != nil && head.Val < key {

		}
	}
	return 1
} 

func (sl *Skiplist) Exist(key int) bool {
	head := sl.Head
	for i := len(head.Nexts) ; i >= 0 ; i -- {
		for head.Nexts[i] != nil && head.Nexts[i].Key < key {
			head = head.Nexts[i]
		} 
		if head.Nexts[i] != nil && head.Nexts[i].Key == key {
			return true
		}
	}
	return false
}

// 返回key值最小的KV对
func (sl *Skiplist) First() (int, int, error) {
	head := sl.Head
	if head.Nexts[0] == nil {
		return -1, -1, errors.New("SKiplist中不存在元素")
	} else {
		return head.Nexts[0].Key, head.Nexts[0].Val, nil
	}
}

// 返回Key值最大的KV对
func (sl *Skiplist) Last() (int, int, error) {
	head := sl.Head
	if head.Nexts[0] == nil {
		return -1, -1, errors.New("Skiplist中不存在元素")
	}
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		for head.Nexts[i] != nil {
			head = head.Nexts[i]
		}
	}
	return head.Key, head.Val, nil
}

func (sl *Skiplist) Empty() {

}

// 返回Key值小于Target的最接近的KV键值对
func (sl *Skiplist) floor(target int) (int, int ,error) {
	if sl.Size == 0 {
		return -1, -1, errors.New("Skiplist 为空")
	}
	if key, _, _ := sl.First() ; key > target {
		return -1, -1, errors.New("不存在小于等于target的key")
	}
	head := sl.Head
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		for head.Nexts[i] != nil && head.Nexts[i].Key < key {
			head = head.Nexts[i]
		}
		if head.Next[i] != nil {
			
		}
	}
}


