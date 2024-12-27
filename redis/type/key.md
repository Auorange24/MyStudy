# Redis Key

## Summary 

+ Redis中的key是string类型。

## Command

+ `KEYS pattern`，根据匹配模式查看当前库的key，*表示查询所有的key，结果以列表的形式返回。
+ `DUMP key`，将某个key进行序列化，返回序列化的值。
+ `EXISTS key [key...]`，查询某个key是否存在，结果返回库中存在key的个数。
+ `EXPIRE key seconds；EXPIREAT key timestamp；PEXPIRE key milliseconds；PEXPIREAT key milliseconds-timestamp`，为给定key设定过期时间，以秒、毫秒、UNIX时间为单位。
+ `PERSIST key`，移除key的过期时间，key将会永久生效。
+ `TTL key；PTTL key`，以秒、毫秒为单位返回key的剩余过期时间。以秒为单位时，-1表示永不过期，-2表示已经过期。
+ `MOVE key dbindex`，将某个key移动到指定库。
+ `DEL key`，删除某个key。`UNLINK key`，惰性删除，仅仅在keyspace中删除，真正的键值对将会由异步线程删除。
+ `TYPE key`，查看key对应的value的类型。
+ `RANDOMKEY`，从当前库中随机返回一个key。
+ `RENAME key newkey`，修改key的名字为newkey。
+ `RENAMENX key newkey`，仅当newkey不存在时，才会把key修改为newkey。
+ `SELECT dbindex`，选择其他数据库。
+ `FLUSHDB; FLUSHALL`，清楚当前数据库或者所有数据库。
+ `help @type`，查询某个类型的相关命令。