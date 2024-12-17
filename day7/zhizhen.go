package main
import "fmt"

func test(num *int){
	*num=20
	fmt.Println(num)
}
func main() {
	var num int = 10
	test(&num)
	fmt.Println(num)
	/* 	0xc000122058
		20 */
}
