package main
import "fmt"
func main() {
	test()
	fmt.Println("正常执行")
}
func test(){
	//recover必须搭配的defer使用
	defer func ()  {
		//捕获错误
		err :=recover()
		//捕获到则返回对应的错误信息
		//没有捕获到错误，返回nil即零值
		if err != nil {
			fmt.Println("出现错误")
			fmt.Println("err：",err)
		}
	//defer后面必须是函数调用，非函数调用（如普通变量）会报错
	// expression in defer must be function call
	}()
	num1 :=10
	num2 :=0 
	// panic: runtime error: integer divide by zero
	fmt.Println(num1/num2)
}
/* 出现错误
err： runtime error: integer divide by zero
正常执行 */