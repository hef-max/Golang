package main

import "fmt"

func main(){
	// cara 1
	const firstName string = "Hefri"
	const phi = 3.14

	fmt.Println(firstName)
	fmt.Println(phi)

	// cara 2 
	const(
		lastName string = "Anesti"
		number = 8000
	)
}