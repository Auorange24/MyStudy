package bitmap

import (
	"testing"
	"fmt"
)

func TestConnectRedis(t *testing.T) {
	rdb := ConnectRedis()
	if rdb != nil {
		fmt.Println("Connect Success")
	}
}

func TestCheckIn(t *testing.T) {
	rdb := ConnectRedis()
	if rdb != nil {
		fmt.Println("Connect Success")
	}
	CheckIn(rdb)
}