package main
import (
	"fmt"
	"time"
)
func main() {
	//返回值类型是结构体类型
	now :=time.Now()
	fmt.Printf("返回值类型：%T，值：%v \n",now,now)
	/* 返回值类型：time.Time，
	值：2024-12-13 17:22:14.6276704 +0800 CST m=+0.008440901 */

	fmt.Println("年：",now.Year())
	fmt.Println("月：",now.Month())
	fmt.Println("月：",int(now.Month()))
	fmt.Println("日：",now.Day())
	fmt.Println("时：",now.Hour())
	fmt.Println("分：",now.Minute())
	fmt.Println("秒：",now.Second())
	/* 	年： 2024
		月： December
		月： 12
		日： 13
		时： 17
		分： 27
		秒： 27 */

	//日期的格式化
	//直接输出
	fmt.Printf("当前年月日：%d-%d-%d 时分秒：%d:%d:%d \n",
	now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second() )
	//将拼接的字符串返回
	datestr :=fmt.Sprintf("当前年月日：%d-%d-%d 时分秒：%d:%d:%d",
	now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second() )
	fmt.Println(datestr)
	/* 	当前年月日：2024-12-13 时分秒：17:34:45
		当前年月日：2024-12-13 时分秒：17:34:45 */
	
	//参数字符串各个数字必须是固定的
	//分割符和数字个数可以改变，但数字不能变
	fmt.Println(now.Format("2006/01/02 15/04/05"))
	// 2024/12/13 17/39/04
	fmt.Println(now.Format("01-02 15:04:05"))
	// 12-13 17:41:19
}