package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/joaovictorcruz/crypin-crawler/internal/crawlers/models"
)

const coinMarketCapUrl = "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing"

func FetchCrypto() (*models.CoinMarketCap, error) {
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(coinMarketCapUrl)
	if err != nil {
		return nil, fmt.Errorf("error in request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var result models.CoinMarketCap
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &result, nil
}

func RunCoinMarketCapCrawler() {
	data, err := FetchCrypto()
	if err != nil {
		log.Printf("Error fetching CoinMarketCap data: %v\n", err)
		return
	}

	log.Println("Successfully fetched data from CoinMarketCap")

	for i, crypto := range data.Data.List {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %s (%s)\n", crypto.ID, crypto.Name, crypto.Domain)
	}
}
