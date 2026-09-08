package main

import "fmt"

func main() {
	var a [2]string
	// an array of size 2 of type string, size of the array is part of it's type
	a[0] = "helu"
	a[1] = "belu"
	// the normal how we access the elements of array.
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,9,11}
	// here if we don't wanna do lengthy part, this is how we enter harcoded elements.
	fmt.Println(primes)
	// Go is a statically typed lang. So take that shi- in mind
	// means we won't be able to resize the array, it works as a constant
	// we won't use this as much, why? cause we have slices much better than arrays tbh(IN GO)
}