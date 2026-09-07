package main
import "fmt"

type Vertex struct {
	x int 
	y int
}
func main() {
	v := Vertex{1,9}
	v.x = 2
	// the ' . ' is used to access the values of a struct. 
	fmt.Println(v.x)
}
