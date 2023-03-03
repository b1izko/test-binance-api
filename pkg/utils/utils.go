package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Pair model
type Pair struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// SendRequestToBinance send request to Binance API and return response for pairs
func SendRequestToBinance(pairs []string) (map[string]string, error) {
	result := make(map[string]string)
	for _, pair := range pairs {
		tokens := strings.Split(pair, "-")
		if len(tokens) != 2 {
			return nil, errors.New("pair parsing error")
		}

		resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + tokens[0] + tokens[1])
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var buffer Pair
		err = json.Unmarshal(body, &buffer)
		if err != nil {
			return nil, err
		}

		if buffer.Symbol == "" {
			var apiErr struct {
				Msg string `json:"msg"`
			}
			err = json.Unmarshal(body, &apiErr)
			if err != nil {
				return nil, err
			}
			return nil, errors.New(apiErr.Msg)
		}

		result[buffer.Symbol] = buffer.Price
	}

	return result, nil
}
