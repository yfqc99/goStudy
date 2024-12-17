package main
import "fmt"
func main() {
	/* var slice []int
	fmt.Println(slice)
	fmt.Println(cap(slice)) */
	// [] 0

	var arr [6]int = [6]int{1,4,7,2,5,8}
	var slice []int= arr[1:4]
	fmt.Println(slice)
	// [4 7 2]
	// index out of range [3] with length 3
	// fmt.Println(slice[3])

	/* var slice =arr[0:end]---->var slice =arr[:end]
	var slice =arr[start:len(arr)]---->var slice =arr[start:]
	var slice =arr[0:len(arr)]---->var slice =arr[:] */
	slice1 :=slice[1:2]
	fmt.Println(slice1)
	// [7]
	slice1[0]=66
	fmt.Println(arr)
	fmt.Println(slice)
	/* 	[1 4 66 2 5 8]
		[4 66 2] */
}