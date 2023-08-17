package main

import "fmt"

func main(){
	type noKTP string
	type married bool

	var nomorktp noKTP = "3212222222"
	var status married = true

	fmt.Println(nomorktp)
	fmt.Println(status)
}