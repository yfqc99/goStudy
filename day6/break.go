package main
import "fmt"
func main(){
	/* var sum int = 0
	for i := 1 ; i <= 100 ; i++ {
		sum += i
		fmt.Print(sum, " ") */
		/*go语言中if后必须有{} syntax error: unexpected if, expected { after for clause
		.\break.go:9:2: syntax error: unexpected }, expected 
		expression */
		// if sum >=300 break
		/* if sum >=300 {
			break
		}  */
		// 1 3 6 10 15 21 28 36 45 55 66 78 91 105 120 136 153 171 190 210 231 253 276 300
	// } 

	lable2:
	for i := 1; i <=5; i++ {
		/* label lable1 defined and not used
		lable1: */
		for j := 2; j <=4; j++ {
			fmt.Printf("i：%v，j：%v \n",i,j)
			if i==2 &&j==2{
				break lable2
			}
			/* 	i：1，j：2 
				i：1，j：3 
				i：1，j：4
				i：2，j：2 */
		}
	}
}