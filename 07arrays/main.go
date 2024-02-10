package main

import "fmt"

func main() {
	fmt.Println("Learning Arrays")

	var friends [3]string

	friends[0] = "a"
	friends[2] = "c"

	fmt.Println("the friends are:", friends)
	fmt.Println("the size of array of friends is:", len(friends))
	/*len function will always print the size which is declare and not the
	size of the actual values stored in the array
	*/

	// another way of declaring and initializing array is

	var closeFriends = [3]string{"x", "y", "z"}
	fmt.Println("the friends are:", closeFriends)
	fmt.Println("the size of array of friends is:", len(closeFriends))
}
