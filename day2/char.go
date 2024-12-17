package main
import "fmt"
func main() {
	//字符类型本质就是一个整数，可以直接参与运算
	//输出字符时会将对应的码值做一个输出
	//字母，数字，标点等字符，底层是ASCII存储
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '(' //40
	fmt.Println(c2+20)
	//汉字字符，底层是Unicode码
	//当溢出时可以使用int类型
	var c3 int = '中'
	fmt.Println(c3)
	//字符对应的是UTF-8编码
	/* 	97
		60
		20013 */
	var c4 byte = 'A'
	//显示对应的字符，必须使用格式化输出
	// c4对应的字符为：A
	fmt.Printf("c4对应的字符为：%c",c4)
}