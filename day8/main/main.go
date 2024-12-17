package main
import (
	"fmt"
	//起别名后就只能通过别名调用
	util "study/main/src/day8/utils"
)
func main() {
	fmt.Println("main函数")
	util.Getutil()
	/* 	main函数
		测试utils */

	/* "study/main/src/day8/utils" 
	imported as util and not used
.\main.go:9:2: undefined: utils */
	utils.Getutil()
}