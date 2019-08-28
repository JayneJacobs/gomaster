package restclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RestClient(link string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
