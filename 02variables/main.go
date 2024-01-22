package main

import "fmt"

func main() {
	var username string = "Aman"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	//default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type:%T \n", anothervariable)

	//implicit type
	var website = "https://random.com"
	fmt.Println(website)

	//walrus operator , no var style

	number := 100
	fmt.Println(number)
}
