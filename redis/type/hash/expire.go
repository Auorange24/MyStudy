package hash

import (
	"context"
	"fmt"
	"math/rand/v2"
	"github.com/redis/go-redis/v9"
)

/*
HEXPIRE key seconds [NX | XX | GT | LT] FIELDS numfields field [field ...]
为hash中的一个字段或多个字段设置过期时间，至少应该指定一个字段。
key Redis中的键
seconds 要设置的过期时间，以秒为单位。
FIELDS 暂时不知道意义在哪
numfields 要设置过期时间的字段的个数
field [field ...]

字段过期之后，字段会被从哈希键中直接删除。
字段的过期时间只有通过删除（HDEL）/覆盖（HSET）操作才能移除字段中的过期时间
如果对字段中的值进行加减操作而不是用新值去替换它，都不会改变过期时间的设置。
可以使用HPERIST操作清楚字段的过期时间使用过期时间（TTL）为零的 “HEXPIRE”（设置哈希表字段的过期时间）
“HPEXPIRE”（以毫秒为单位设置哈希表字段的过期时间）命令，
或者使用指定时间为过去时间的 “HEXPIREAT”（设置哈希表字段过期的时间戳）
“HPEXPIREAT”（以毫秒为单位设置哈希表字段过期的时间戳）命令，将会导致哈希字段被删除。
*/

func HEXPIRE(rdb *redis.Client) {
	// 1.准备数据
	insertData := make([]string, 0)
	for i := 0 ; i < rand.IntN(20) ; i ++ {
		insertData = append(insertData, "key" + fmt.Sprint(i))
		insertData = append(insertData, "value" + fmt.Sprint(i))
	}
	// 2. 插入数据
	if result, err := rdb.HSet(context.Background(), "hexpire_test", insertData).Result() ; err != nil {
		FatalError(err)
	} else {
		fmt.Println("成功插入", result, "条数据")
	}
	/*
	3. 设置过期时间
		（1）
	*/
	rdb.HExpire()
}