package instruments

//import errors to log errors when they occur
import (
	"errors"
	"fmt"
	"gotrain/goMaster/FactoryPattern/ClassFactoryTutorial/appliances"
)

//The main interface used to describe instruments
type Instrument interface {
	PlayMusic() string
	appliances.Appliance
}

//Our instrument types
const (
	GUITAR = iota
	FLUTE
	//Now we support drums
	DRUMS
)

func CreateInstrument() (Instrument, error) {
	//Request the user to enter the instrument type
	fmt.Println("Enter preferred instrument type")
	fmt.Println("0: Guitar ")
	fmt.Println("1: Flute")
	fmt.Println("2: Drums")

	//use fmt.scan to retrieve the user's input
	var myType int
	fmt.Scan(&myType)

	switch myType {
	case GUITAR:
		return new(Guitar), nil
	case FLUTE:
		return new(Flute), nil
		//new case added for microwaves
	case DRUMS:
		return new(Drums), nil

	default:
		return nil, errors.New("Invalid instrument Type")
	}
}
