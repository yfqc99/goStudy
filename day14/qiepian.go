package main
import "fmt"
func main() {
	var arr [6]int = [6]int{3,6,9,1,4,7}
    // 切片构建在数组之上
    //动态变化的数组，所以不用指定长度
    //切出从索引1到索引3的片段（不包含3）
    var slice []int = arr[1:3]
    slice1 := arr[2:5]
    fmt.Println("arr：",arr)
    fmt.Println("slice：",slice)
    fmt.Println("slice1：",slice1)
    fmt.Println("slice的元素个数：",len(slice))
    //底层切片的长度，可以动态变化，一般大约是元素的2倍
    fmt.Println("slice的容量：",cap(slice))
    /*  arr： [3 6 9 1 4 7]
        slice： [6 9]
        slice1： [9 1 4]
        slice的元素个数： 2
        slice的容量： 5 */
    fmt.Printf("slice[0]的地址：%p \n",&slice[0])
    fmt.Printf("arr[1]的地址：%p \n",&arr[1])
    /*  slice[0]的地址：0xc00013a068
        arr[1]的地址：0xc00013a068 */
    slice[1]=16
    fmt.Println("arr：",arr)
    fmt.Println("slice：",slice)
    /*  arr： [3 6 16 1 4 7]       
        slice： [6 16] */
}