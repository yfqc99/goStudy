package main
import (
	"fmt"
	"errors"
)
func main() {
	err :=test()
	if err !=nil{
		fmt.Println(err)
		//中止程序
		panic(err)
	}
	fmt.Println("正常执行")
	/* 	除数不能为0
		panic: 除数不能为0 */
}
func test() (err error){
	num1 :=10
	num2 :=0 
	/* 	除数不能为0
		正常执行 */
	// num2 :=5
	/* 	2
		正常执行 */
	//自定义错误
	if(num2 == 0){
		//New函数会返回err类型的值
		return errors.New("除数不能为0")
	}
	fmt.Println(num1/num2)
	//没有错误返回零值
	return nil
}