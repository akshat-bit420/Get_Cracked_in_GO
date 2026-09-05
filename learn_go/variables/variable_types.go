package main 
// import (
// 	"fmt"
// 	"math/cmplx"
// )
// var (
// 	toBe   bool       = false
// 	// Go has 4 major data types: 
// 	// 1). Basic - Numbers, bool, String
// 	// INT contains- int, int8, int16, int32, int64 [uint, uint8, uint16, uint32, uint64, uintptr]
// 	// These numbers decides the space of a variable or something
// 	// eg- int8, a variable can have numbers b/w (-127,128) [means it can use 8 bits to represent a number]
// 	// 8bits mean - 2^8 = 256, meaning it can contain numbers in [-127, 128] in this range.
// 	// in the same way we can determine the ranges of other int** variables.
// 	// UINT means (usigned int as 'u' suggests), so only positive numbers,
// 	// uint8 can have numbers from [0,256]. 
// 	// if our os is of 32-bit architecture then int is of int32, int64 in case of 64-bit architecture.
// 	// bytes -> an alias for uint8(but it can also contain strings, will see that later on)
// 	// float32 and float64 , complex64 and complex128 [ i hate complex numbers if this is that].
// 	// 2). Aggregate - 1. Array(type has to be same, if it has int then only integers are allowed) , 
// 	// 2.  Struct it is a collecction of hetrogeneous data types. Called as aggeregate cause they
// 	// consist of basic data types.
// 	// 3). Reference - 1. Pointers , 2. Slices , 3. Functions , 4. Channels , 5. Maps
// 	// called as reference cause they themselves does not contain any data(as they refer to
// 	// another data types)
// 	// 4). Interfaces - To be covered yet...

// 	maxInt uint64     = 1<<64 -1
// 	z      complex128 = cmplx.Sqrt(-3 + 15i)

// )
// func main() {
// 	fmt.Printf("Type: %T value: %v\n", toBe,toBe)
// 	// %T verb is used to type the print of a variable, %v is used to print the value
// 	fmt.Printf("Type: %T value: %v\n",maxInt,maxInt)
// 	fmt.Printf("Type: %T value: %v\n",z,z)
// }

import "fmt"
func main() {
	var i int
	var j float64
	var k bool 
	var l string
	// if we declare our variable without a value, Go automatically takes it has 0,empty.
	// and 1 called nil(from reference data types).
	// %q is basically printing a string but with quotes, "", if %v then we won't see any quotes,
	// so if we need a empty sting to be visible use %q
	fmt.Printf("%v %v %v %q\n", i,j,k,l) 
}