package main
import "fmt"
func main() {
	//只声明map内存没有分配空间，赋值后才会分配空间
	//数组声明后就会分配空间
	var a map[int]string
	//通过make函数初始化后分配空间
	//可以存放10个键值对
	//map中按照key进行从小到大排序
	//map中key不能重复，重复的key不会报错，但会将value进行修改
	//value可以重复
	//make函数的第二个参数size可以省略，默认分配1个起始空间大小
	a=make(map[int]string,10)
	a[20221554]="张三"
	a[20221556]="李四"
	a[20221555]="王五"
	a[20221556]="周六"
	a[20220011]="张三"
	fmt.Println(a)
	// map[20220011:张三 20221554:张三 20221555:王五 20221556:周六]
}