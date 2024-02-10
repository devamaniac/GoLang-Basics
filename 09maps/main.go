package main

import "fmt"

func main() {
	fmt.Println("Learning Map")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("All the values of map are as following:", languages)
	fmt.Println("The value js is:", languages["js"])

	delete(languages, "js")
	fmt.Println("All the values of map are as following:", languages)
	fmt.Println("The value js is:", languages["js"])

	//loops in golang
	for key, value := range languages {
		fmt.Printf("for key %v the value is %v\n", key, value)
	}

}
