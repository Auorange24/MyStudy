# hash 过期命令相关

## TTL 查询过期时间

> HTTL 以秒为单位返回
> HPTTL 以毫秒为当单位返回 

```shell
HTTL key FIELDS numfields field [field...]
HPTTL key FIELDS numfields field [field...]
```

**时间复杂度**

为`O(N)`，N为哈希键的字段个数。

**返回值**

返回一个数据，对于每个字段，
+ `-2`表示key不存在或者字段不存在
+ `-1`表示字段未设置过期时间
+ 其他的代表字段的过期时间

## EXPIRE 设置过期时间

```shell
HEXPIRE key seconds [NX | XX | GT | LT] FIELDS numfields field [field ...]
HEXPIRE key unix-time-seconds [NX | XX | GT | LT] FIELDS numfields field [field ...]
HEXPIRETIME key FIELDS numfields field [field ...]
HPEXIPRE key milliseconds [NX | XX | GT | LT] FIELDS numfields field [field ...]

```