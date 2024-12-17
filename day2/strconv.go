package main
import (
	"fmt"
	"strconv"
)
func main() {
	var n1 int = 16
	//参数一：整数必须是int64
	//参数二：以什么进制方式转换为字符串
	var s1 string = strconv.FormatInt(int64(n1),10)
	fmt.Printf("s1的类型：%T，s1 = %q \n",s1,s1)
	var n2 float64=4.29
	//参数一：浮点数必须是float64
	//参数二：以什么进制方式转换为字符串 'f'小数的十进制表示
	//参数三：保留小数后几位
	//参数四：浮点数类型
	var s2 string = strconv.FormatFloat(n2,'f',9,64)
	fmt.Printf("s2的类型：%T，s2 = %q \n",s2,s2)

	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3的类型：%T，s3 = %q \n",s3,s3)
	/* 	s1的类型：string，s1 = "16" 
		s2的类型：string，s2 = "4.290000000" 
		s3的类型：string，s3 = "true" */

}