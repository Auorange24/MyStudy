# Skiplist

## 有序表

> 内部元素按照规则组织好顺序的数据结构

```go
type OrderedList interface {
    put() // 将一组数key-value写入到有序表中，时间复杂度为O(log(N))
    get() // 根据 key 从表中获取对应的 value，时间复杂度 <= O(logN)
    del() // 根据 key 从表中删除 key-value 对，时间复杂度O(logN)
    exist() // 判断 key 在表中是否存在，时间复杂度 <= O(logN)
    first() // 返回 key 值最小的 key-value 对，时间复杂度 <= O(logN)
    last() // 返回 key 值最大的 key-value 对，时间复杂度 <= O(logN)
    floor() // 返回key值 <= target 时间复杂度为O(log(N))
    ceiling() // 返回key值 >= target 时间复杂度为O(log(N)) 二分查找
}
```

**常见的有序表**

+ 平衡二叉搜索树
+ 红黑树
+ 节点大小平衡树 
+ 跳表

## 跳表

> 跳表的最底层就是有序链表
> 第二层到最顶层是最底层的索引
> 基于二分查找的方式实现O(log(N))级别时间复杂度下的增删改查操作

## 有序数据结构

### 平衡搜索二叉树

**什么是搜索树**

对于父节点，如果其左子树的节点的值都小于父节点，右子树的节点的值都大于父节点

**什么是平衡树**

对于parent节点，其左子树和右子树的高度最大相差为1，配合搜索树的特性，每次检索都可以筛掉一半的数据。
但是二叉搜索树每次插入新的数据需要调整子树，时间复杂度高。

### 红黑树

1. 红黑树的本质是一颗搜索树
2. 根节点为黑色
3. 所有叶子节点也是黑色
4. 从根节点到叶节点的路径上，不允许出现两个连续红色的节点
5. 从任意一个根节点到叶节点的路径中，都包含相同数目的黑色节点

parent节点的左右高度差距不会超过一倍，也能维持在O(log(N))的时间复杂度。

### 红黑树对比跳表

**时间复杂度**

有序表结构经常使用的方法在跳表和红黑树中均能维持O(log(N))的水平

**空间复杂度**

对于红黑树，每个节点存储值、颜色、只想下一个节点的指针。
对于跳表，每个节点可能由多层，会使用更多的空间来存储节点。

**实现难易度**

对于红黑树，涉及到颜色的选择，实现起来复杂。
对于跳表，实现简单。

**range操作**

红黑树是树结构，难以实现按照范围取值。
跳表是链表结构，容易实现按照范围取值。

**并发性**

红黑树是粗锁
跳表可以支持细锁

> 粗锁和细锁的区别

## Skiplist的函数实现

### PUT

**个人实现**

```go
func (sl *Skiplist) Put(key, value int) {
	head := sl.Head
	node := NewNode(key, value, calLevel(1))
	for i := len(head.Nexts) - 1 ; i >= 0 ; i -- {
		if i >= len(node.Nexts) {
			continue
		}
		for head.Nexts[i] != nil && head.Nexts[i].Key < key {
			head = head.Nexts[i]
		}
		node.Nexts[i] = head.Nexts[i]
		head.Nexts[i] = node
	}
	for i := len(head.Nexts) ; i < len(node.Nexts) ; i ++ {
		head.Nexts = append(head.Nexts, node)
	}
}
```

> Note: 先更新还是先扩展头节点
> 先扩展头节点，在更新时可以直接按照新的子节点的高度进行更新
> 先扩展头节点，Redis源码中使用数组存储了每一层的前驱节点，但是不一样能用来，个人实现需要进行判断

### GET

**个人实现**

```go
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
```

