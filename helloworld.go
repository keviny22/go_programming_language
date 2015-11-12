package main

import (
	"./greeting"
)

func main() {
	//var s = greeting.Salutation{"Bob", "Hello"}

	slice := []greeting.Salutation{
		{"Bob", "hello"},
		{"Joe", "hi"},
		{"Mary", "whatup"},
	}
	greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 6)
}
