package main

import (
	"gotrain/goMaster/webex/clienttest"
)

func main() {
	// var Link string
	// fmt.Println("Enter a url")
	// fmt.Scan(&Link)
	// restclient.RestClient(Link)
	// //curl -v  -i -X GET http://quotes.rest/qod.json
	//******Slices********
	// s := make(set.Set)
	// var Item string
	// fmt.Println("Enter an Item when finished press enter")
	// fmt.Scan(&Item)
	// //Slices
	// for Item != "." {
	// 	fmt.Println("Enter an Item when finished enter a period")
	// 	fmt.Scan(&Item)
	// 	s[Item] = struct{}{}
	// }
	// fmt.Println(set.GetSetValues(s))
	//******add Items******
	//sliceexamples.SliceExamples()

	///****** Methods and Interfaces******
	//sliceexamples.SliceAppend()
	//******Interfaces******
	// methodsinterfaces.GetSLLNode()
	// methodsinterfaces.SingleLinkedList()
	// printtypes.PrintTypes("Jayne")
	// printtypes.PrintTypes(s)
	// printtypes.PrintTypes(Link)
	// printtypes.PrintTypes(Item)
	//magicstore.MagicStore("Hello", 5)
	//customrand.CustomRandEx()
	//hydra.Hydra("Start the Server")
	// x2 := 5.23477998797
	// reflectex.ReflectEx(x2)
	// type myFloat float32
	// var x3 myFloat = 6.7
	// reflectex.ReflectEx(x3)

	// port := ":2300"
	// reflectex.DialEx(port)
	// // x1 := "Test STring" //This does not work
	// reflectex.ReflectEx(&x1)
	// var x1 float32 = 5.8
	// reflectex.Settable(x1)
	// x2 := 5
	// reflectex.SimpleSetter(x2)

	// myStruct := reflectex.MyStruct{2, "Hello", 2.4}

	// reflectex.MyType(myStruct)
	// reflectex.MyType(&myStruct)

	// p := new(reflectex.PStruct)
	// reflectex.InspectType(p)

	//tcpserver.Ex()
	// udpserver.Ex()
	// protobuf.Ex()
	//webex.Run()
	clienttest.Ex()
}
