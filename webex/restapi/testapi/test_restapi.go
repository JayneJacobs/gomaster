package main

import (
	"fmt"
	"gotrain/goMaster/webex/restapi"
)

func main() {
	err := restapi.RunAPI()
	fmt.Println("Error in runAPI,", err)
}
