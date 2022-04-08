package structs

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
