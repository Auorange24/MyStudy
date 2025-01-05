package lock

import (
	"context"
	"time"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type DistributedLock struct {}

const LockKey string = "product:1"
const LockValue string = "lock:value"

var rdb *redis.Client

func (l *DistributedLock) Lock(rdb *redis.Client) {
	if flag, err := rdb.SetNX(context.Background(), LockKey, LockValue, time.Second * 5).Result() ; err != nil {
		// key 对应错误，要防止这种情况
		logrus.Error("锁的类型不正确")
		panic(err)
	} else {
		if flag {
			// 设置成功，表示获得锁。
			logrus.Info("获取分布式锁成功")
		} else {
			logrus.Info("获取分布式锁失败")
		}
	}
}

func (l *DistributedLock) UnLock(rdb *redis.Client) {
	if _, err := rdb.GetDel(context.Background(), LockKey).Result() ; err != nil {
		logrus.Error("锁的类型错误")
	} else {
		logrus.Info("释放分布式锁成功")
	}
}

/*
简易分布式锁
背景：多个服务需要访问库存
*/

var store int = 100


type Service struct {}

func InitDistributedLock() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
		Password: "",
		DB: 0,
	})
	return rdb
}

func (s *Service) GetProduct() {
	// 获取分布式锁
	
}

func startService() {
	s := &Service{}
	if 
	s.GetProduct()
}


func main() {
	// 开始商品库存服务
	// Load Redis
	rdb = InitDistributedLock()
	
}