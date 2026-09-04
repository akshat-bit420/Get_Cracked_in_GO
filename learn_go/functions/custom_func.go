package main
import "fmt"

// func add(x int, y int) int {
// 	 func is the keyword used to make a function in go,
// 	 more importantly the declaration of variable type is
// 	 very different in Go, in comparison with it's parent
// 	 language "C", c wants to know the variable(parameter in this case) type before 
// 	 it's being made where as in Go, it's opposite.
// 	 and it's mandatory to explicitly mentions the return type.
// 	return x + y
// }
// func swap(x,y string) (string, string) {
// 	return y,x
// // Go can also return multiple values from inside, it takes 2 parameter.
// // as both the parameters are same we have merged it via single string for declaration
// // just takes x and y and return y inplace of x and x in place of y.
// }
func split(sum int) (x,y int) {
	// we have a func split(takes 1 parameter and returns 2 vals of type int)
	// we have named the return values by x and y, and they are now initialized 
	// by a value 0 implicitily, meaning we can directly use those values,
	// which helps us not by writing return x , y (just return is sufficient)
	x = sum * 4 / 9
	y = sum - x
	return 
}

func main() {	
	// a , b := swap("Hello", "world")
	// the  " := " is used when we don't have to declare the variable type,
	// as Go does it automatically, just like python.
	fmt.Println(split(17))
}