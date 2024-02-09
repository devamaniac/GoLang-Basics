package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")

	fmt.Println("Rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	stringRating, _ := reader.ReadString('\n')

	fmt.Println("The rating is :", stringRating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(stringRating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to rating: ", numRating+1)
	}
}
