package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learning Slices")

	var friends = []string{"A", "B", "C"}
	fmt.Printf("The type of friends is:%T\n", friends)
	fmt.Println("The values in slice friends are:", friends)

	friends = append(friends, "D")
	fmt.Println(friends)

	friends = append(friends[1:])
	fmt.Println(friends)
	friends = append(friends[1:3])
	fmt.Println(friends)

	//syntax for make

	var highscore = make([]int, 4)
	highscore[0] = 1
	highscore[1] = 3
	highscore[2] = 4
	highscore[3] = 2

	fmt.Println("the Highscore are:", highscore)
	//highscore[4]=5 will give error index out of bounds

	//but if we use append a new memory allocation takes place and we can add
	//any number of values to our slice
	sort.Ints(highscore)
	fmt.Println("the Highscore in sorted form are:", highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	//removing a value from slice based on index

	var subjects = []string{"maths", "science", "english", "hindi", "french"}
	fmt.Println(subjects)
	var index int = 2
	subjects = append(subjects[:index], subjects[index+1:]...)
	fmt.Println("slice after removing english:", subjects)

}
