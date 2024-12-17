package main
import "fmt"
func main() {
	//变量声明
	var age int
	//变量赋值
	age = 10
	//变量使用
	fmt.Println("age=",age)

	//声明和赋值合成一句
	// .\various.go:14:6: age2 declared and not used
	var age2 int = 19

	//变量重复声明会报错
	/* .\various.go:17:6: age redeclared in this block      
        .\various.go:5:6: other declaration of age  */
	var age int = 19

	//go语言中无法自动转换类型
	/*  .\various.go:20:16: cannot use 12.56 (untyped float constant) 
	as int value in variable declaration (truncated) */
	var num int = 12.56
}