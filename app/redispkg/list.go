/**
 * @Description 列表操作
 **/
package redispkg

import (
	"context"
	"fmt"
)

// 列表插入
func InsertList() {
	// 连接redis
	client, _ := ConnectSingle()
	ctx := context.Background()
	key := "insertList"
	//从列表头部插入,不存在则新建并插入数据
	for _, val := range []string{"php", "go"} {
		// 插入头部(左边)
		client.LPush(ctx, key, val)
		// 获取
		result, _ := client.LRange(ctx, key, 0, -1).Result()
		fmt.Printf("LPush 从头部插入【%v】: %v\n", val, result)
	}
	// 从列表尾部插入
	for _, val := range []string{"张三", "李四"} {
		// 插入尾部(右边)
		client.RPush(ctx, key, val)
		// 获取
		result, _ := client.LRange(ctx, key, 0, -1).Result()
		fmt.Printf("RPush 从尾部插入【%v】: %v\n", val, result)
	}
	result, _ := client.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("当前列表所有值: %+v\n", result)

	// 在指定的值前插入
	client.LInsertBefore(ctx, key, "php", "php5.6")
	result, _ = client.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("在php前插入%v,当前列表所有值: %v\n", "php5.6", result)
	// 在指定的值后插入
	client.LInsertAfter(ctx, key, "go", "go1.0")
	result, _ = client.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("在go后插入%v,当前列表所有值: %v\n", "go1.0", result)
}

// 列表读取
func ReadList() {
	// 连接redis
	client, _ := ConnectSingle()
	ctx := context.Background()
	key := "language-list"
	// 插入元素
	client.LPush(ctx, key, "php", "go", "java", "c", "c++", "python")
	fmt.Println("插入列表: ", "php", "go", "java", "c", "c++", "python")
	// 获取长度
	result, _ := client.LLen(ctx, key).Result()
	fmt.Println("列表长度: ", result)
	// =========获取指定key的值=======
	val, _ := client.LIndex(ctx, key, 0).Result()
	fmt.Println("获取索引为0的值: ", val)

	// =========获取指定范围的值=======
	// 获取[0,2]的元素
	strings, _ := client.LRange(ctx, key, 0, 2).Result()
	fmt.Printf("获取列表位置为[0,2]的元素：%v\n", strings)
	// 获取全部元素[0,-1]
	all, _ := client.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("获取列表所有[0,-1]元素：%v\n", all)
}

func DelList() {
	// 连接redis
	client, _ := ConnectSingle()
	ctx := context.Background()
	key := "language-list"
	// 插入元素
	client.LPush(ctx, key, "php", "go", "java", "c", "c++", "python")
	all, _ := client.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("当前列表所有值：%v\n", all)

	// 移出并获取列表的第一个元素
	first, _ := client.LPop(ctx, key).Result()
	fmt.Printf("LPop,移出并获取列表的第一个元素：%v\n", first)
	// 当前列表所有值
	all, _ = client.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("当前列表所有值：%v\n", all)
	// 移出并获取列表的最后一个元素
	last, _ := client.RPop(ctx, key).Result()
	fmt.Printf("RPop,移出并获取列表的最后一个元素：%v\n", last)
	// 当前列表所有值
	all, _ = client.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("当前列表所有值：%v\n", all)
}
