package tickers

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Stock struct {
	Ticker string
	Watch  bool
	Bought bool
}

func ReadTickersFromDB(conn string) (stocks []Stock) {

	var ticker string
	var watch bool
	var bought bool
	var stock Stock // slice que recebe a linha para possibilitar append no slice stocks do tipo Stock

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	stocksQuery, err := db.Query("SELECT ticker, watch, bought FROM stocks.tickers")
	if err != nil {
		log.Fatal(err)
	}
	defer stocksQuery.Close()

	for stocksQuery.Next() {
		err := stocksQuery.Scan(&ticker, &watch, &bought)
		if err != nil {
			log.Fatal(err)
		}
		stock.Ticker = ticker
		stock.Watch = watch
		stock.Bought = bought
		stocks = append(stocks, stock)
	}
	return stocks
}
