package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"frontendmasters.com/go/crypto/datatypes"
)

const API_URL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {

	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required : %d recieved", len(currency))
	}
	res, err := http.Get(fmt.Sprintf(API_URL, strings.ToUpper(currency)))
	var response CEXResponse

	if err != nil {
		return nil, err // the domain doesn't work
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {

			return nil, err // we couldn't get the body

		}
		err = json.Unmarshal(bodyBytes, &response)

	} else {

		// the body is not STATUS OK
		return nil, fmt.Errorf("status code received : %v", res.StatusCode)

	}

	rate := &datatypes.Rate{Currency: currency, Price: response.Ask}
	return rate, nil

}
