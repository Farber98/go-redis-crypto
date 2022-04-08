package structs

import (
	"time"
)

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
