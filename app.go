package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	clientKey := os.Getenv("MET_CLIENT_ID")
	url := "https://data.met.no/tests/secureHello"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(clientKey, "")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(body))

}