// In this program we are taking input from StandardIO
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza")

	//comma ok

	// error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)

	input2, _ := reader.ReadString('\n')
	fmt.Println("The output of 2nd is:", input2)
	fmt.Printf("The type of variable is:", input2)
}
