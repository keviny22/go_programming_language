package greeting

import (
	"fmt"
)

type Printer func(string)

type Salutation struct {
	Name1     string
	Greeting1 string
}

type Renameable interface {
	Rename(newName string)
}

func (salutation *Salutation) Rename(newName string) {
	salutation.Name1 = newName
}

type Salutations []Salutation

//both are strings as input, output is also returns two strings
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey " + name
	return

}

func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {

	//loop through ever itme in the solutation colletion(slice)
	for _, s := range salutations {
		message, alternate := CreateMessage(s.Name1, s.Greeting1)

		if prefix := GetPrefix(s.Name1); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}

}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Mary": "Mrs ",
		"Amy":  "Dr ",
	}

	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	prefixMap["Kevin"] = "Jr "

	//map with key of string and value of string type
	//	var prefixMap map[string]string
	//	prefixMap = make(map[string]string)
	//
	//	prefixMap["Bob"] = "Mr "
	//	prefixMap["Joe"] = "Dr "
	//	prefixMap["Mary"] = "Mrs "
	//	prefixMap["Amy"] = "Dr "

	//return prefixMap[name]
	return "Dude "
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
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
	}
}

//implaments write interface, pass in byte, return integer and error
func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}
