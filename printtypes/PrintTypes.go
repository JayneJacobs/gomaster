package printtypes

import "fmt"

// PrintTypes takes any type and Prints it
func PrintTypes(i interface{}) {
	printType("text")
	printType(3)
	printType(4.005)
	printType(i)
}

func printType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("This is a string type", i)
	case int:
		fmt.Println("this is an int type", i)
	case float32:
		fmt.Println("this is an float type", i)
	default:
		fmt.Printf("This is a %t %s %v", i, i, i)
	}
}
