package greeting

import (
	"fmt"
)

//type Salutation string

type Salutation struct {
	name1     string
	greeting1 string
}

func main() {
	message := "Hello Go"
	var greeting *string = &message

	fmt.Println(message, greeting)
	fmt.Println(&message, *greeting)
	fmt.Println(message, *greeting)

	//reassign memory address for greeting and message to a different value
	*greeting = "hi"

	var s = Salutation{"Joe", "hello"}
	s.name1 = "Bob"

	Greet(s, CreatePrintFunction("!!!!"), true)
}

//both are strings as input, output is also returns two strings
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey " + name
	return

}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {

	message, alternate := CreateMessage(salutation.name1, salutation.greeting1)

	if isFormal {
		do(message)
	}

	do(alternate)

}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func PrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}
