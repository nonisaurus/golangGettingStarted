package main

import "fmt"

func main() {
	fmt.Println("Hello")

	fmt.Println("What's your name?")
	var yourName string
	fmt.Scan(&yourName)
	fmt.Printf("Hi there %v! \n", yourName)

	template := "I wish I had a %v."
	pet := "puppy"
	var wish string
	wish = fmt.Sprintf(template, pet)
	fmt.Println(wish)
}
