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
Put ToDo:
将一组key-value写入到有序表中，时间复杂度为O(log(N))
*/
func (sl *Skiplist) Put(key, value int) {
	// TODO: 先判断一下KV对是否存在

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
		for head.Nexts[i] != nil && head.Nexts[i].Key < key {
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

// 根据Key获取value的值

func (sl *Skiplist) Get(key int) (int, error) {
	// 1. check skiplist is empty
	if flag := sl.Empty() ; flag {
		return 0, errors.New("skiplist is empty")
	}
	head := sl.Head 
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		for head.Nexts[i] != nil && head.Val < key {
			head = head.Nexts[i]
		}
		if head.Nexts[i] != nil && head.Nexts[i].Key == key {
			return head.Nexts[i].Val, nil
		}
	}
	return -1, nil
}

func (sl *Skiplist) Del(key int) (int, error) {
	// check skiplist is empty
	if flag := sl.Empty() ; flag {
		return 0, errors.New("Skiplist is empty")
	}
	// check key is exist
	head := sl.Head 
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		for head.Nexts[i] != nil && head.Nexts[i].Key < key {
			head = head.Nexts[i]
		}
		if head.Nexts[i] != nil && head.Nexts[i].Key == key {
			val := head.Nexts[i].Val
			head.Nexts[i] = head.Nexts[i].Nexts[i]
			return val, nil
		}
	}
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

/*
判断跳表是否为空
为空  返回true
不为空返回false
*/
func (sl *Skiplist) Empty() bool {
	if sl.Head.Nexts[0] == nil {
		return true
	} else {
		return false
	}
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
		if head.Nexts[i] != nil {
			
		}
	}
}

// 返回key值大于等于target 的第一个元素
func (sl *Skiplist) ceiling(target int) (int, int, error) {
	// 判断当前跳表是否为空
	if flag := sl.Empty() ; flag == true {
		return 0, 0, errors.New("skiplist is empty")
	}
	/*
	判断当前跳表的最大key是否小于target
	如果小于target，返回error
	*/
	if k, _, _ := sl.Last() ; k < target {
		return 0, 0, errors.New("no return value")
	}

	// 查询结果
	head := sl.Head
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		for head.Nexts[i] != nil && head.Nexts[i].Key < target {
			head = head.Nexts[i]
		}
	}
	return head.Key, head.Val, nil
}
