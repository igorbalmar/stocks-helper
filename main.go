package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	stocks "stocks-helper/stocks"
)

func main() {

	notifierUrl := fmt.Sprintf("http://%s", os.Args[3])

	resp, err := http.Post(notifierUrl+"/telegram", "application/json", stocks.PrepareStockPayload())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

}
