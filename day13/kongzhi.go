package main
import "fmt"
func main() {
	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Println("学生",i+1,"的成绩：")
		fmt.Scanln(&scores[i])
	}
	var sum =0
	for i := 0; i < len(scores); i++ {
		sum+=scores[i]
	}
	fmt.Println(sum)
	fmt.Println(sum/5)
	/* 	学生 1 的成绩：99
		学生 2 的成绩：66
		学生 3 的成绩：77
		学生 4 的成绩：88
		学生 5 的成绩：44
		374
		74 */
	
		for key,value := range scores {
			fmt.Println("学生",key+1,"的成绩：",value)
		}
		/* 	学生 1 的成绩： 99
			学生 2 的成绩： 66
			学生 3 的成绩： 77
			学生 4 的成绩： 88
			学生 5 的成绩： 44 */

		//不接收key
		for _,value := range scores {
			fmt.Println("学生的成绩：",value)
		}
}