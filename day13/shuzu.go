package main
import "fmt"
func main() {
	//连续的空间地址
	var scores [5]int
	scores[0]=66
	scores[1]=99
	scores[2]=88
	scores[3]=44
	scores[4]=77
	var sum =0
	for i := 0; i < len(scores); i++ {
		sum+=scores[i]
	}
	fmt.Println(sum)
	fmt.Println(sum/5)
	/* 	374
		74 */
}