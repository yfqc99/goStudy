package main
import "fmt"
func main() {
	var arr [6]int = [6]int{3,6,9,1,4,7}
	slice := arr[1:4]
	fmt.Println(slice)
	// [6 9 1]
	//不是直接追加到当前数组，而是创建一个新的数组
	//将slice当前的元素复制到新数组后再增加追加的元素
	/* append(slice, 88, 99) (value of type []int) is not used
	append(slice,88,99) */
	//slice1指向的是新数组
	slice1 :=append(slice,88,99)
	fmt.Println(slice)
	fmt.Println(slice1)
	// [6 9 1]
	// [6 9 1 88 99]
	slice =append(slice,88,99)
	fmt.Println(slice)
	// [6 9 1 88 99]

	//切片追加给切片
	slice2 :=[]int {66,44}
	/* cannot use slice2 (variable of type []int) 
	as int value in argument to append
	slice =append(slice,slice2) */
	//...代表追加的是切片
	slice =append(slice,slice2...)
	fmt.Println(slice)
	// [6 9 1 88 99 66 44]
}