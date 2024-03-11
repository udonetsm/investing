package myredis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/udonetsm/investing/models"
)

type MyRedis models.Redis

const (
	REQUESTS = iota
	TRANSACTIONS
)

func NewRedis(database int) *MyRedis {
	return &MyRedis{
		Client: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       database,
		}),
	}
}

func (r *MyRedis) SetLastSeen(key string, value int64) {
	ctx := context.Background()
	defer ctx.Done()
	r.Client.Set(ctx, key, value, 0)
}

func (r *MyRedis) GetLastSeen(key string) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res := r.Client.Get(ctx, key)
	ls, err := res.Int64()
	if err != nil && err == redis.Nil {
		return 0
	}
	if err != nil {
		log.Println(err)
	}
	return ls
}

func (r *MyRedis) SaveTransaction(transaction *models.Transaction) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	fmt.Println(r.Client.Set(ctx, transaction.Tid, transaction, 0).Err())
}

func (r *MyRedis) RestoreTransaction(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	fmt.Println(r.Client.Get(ctx, key))
}
