# Command

## GEOADD

```shell
GEOADD key [NX | XX] [CH] longitude latitude member [longitude latitude member...]
```

**时间复杂度**

`O(log(N))`

**命令解析**

+ `NX`表示仅当member不存在的时候才进行更新
+ `XX`表示仅当member存在的时候才进行更新