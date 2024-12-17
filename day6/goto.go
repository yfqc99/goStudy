package main
import "fmt"
func main() {
	fmt.Println("ceshi1")
	fmt.Println("ceshi2")
	goto label
	fmt.Println("ceshi3")
	fmt.Println("ceshi4")
	fmt.Println("ceshi5")
	label:
	fmt.Println("ceshi6")
	/* 	ceshi1
		ceshi2
		ceshi6 */
}