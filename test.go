package main

import "fmt"

func main() {
	var name string
	name = "Bob"

	fmt.Println(name) //=> Bob

	var name2 string = "Alice"

	fmt.Println(name2) //=> Alice

	var num, num2 int = 1, 2

	fmt.Println(num, num2) //=> 1 2

	isChecked := true

	fmt.Printf("%v", isChecked) //=> true
}
