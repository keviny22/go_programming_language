package main

import (
	"fmt"
)

//type Salutation string

type Salutation struct {
	name1     string
	greeting1 string
}

const (
	D = iota

	PI       = 3.14 //types are inferred
	Language = "Go"
	//iota is successvie untyped integer constangs
	A = iota
	B = iota
	C = iota
	// becuase you dont declare, it will assume for previous
	E = iota
	G
	F
)

const (
	//new const will set iota to zero
	H = iota
	I
	J
)

func main() {
	//var message string = "hello go"
	//var message = "hello go"
	message := "Hello Go"
	var greeting *string = &message

	fmt.Println(message, greeting)
	fmt.Println(&message, *greeting)
	fmt.Println(message, *greeting)

	//reassign memory address for greeting and message to a different value
	*greeting = "hi"
	fmt.Println(message, *greeting)

	//var message1 Salutation = "Hello World"
	//fmt.Println(message1)

	var s = Salutation{"Joe", "hello"}
	s.name1 = "Bob"
	fmt.Println(s.name1)
	fmt.Println(s.greeting1)

	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A, B, C, D, E, F, G, H, I, J)

	//var a, b, c int = 1, 2, 3
	//var a, b, c  = 1, 2, 3
	//a, b, c := 1, false, "foo"
	//a = 1
	//message = "hello go world"
	//fmt.Println(message, a, b, c)

	Greet(s, CreatePrintFunction("!!!!"), true)
}

//both are strings as input, output is also returns two strings
func CreateMessage(name, greeting string) (message string, alternate string) {
	//fmt.Println(len(greeting))
	message = greeting + " " + name
	alternate = "Hey " + name
	return

	//fmt.Println(name, greeting)
}

type Printer func(string)

//func Greet(salutation Salutation, do func(string)) {
func Greet(salutation Salutation, do Printer, isFormal bool) {
	//fmt.Println(salutation.name1)
	//fmt.Println(salutation.greeting1)
	//fmt.Println(CreateMessage(salutation.name1, salutation.greeting1))
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
