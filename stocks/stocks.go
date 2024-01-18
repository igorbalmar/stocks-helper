package stock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetStockPrice(t string, w bool, b bool, target_b float32, token string) (cotacao StockProps) { //t is ticker, b stands for bought and w for watch
	brApiStockEndpoint := "https://brapi.dev/api/quote"
	ticker := t
	stockUrl := fmt.Sprintf("%s/%s?token=%s", brApiStockEndpoint, ticker, token)
	stock, err := http.Get(stockUrl)

	if err != nil {
		log.Fatal("Falha no GET da url ", stockUrl)
	}

	if stock.StatusCode != 200 {
		log.Fatalf("Erro na chamada do enpoint %s.\nVerifique se o token está correto ou se houve alterações no retorno", brApiStockEndpoint)
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
		log.Fatalf("Não foi possível tratar o JSON via json.Unmarshal para %s\n%s", t, error)
	}

	for _, rec := range result.Data {
		if (w || b) && rec.FiftyTwoWeekLow != 0.00 && rec.RegularMarketPrice <= rec.FiftyTwoWeekLow { //b stands for bought and w for watch
			cotacao.Status = rec.Symbol + " abaixo da mínima em 52 semanas!\n"
			cotacao.Notify = true
			log.Println("Preparando payload para  ", rec.Symbol)
		} else if b && rec.FiftyTwoWeekHigh != 0.00 && rec.RegularMarketPrice >= rec.FiftyTwoWeekHigh {
			cotacao.Status = rec.Symbol + " acima da máxima em 52 semanas\n"
			cotacao.Notify = true
			log.Println("Preparando payload para  ", rec.Symbol)
		} else if rec.RegularMarketPrice != 0.00 && target_b != 0 && rec.RegularMarketPrice <= target_b {
			cotacao.Status = rec.Symbol + " está no valor definido para compra!"
			cotacao.Notify = true
		} else if rec.RegularMarketPrice == 0.00 {
			cotacao.Status = "Dados insuficientes para " + rec.Symbol
			cotacao.Notify = false
			log.Printf("%s\nCotação: %.2f\n52 High: %.2f\n52 Low: %.2f\n%s UTC", rec.Symbol, rec.RegularMarketPrice, rec.FiftyTwoWeekHigh, rec.FiftyTwoWeekLow, rec.UpdatedAt)
		} else {
			cotacao.Notify = false
			log.Println("Oportunidade não identificada para", rec.Symbol)
		}
		cotacao.Ticker = rec.Symbol
		cotacao.Price = fmt.Sprintf("%s %.2f", rec.Currency, rec.RegularMarketPrice)
		//cotacao.Hora = rec.RegularMarketTime
		cotacao.Low52 = fmt.Sprintf("%.2f", rec.FiftyTwoWeekLow)
		cotacao.High52 = fmt.Sprintf("%.2f", rec.FiftyTwoWeekHigh)
	}
	return cotacao
}

func PrepareStockPayload(r StockProps, g int64) *bytes.Buffer {

	cotacao := r.Price
	ticker := r.Ticker
	low52 := r.Low52
	high52 := r.High52
	status := r.Status
	avg200 := r.Avg200
	hora := r.UpdatedAt
	//hora := r.Hora.Local().Format("Mon Jan 02 15:04:05 2006")
	content := fmt.Sprintf("%s - %s\n\n%s\n\nHigh 52: %s\nLow 52: %s\nAvg 200: %s\n%s UTC",
		ticker,
		cotacao,
		status,
		high52,
		low52,
		avg200,
		hora)

	message := TelegramPost{
		Text:    content,
		GroupId: g,
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
	Ticker    string
	Price     string
	UpdatedAt string
	Low52     string
	High52    string
	Avg200    string
	Status    string
	Notify    bool
}

type StockData struct {
	Data []struct {
		Symbol                            string  `json:"symbol"`
		Currency                          string  `json:"currency"`
		TwoHundredDayAverage              float32 `json:"twoHundredDayAverage"`
		TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
		TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
		MarketCap                         int64   `json:"marketCap"`
		ShortName                         string  `json:"shortName"`
		LongName                          string  `json:"longName"`
		RegularMarketChange               float64 `json:"regularMarketChange"`
		RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
		RegularMarketTime                 string  `json:"regularMarketTime"`
		RegularMarketPrice                float32 `json:"regularMarketPrice"`
		RegularMarketDayHigh              float32 `json:"regularMarketDayHigh"`
		RegularMarketDayRange             string  `json:"regularMarketDayRange"`
		RegularMarketDayLow               float32 `json:"regularMarketDayLow"`
		RegularMarketVolume               int     `json:"regularMarketVolume"`
		RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
		RegularMarketOpen                 float64 `json:"regularMarketOpen"`
		AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
		AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
		FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
		FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
		FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
		FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
		FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
		FiftyTwoWeekLow                   float32 `json:"fiftyTwoWeekLow"`
		FiftyTwoWeekHigh                  float32 `json:"fiftyTwoWeekHigh"`
		PriceEarnings                     float64 `json:"priceEarnings"`
		EarningsPerShare                  float64 `json:"earningsPerShare"`
		Logourl                           string  `json:"logourl"`
		UpdatedAt                         string  `json:"updatedAt"`
	} `json:"results"`
	RequestedAt time.Time `json:"requestedAt"`
	Took        string    `json:"took"`
}
