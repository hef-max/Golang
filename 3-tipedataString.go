package main

import "fmt"

func main(){
	var name string
	name = "Hefri"
	fmt.Println(name)

	name = "John"
	fmt.Println(name)

	var kota = "Indramayu"
	fmt.Println(kota)

	country := "Indonesia"
	fmt.Println(country)

	country = "Belanda"
	fmt.Println(country)

	// mengambil index dan panjang
	fmt.Println(len(name))
	fmt.Println(name[0])
}