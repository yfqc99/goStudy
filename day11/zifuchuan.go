package main
import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
	str :="golang你好"
	//内置函数，不需要导包
	//按字节统计字符串的长度
	fmt.Println(len(str))
	// 12
	//切片遍历
	r :=[]rune(str)
	for i := 0; i < len(r); i++ {
		// fmt.Print(r[i]," ")
		// 103 111 108 97 110 103 20320 22909
		fmt.Printf("%c \u0020",r[i])
		// g  o  l  a  n  g  你  好
	}
	fmt.Println()
	//字符串转整数
	num,_ :=strconv.Atoi("666")
	fmt.Printf("num的类型：%T，值：%v \n",num,num)
	// num的类型：int，值：666
	//整数转字符串
	str1 :=strconv.Itoa(99)
	fmt.Printf("str1的类型：%T，值：%v \n",str1,str1)
	// str1的类型：string，值：99

	//统计字符串中有几个指定的子串
	count :=strings.Count("golangandjava","go")
	fmt.Println(count)
	// 1
	//不区分大小写进行字符串比较
	flag :=strings.EqualFold("hello","HEllo")
	fmt.Println(flag)
	//区分大小写进行字符串比较
	fmt.Println("hello"=="HEllo")
	/* 	true
		false */

	//返回子串在字符串中第一次出现的索引值，没有返回-1
	fmt.Println(strings.Index("golangandjava","ga"))
	fmt.Println(strings.Index("golangandjava","ga0"))
	/* 	5
		-1 */
	//参数四：要替换的数量，-1表示全部替换
	fmt.Println(strings.Replace("goandjavago","go","goland",-1))
	fmt.Println(strings.Replace("goandjavago","go","goland",1))
	/* 	golandandjavagoland
		golandandjavago */

	/* cannot use '-' (untyped rune constant 45) as 
	string value in argument to strings.Split */
	//根据指定字符切割字符串，形成一个数组
	arr :=strings.Split("go-python-java","-")
	fmt.Println(arr)
	// [go python java]

	//将字符串的字符进行大小写转换
	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("go"))
	/* 	go
		GO */

	//将字符串左右两边的空格去掉
	fmt.Println(strings.TrimSpace("   go and java   "))
	// go and java

	//将字符串左右两边指定的字符去掉
	fmt.Println(strings.Trim("-golang-java-","-"))
	// golang-java

	//将字符串左边指定的字符去掉
	fmt.Println(strings.TrimLeft("-golang-java-","-"))
	// golang-java-

	//将字符串右边指定的字符去掉
	fmt.Println(strings.TrimRight("-golang-java-","-"))
	// -golang-java

	//是否以指定字符串开头
	fmt.Println(strings.HasPrefix("https://golang","https"))
	// true

	//是否以指定字符串结尾
	fmt.Println(strings.HasSuffix("golang.png",".png"))
	// true
}