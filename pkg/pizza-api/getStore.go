package pizza_api

import (
	"encoding/json"
	"fmt"
	"github.com/tony-mw/pizza-api-client/pkg/pizza-api/responseTypes"
	"io/ioutil"
	"log"
	"net/http"
)

var storeResponse responseTypes.StoreResponse

type Address struct {
	Street       string `json:"Street,omitempty"`
	StreetNumber string `json:"StreetNumber,omitempty"`
	StreetName   string `json:"StreetName,omitempty"`
	UnitType     string `json:"UnitType,omitempty"`
	UnitNumber   string `json:"UnitNumber,omitempty"`
	City         string `json:"City,omitempty"`
	Region       string `json:"Region,omitempty"`
	PostalCode   string `json:"PostalCode,omitempty"`
}

func (c Client) findStore(a Address) (*responseTypes.StoreResponse, error) {

	resp, err := http.Get(fmt.Sprintf("%s?s=%s&c=%s&s=Delivery", store_locator, a.Street, a.PostalCode))
	if err != nil {
		log.Fatal("Error making api call: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading the response: ", err)
	}

	err = json.Unmarshal(body, &storeResponse)
	if err != nil {
		log.Fatal("Failed to unmarshal into struct: ", err)
	}

	return &storeResponse, nil
}
