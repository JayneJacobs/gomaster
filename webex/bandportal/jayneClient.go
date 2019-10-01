package bandportal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Ex tests an API endpoint
func Ex() {
	url := "https://en6yp851rhjlw.x.pipedream.net/"
	resp, err := http.Get(url)
	inspectResponse(resp, err)
	data, err := json.Marshal(struct {
		X int
		Y float32
	}{X: 4, Y: 3.8})

	if err != nil {
		log.Fatal("Error occured while marshaling json", err)
	}
	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	inspectResponse(resp, err)

	client := http.Client{
		Timeout: 3 * time.Second,
	}
	client.Get(url)
	req, err := http.NewRequest(http.MethodPut, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("x-testheader", "learning go header")
	req.Header.Set("User-agent", "Go learning HTTP/1.1")
	resp, err = client.Do(req)
	//ipify returns the ip address
	inspectResponse(resp, err)
	url = "https://api.ipify.org?format=json"

	resp, err = http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	v := struct {
		IP string `json:"ip"` //Name has to be exportable for json
	}{}

	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {

		log.Fatal(err)
	}
	log.Println(v.IP)
}

func inspectResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal("Error while marhalling json", err)
		defer resp.Body.Close()
	}
	b, d := ioutil.ReadAll(resp.Body)
	log.Println(string(b), d)
}
