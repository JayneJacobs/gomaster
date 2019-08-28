package reflectex

import (
	"fmt"
	"reflect"
)

// Printer interface  method takes a string and prints the contents
type Printer interface {
	Print(s string)
}

//PStruct is a string to be printed
type PStruct struct {
	s string
}

// Print prints the contents fo pStruct
func (p *PStruct) Print(s string) {
	p.s = s
	fmt.Println(s)
}

func InspectType(obj interface{}) {
	// fmt.Println(obj)
	v := reflect.ValueOf(obj)
	// fmt.Println(v)
	t := v.Type()
	// fmt.Println(t)
	myInterface := reflect.TypeOf((*Printer)(nil)).Elem()
	// fmt.Println(myInterface)
	fmt.Println("obj implemnts Printer?", t.Implements(myInterface))

	if t.Implements(myInterface) {
		printFunc := v.MethodByName("Print")
		args := []reflect.Value{reflect.ValueOf("Printing Hello from a reflection object")}
		printFunc.Call(args)
	}

}
