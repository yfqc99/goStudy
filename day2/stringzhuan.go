package main
import "fmt"
func main() {
	var n1 int = 19
	//根据%d等参数生成格式化字符串，并返回
	var s1 string =fmt.Sprintf("%d",n1)
	//%V将原始字面值填入对应的位置
	fmt.Printf("s1的类型：%T，s1 = %v \n",s1,s1)
	var n2 float32 = 4.78
	var s2 string =fmt.Sprintf("%f",n2)
	//%q对应单引号括起来的go语法字符字面值
	fmt.Printf("s2的类型：%T，s2 = %q \n",s2,s2)
	var n3 bool = false
	var s3 string =fmt.Sprintf("%t",n3)
	fmt.Printf("s3的类型：%T，s3 = %q \n",s3,s3)
	var n4 byte = 'a'
	var s4 string =fmt.Sprintf("%c",n4)
	fmt.Printf("s4的类型：%T，s4 = %q \n",s4,s4)
	/* 	s1的类型：string，s1 = 19 
		s2的类型：string，s2 = "4.780000" 
		s3的类型：string，s3 = "false"
		s4的类型：string，s4 = "a" */
}