package pizza_api

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestGetStore(t *testing.T) {
	apiClient := Client{http.Client{}}
	address := Address{
		Street:       "Terrace",
		StreetNumber: "464",
		StreetName:   "Pyrite",
		UnitType:     "Home",
		UnitNumber:   "",
		City:         "Colorado Springs",
		Region:       "Colorado",
		PostalCode:   "80905",
	}

	stores, err := apiClient.findStore(address)
	if err != nil {
		log.Fatal("An error occurred:", err)
	}

	for _, v := range stores.Stores {
		fmt.Printf("Address is: %s\n", v.AddressDescription)
		fmt.Printf("Open Now: %t\n\n", v.IsOpen)
	}
}
