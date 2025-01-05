package bitmap

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/redis/go-redis/v9"
)

/*
Example 1
提交1000辆自行车在一个小时内的签到情况
*/

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
		Password: "",
		DB: 0,
	})
	return rdb
}

const checkIn string = "bicyle:19:00-20:00"

func CheckIn(rdb *redis.Client) {
	for i := 0 ; i < 1000 ; i ++ {
		if rand.IntN(10) >= 5 {
			if _, err := rdb.SetBit(context.Background(), checkIn, int64(i), 1).Result() ; err != nil {
				fmt.Println("err: ", err)
			} 
		}
	}

	ans, err := rdb.BitCount(context.Background(), checkIn, &redis.BitCount{
		Start: 0,
		End: 1000,
		Unit: "BIT",
	}).Result()

	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("bicycle: ", ans)
	}
} 