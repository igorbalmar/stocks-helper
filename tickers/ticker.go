package tickers

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Stock struct {
	Ticker    string
	Watch     bool
	Bought    bool
	TargetBuy float32
}

func ReadTickersFromDB(conn string) (stocks []Stock) {

	var ticker string
	var watch bool
	var bought bool
	var target_b float32
	var stock Stock // slice que recebe a linha para possibilitar append no slice stocks do tipo Stock

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	stocksQuery, err := db.Query("SELECT ticker, watch, bought, target_b FROM stocks.tickers")
	if err != nil {
		log.Fatal(err)
	}
	defer stocksQuery.Close()

	for stocksQuery.Next() {
		err := stocksQuery.Scan(&ticker, &watch, &bought, &target_b)
		if err != nil {
			log.Fatal(err)
		}
		stock.Ticker = ticker
		stock.Watch = watch
		stock.Bought = bought
		stock.TargetBuy = target_b
		stocks = append(stocks, stock)
	}
	return stocks
}
