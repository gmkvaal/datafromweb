package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"errors"
	"encoding/json"
	)

// {'sources': 'SN18700', 'elements': 'mean(wind_speed P1D)', 'referencetime': '2010-04-01/2010-06-01/'}

const baseURL = "https://data.met.no"

type Response struct {
	Context string `json:"@context"`
	Type string `json:"@type"`
	ApiVersion string `json:"apiVersion"`
	License string `json:"license"`
	CreatedAt string `json:"license"`
	QueryTime float64 `json:"queryTime"`
	ItemsPerPage int `json:"itemsPerPage"`
	Offset int `json:"offset"`
	TotalItemCount int `json:"totalItemCount"`
	CurrentLink int `json:"currentLink"`
	Data []Data `json:"data"`
}

type Data struct {
	SourceID string `json:"sourceId"`
	ReferenceTime string `json:"referenceTime"`
	Observations []Observations `json:"observations"`
}

type Observations struct {
	ElementID string `json:"elementId"`
	Value float64 `json:"value"`
	PerformanceCategory string `json:"performanceCategory"`
	ExposureCategory string `json:"exposureCategory"`
	qualityCode string `json:"qualityCode"`
}

func fetchClientID () (string, error) {
	clientID := os.Getenv("MET_CLIENT_ID")
	if clientID == "" {
		return clientID, errors.New("OS env var MET_CLIENT_ID not set: unable to authenticate")
	}

	return clientID, nil
}

func newRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return req, err
}

func authenticate(req *http.Request) (*http.Response, error) {
	clientID, err := fetchClientID()
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req.SetBasicAuth(clientID, "")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("unable to authenticate")
	}

	return resp, err
}

func getBody(endpoint string) ([]byte, error) {
	req, err := newRequest(endpoint)
	if err != nil {
		return nil, err
	}

	resp, err := authenticate(req)
	if err != nil {
		return nil, err
	}

	//defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

/*
func getDataFromBody(endpoint string) ([]Data, error) {
	body, err := getBody(endpoint)
	if err != nil {
		return []Data{}, err
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Data{}, err
	}

	fmt.Print(response.TotalItemCount)

	return response.Data, nil

}
*/

//func createEndpoint(baseURL string) (string, error) {

//	endpoint := strings.join()
//}


func main() {


	endpoint := "https://data.met.no/tests/secureHello"
	//endpoint = "https://data.met.no/observations/v0.jsonld"
	endpoint = "https://data.met.no/observations/v0.jsonld?sources=SN18700&referencetime=2010-04-01%2F2010-06-01&elements=mean(wind_speed%20P1D)"
	endpoint = "https://data.met.no/observations/v0.jsonld?elements=sum%28precipitation_amount+P1M%29&referencetime=2016-01-01T00%3A00%3A00.000Z%2F2016-12-21T00%3A00%3A00.000Z&sources=SN18700"


	body, err := getBody(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)

	fmt.Printf("%s\n", body)
	fmt.Println(response.Data)

	for _, d := range response.Data {
		obs := d.Observations[0]
		fmt.Println(obs.Value)
	}

}