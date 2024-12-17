package main
// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)
//全局变量
var n7 = 100
var n8 = 9.8

//一次性声明
var (
	n9 = 600
	n10 = "netty"
)
func main() {
	/* var num int = 18
	fmt.Println("num=",num)

	//在没有进行赋值操作时会输出默认值
	var num1 int
	fmt.Println("num1=",num1)

	//没有写变量的类型，那么会自动根据后面的值进行类型的推断
	var num2 = "tom"
	fmt.Println("num2=",num2)

	//省略var
	sex := "男"
	fmt.Println("sex=",sex)*/
	/* num= 18
		num1= 0
		num2= tom
		sex= 男 */
	//局部变量
	//声明多个变量
	var n1,n2,n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	//多个变量赋值
	var n4,name,n5 = 10,"jack",6.6
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	n6,height := 9.9,100.60
	fmt.Println(n6)
	fmt.Println(height)
	/*  0
		0
		0
		10
		jack
		6.6
		9.9
		100.6 在小数点后最后一位是0的话会被省略 */	

		//Printf函数：格式化的，把变量的类型填充到%T的位置上
		// name的类型：string
		fmt.Printf("name的类型：%T",name)
		fmt.Println()
		//输出该变量类型的空间大小
		//16
		fmt.Println(unsafe.Sizeof(name))

}