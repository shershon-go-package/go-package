/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/25 6:35 PM
 */

package main

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	// 建立链接
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println("error ", err)
	}
	var result int
	for i := 0; i < 10; i++ {
		err = client.Call("MathService.Multi", i, &result)
		fmt.Printf("i:%v result:%v \n", i, result)
		time.Sleep(time.Second)
	}
}
