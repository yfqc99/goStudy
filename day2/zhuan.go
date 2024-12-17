package main
import "fmt"
func main() {
	// \n 换行
	fmt.Println("aaa\nbbb")
	// \b 退格
	fmt.Println("aaa\bbbb")
	// \r 光标回到本行的开头，后续输入会替换原有字符
	fmt.Println("aaaa\rbbb")
	// \t 制表符 8个字符为一组
	fmt.Println("aaaaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbb")
	fmt.Println("aaaa\tbbbbb")
	fmt.Println("aaaaaaaa\tbbbbb")
	// \"
	fmt.Println("\"Golang\"")
	/* 	aaa
		bbb
		aabbb
		bbba
		aaaaaaaaaaaaaa
		aaaaa   bbbbb
		aaaa    bbbbb
		aaaaaaaa        bbbbb
		"Golang" */
}