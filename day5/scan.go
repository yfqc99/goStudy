package main
import "fmt"
func main() {
	/* var age int
	fmt.Println("输入学生年龄：")
	//通过获得地址值，直接改变地址空间里的值，达到输入的效果
	fmt.Scanln(&age)
	var name string
	fmt.Println("输入学生姓名：")
	//录入数据，类型要匹配，底层会自动判定数据
	//不匹配的类型不会报错，会输出默认值
	fmt.Scanln(&name)
	var score float32
	fmt.Println("输入学生成绩：")
	fmt.Scanln(&score)
	var isVIP bool
	fmt.Println("是否是VIP：")
	fmt.Scanln(&isVIP)
	fmt.Printf("年龄：%v，姓名：%v，成绩：%v，是否是VIP：%v",age,name,score,isVIP) */
/* 	输入学生年龄：
	1
	输入学生姓名：
	李四
	输入学生成绩：
	60
	是否是VIP：
	false
	年龄：1，姓名：李四，成绩：60，是否是VIP：false */


	var age int
	var name string
	var score float32
	var isVIP bool
	fmt.Printf("录入年龄，姓名，成绩，是否是VIP，用空格分隔 \n")
	fmt.Scanf("%d %s %f %t",&age,&name,&score,&isVIP)
	fmt.Printf("年龄：%v，姓名：%v，成绩：%v，是否是VIP：%v",age,name,score,isVIP)
	/* 	录入年龄，姓名，成绩，是否是VIP，用空格分隔 
		1 王五 55 false 
		年龄：1，姓名：王五，成绩：55，是否是VIP：false */

}