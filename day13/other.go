package main
import "fmt"
func main() {
	var arr [3]int = [3]int{3,6,9}
	test(arr)
	fmt.Println(arr)
	// [3 6 9]

	test1(&arr)
	fmt.Println(arr)
	[10 6 9]
}
func test(arr [3]int)  {
	arr[0]=10
}
func test1(arr *[3]int)  {
	(*arr)[0]=10
}