package pizza_api

import (
	"fmt"
	"net/http"
	"log"
	)

func (c Client) getMenu(storeid string) (string, err) {

	resp, err := http.Get(fmt.Sprintf("%s%s%s", menu["prefix"], storeid, menu["suffix"]))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(resp), nil

}
