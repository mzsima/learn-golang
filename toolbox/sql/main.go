package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type CandleStick struct {
	open  float64
	close float64
	high  float64
	low   float64
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string

	err = db.QueryRow("select version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)

	res, err := db.Query("select open, high, low, close from candlestick")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	fmt.Printf("+%s+\n", strings.Repeat("-", 74))
	fmt.Printf("| %-16s | %-16s | %-16s | %-16s |\n", "open", "high", "low", "close")
	fmt.Printf("+%s+\n", strings.Repeat("-", 74))

	for res.Next() {
		var candleStick CandleStick
		err = res.Scan(&candleStick.open, &candleStick.high, &candleStick.low, &candleStick.close)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("| %16.5f | %16.5f | %16.5f | %16.5f |\n", candleStick.open, candleStick.high, candleStick.low, candleStick.close)
		fmt.Printf("+%s+\n", strings.Repeat("-", 74))
	}

}
