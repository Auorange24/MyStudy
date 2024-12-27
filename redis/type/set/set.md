# SET
## Summary
1. 无序且唯一得集合，存储顺序不会按照插入的先后顺序
2. 最多可以存储`2^31-1`个元素，怎么计算出来的？

## How to use Set
1. 记录访问的唯一IP
2. 表示关系，比如某个角色（如管理员）中的所有用户
3. 并、交、差集合操作

## Set Command

**SADD** 

```redis
SADD key member [member ...]
```

1. 时间复杂度为`O(1)`
2. 添加`member`，如果`member`已经存在则会忽略，如果`key`不存在，则会创建新的`set`。如果`key`对应得`value`不是`set`类型，则会返回`error`。

**SCARD**
```redis
SCARD key 
```
1. 时间复杂度为`O(1)`
2. 返回集合中的元素，不存在对应的集合则返回0

**SDIFF SDIFFSTORE**

```redis
SUNION key [key ...]
SUNIONSTORE destination key [key ...]
```

**SINTER SINTERSTORE SINTERCARD**

```redis
SUNION key [key ...]
SUNIONSTORE destination key [key ...]
```

**SUNION SUNIONSTORE**

```redis
SUNION key [key ...]
SUNIONSTORE destination key [key ...]
```




## 底层实现
Set类型的底层数据结构是由`hashtable or intset`实现。
+ 如果Set中的元素都是整数且元素个数小于512个，Redis会使用intset惊醒存储。
```yaml
set-maxintset-entries配置
```
+ 如果Set不满足条件，则会使用hashtable进行存储。
