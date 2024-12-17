package main
import "fmt"

type Teacher struct {
	Name string
	Age int
	School string
}

func main()  {
	//不赋值会有默认的初始值
	var t1 Teacher
	fmt.Println(t1)
	t1.Name="王五"
	t1.Age=18
	t1.School="家里蹲"
	fmt.Println(t1)
	fmt.Println(t1.Age+10)
	/* 	{ 0 }
		{王五 18 家里蹲}
		28 */
	var t2 Teacher = Teacher{}
	t2.Name="张三"
	t2.Age=20
	t2.School="牛顿大学"
	fmt.Println(t2)
	var t3 Teacher = Teacher{"周六",30,"清北大学"}
	fmt.Println(t3)
	/* 	{张三 20 牛顿大学}
		{周六 30 清北大学} */

	//返回的结构体指针
	var t4 *Teacher =new(Teacher)
	// &{ 0 }
	fmt.Println(t4)
	(*t4).Name="麻子"
	(*t4).Age=50
	(*t4).School="东北大学"
	fmt.Println(*t4)
	//go语言简化代码，底层对其转化为(*t4).School
	t4.School="华北大学"
	fmt.Println(*t4)
	/* 	{麻子 50 东北大学}
		{麻子 50 华北大学} */

	var t5 *Teacher =&Teacher{}
	(*t5).Name="溜子"
	(*t5).Age=66
	(*t5).School="大学1"
	t5.School="大学2"
	fmt.Println(*t5)
	// {溜子 66 大学2}
	var t6 *Teacher =&Teacher{"名字",77,"学校"}
	fmt.Println(*t6)
	// {名字 77 学校}
}