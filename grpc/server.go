/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/28 3:18 PM
 */

package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"shershon1991/go-study-example/grpc/sever/hello"
)

func main() {
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 注册服务
	hello.RegisterUserServiceServer(grpcServer, new(hello.UnimplementedUserServiceServer))
	// 监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("服务启动失败", err)
		return
	}
	grpcServer.Serve(listen)
}
