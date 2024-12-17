package main
import "fmt"
func main() {
	/* var count int = 100
	//if后一定要有空格，和条件表达式分隔
	//左右的括号可以写 
	if count <30{
		fmt.Println("进入条件语句")
	}
	fmt.Println("不满足条件")
	// 不满足条件

	//可以将变量的定义和赋值拼接在条件语句中
	if count1 :=20;count1 <30{
		fmt.Println("进入条件语句")
	}
	// 进入条件语句 */

	/* var count int = 100
	//else语句必须跟在if语句的最后}后面
	//否则报错unexpected else, expected }
	if count<30{
		fmt.Println("进入条件语句")
	}else{
		fmt.Println("不满足条件")
	}
	// 不满足条件 */

	/* var count int = 100
	if count<30{
		fmt.Println("进入条件语句")
	}else if count<50{
		fmt.Println("满足条件20")
	}else if count<60{
		fmt.Println("满足条件10")
	}else {
		fmt.Println("满足条件大于60")
	}
	// 满足条件大于60 */

	var count int = 100
	//不需要写break
	/* 分支相同报错.\fenzhi.go:54:8: duplicate case 100 
	(constant of type int) in expression switch
        .\fenzhi.go:52:8: previous case */
	switch count {
		case 110,120:
			fmt.Println("满足条件")
		case 20:
			fmt.Println("满足条件20")
		case 30:
			fmt.Println("满足条件30")
		case 40:
			fmt.Println("满足条件40")
		case 50:
			fmt.Println("满足条件50")
		/* 	满足条件100
			满足条件130 */
		case 100:
			fmt.Println("满足条件100")
			fallthrough
		case 130:
			fmt.Println("满足条件130")
		/* cannot convert "shishi" 
		(untyped string constant) to type int */
		/* case "shishi":
			fmt.Println("满足条件") */
		/* case 100:
			fmt.Println("满足条件100") */
		default :
			fmt.Println("不满足条件")
	}
	// 满足条件100

	/* switch {
		case count==50:
			fmt.Println("满足50")
		case count==100:
			fmt.Println("满足100")
	}
	// 满足100

	//var定义和不加；会报错，
	switch b :=7; {
		case b>6:
			fmt.Println("大于6")
		case b<=6:
			fmt.Println("小于等于6")
	}
	// 大于6 */
	
}