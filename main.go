package main 
// just like in C, we need atleast 1 library like, stdio,stdlib etc. the main package
// works like the same, which should also contain 1 main function for any executable.
import "fmt"
// functions related to print(string manipulations).
func main() {
	fmt.Println("hello, 世界")
	// case sensitive P(something new that i learnt, for now don't know why it's capital though).
	// Most of the older languages mainly supported ASCII(character encoding), but GO is based on ,
	// UTF-8(falls under the unicode encoding system), it suppports over 128k characters in different
	// languages.
}
// go build <file name whatever you named the file>, it will generate a binary executable for it.
// GO provies a feature of 'cross platform single binary executables', means a go program can , 
// generate binary executable. (meaning a executable that can run in any os that is "SAME" as 
// mine.), without mentioning it, let's say if i want to make a executable for another os(we 
// need to specify in the command.) so a system does not have to download go explicitly.