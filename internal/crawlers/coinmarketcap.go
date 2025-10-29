package crawlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const coinMarketCapUrl = "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing"

type CoinMarketCapResponse struct {
Data struct {
		CryptoCurrencyList []struct {
			ID     int     `json:"id"`
			Name   string  `json:"name"`
			Symbol string  `json:"symbol"`
			Slug   string  `json:"slug"`
			Price  float64 `json:"price,omitempty"` 
		} `json:"cryptoCurrencyList"`
	} `json:"data"`
	Status struct {
		ErrorCode int    `json:"error_code"`
		Message   string `json:"error_message"`
	} `json:"status"`
}

func FetchCrypto() (*CoinMarketCapResponse, error) {
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(coinMarketCapUrl)
	if err != nil {
		return nil, fmt.Errorf("erro in request")
	}
	defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var result CoinMarketCapResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &result, nil
}

func RunCoinMarketCapCrawler() {
	data, err := FetchCrypto()
	if err != nil {
		log.Printf("erro fetching coinMarketCap data: %v\n", err)
		return
	}

	log.Println("successfully fetched data from CoinMarketCap")
	for i, crypto := range data.Data.CryptoCurrencyList {
		if i >= 10 {
			break 
		}
		fmt.Printf("%d. %s (%s) — slug: %s\n", crypto.ID, crypto.Name, crypto.Symbol, crypto.Slug)
	}
}