/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/25 6:29 PM
 */

package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type MathService struct {
}

func (m *MathService) Multi(a int, sum *int) error {
	*sum = a * a
	return nil
}

func main() {
	userService := new(MathService)
	// 注册服务
	err := rpc.Register(userService)
	if err != nil {
		return
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
