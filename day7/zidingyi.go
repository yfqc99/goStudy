package main
import "fmt"


func test (num1 int,num2 int) (int,int){
	sum :=num1+num2
	sub :=num1-num2
	return sum,sub
}
//会跟据变量名称，自动传入对应位置
func test1 (num1 int,num2 int) (sum int,sub int){
	sum :=num1+num2
	sub :=num1-num2
	return
}
func main() {
	//给int类型起别名
	type myInt int
	var num myInt = 30
	fmt.Printf("num的类型：%T",num)
	// num的类型：main.myInt

	//go在识别的时候不认为两者是同一个数据类型
	/* cannot use num (variable of type 
		myInt) as int value in assignment */
	/* var num1 int = 20
	num1=num */
}