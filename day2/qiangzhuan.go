package main
import "fmt"
func main() {
	var n1 int = 100
	/* cannot use n1 (variable of type int) 
	as float32 value in variable declaration */
	//var n2 float32=n1 //隐式转换报错
	fmt.Println(n1)
	//强制转换
	var n2 float32=float32(n1)
	fmt.Println(n2)

	//将int64转为int8的时候，编译不会报错，但数据有可能溢出
	//精度丢失
	var n3 int64 = 888888
	var n4 int8 = int8(n3)
	fmt.Println(n4)
	/* 	100
		100
		56 */
	var n5 int32 = 12
	// var n6 int64 = n5+32
	/* cannot use n5 + 32 (value of 
	type int32) as int64 value in variable declaration */
	// fmt.Println(n6)
	var n6 int64 = int64(n5)+32
	fmt.Println(n6)

	var n7 int64 = 12
	//因为+127没有超出int8的范围所以编译不会报错但数据可能溢出
	var n8 int8 = int8(n7)+127
	//因为+128超出int8的范围所以编译报错
	// 128 (untyped int constant) overflows int8
	var n9 int8 = int8(n7)+128
	fmt.Println(n8)
	fmt.Println(n9)
}
