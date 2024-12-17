package main
import "fmt"
func main() {
	var num1 float32 = 3.14
	var num2 float32 = -3.14
	//可以用十进制和科学计数法表示
	var num3 float32 = 314E-2
	//314×10的2次方
	var num4 float32 = 314e+2
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	//64位比32位更准确
	//浮点数可能会有精度丢失的状况
	//建议使用float64，默认的浮点类型也是float64
	var num5 float32 = 256.000000916
	var num6 float64 = 256.000000916
	fmt.Println(num5)
	fmt.Println(num6)
/* 	3.14
	-3.14
	3.14
	31400
	256
	256.000000916 */
}