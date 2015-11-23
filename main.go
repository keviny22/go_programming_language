package main
import ("fmt")

type Dog struct {
	Name string
	Age int
}

type Targets struct {
  target string
}

func main() {
	jackie := Dog{
		Name: "Jackie",
		Age: 19,
	}
	dogs := []Dog{jackie}



	var targets []string
	for k, v := range dogs {
		fmt.Println(k)
		fmt.Println(v)
		targets = append(targets, "arn:aws:iam::")
		fmt.Println(targets)
	}
	//principals := map[string]string{
	//	"AWS": targets,
	//}

	//fmt.Println(principals)
}
