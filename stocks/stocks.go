package stock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func StockPrice(t string) StockData {
	brApiToken := os.Getenv("BRAPI_TOKEN")
	brApiStockEndpoint := "https://brapi.dev/api/quote"
	ticker := t
	stockUrl := fmt.Sprintf("%s/%s?token=%s", brApiStockEndpoint, ticker, brApiToken)
	stock, err := http.Get(stockUrl)

	if err != nil {
		log.Fatal("Falha no GET da url ", stockUrl)
	}
	defer stock.Body.Close()
	body, err := io.ReadAll(stock.Body)

	if err != nil {
		log.Fatal(err)
		panic("Falha no parse do JSON via io.ReadAll")

	}

	var result StockData
	error := json.Unmarshal(body, &result)
	if error != nil {
		log.Fatal("Não foi possível tratar o JSON via json.Unmarshal:\n", error, "\nTicker: ", t)
	}
	return result
}

func OportunityCheck() StockProps {
	var cotacao StockProps

	for _, rec := range StockPrice(os.Args[1]).Data {
		if rec.RegularMarketPrice <= rec.FiftyTwoWeekLow {
			cotacao.Status = "Oportunidade de Compra!\nValor abaixo da mínima de 52 semanas!\n"
			log.Println("Enviando mensagem de oportunidade de compra para", rec.Symbol)
		} else if rec.RegularMarketPrice <= rec.TwoHundredDayAverage {
			cotacao.Status = "Valor abaixo da média de 200 dias"
			log.Println("Enviando mensagem de oportunidade para", rec.Symbol)
		} else if rec.RegularMarketPrice >= rec.FiftyTwoWeekHigh {
			cotacao.Status = "Oportunidade de Venda!\nValor acima  da máxima de 52 semanas!\n"
			log.Println("Enviando mensagem de oportunidade de venda para", rec.Symbol)
		} else {
			log.Println("Oportunidade não identificada para", rec.Symbol)
			os.Exit(0)
		}
		cotacao.Ticker = rec.Symbol
		cotacao.Price = fmt.Sprintf("%s %.2f", rec.Currency, rec.RegularMarketPrice)
		cotacao.Hora = rec.RegularMarketTime
		cotacao.Low52 = fmt.Sprintf("%.2f", rec.FiftyTwoWeekLow)
		cotacao.High52 = fmt.Sprintf("%.2f", rec.FiftyTwoWeekHigh)
		cotacao.Avg200 = fmt.Sprintf("%.2f", rec.TwoHundredDayAverage)
	}
	return cotacao
}

func PrepareStockPayload() *bytes.Buffer {

	TelegramGroupId, _ := strconv.ParseInt(os.Args[2], 10, 64)

	r := OportunityCheck()

	cotacao := r.Price
	hora := r.Hora.Local().Format("Mon Jan 02 15:04:05 2006")
	ticker := r.Ticker
	low52 := r.Low52
	high52 := r.High52
	avg200 := r.Avg200
	status := r.Status

	content := fmt.Sprintf("%s - %s\n\n%s\n\nHigh 52 weeks: %s\nLow 52 weeks: %s\n200 avg: %s\nLast updated: %s",
		ticker,
		cotacao,
		status,
		high52,
		low52,
		avg200,
		hora)

	message := TelegramPost{
		Text:    content,
		GroupId: TelegramGroupId,
	}
	body, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	payload := bytes.NewBuffer(body)

	return payload
}

type TelegramPost struct {
	Text    string `json:"text"`
	GroupId int64  `json:"group_id"`
}

type StockProps struct {
	Ticker string
	Price  string
	Hora   time.Time
	Low52  string
	High52 string
	Avg200 string
	Status string
}

type StockData struct {
	Data []struct {
		Symbol                            string    `json:"symbol"`
		Currency                          string    `json:"currency"`
		TwoHundredDayAverage              float32   `json:"twoHundredDayAverage"`
		TwoHundredDayAverageChange        float64   `json:"twoHundredDayAverageChange"`
		TwoHundredDayAverageChangePercent float64   `json:"twoHundredDayAverageChangePercent"`
		MarketCap                         int64     `json:"marketCap"`
		ShortName                         string    `json:"shortName"`
		LongName                          string    `json:"longName"`
		RegularMarketChange               float64   `json:"regularMarketChange"`
		RegularMarketChangePercent        float64   `json:"regularMarketChangePercent"`
		RegularMarketTime                 time.Time `json:"regularMarketTime"`
		RegularMarketPrice                float32   `json:"regularMarketPrice"`
		RegularMarketDayHigh              float32   `json:"regularMarketDayHigh"`
		RegularMarketDayRange             string    `json:"regularMarketDayRange"`
		RegularMarketDayLow               float32   `json:"regularMarketDayLow"`
		RegularMarketVolume               int       `json:"regularMarketVolume"`
		RegularMarketPreviousClose        float64   `json:"regularMarketPreviousClose"`
		RegularMarketOpen                 float64   `json:"regularMarketOpen"`
		AverageDailyVolume3Month          int       `json:"averageDailyVolume3Month"`
		AverageDailyVolume10Day           int       `json:"averageDailyVolume10Day"`
		FiftyTwoWeekLowChange             float64   `json:"fiftyTwoWeekLowChange"`
		FiftyTwoWeekLowChangePercent      float64   `json:"fiftyTwoWeekLowChangePercent"`
		FiftyTwoWeekRange                 string    `json:"fiftyTwoWeekRange"`
		FiftyTwoWeekHighChange            float64   `json:"fiftyTwoWeekHighChange"`
		FiftyTwoWeekHighChangePercent     float64   `json:"fiftyTwoWeekHighChangePercent"`
		FiftyTwoWeekLow                   float32   `json:"fiftyTwoWeekLow"`
		FiftyTwoWeekHigh                  float32   `json:"fiftyTwoWeekHigh"`
		PriceEarnings                     float64   `json:"priceEarnings"`
		EarningsPerShare                  float64   `json:"earningsPerShare"`
		Logourl                           string    `json:"logourl"`
		UpdatedAt                         time.Time `json:"updatedAt"`
	} `json:"results"`
	RequestedAt time.Time `json:"requestedAt"`
	Took        string    `json:"took"`
}
