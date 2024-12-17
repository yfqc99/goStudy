package main
import (
	"fmt"
	"study/main/src/day9/utils"
)
var num int =test()
func test() int {
	fmt.Println("Test函数")
	return 10
}
func init(){
	fmt.Println("Init函数")
}
func main() {
	fmt.Println("Main函数")
	fmt.Println("Age=",utils.Age,"Name=",utils.Name)
	/* 	utils的init
		Test函数
		Init函数
		Main函数
		Age= 18 Name= 李四 */
	/* 	Test函数
		Init函数
		Main函数 */
	/* 	Init函数
		Main函数 */
}