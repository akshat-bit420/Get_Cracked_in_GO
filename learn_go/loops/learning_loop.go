package main
import "fmt"
func main() {
	sum := 1
	// in Go, it's crazy that it only has 1 loop, The for loop
	// no while, no do while.
	// for i := 0; i <= 10; i++ {
	// 	sum += i
	// }
	for sum < 1000 {
		sum += sum
	}
	// using for loop as a while loop. Go being crazy
	fmt.Println(sum)
}