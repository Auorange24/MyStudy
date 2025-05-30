# Hash


## Commands

### Summary

+ `HDEL`，删除哈希对应的一个或者多个字段。
+ `HEXISTS`，测试键中某个字段是否存在。
+ `HGET`，获取哈希中某个字段的对应值。
+ `HGETALL`，获取哈希的所有内容，以列表形式输出[k1, v1, k2, v2]
+ `HKEYS`，获取哈希的所有字段内容。
+ `HLEN`，计算哈希中键值对的个数。
+ `HMGET`，计算哈希中一个或者多个字段对应的值，以列表的形式进行输出，[v1, v2, v3 ...]
+ `HSET`，向hash中添加字段。

### Edit Hash

**HDEL**

```redis
HDEL key field [field ...]
```
时间复杂度:`O(N)`，N与被删除的字段的个数有关。删除哈希中对应的字段。字段存在会删除，不存在则会忽略这个键，删除键之后哈希为空则删除哈希，键对应的值不是哈希类型直接忽略。

返回值
+ 被删除的键的数量，不包括不存在的键。

**HEXISTS**

```redis
HDEL key field
```
时间复杂度:`O(1)`，判断某个键对应的字段是否存在。

返回值
+ 键对应的字段存在，返回1。
+ 键不存在或者字段不存在，返回0。

HDEL 删除hash中的某个字段
HEXISTS 测试某个字段是否存在
HEXPIRE 设置字段的过期时间
 