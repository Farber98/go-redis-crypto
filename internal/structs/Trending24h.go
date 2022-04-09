package structs

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
