package reflectex

import (
	"fmt"
	"reflect"
)

//Settable takes a value and determines if it is settable according to the reflect pacakge SetFloat
func Settable(x1 float32) {
	// var x1 float32 = 5.8  // value to use as an argument in main
	v := reflect.ValueOf(x1)
	//v.SetFloat(2.2)// Will not work becaus v is not settable;
	setable := v.CanSet()
	testSettable(setable)

	t := reflect.ValueOf(&x1) // value is a *float32 ==> x1
	//v.SetFloat(2.2)// Will not work becaus v is not settable;
	setable = t.CanSet()
	testSettable(setable)

	vpElem := t.Elem()
	setable = vpElem.CanSet()
	testSettable(setable)
}

// SimpleSetter takes an integer value and prints the results of changing it to 10
func SimpleSetter(X int) {
	fmt.Println("Oringinally:", X)
	ChangeX(&X)
	fmt.Println("Changed to ", X)
}

//ChangeX takes a pointer to an int and changes it to 10
func ChangeX(X *int) {
	*X = 10
}

func testSettable(s bool) {
	switch s {
	case false:
		fmt.Println("v settable?", s, "v is a pointer to a value of the type")
	case true:
		fmt.Println("v settable?", s, "&v.Elem() is address location of the value")
	}
}
