package main
import "fmt"
func main() {
	var num int = 10
	fmt.Println(num)
	var ptr *int = &num
	*ptr = 20
	fmt.Println("通过ptr修改后的值：",num)
	/* 	10
		通过ptr修改后的值： 20 */
	/* cannot use num (variable of type int) 
	as *int value in variable declaration */
	var ptr1 *int = num
	fmt.Println(*ptr1)
	/* cannot use &num (value of type *int) as 
	*float32 value in variable declaration */
	var pte2 *float32 = &num
	fmt.Println(*ptr2)
}