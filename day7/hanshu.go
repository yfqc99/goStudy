package main
import "fmt"
/* too many return values
        have (int, int)
        want (int)
.\hanshu.go:14:12: assignment mismatch: 
2 variables but cal returns 1 value */
//返回值只有一个，可以省略第二个括号不写
//没有顺序要求，main后面也可以
func cal (num1 int,num2 int) (int)  {
	var sum int = 0
	sum = num1+num2
	return sum
}
/* cal redeclared in this block       
.\hanshu.go:10:6: other declaration of cal */
func cal (num1 int){
	var sum int = 0
}
func main() {
	var num1 int =10
	var num2 int = 20
	/* not enough arguments in call to cal
        have (int, int)
        want (int, int, int) */
	sum :=cal(num1,num2)
	fmt.Println(sum)
	// 30
}