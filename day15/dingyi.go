package main
import "fmt"
func main() {
	var a map[int]string
	a=make(map[int]string,10)
	a[20221554]="张三"
	a[20221556]="李四"
	fmt.Println(a)
	b :=make(map[int]string)
	b[20221554]="张三"
	b[20221556]="李四"
	fmt.Println(b)

	//每个数据后面都要加逗号，否则会报错
	/* unexpected newline in composite 
	literal; possibly missing comma or } */
	c :=map[int]string{
		20221554 : "张三",
		20221556 : "李四",
	}
	c[20221553]="王五"
	fmt.Println(c)
	/* 	map[20221554:张三 20221556:李四]
		map[20221554:张三 20221556:李四]
		map[20221553:王五 20221554:张三 20221556:李四] */
}