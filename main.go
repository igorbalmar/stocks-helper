package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	stocks "stocks-helper/stocks"
	tickers "stocks-helper/tickers"

	"strconv"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaHandler)
}

type LambdaResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	notifierUrl := fmt.Sprintf("http://%s", os.Getenv("NOTIFIER_ADDR"))
	dbString := os.Getenv("STOCKS_DB_STRING")
	brApiToken := os.Getenv("BRAPI_TOKEN")

	telegramGroup := os.Getenv("TELEGRAM_GROUP_ID")

	if notifierUrl == "" || dbString == "" || brApiToken == "" || telegramGroup == "" {
		err := fmt.Sprintf("Verfique se as variáveis de ambiente foram definidas:\nNOTIFIER_ADDR: %s\nSTOCKS_DB_STRING: %s\nBRAPI_TOKEN: %x\nTELEGRAM_GROUP_ID: %s",
			notifierUrl,
			dbString,
			&brApiToken,
			telegramGroup)
		log.Fatal(err)
		lambdaReturn := events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "\"Environment not properly set!\"",
		}
		return lambdaReturn, fmt.Errorf(err)
	}

	telegramGroupId, err := strconv.ParseInt(telegramGroup, 10, 64)
	if err != nil {
		log.Fatal(err)
		lambdaReturn := events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "\"Failed to parse TELEGRAM_GROUP_ID\"",
		}
		return lambdaReturn, err
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
				lambdaReturn := events.APIGatewayProxyResponse{
					StatusCode: 500,
					Body:       "\"Failed to send message!\"",
				}
				return lambdaReturn, err
			}
			defer resp.Body.Close()
		} else {
			log.Println(ret.Status)
		}
	}
	lambdaReturn := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Ok\"",
	}
	return lambdaReturn, nil
}
