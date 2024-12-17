package main
import "fmt"

func test(num int){
	fmt.Println("test")
}

type myFunc func(int)
func test1(num int,testFunc myFunc){
	fmt.Println("test1")
}
func main() {
	/* a :=test
	fmt.Printf("a的类型：%T，test的类型：%T",a,test) */
	// a的类型：func(int)，test的类型：func(int)
	test1(10,test)
	// test1
}