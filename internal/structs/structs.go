package structs

import "time"

type Cryptop10 []struct {
	MarketCapRank     int       `json:"market_cap_rank,omitempty"`
	ID                string    `json:"id,omitempty"`
	Symbol            string    `json:"symbol,omitempty"`
	Name              string    `json:"name,omitempty"`
	CurrentPrice      float32   `json:"current_price,omitempty"`
	DayHigh           float32   `json:"high_24h,omitempty"`
	DayLow            float32   `json:"low_24h,omitempty"`
	CirculatingSupply float32   `json:"circulating_supply,omitempty"`
	TotalSupply       float32   `json:"total_supply,omitempty"`
	Ath               float32   `json:"ath,omitempty"`
	AthDate           time.Time `json:"ath_date,omitempty"`
	Atl               float32   `json:"atl,omitempty"`
	AtlDate           time.Time `json:"atl_date,omitempty"`
	LastUpdate        time.Time `json:"last_updated,omitempty"`
}

type Trending24h struct {
	Coins []struct {
		Item struct {
			ID            string  `json:"id"`
			CoinID        int     `json:"coin_id"`
			Name          string  `json:"name"`
			Symbol        string  `json:"symbol"`
			MarketCapRank int     `json:"market_cap_rank"`
			PriceBtc      float64 `json:"price_btc"`
			Score         int     `json:"score"`
		} `json:"item"`
	} `json:"coins"`
}

type FiatCurPrice struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	MarketData struct {
		CurrentPrice struct {
			Aud  float32 `json:"aud"`
			Cad  float32 `json:"cad"`
			Chf  float32 `json:"chf"`
			Cny  float32 `json:"cny"`
			Eur  float32 `json:"eur"`
			Gbp  float32 `json:"gbp"`
			Hkd  float32 `json:"hkd"`
			Jpy  float32 `json:"jpy"`
			Nzd  float32 `json:"nzd"`
			Usd  float32 `json:"usd"`
			Sats float32 `json:"sats"`
		} `json:"current_price"`
	} `json:"market_data"`
}
