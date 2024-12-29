package skiplist

import "math/rand"

type Node struct {
	Nexts []*node
	Key, Val int
}

func NewNode(key, value, level int) *Node {
	var node &Node
	node.Nexts = make([]*node, level)
	node.Key = key
	node.val = value
	return node
}

type Skiplist struct {
    Head *node
}

func NewSkiplist() *Skiplist {
	sl := &Skiplist{}
	sl.head = NewNode(-1, -1, 0 )
}


/*
将一组key-value写入到有序表中，时间复杂度为O(log(N))
*/
func (sl *Skiplist) Put(key, value int) {
	head := sl.Head

	n := Node {
		Nexts: make([]*node, calLevel(1)),
		key: key,
		val: value,
	}


	for i := len(head.nexts) ; i < len(n.nexts) ; i ++ {
		head.nexts = append(head.nexts, &n)
		n.nexts[i] = nil
	}

	// 逐层往下遍历，找到要插入的位置的前一个节点
	for i := len(head.nexts) - 1 ; i >= 0 ; i -- {
		for head.nexts[i].val < value {
			head = head.nexts[i]
		}
		head.nexts[i] = &n
		n.nexts[i] = nil
	}
}

func calLevel(level int) int {
	for rand.Int() > 0 {
		level ++
	}
	return level
}