package greeting

import (
	"fmt"
)

type Printer func(string)

type Salutation struct {
	Name1     string
	Greeting1 string
}

//both are strings as input, output is also returns two strings
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey " + name
	return

}

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {

	//loop through ever itme in the solutation colletion(slice)
	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name1, s.Greeting1)

		if prefix := GetPrefix(s.Name1); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}

}

func GetPrefix(name string) (prefix string) {

	//map with key of string and value of string type
	var prefixMap map[string]string
	prefixMap = make(map[string]string)

	prefixMap["Bob"] = "Mr "
	prefixMap["Joe"] = "Dr "
	prefixMap["Mary"] = "Mrs "
	prefixMap["Amy"] = "Dr "

	switch {
	case name == "Bob":
		prefix = "Mr "
	case name == "Joe", name == "Amy", len(name) == 10:
		prefix = "Dr "
	case name == "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "
	}
	return
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

//create a empty interface of any type, return the type
func TypeSwitchTest(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
		_ = t

	}
}
