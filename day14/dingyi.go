package main
import "fmt"
func main() {
	var arr [6]int = [6]int{3,6,9,1,4,7}
    slice := arr[2:5]
	fmt.Println(slice)

	//make 底层创建一个数组，外界不能访问，不能直接操作
	//通过slice引用数组对数组进行操作
	//参数一：切片类型
	//参数二：切片长度
	//参数三：切片容量
	slice1 := make([]int,4,20)
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	slice1[0]=66
	slice1[1]=88
	fmt.Println(slice1)
	/* 	[0 0 0 0]
		20
		[66 88 0 0] */

	//直接指定具体数组，类似于make创建
	slice2 :=[]int{1,4,7}
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	/* 	[1 4 7]
		3
		3 */
}