// these are of 5 types :- functions,maps,pointers,slices and channels.
package main
import "fmt"
func main() {
	i,j := 4,2701
	p := &i // pointer to i

	fmt.Println(*p) // read i through the pointer
	*p = 21 // see i through the pointer
	fmt.Println(i) // see new value of i
	p = &j // pointer to j
    *p = *p / 37 // divide j through the pointer
	fmt.Println(*p) // new value of j
}
/* pointer stores the address of an variable,
& is used to get to the address and *(dereferencing) is used to find
the value at that particular address
*/
