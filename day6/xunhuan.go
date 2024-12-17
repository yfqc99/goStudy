package main
import "fmt"
func main() {
	/* var num int = 0 */
	//for 初始表达式，不能用var声明
	/* for i :=1;i<50;i++{
		num +=i
	}
	fmt.Println(num) */
	// 1225


	/* i :=1
	for i<50{
		num +=i
		i++
	}
	fmt.Println(num) */
	// 2450
	
	/* for {
		fmt.Println("死循环")
	}
	for ;; {
		fmt.Println("死循环")
	} */

	var str string = "hello golang你好"
	//按照字节遍历输出，中文会出现乱码，因为中文占三个字节
	for i := 0; i <len(str); i++ {
		fmt.Printf("%c",str[i])
	}
	// hello golangä½ å¥½
	fmt.Println()
	//这里中文不会乱码，按照字符进行遍历
	for i, value := range str {
		fmt.Printf("索引为：%d，具体数值为：%c \n",i,value)
	}
	/* 	索引为：0，具体数值为：h
		索引为：1，具体数值为：e 
		索引为：2，具体数值为：l
		索引为：3，具体数值为：l
		索引为：4，具体数值为：o
		索引为：5，具体数值为：
		索引为：6，具体数值为：g
		索引为：7，具体数值为：o
		索引为：8，具体数值为：l 
		索引为：9，具体数值为：a 
		索引为：10，具体数值为：n
		索引为：11，具体数值为：g
		索引为：12，具体数值为：你
		索引为：15，具体数值为：好 */
}