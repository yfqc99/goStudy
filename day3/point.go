package main
import "fmt"
func main() {
	var age int = 18
	//输出该空间在内存中的地址
	// 0xc00001a078
	fmt.Println(&age)

	//定义指针变量
	//*int int类型的指针变量
	//&age 具体地址
	var ptr *int = &age
	fmt.Println("ptr内存空间存储的地址值：",ptr)
	fmt.Println("ptr本身的内存空间地址：",&ptr)

	//获取ptr指针指向的数据值
	fmt.Println("ptr指向的数值：",*ptr)
	/* 	0xc0000a6038
		ptr内存空间存储的地址值： 0xc0000a6038
		ptr本身的内存空间地址： 0xc0000c2018
		ptr指向的数值： 18 */
}