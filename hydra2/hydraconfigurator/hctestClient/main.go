package main

import (
	"fmt"
	"gotrain/goMaster/hydra/hydraconfigurator"
)

//Confs Struct is populated with the parsed(Marshalled) file values
type Confs struct {
	TS      string  `name:"testString"`
	TB      bool    `name:"testBool"`
	TF      float64 `name:"testFloat"`
	TestInt int
}

func main() {

	configstruct := new(Confs)
	hydraconfigurator.GetConfiguration(hydraconfigurator.CUSTOM, configstruct, "configfile.conf")
	fmt.Println(*configstruct)
	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))
	fmt.Println(5 * configstruct.TestInt)
	fmt.Println(configstruct.TS)
}
