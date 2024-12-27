# List

## Summary

1. List就是一个双端列表，可以在表头或者表尾插入。
<<<<<<< HEAD
2. List最多可以包含2^32 - 1个元素。




## Commands

### Push、Pop、Move

+ `LPUSH key element [element ...]`，从列表左侧插入元素，像入栈操作一样，时间复杂度和插入的元素的个数有关，返回值为插入元素后当前列表的长度。`LPUSHX`，表示只有当前列表存在才会进行插入操作。`RPUSH`、`RPUSHX`操作同理。
+ `LPOP key [count]`，从列表左侧删除如干个元素，如果value类型不匹配，则返回错误，如果key不存在，则返回0，如果未设置count，则返回首元素，设置`count > 1`，则返回元素列表。`RPOP`同理。
+ `LINDEX key index`，根据index获取列表中元素。索引从0开始，-1、-2也可以被使用。如果key对应的类型不是列表，则返回错误。当index超出范围时，返回空值，index有效时，返回element。
+ `LLEN key`，返回列表中包含的元素。

### Blocking Command
=======
2. List最多可以包含2^32 - 1个元素。
>>>>>>> 91165c7c2a8e20c921f9a0aaecc9eb0263b1c5c2
