package zset

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/redis/go-redis/v9"
)

func ZADD(rdb redis.Client) {
	// 准备要插入的数据
	insertData := make([]redis.Z, 0)
	for i := 0 ; i < rand.IntN(20) ; i ++ {
		insertData = append(insertData, redis.Z{
			Score: rand.Float64(),
			Member: "member" + fmt.Sprint(i),
		})
	}
	// 使用zadd 插入数据
	for i := 0 ; i < len(insertData) ; i ++ {
		result, err := rdb.ZAdd(context.Background(), "zadd_test1", insertData[i]).Result()
		if err != nil {
			fmt.Println("err: ", err)
		} else {
			fmt.Println("成功插入", result, "条数据")
		}
	}
}