# ConcurrentHashMap

## ConcurrentHashMap Field

1. `volatile Node<K, V>[] table`

存储Map数据的变量，采用lazy load的方式，第一次插入数据才会进行初始化。

2. `volatile Node<K, V>[] nextTable`

扩容时使用，平时为null，扩容时非null。

3. `volatile int sizeCtl`

sizeCtl控制table数组的大小

+ 当值为负数时，如果为-1表示正在初始化，如果为-N表示当前N-1个线程正在进行扩容操作。

+ 当值为正数时
    * 如果当前table为null的话表示table正在初始化中。
    * 如果table不为null，表示当前table的可用容量，即table节点的临界值，超过了临界值就需要扩容。
    临界值可以由 table长度 * loadfactor计算。

+ 当sizeCtl = 0时，table长度为默认初始值。

4. `sum.misc.Unsafe U`

用于进行CAS操作

> unsafe的获取在静态代码块中

```java
static {
    try {
        U = sun.misc.Unsafe.getUnsafe();
		.......
    } catch (Exception e) {
        throw new Error(e);
    }
}
```

## ConcurrentHashMap Inner class

1. **Node**

Node<K, V> class extend Map.Entry<K, V>

Many field is descried by volatile.

volatile can't describe class, only field.

```java
static class Node<K,V> implements Map.Entry<K,V> {
    final int hash; // hashcode
    final K key; // key
    volatile V val; // value
    volatile Node<K,V> next; // next node
    ......
}
```

2. **TreeNode**

TreeNode, serve for red-black tree.

```java
/*
 * Nodes for use in TreeBins
 */
static final class TreeNode<K,V> extends Node<K,V> {
        TreeNode<K,V> parent;  // red-black tree links
        TreeNode<K,V> left;
        TreeNode<K,V> right;
        TreeNode<K,V> prev;    // needed to unlink next upon deletion
        boolean red;
		......
}
```

## Concurrent Method

### Constructor Method

* `ConcurrentHashMap()`

Empty Constructor Method, default length is 16, 16 hash slot.

* `ConcurrentHashMap(int initialCapacity)`

set ConcurrentHashMap capacity.

```java
public ConcurrentHashMap(int initialCapacity) {
	//1. 小于0直接抛异常
    if (initialCapacity < 0)
        throw new IllegalArgumentException();
	//2. 判断是否超过了允许的最大值，超过了话则取最大值，否则再对该值进一步处理
    // >>> 无符号右边移动操作符
    int cap = ((initialCapacity >= (MAXIMUM_CAPACITY >>> 1)) ?
               MAXIMUM_CAPACITY :
               tableSizeFor(initialCapacity + (initialCapacity >>> 1) + 1));
	//3. 赋值给sizeCtl
    this.sizeCtl = cap;
}
```
1. 如果容量不规范，抛出异常。
2. 如果初始化容量小于最大容量的一半，则设置初始化容量为最大值的一半加上设置的初始化容量。tablesize()
函数将初始化的转换成2的N次幂。
3. 赋值初始化容量，只有当第一次插入数据的时候才会进行concurrentHashMap的初始化。

* `ConcurrentHashMap(Map<? extends K, ? extends V>)`

give a map, key is father class of given key, 
value is father class of given value.

* `ConcurrentHashMap(int initialCapacity, float loadFactor)`

Initial capacity and loadFactor of ConcurrentHashMap.

* `ConcurrentHashMap(int initialCapacity, float loadFactor, int concurrencyLevel)`

concurrencyLevel is expected running thread.