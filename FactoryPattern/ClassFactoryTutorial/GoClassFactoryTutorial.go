package main

import (
	"fmt"
	"gotrain/goMaster/FactoryPattern/ClassFactoryTutorial/appliances"
	"gotrain/goMaster/FactoryPattern/ClassFactoryTutorial/instruments"
)

func main() {
	//Use the class factory to create an appliance of the requested type

	myAppliance, err := appliances.CreateAppliance()
	//if no errors start the appliance then print it's purpose
	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		//if error encountered, print the error
		fmt.Println(err)
	}

	myInstrument, err := instruments.CreateInstrument()
	//if no errors start the appliance then print it's purpose
	if err == nil {
		myInstrument.Start()
		myInstrument.PlayMusic()

		fmt.Println(myInstrument.GetPurpose())
	} else {
		//if error encountered, print the error
		fmt.Println(err)
	}

}
