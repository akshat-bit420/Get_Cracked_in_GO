package main
import "fmt"

type Vertex struct {
	x int 
	y int
}
// func main() {
// 	v := Vertex{1,9}
// 	v.x = 2
// 	// the ' . ' is used to access the values of a struct. 
// 	fmt.Println(v.x)
// }
/*How pointers and structs will work together*/
func main() {
	v  := Vertex{7,8}
	p  := &v
	p.x = 2e6
	// 1e notation works as power but only in terms of 10,
	// 2e6 means 6 times 0 after 2
	fmt.Println(v)
}
