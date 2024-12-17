package main
import "fmt"
func main() {
	fmt.Println(add(30,36))
}
func add(num1 int,num2 int) int {
	//程序遇到defer后不会立即执行
	// 而是将后面的语句压入栈内
	//在函数执行完毕后依次弹出执行
	defer fmt.Println("num1=",num1)
	defer fmt.Println("num2=",num2)
	//在入栈的同时会将当时变量的字面值同时入栈，
	//之后改变变量不会修改栈里的数据
	num1 +=300
	num2 +=300
	var sum int =num1+num2
	fmt.Println("sum=",sum)
	return sum
}
/* 	sum= 666
	num2= 36
	num1= 30
	666 */
/* 	sum= 66
	num2= 36
	num1= 30
	66 */