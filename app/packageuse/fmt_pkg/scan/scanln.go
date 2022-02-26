/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/02/23 10:38 PM
 */

package main

import "fmt"

func main() {
	var (
		name   string
		age    int64
		isBody bool
	)
	_, err := fmt.Scanln(&name, &age, &isBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("name:%s age: %d isBody:%t", name, age, isBody)
}
