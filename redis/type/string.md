# String

## Summary

1. string 是Redis最基础的数据类型，Redis中的key都是string类型，value也可以是string类型。
2. string是二进制安全的，可以包含任何数据，包括jpg图片或者序列化数据。 二进制安全什么意思？
3. 一个Redis string的最大长度为512M。

## Command

### SET

### Append

```shell
APPEND key value
```

**时间复杂度**

时间复杂度为`O(1)`，假设追加的值比较小，平均时间复杂度为`O(1)`，因为不会发生空间的扩容。

**命令的返回值**

进行完`APPEND`操作后字符串的长度。

**应用案例**

构造时间序列，记录每个时刻到来的数据，添加到value的末尾。

> 如何从时间序列中提取信息

1. `APPEND`，构造时间序列数据。
2. `STRLEN`，获取时间序列数据的长度。
3. `GETRANGE`，按照范围获取数据，可以和`lua`引擎、二分算法实现。
4. `SETRANGE`，按照范围覆写数据。

> 使用APPEND命令创建时间序列的优点和缺点

1. `APPEND`命令只在尾部进行操作，并且Redis不存在可以修剪字符串的命令，不方便操作开头和中间的数据，难以直接将整个时间序列的长度控制在某个特定规模。
2. 可以根据UNIX时间戳作为Redis的key，避免某个key存储的时间过多，数据分散有助于把数据分散到多个Redis Instance中。
3. 使用string可以用二进制存储数据，存储效率更高。

### DECR DECRBY INCR INCRBY INCRBYFLOAT

```shell
DECR key 
DECRBY key decrement
INCR key
INCRBY key increment
INCRBYFLOAT key increment 
```

**时间复杂度**

时间复杂度为`O(1)`。

**返回值**

对value进行操作后的value。


> INCR INCRBY DECR DECRBY INCRBYFLOAT
>
> 都是把string当成整数、浮点数来处理
> 
> 当key不存在时，会创建0字符串，对其进行操作。
> 当key的类型错误或者value不能被解析为整数、浮点数时会返回错误。
>
> 对于整数，Redis将其转换为64位有符号整数进行存储。


**使用案例**

Counter

1. 使用INCR、EXPIRE命令计算一定的时间段内用户的浏览次数。
2. 使用GETSET

Rate Limiter


<<<<<<< HEAD
### SETRANGE

```shell
SETRANGE key offset value
```

### GETDEL GETSET



### STRLEN

```shell
STRLEN key
```

**时间复杂度**

`O(1)`

**返回值**

如果不存在key，则会返回0.
如果key对应的value类型不是string，则会返回error。
=======
### GETRANGE SETRANGE
### GETDEL GETSET

### 
>>>>>>> 91165c7c2a8e20c921f9a0aaecc9eb0263b1c5c2
