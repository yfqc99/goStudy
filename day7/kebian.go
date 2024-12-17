package main
import "fmt"

//可以传入任意数量的int类型数据
func test (nums...int){
	//内部处理可变参数时，当做切片来处理
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i]," ")
	}
	fmt.Println()
}
func main() {
	test()
	test(3)
	test(4,5)
	test(7,8,9)
/* 		
	3 
	4 5
	7 8 9 */
}