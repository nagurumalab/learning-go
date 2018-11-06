package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, err := strconv.ParseBool("True 1")
	// if err == nil {
	// 	/** displayg the type of the b variable */
	// 	fmt.Printf("Type: %T \n", b)

	// 	/** displaying the string variable into the console */
	// 	fmt.Println("Value:", b)
	// } else {
	fmt.Println("Its an error - ", err)
	fmt.Println("Its a success - ", b)
	//	}
}
