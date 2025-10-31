package main

import (
	"fmt"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	rate, err := api.GetRate("BTC") // the rate here is a pointer
	//fmt.Println(rate.Price, err)    // the . here acts as ->
	if err == nil {
		fmt.Printf("The rate for %v is %.2f dollars \n", rate.Currency, rate.Price)
	}

}
