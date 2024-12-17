package main
import "fmt"
func main() {
	var s1 string="试试就好"
	fmt.Println(s1)
	//可以实现指向地址的改变
	var s2 string="abc"
	s2="def"
	//字符串一旦定义完成，其中的字符的值不能改变
	// cannot assign to s2[0] (value of type byte)
	// s2[0]='t'
	fmt.Println(s2)
	//表示形式
	//没有特殊字符
	var s3 string="abc"
	fmt.Println(s3)
	//有特殊字符，用反引号
	var s4 string=`
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '(' //40
	fmt.Println(c2+20)
	`
	fmt.Println(s4)
	/* 	试试就好
		def
		abc

				var c1 byte = 'a'
				fmt.Println(c1)
				var c2 byte = '(' //40
				fmt.Println(c2+20) */
	//字符串的拼接
	var s5 string = "abc"+"def"
	fmt.Println(s5)
	s5 +="hijk"
	fmt.Println(s5)
	//当字符串过长时，注意+保留在一行的末尾
	var s6 string = "abc"+"def"+
	"abc"+"def"+
	"abc"+"def"+
	"abc"+"def"
	fmt.Println(s6)
	/* 	abcdef
		abcdefhijk
		abcdefabcdefabcdefabcdef */
}
