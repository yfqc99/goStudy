package main
import "fmt"
func main() {
	var arr []int = []int{3,6,9,1,4,7}
	var b []int =make([]int,10)
	//将arr里面的元素拷贝给b
	//将arr元素复制到b指向的数组中
	copy(b,arr)
	fmt.Println(b)
	// [3 6 9 1 4 7 0 0 0 0]
}