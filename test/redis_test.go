/**
 * @Description https://github.com/go-redis/redis
 **/
package test

import (
	"context"
	"fmt"
	"shershon1991/go-utils/app/redispkg"
	"testing"
	"time"
)

// Get&Set
func TestGetAndSet(t *testing.T) {
	// 使用单机模式连接redis
	client, err := redispkg.ConnectSingle()
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()
	// 设置缓存
	err = client.Set(ctx, "abc1", "hello word!", time.Minute*10).Err()
	if err != nil {
		t.Error("设置缓存失败" + err.Error())
		return
	}
	// 获取缓存
	result, err := client.Get(ctx, "abc1").Result()
	if err != nil {
		t.Error("获取缓存失败" + err.Error())
		return
	}
	fmt.Println("Get获取结果: ", result)
}

// SetNX: 指定的 key 不存在时，为 key 设置指定的值
func TestSetNx(t *testing.T) {
	// 连接redis
	client, err := redispkg.ConnectSingle()
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		res, err := client.SetNX(ctx, "abc", time.Now().Unix(), time.Hour).Result()
		if err != nil {
			t.Error(err)
			break
		}
		fmt.Println("SetNX abc success :", res)
	}
}

// 批量设置和获取
func TestMGetSet(t *testing.T) {
	_ = redispkg.MGetSet()
}

// 自增和自减
func TestIncrAndDecr(t *testing.T) {
	redispkg.UseIncrAndDecr()
}

// 删除和追加
func TestDelAndAppend(t *testing.T) {
	redispkg.DelAndAppend()
}

// 列表插入
func TestInsertList(t *testing.T) {
	redispkg.InsertList()
}

// 列表查找
func TestReadList(t *testing.T) {
	redispkg.ReadList()
}

// 列表删除
func TestDelList(t *testing.T) {
	redispkg.DelList()
}
