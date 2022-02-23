/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/02/23 10:38 PM
 */

package fmt_pkg

import "fmt"

func main() {
	var (
		name   string
		age    int64
		isBody bool
	)
	_, err := fmt.Scan(&name, &age, &isBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("name:%s age: %d isBody:%t", name, age, isBody)
}
