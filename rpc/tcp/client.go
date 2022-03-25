/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/25 6:08 PM
 */

package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	// 建立链接
	dial, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dial error ", err)
	}
	var result string
	for i := 0; i < 5; i++ {
		// 发起请求
		_ = dial.Call("HelloService.Say", "go", &result)
		fmt.Println(result)
		time.Sleep(time.Second * 2)
	}
}
