package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Learning if and else ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")
	count, _ := reader.ReadString('\n')
	fmt.Println("the value in count is:", count)
	fmt.Printf("the type of count is:%T\n", count)
	countInInt, _ := strconv.ParseUint(strings.TrimSpace(count), 10, 64)
	fmt.Printf("the type of countInInt is:%T\n", countInInt)
	fmt.Println("the value in countInInt is:", countInInt)

	if countInInt%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	if num := 5; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

}
