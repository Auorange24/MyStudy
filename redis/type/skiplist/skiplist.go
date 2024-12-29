package skiplist

import "math/rand"

type node struct {
	nexts []*node
	key, val int
}

type Skiplist struct {
    head *node
}

/*
将一组key-value写入到有序表中，时间复杂度为O(log(N))
*/
func (sl *Skiplist) put(key, value int) {
	head := sl.head

	n := node {
		nexts: make([]*node, calLevel(1)),
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