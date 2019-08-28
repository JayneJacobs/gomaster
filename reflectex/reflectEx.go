package reflectex

import (
	"fmt"
	"reflect"
)

//ReflectEx takes an integer or float
func ReflectEx(i interface{}) {

	//...
	//...

	v := reflect.ValueOf(i)
	fmt.Printf("This is the initial type: %s\n", v.Type())

	inspectIfTypeFloat(i)
	fmt.Println(v.Kind() == reflect.Float64)
}

func inspectIfTypeFloat(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println("initial type: ", v.Type()) // eq to reflect.TypeOf(i)
	fmt.Println("Float Value:", v.Float())
	x2 := v.Float()
	v2 := reflect.ValueOf(x2)
	fmt.Println("is v type float64?", v.Kind() == reflect.Float64)
	fmt.Println("Type of v2", v2.Type())
	interfaceValue := v.Interface()
	switch t := interfaceValue.(type) {
		case float32:
			fmt.Printf("This is the original type %T value %v ", t, t)
		case float64:
			fmt.Printf("This is the converted type %T value %v ", t, t)
	}
}
