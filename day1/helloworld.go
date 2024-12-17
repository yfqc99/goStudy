package main //声明文件所在的包，每个go文件必须有归属的包
import "fmt" //引入程序需要用的包，使用包内的函数
func main() { //main 主函数 程序入口
   //注意go编译器是一行行编译，所以一行只能有一条语句
   fmt.Println("Hello, World!") //在控制台打印输出

   //当输入内容过长，要提高代码阅读性进行换行时，要使用下述格式
   //直接换行会报错
   fmt.Println("Hello, World!",
   "Hello, World!",
   "Hello, World!",
   "Hello, World!",
   "Hello, World!")
   
   //可以通过添加 ；的方式实现一行多条语句
   fmt.Println("Hello, World!");fmt.Println("Hello, World!");

   //报错，因为没有使用
   var test;
}