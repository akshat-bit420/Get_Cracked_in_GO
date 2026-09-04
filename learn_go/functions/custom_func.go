package main
import "fmt"

func add(x int, y int) int {
	// func is the keyword used to make a function in go,
	// more importantly the declaration of variable type is
	// very different in Go, in comparison with it's parent
	// language "C", c wants to know the variable(parameter in this case) type before 
	// it's being made where as in Go, it's opposite.
	//  and it's mandatory to explicitly mentions the return type.
	return x + y
}

func main() {	
	fmt.Println(add(12,13))
	// called the function 
}