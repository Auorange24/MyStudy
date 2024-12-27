# union inter diff



# ZMPOP

> 同时从多个key中按照（min or max 的排序）删除

# BZMPOP

> blocking zset multiply pop
> 阻塞删除多个


# ZUNIONSTORE

```shell
ZDIFF numkeys key [key...] [WITHSCORES]
ZDIFFSTORE destination numkeys key [key...]
ZINTER numkeys key [key...] [WEIGHTS weight [weight...]] [AGGREGATE <SUM | MIN | MAX>]
ZINTERCARD numkeys key [key...] [LIMIT limit]
ZINTERSTORE destination numkeys key [key...] [WEIGHTS weight [weight...]] [AGGREGATE <SUM | MIN | MAX>]
ZUNIONSTORE destination numkeys key [key...] [WEIGHTS weight [weight...]] [AGGREGATE <SUM | MIN | MAX>]
ZUNION numkeys key [key...] [WEIGHTS weight [weight...]] [AGGREGATE <SUM | MIN | MAX>]
```

# ZRANGE

> 根据条件返回指定范围内的元素

```shell
ZRANGE key start stop [BYSCORE | BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
```

**命令解析**

1. 默认表示按照索引进行返回，即`RANK`，`BYSCORE`表示按照分数进行返回，`BYLEX`，表示按照字典序进行返回。
2. 默认表示从小到大进行返回，当使用`REV`参数表示从大到小进行返回。
3. `LIMIT`
4. 默认只返回member的列表，`[member1, member2, member3 ...]`，`[WITHSCORES]`，表示返回值携带每个member的score，`[member1, score1, member2, score2 ...]`

**索引范围**


# ZSCORE

```shell
ZSCORE key member
```

**时间复杂度**

`O(1)`，`ZSET`使用哈希存储`member`与`score`，可以在`O(1)`的时间内获取指定`key`和`member`对应`score`。

**返回值**

+ 类型错误返回`error`
+ `key`不存在或者`member`不存在返回`nil`
+ 正常查询到返回`score`

# ZRANK && ZREVRANK

```shell
ZRANK key member [WITHSCORE]
ZREVRANK key member [WITHSCORE]
```

**命令解析**

+ `ZRANK`返回指定key和`member`的排序(从小到大)，可以选择`score`是否一块返回，`rank`从0开始。
+ `ZREVRANK`的排序为从大到小。

**时间复杂度**

`O(log(N))`，需要在跳表中查询指定key和member对应的rank，需要从最高层开始搜索。

**返回值**

+ 如果类型不正确，返回`error`。
+ 如果`key`或者`member`不存在，返回`nil`。
+ 如果`[WITHSCORE]`不存在，则只返回`Integer`，否则返回`Array reply`。

# ZCARD && ZCOUNT

```shell
ZCARD key
```

**命令解析**




