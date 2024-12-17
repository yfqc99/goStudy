package main
import "fmt"
func main() {
	//key不存在增加
	//key存在更新
	b :=make(map[int]string)
	b[20221554]="张三"
	b[20221556]="李四"
	b[20221554]="王五"
	b[20221553]="周六"
	fmt.Println(b)
	// map[20221553:周六 20221554:王五 20221556:李四]

	//delete删除，key不存在也不会报错
	delete(b,20221554)
	delete(b,20221551)
	fmt.Println(b)
	// map[20221553:周六 20221556:李四]

	//清空
	//将map遍历逐清空
	//直接make一个新的，原来的成为垃圾，被gc回收

	//查找
	value,flag :=b[20221553]
	fmt.Println(value," ",flag)
	v,f :=b[20221550]
	fmt.Println(v," ",f)
	/* 	周六   true
		   false */
}
