package main

import "fmt"

func main() {
	// no inheritence in golang; No super or parent

	amanDetails := User{"aman", "manaiac@go.dev", true, 24}
	fmt.Println(amanDetails)

	//+v gives all deatils along with parameter
	fmt.Printf("Aman details are as following: %+v\n", amanDetails)

}

type User struct {
	//these can be accessed by anybody so we keeping first letter as capital
	Name   string
	Email  string
	Status bool
	Age    int
}
