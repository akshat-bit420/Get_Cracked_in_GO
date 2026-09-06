package main
import (
	"fmt"
	"runtime"
)
// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim{
// 		// in this condition we can see we made a variable that has a value,
// 		// and in the same line we also had our condition separated by " ; ",
// 		// v is only initialized in the "if" block so it can't accessed outside the func.
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n",v, lim)
// 	}
// 	// won't be able to use v here though, only in a true branch of if-else conditions.

// 	return lim
// }

// func main() {
// 	fmt.Println(
// 		pow(3,2,10),
// 		pow(3,3,10),
// 	)
// }

func main() {
	fmt.Print("Go runs on:- ")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Printf("%s. But I wanna use arch though",os)
	case "arch":
		fmt.Println("I use arch btw.")
	default:
		fmt.Printf("%s.\n", os)		
	}
}