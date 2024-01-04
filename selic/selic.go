package main

import (
	"fmt"
	"net/http"
	"os"
)

func SelicPrice() float64 {
	brapiToken := os.Getenv("BRAPI_TOKEN")
	brapiSelicEndpoint := os.Getenv("BRAPI_SELIC_URL")
	selicUrl := fmt.Sprintf("%s/token=%s", brapiSelicEndpoint, brapiToken)
	selic, err := http.Get(selicUrl)
}
