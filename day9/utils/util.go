package utils
import "fmt"

var Name string
var Age int
func init() {
	Name="李四"
	Age=18
	fmt.Println("utils的init")
}