package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	stocks "stocks-helper/stocks"
	tickers "stocks-helper/tickers"

	"strconv"
)

func main() {
	notifierUrl := fmt.Sprintf("http://%s", os.Getenv("NOTIFIER_ADDR"))
	dbString := os.Getenv("STOCKS_DB_STRING")
	brApiToken := os.Getenv("BRAPI_TOKEN")

	telegramGroup := os.Getenv("TELEGRAM_GROUP_ID")

	if notifierUrl == "" || dbString == "" || brApiToken == "" || telegramGroup == "" {
		log.Fatalf("Verfique se as variáveis de ambiente foram definidas:\nNOTIFIER_ADDR: %s\nSTOCKS_DB_STRING: %s\nBRAPI_TOKEN: %x\nTELEGRAM_GROUP_ID: %s",
			notifierUrl,
			dbString,
			&brApiToken,
			telegramGroup)
	}

	telegramGroupId, err := strconv.ParseInt(telegramGroup, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	stocksList := tickers.ReadTickersFromDB(dbString)
	for _, stock := range stocksList {
		log.Println("Checando", stock.Ticker)
		ret := stocks.GetStockPrice(stock.Ticker, stock.Watch, stock.Bought, brApiToken)

		if ret.Notify {
			payload := stocks.PrepareStockPayload(ret, telegramGroupId)
			resp, err := http.Post(notifierUrl+"/telegram", "application/json", payload)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
		} else {
			log.Println(ret.Status)
		}
	}
}
