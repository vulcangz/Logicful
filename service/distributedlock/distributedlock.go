package distributedlock

import (
	"logicful/service/cache"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

func Acquire(key string) (*redsync.Mutex, error) {
	client, err := cache.Instance()
	if err != nil {
		return nil, err
	}
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	mutex := rs.NewMutex(key)
	if err := mutex.Lock(); err != nil {
		return nil, err
	}
	return mutex, nil
}

func Release(mutex *redsync.Mutex) error {
	if _, err := mutex.Unlock(); err != nil {
		return err
	}
	return nil
}
