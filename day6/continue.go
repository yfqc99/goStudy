package main
import "fmt"
func main() {

	/* for i := 1; i <=100; i++ {
		if i%6 !=0{
			continue
		}
		fmt.Print(i," ")
	} */
	// 6 12 18 24 30 36 42 48 54 60 66 72 78 84 90 96
	
	
	lable:
	for i := 1; i <=5; i++ {
		for j := 2; j <=4; j++ {
			if i==2 &&j==2{
				continue lable
			}
			fmt.Printf("i：%v，j：%v \n",i,j)
		}
		/* 	i：1，j：2 
			i：1，j：3 
			i：1，j：4
			i：3，j：2
			i：3，j：3
			i：3，j：4
			i：4，j：2
			i：4，j：3
			i：4，j：4
			i：5，j：2
			i：5，j：3
			i：5，j：4 */
	}
}