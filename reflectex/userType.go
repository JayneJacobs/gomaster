package reflectex

import (
	"fmt"
	"io"
	"net"
	"reflect"
)

//MyStruct is a type with an int, string, and float64
type MyStruct struct {
	Field1 int     `alias:"f1" desc:"field number1"`
	Field2 string  `alias:"f2" desc:"field number2"`
	Field3 float64 `alias:"f3" desc:"field number3"`
}

//DialEx c is of type conn which implemnets io.Reader
func DialEx(port string) {
	c, _ := net.Dial("tcp", port)
	var r io.Reader
	r = c //r now stores (value:c , type descriptor: net.Conn)

	//that's why we can also do this:
	if _, ok := r.(io.Writer); ok {
		/*
		   even though r in theory is only of type io.Reader,
		   the underlying value stored also implements the io.writer interface
		*/
		fmt.Printf("We didn't forget there is a writer inside value r: %v \n", r)
	}

	if _, ok := c.(io.Writer); !ok {
		/*
		   even though r in theory is only of type io.Reader,
		   the underlying value stored also implements the io.writer interface
		*/
		fmt.Printf("We didn't forget there is a writer inside value c when it is !ok: %v \n", c)
		return
	}
	fmt.Printf("We didn't forget there is a writer inside value r: %v \n", r)
}

//MyType takes a Struct of the type MyStruct and
//analyzes the field types and values.
func MyType(i interface{}) {
	mysRValue := reflect.ValueOf(i)
	if mysRValue.Kind() != reflect.Ptr {
		return
	}
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return
	}
	mysRValue.Field(0).SetInt(15)
	mysRType := mysRValue.Type() // reflect.TypeOf(i)

	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i) //datatype: StructField
		fieldRValue := mysRValue.Field(i)
		fmt.Printf("FieldName: '%s', field type: '%s' , field value: '%v' \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface())
		fmt.Println("Struct tags, alias: ", fieldRType.Tag.Get("alias"), " description: ", fieldRType.Tag.Get("desc"))
	}
}
