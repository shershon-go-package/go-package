/**
 * @Description go-redis使用
 **/
package goredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// 单机连接redis
func ConnectSingle() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:      "127.0.0.1:6379",
	})
	// 检测是否建立连接(需要传递上下文)
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	// 检测
	_, err := client.Ping(timeoutCtx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 连接新连接时回调
func onConnectFunc(ctx context.Context, cn *redis.Conn) error {
	fmt.Println("新连接回调")
	return nil
}

// 哨兵模式连接
func ConnectSentinel() (*redis.Client, error) {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master-name",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
	// 检测是否建立连接(需要传递上下文)
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	// 检测
	_, err := client.Ping(timeoutCtx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 集群模式连接
func ConnectCluster() (*redis.ClusterClient, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003"},
	})
	// 检测是否建立连接(需要传递上下文)
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	// 检测
	_, err := client.Ping(timeoutCtx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
