package datafromweb

import (
	"net/http"
	"io/ioutil"
	"os"
	"errors"
	)

// {'sources': 'SN18700', 'elements': 'mean(wind_speed P1D)', 'referencetime': '2010-04-01/2010-06-01/'}



func fetchClientID () (string, error) {
	clientID := os.Getenv("MET_CLIENT_ID")
	if clientID == "" {
		return clientID, errors.New("OS env var MET_CLIENT_ID not set: unable to authenticate")
	}

	return clientID, nil
}

func fetchClientSecret () (string, error) {
	clientID := os.Getenv("MET_CLIENT_SECRET")
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



