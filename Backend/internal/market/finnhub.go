package market

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Quote struct {
	CurrentPrice  float64 `json:"current_price"`
	Change        float64 `json:"change"`
	PercentChange float64 `json:"percent_change"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Open          float64 `json:"open"`
	PreviousClose float64 `json:"previous_close"`
}
type FinnhubClient struct {
	apiKey string
	client *http.Client
}

func NewFinnhubClient(apiKey string) *FinnhubClient {
	return &FinnhubClient{apiKey: apiKey, client: &http.Client{Timeout: 10 * time.Second}}
}
func (c *FinnhubClient) Quote(ctx context.Context, symbol string) (Quote, error) {
	if c.apiKey == "" {
		return Quote{}, errors.New("FINNHUB_API_KEY is not configured")
	}
	symbol = strings.ToUpper(strings.TrimSpace(symbol))
	if symbol == "" {
		return Quote{}, errors.New("symbol is required")
	}
	endpoint := "https://finnhub.io/api/v1/quote?symbol=" + url.QueryEscape(symbol) + "&token=" + url.QueryEscape(c.apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return Quote{}, fmt.Errorf("create quote request: %w", err)
	}
	response, err := c.client.Do(req)
	if err != nil {
		return Quote{}, fmt.Errorf("request quote: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Quote{}, fmt.Errorf("market provider returned %s", response.Status)
	}
	var payload struct {
		CurrentPrice  float64 `json:"c"`
		Change        float64 `json:"d"`
		PercentChange float64 `json:"dp"`
		High          float64 `json:"h"`
		Low           float64 `json:"l"`
		Open          float64 `json:"o"`
		PreviousClose float64 `json:"pc"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return Quote{}, fmt.Errorf("decode quote: %w", err)
	}
	return Quote{payload.CurrentPrice, payload.Change, payload.PercentChange, payload.High, payload.Low, payload.Open, payload.PreviousClose}, nil
}
