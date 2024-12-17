package main
import (
	"fmt"
	"strconv"
)
func main() {
	var s1 string = "16"
	//参数一：要转换的字符串
	//参数二：要转成的进制类型
	//参数三：数据的类型
	var n1 , _ = strconv.ParseInt(s1,10,64)
	fmt.Printf("n1的类型：%T，n1 = %v \n",n1,n1)

	var s2 string = "4.29"
	//参数一：要转换的字符串
	//参数二：数据的类型
	var n2 , _ = strconv.ParseFloat(s2,64)
	fmt.Printf("n2的类型：%T，n2 = %v \n",n2,n2)

	var s3 string = "true"
	/* multiple-value strconv.ParseInt(s1, 10, 64) 
	(value of type (i int64, err error)) in single-value context */
	//ParseBool返回值有两个：（value,err）
	//value就是要得到的卡类型数据，err出现的错误
	//用_直接忽略
	var n3,_ = strconv.ParseBool(s3)
	fmt.Printf("n3的类型：%T，n3 = %v \n",n3,n3)

	//当字面值无法被正确转换时，不会报错，直接输出该类型的默认值
	var s4 string = "golang"
	var n4,_ = strconv.ParseBool(s4)
	fmt.Printf("n4的类型：%T，n4 = %v \n",n4,n4)
	/* 	n1的类型：int64，n1 = 16 
		n2的类型：float64，n2 = 4.29 
		n3的类型：bool，n3 = true
		n4的类型：bool，n4 = false */

}