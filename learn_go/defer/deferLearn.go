// Kind of interesting concept ngl.
// WHY DO WE NEED DEFERS?
// ANS - Cleanup that must run when the function exits, no matter which
// return path you take (normal return, early return, or even a panic).
// You register the cleanup next to the setup, so you do not forget it
// at every return. Typical uses: close a file, unlock a mutex, cancel a
// context. Without defer you would repeat that close/unlock at every
// return, and miss it if you add a new return later.
package main
import "fmt"

// func main() {
// 	// defer queues a function call (not any statement). Arguments are
// 	// evaluated immediately; only the call itself is delayed.
// 	// Deferred calls run when this function is about to return, last
// 	// deferred first (stack / LIFO). Then the function actually returns.
// 	defer fmt.Println("Gozaimaz")

// 	fmt.Println("Konichiva")
// }

func main() {
	fmt.Println("counting")

	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
/*in the output we might thing that both print statement runs first,
first counting gets printed and then done, 
but actually the first print happens immediately and then the for loop runs,
sees the defer so nothing gets printed at first cause main func has not finished yet,
after getting all the elements in the stack | done gets printed and then following 
the LIFO, 10 gets printed first, as we might think why the sequence is in reverse. */
// counting
// done
// 10
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0