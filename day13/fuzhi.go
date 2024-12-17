package main
import "fmt"
func main() {
	var arr [3]int = [3]int{3,6,9}
	fmt.Println(arr)

	var arr2 = [3]int{1,4,7}
	fmt.Println(arr2)

	var arr3 = [...]int{5,8,2}
	fmt.Println(arr3)

	var arr4  = [...]int{1:22,5:88,6:66}
	fmt.Println(arr4)
	/* 	[3 6 9]
		[1 4 7]
		[5 8 2]
		[0 22 0 0 0 88 66] */

	fmt.Printf("arr4的类型：%T",arr4)
	// arr4的类型：[7]int
	
	
	test(arr)
	fmt.Println(arr)
	// [3 6 9]
}

func test(arr [3]int)  {
	arr[0]=10
}