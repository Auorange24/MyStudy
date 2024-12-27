/*
查询hash中key和value
HVALS key 
获取对应key的所有value
时间复杂度为O(N)，N对应hash中的字段个数
返回值以列表的形式进行返回。
value类型不对应则会返回err
不存在value则返回为nil

HKEYS key 同理
*/

package hash

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/redis/go-redis/v9"
)

// handle error

func FatalError(err error) {
	fmt.Println("err: ", err)
}

func TestHset(rdb *redis.Client) {
	rdb.HSet()
}

func HVALS (rdb *redis.Client) {
	// 1. 插入数据
	insertData := make([]string, 0)
	for i := 0 ; i < rand.Intn(20) ; i ++ {
		insertData = append(insertData, "field" + fmt.Sprint(i))
		insertData = append(insertData, "value" + fmt.Sprint(i))
	}
	if res, err := rdb.HSet(context.Background(), "test_hvals", insertData).Result() ; err != nil {
		FatalError(err)
	} else {
		fmt.Println("成功插入", res, "条数据")
	}
	// 2. 获取hash中的value
	if strs, err := rdb.HVals(context.Background(), "test_hvals").Result() ; err != nil {
		FatalError(err)
	} else {
		for i := 0 ; i < len(strs) ; i ++ {
			fmt.Println("value" + fmt.Sprint(i), " = ", strs[i])
		}
	}
}