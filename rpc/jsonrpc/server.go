/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/25 6:50 PM
 */

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"shershon1991/go-study-example/rpc/jsonrpc/dto"
	"time"
)

type MathService struct {
}

func (m *MathService) Sum(arg *dto.SumParam, res *dto.SumRes) error {
	fmt.Printf("arg: %#v\n", arg)
	*res = dto.SumRes{
		Sum:  arg.A + arg.B,
		Time: time.Now(),
	}
	return nil
}

func main() {
	// 注册服务
	err := rpc.RegisterName("MathService", new(MathService))
	if err != nil {
		fmt.Println("rpc RegisterName error ", err)
		return
	}
	// 监听端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("listen error ", err)
		return
	}
	// 监听
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn error ", err)
			return
		}
		// 使用json编码
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
