package main
import "fmt"
type Student struct{
	Age int
}
type Stu Student
func main() {
	var s1 Student = Student{19}
	var s2 Stu = Stu{22}
	/* cannot use s2 (variable of type Stu)
	 as Student value in assignment */
	// s1=s2
	s1=Student(s2)
	fmt.Println(s1)
	fmt.Println(s2)
	/* 	{22}
		{22} */
}