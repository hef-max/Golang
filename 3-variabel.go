package main

import "fmt"

func main(){
	// cara 1
	var name string

	name = "Hefri"
	fmt.Println(name)

	name = "John"
	fmt.Println(name)

	// cara 2
	var age = 20
	fmt.Println(age)

	// cara 3
	country := "Indonesia"
	fmt.Println(country)

	country = "Belanda"
	fmt.Println(country)

	var (
		firstName = "Hefri"
		lastName = "Anesti"
	)

	fmt.Println(firstName) 
	fmt.Println(lastName) 

	// mengambil index dan panjang
	fmt.Println("panjang nama", name, "=", len(name)) //"Jhon"
	fmt.Println(name[0])

}