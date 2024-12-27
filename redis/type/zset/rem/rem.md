# Redis Delete Command

## Rem

```redis
ZREM key member [member...]
```

**时间复杂度**

`O(M * log(N))`，`M`是要删除的`member`的个数，`N`是整个`ZSET`存在的`member`的个数。每删除一个`member`的时间复杂度为`O(log(N))`。

**命令解析**

删除一个或者多个member，当member不存在，直接忽略该member。

**返回值**

+ `err`，对应的key不是ZSET类型。
+ `Integer`，成功删除的member的个数

## ZREMRANGEBY <LEX | RANK | SCORE>

```shell
ZREMRANGEBY<LEX | RANK |SCORE> key start stop
```

**时间复杂度**

`O(log(N) + M)`，`O(log(N))`，查找第一个member的位置，然后直接在最低level往后遍历即可。

**命令解析**

按照字典序、排序或者分数的形式删除元素。首先根据key和start找到起止点


## ZREMRANGEBYRANK



## ZREMRANGEBYSCORE