package hash

import (
	"goredis/conf"
	"testing"
)

func TestHvals(t *testing.T) {
	HVALS(conf.InitRedis())
}