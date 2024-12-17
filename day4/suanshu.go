package main
import "fmt"
func main() {
	//+
	//表示正数
	var n1 int = +10
	fmt.Println(n1)
	//相加操作
	var n2 int = 4+5
	fmt.Println(n2)
	//字符串拼接
	var n3 string = "abc"+"def"
	fmt.Println(n3)

	// /
	//int类型数据运算结果为int
	fmt.Println(10/3)
	//浮点类型运算结果为浮点型
	fmt.Println(10.0/3)

	// %
	//等价：a%b=a-a/b*b
	fmt.Println(10%3)
	fmt.Println(-10%3)
	fmt.Println(10%-3)
	fmt.Println(-10%-3)

	// ++
	//go语言中，++，--操作只能单独使用，不能参与运算
	//只是数值加1或者减1，只能写在变量后面
	var a int = 10
	// unexpected ++, expected expression
	/* fmt.Println(a++)
	fmt.Println(++a) */
	a++
	fmt.Println(a)
	// syntax error: unexpected ++, expected }
	/* ++a
	fmt.Println(a) */
	
	/* 	10
		9
		abcdef
		3
		3.3333333333333335
		1
		-1
		1
		-1
		11 */
}