package redis

import (
	"context"
	"errors"
	"project/xihe-statistics/infrastructure/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	atomicFlag       = "atomic-"
	atmoicRetryTimes = 3
	atmoicWaitTime   = 3 * time.Millisecond
)

type dbRedis struct {
	Expiration time.Duration
}

func NewDBRedis(expiration int) dbRedis {
	return dbRedis{
		Expiration: time.Duration(expiration),
	}
}

func (r dbRedis) Create(
	ctx context.Context, key string, value interface{},
) *redis.StatusCmd {
	return client.Set(ctx, key, value, r.Expiration*time.Second)
}

func (r dbRedis) Get(
	ctx context.Context, key string,
) *redis.StringCmd {
	return client.Get(ctx, key)
}

func (r dbRedis) Set(
	ctx context.Context, key string, value interface{},
) *redis.StatusCmd {
	return client.Set(ctx, key, value, r.Expiration)
}

func (r dbRedis) Delete(
	ctx context.Context, key string,
) *redis.IntCmd {
	return client.Del(ctx, key)
}

func (r dbRedis) Expire(
	ctx context.Context, key string, expire time.Duration,
) *redis.BoolCmd {
	return client.Expire(ctx, key, 3*time.Second)
}

func (r dbRedis) AtomicOpt(
	ctx context.Context, key string, val interface{},
	f func(context.Context, interface{}) error,
) error {
	atomicKey := atomicFlag + key

	for i := 0; i <= atmoicRetryTimes; i++ {
		// get atmoic flag
		if resCmd := r.Get(ctx, atomicKey); resCmd.Err() == nil {
			time.Sleep(atmoicWaitTime)

			continue
		}

		// lock
		r.Create(ctx, atomicKey, 1)
		defer r.Delete(ctx, atomicKey)

		// do
		if err := f(ctx, val); err != nil {
			return err
		} else {
			return nil
		}
	}

	return repositories.NewErrorConcurrentUpdating(errors.New("cannot do atomic operation"))
}

func DB() *redis.Client {
	return client
}
