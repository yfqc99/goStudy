package main
import "fmt"
func main() {
	var arr [3]int
	fmt.Println(len(arr))
	// %p以十六进制值形式输出
	fmt.Printf("arr的地址：%p \n",&arr)
	fmt.Printf("第一个空间的地址：%p\n",&arr[0])
	/* 	3
		arr的地址：0xc000018180
		第一个空间的地址：0xc000018180 */
		fmt.Printf("第二个空间的地址：%p \n",&arr[1])
		fmt.Printf("第三个空间的地址：%p \n",&arr[2])
/* 	第二个空间的地址：0xc000018188
	第三个空间的地址：0xc000018190 */

	var num int
	fmt.Println(&num)
	// 0xc00001a0b8
	fmt.Println(&arr)
	fmt.Println(arr)
	/* &[0 0 0]
		[0 0 0] */
}