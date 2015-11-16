package main

import (
	"./greeting"
	"fmt"
)

func RenameToFrog(r greeting.Renameable) {
	r.Rename("Frog")
}

func main() {
	//var s = greeting.Salutation{"Bob", "Hello"}

	//var s []int = make([]int, 3, 5)
	//var s []int
	//s = make([]int, 3)
	//s[0] = 1
	//s[1] = 10
	//s[2] = 500

	//s := []int { 1, 10, 500, 25 }

	//fmt.Println(s)

	salutations := greeting.Salutations{
		{"Bob", "hello"},
		{"Joe", "hi"},
		{"Mary", "whatup"},
		{"Kevin", "whatup"},
	}
	salutations[0].Rename("john")

	//& means to dereference, get address of the value so it will act as a pointer
	//RenameToFrog(&salutations[0])

	fmt.Fprintf(&salutations[0], "the cound is %d", 10)

	salutations.Greet(greeting.CreatePrintFunction("?"), true, 6)
}
