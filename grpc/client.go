/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/28 3:25 PM
 */

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shershon1991/go-study-example/grpc/sever/hello"
)

func main() {
	// 建立链接
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial err ", err)
		return
	}
	// 延迟关闭链接
	defer conn.Close()
	// 实例化客户端
	client := hello.NewUserServiceClient(conn)
	// 发起请求
	reply, err := client.Say(context.TODO(), &hello.Request{Name: "张三"})
	if err != nil {
		return
	}
	fmt.Println("返回：", reply.Result)
}
