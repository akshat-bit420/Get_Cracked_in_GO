// a slice has always a backing array, it mainly consist of 3 things.
// 1. length , 2. capacity , 3. pointers(to the first element/index where slice index is pointing to.)
/*
*/
package main
import "fmt"

// func main() {
// 	q := []int{2,3,5,7,9,11}
// 	// it's a slice not an array, array will always have a number in them	
// 	fmt.Println(q)

// 	r := []bool{true,false,true,true,false,false}
// 	fmt.Println(r)

// 	s := []struct {
// 		i int
// 		b bool
// 	}{
// 		{2,true},
// 		{3,false},
// 		{5,true},
// 		{7,true},
// 		{11,false},
// 		{13,false},
// 	}
// 	fmt.Println(s)

// }


// func main() {
// 	// primes := [6]int{2,3,5,7,9,11}
// 	names := [4]string{
// 		"ram",
// 		"sharan",
// 		"kulfi",
// 		"hoe",
// 	}
// 	fmt.Println(names)
// 	// we have an array which works as a constant, but we will change it now
// 	a := names[0:2]
// 	b := names[1:3]

// 	fmt.Println(a, b)
// 	b[0] = "XXX"
// 	// here we just change b, so in that manner we should see b being updated
// 	// but the value got reflected to all, the backing array(names), a and b.
// 	fmt.Println(a, b)
// 	fmt.Println(names)
// 	// var s []int = primes[1:3]
// 	// // [low:high] is half-open: low is included, high is excluded.
// 	// // primes[1:3] takes indices 1 and 2 → [3, 5].
// 	// // 5 is at index 2. Index 3 (the 7) is the first value left out.
// 	// fmt.Println(s)
// }

func main() {
	num :=[]int{2,3,5,7,11,13}

	num = num[1:4]
	// here the slice will begin from the second index which is 1 and until it,
	// reaches to the third index, cause 4 is seen as 4 - 1, not 4 exactly.
	// which means num = [3,5,7]
	fmt.Println(num)

	num = num[:2]
	// here the slices is from index 0 until it reaches index 2. 
	// so num = [3,5]
	fmt.Println(num)

	num = num[1:]
	// here the element after index 1 is taken and after 1.
	// so num = [5]
	fmt.Println(num)
}
/* like this,
[3 5 7]
[3 5]
[5]
*/