package main
import "fmt"
func main() {
	var arr [2][3]int
	fmt.Println(arr)
	// [[0 0 0] [0 0 0]]
	fmt.Printf("arr的地址：%p \n",&arr)
	fmt.Printf("arr[0]的地址：%p \n",&arr[0])
	fmt.Printf("arr[0][0]的地址：%p \n",&arr[0][0])
	/* 	arr的地址：0xc00000e540 
		arr[0]的地址：0xc00000e540
		arr[0][0]的地址：0xc00000e540 */

	arr[0][0]=44
	arr[0][2]=66
	arr[1][1]=88
	fmt.Println(arr)
	var arr1 [2][3]int = [2][3]int{{1,4,7},{2,5,8}}
	fmt.Println(arr1)
	/* 	[[44 0 66] [0 88 0]]
		[[1 4 7] [0 0 0]] */
	
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Print(arr1[i][j]," ")
			// 1 4 7 2 5 8
		}
		fmt.Println()
	}
	
	for _,value := range arr1 {
		for _,v := range value {
			fmt.Print(v,"\t")
		}
		/* 	1       4  7
			2       5  8 */
		fmt.Println()
	}
}