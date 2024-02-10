package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	presentTime := time.Now()
	fmt.Println("The date is:", presentTime.Format("01-02-2006"))
	fmt.Println("The date and time is:", presentTime.Format("01-02-2006 15:04:05"))
	fmt.Println("The date and day is:", presentTime.Format("01-02-2006 Monday"))
}
