/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/25 6:55 PM
 */

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"shershon1991/go-study-example/rpc/jsonrpc/dto"
)

func main() {
	// 建立链接
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		fmt.Println("rpc.Dial error ", err)
		return
	}
	// 使用json编码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	// 参数rpc
	arg := dto.SumParam{A: 20, B: 80}
	// 结果
	var sum dto.SumRes
	err = client.Call("MathService.Sum", &arg, &sum)
	if err != nil {
		fmt.Println("client.Call err ", err)
		return
	}
	fmt.Printf("res:%+v \n", sum)
}
