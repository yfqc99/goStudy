package main
import "fmt"

//函数参数为空，返回值为一个参数列表为int，返回值为int的函数
func getSum() func (int) int {
	var sum int = 0
	return func (num int) int {
		sum = sum+num
		return sum
	}
}
// 返回的匿名函+匿名函数以外的变量形成一个闭包
func main() {
	f := getSum()
	//匿名函数中引用的变量会一直保存在内存中，可以一直使用
	fmt.Println(f(30))
	// 30
	fmt.Println(f(36))
	// 66
}