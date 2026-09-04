package main
import (
	"fmt"
	"math/rand"
	// here rand represents that random package will be used, so we have to provide the path.
	// the package name has to be in lower case.
	// the package name should be short cause it easy to use.\
	// package name should match the folder name(helps to write good go.)
)
func main() {
	fmt.Println("My fav No. is", rand.Intn(10))
	// works same as python random function, gives random number between 0-9.
}