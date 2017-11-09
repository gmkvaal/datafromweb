package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func GetBody (url string) []byte {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	return body
}


func main() {

	key := "b0ae4566-93c5-4ed0-93da-41ac70df922f"

	url := "https://data.met.no/observations/availableQualityCodes/v0.jsonld?lang=en-US"
	//body := GetBody(url)
	//fmt.Println(string(body))

	req, err := http.NewRequest("GET", url, nil)
	fmt.Println(req, err)
	req.Header.Add("authorization", key)


	body, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(body))

}