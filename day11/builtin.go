package main
import "fmt"
func main() {
	//new分配内存，实参是具体类型，返回对应类型的指针
	num :=new(int)
	fmt.Printf("num的类型：%T，num的值：%v"+
	"，num的地址：%v，num指针指向的值：%v",num,num,&num,*num)
	/* num的类型：*int，num的值：0xc00001a098，
	num的地址：0xc00000a028，num指针指向的值：0 */

}