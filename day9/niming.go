package main
import "fmt"

/* method has multiple receivers      
.\niming.go:5:2: syntax error: 
unexpected return, expected ) */
/* Func1 :=func (num1 int,num2 int) int {
	return num1+num2
} */
var Func1 =func (num1 int,num2 int) int {
	return num1+num2
}
func main() {
	//匿名函数，定义的同时调用
	result := func (num1 int,num2 int) int {
		return num1+num2
	}(10,20)
	fmt.Println(result)
	// 30

	//将匿名函数赋给变量
	sub :=func (num1 int,num2 int) int {
		return num1+num2
	}
	fmt.Printf("result的类型：%T，sub的类型：%T \n",result,sub)
	fmt.Println(sub(30,70))
	/* 	result的类型：int，sub的类型：func(int, int) int 
		100 */

	fmt.Println(Func1(90,9))
	// 99
}