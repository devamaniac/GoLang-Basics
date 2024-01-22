package main

import "fmt"

func main() {
	var username string ="Aman"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool =true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallValue uint8 =256
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n",smallValue)
}
