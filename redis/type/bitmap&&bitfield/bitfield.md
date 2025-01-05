# bitmap && bitfield

## BitMap Summary

+ 位图不是实际的数据类型，底层结构还是字符串。
+ 位图是一组对bit操作的集合。
+ 字符串是二进制安全的大对象，最大长度为512M，位图操作能设置2的32次方个不同的位置。
+ 位图操作可以对一个或者多个字符串进行操作。

> 使用位图的常见场景
>
> 1. 能使用一个bit表示多个用户动作的情况，例如签到。
> 2. 每个对象的权限

## 

## Summary

+ `Bitfield`能够高效的在一个字符中操作多个计数器。
+ `Bitfield`提供了原子性的GET、SET、INCR操作。
+ `Bitfield`支持多种溢出策略。
+ `Bitfield`可以设置、增加、获取任意比特长度的值，操作范围从无符号1位比特到63位有符号数。

## Commands

### BITFIELD

```shell
BITFIELD key
[GET encoding offset | [OVERFLOW <WARP | SAT | FIAL>]


]
```