# 


# Command

## SETBIT

```bash
SET key offset value
```

+ 当键不存在时，会创建一个新的字符串值。
+ Redis字符串底层的SDS会根据当前长度自动扩充。
+ 偏移量的大小从0 - 2^32次方。
+ 当字符增长时，未被设置的bit位会设置为0。
+ 字符串的命令也可以使用到位图上。

> Warning
>
> 当设置最后一个可能的位（偏移量等于 2 的 32 次方 - 1）并且存储在键中的字符串值尚未存储字符串值，或者存储的是一个较小的字符串值时，Redis 需要分配所有的中间内存，这可能会阻塞服务器一段时间。