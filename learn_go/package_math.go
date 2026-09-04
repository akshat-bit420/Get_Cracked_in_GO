package main
import (
	"fmt"
	"math"
)
// This functions teaches us about the go encapsulation technique, 
// that how it uses encapsulation.
func main() {
	fmt.Println(math.Pi)
	// here as we can see if we use pi the compiler will show an error
	// cause all the varibals or any data type, whose name starts with 
	// lower case will be hidden from external access, only variables with 
	// upper case will be able to access from external package.
}