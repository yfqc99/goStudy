package main
import "fmt"
func main() {
	slice := make([]int,4,20)
	slice[0]=66
	slice[1]=88
	slice[2]=99
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i]," ")
	}
	fmt.Println()
	/* 	66 88 99 0 
		66 88 99 0 */
	for _, v := range slice {
		fmt.Print(v," ")
	}
}