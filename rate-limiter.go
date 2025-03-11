// Q3: Implement a rate limiter in Go that allows N requests per second using redid

package main

import (
    "context"
    "fmt"
    "log"
    "time"

    // "github.com/go-redis/redis"
    redis "github.com/redis/go-redis/v9"
)

type RedisRateLimiter struct{
    client *redis.Client
    limit int
    window time.Duration
}

func NewRateLimiter(redisAddr string,limit int,window time.Duration) *RedisRateLimiter{
    rdb := redis.NewClient(&redis.Options{
        Addr: redisAddr,
    })

    return &RedisRateLimiter{
        client: rdb,
        limit: limit,
        window: window,
    }
}

func(r *RedisRateLimiter)Allow(userID string)bool{
    ctx := context.Background()
    key := fmt.Sprintf("rate_limit:%s", userID)
    // now := time.Now().Unix()

    pipe := r.client.TxPipeline()
    count := pipe.Incr(ctx,key)
    pipe.Expire(ctx,key,r.window)

    _,err := pipe.Exec(ctx)
    if err != nil{
        log.Printf("Redis transaction error: %v", err)
        return false
    }

    return count.Val() <= int64(r.limit)

}

func main(){
    limiter := NewRateLimiter("localhost:6379",5,time.Second)

    userID := "user_123"

    for i:=1;i<=10;i++{
        if limiter.Allow(userID){
            fmt.Println("Request",i,"processed ")
        }else{
            fmt.Println("Request",i," did not processed (Rate limit exceeded)")
        }
        time.Sleep(200 *time.Millisecond)
    }
}