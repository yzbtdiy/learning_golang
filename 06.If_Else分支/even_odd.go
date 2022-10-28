package main

import "fmt"

func main() {
	var num1 int
	fmt.Println("Input a number:")
	fmt.Scan(&num1)
	if num1%2 == 0 {
		fmt.Printf("%d is even", num1)
	} else {
		fmt.Printf("%d is odd", num1)
	}

}
