/**
 * @Description TODO
 **/
package goredis

import (
	"context"
	"fmt"
	"time"
)

// 自增和自减
func UseIncrAndDecr() {
	// 连接redis
	client, _ := ConnectSingle()
	ctx := context.Background()
	// 自增
	for i := 1; i <= 5; i++ {
		// Incr: 每次调用+1
		client.Incr(ctx, "incr1")
		// IncrBy: 每次调用+5
		client.IncrBy(ctx, "incr2", 5)
		// IncrByFloat: 每次调用 +0.5
		client.IncrByFloat(ctx, "incr3", 0.5)
		// 查询缓存
		result, _ := client.MGet(ctx, "incr1", "incr2", "incr3").Result()
		fmt.Printf("第%d次自增后查询: %v \n", i, result)
	}
	// 自减
	for i := 1; i <= 5; i++ {
		// Decr: 每次调用-1
		client.Decr(ctx, "decr1")
		// DecrBy: 每次调用-5
		client.DecrBy(ctx, "decr2", 5)
		// 查询缓存
		result, _ := client.MGet(ctx, "decr1", "decr2").Result()
		fmt.Printf("第%d次自减后查询: %v \n", i, result)
	}
}

// 批量设置和获取
func MGetSet() error {
	// 连接redis
	client, _ := ConnectSingle()
	ctx := context.Background()
	// MSet: 同时设置一个或多个 key-value 对。
	err := client.MSet(ctx, "key1", "val1", "key2", "val2", "key3", "val3").Err()
	if err != nil {
		return err
	}
	// 使用Get获取缓存
	for _, key := range []string{"key1", "key2", "key3"} {
		result, _ := client.Get(ctx, key).Result()
		fmt.Printf("Get获取,键: %s 值: %v \n", key, result)
	}
	// 使用MGet获取
	result, _ := client.MGet(ctx, "key1", "key2", "key3").Result()
	fmt.Println("MGet批量获取:", result)
	return nil
}

// 删除和追加
func DelAndAppend() {
	client, _ := ConnectSingle()
	ctx := context.Background()
	// 删除单个
	client.Del(ctx, "key1")
	// 删除多个
	client.Del(ctx, "incr1", "incr2", "incr3")

	// === 追加value 测试使用 ===
	client.Set(ctx, "key", "hello", time.Hour)
	// 获取值
	res1, _ := client.Get(ctx, "key").Result()
	fmt.Println("追加前的值:", res1)
	// 追加
	client.Append(ctx, "key", " word")
	res2, _ := client.Get(ctx, "key").Result()
	fmt.Println("追加后的值:", res2)
}
