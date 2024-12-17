package main
import "fmt"

type Student struct{
	Age int
}
type Person struct{
	Age int
}
func main() {
	var p Person = Person{10}
	// var p Person = Person{18}
	var s Student = Student{18}
	/* cannot use p (variable of type Person) 
	as Student value in assignment */
	// s=p
	//进行强制转换
	s=Student(p)
	fmt.Println(s)
	// {10}
}