package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

func getCurrencyData(currency string) {

	rate, err := api.GetRate(currency) // the rate here is a pointer
	//fmt.Println(rate.Price, err)    // the . here acts as ->

	if err == nil {
		fmt.Printf("The rate for %v is %.2f dollars \n", rate.Currency, rate.Price)
	}

}
func main() {

	var wg sync.WaitGroup // wg == WaitGroup
	currencies := []string{"BTC", "ETH", "BCH"}

	for _, currency := range currencies {
		wg.Go(
			func() {

				getCurrencyData(currency)

			})
	}

	wg.Wait()

}
