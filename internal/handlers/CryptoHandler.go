package handlers

import (
	"encoding/json"
	"go-redis-crypto/internal/helpers"
	"go-redis-crypto/internal/infraestructure"
	"go-redis-crypto/internal/structs"
	"net/http"

	"github.com/labstack/echo"
)

type CryptoHandler struct {
	Redis *infraestructure.Redis
}

func (ch *CryptoHandler) Top10(c echo.Context) error {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"
	responseByte := helpers.GetData(url)
	crypto := &structs.Cryptop10{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, struct{ Msg string }{"Unable to unmarshall response."})
	}
	return c.JSON(http.StatusOK, crypto)
}

func (ch *CryptoHandler) FiatCurPrice(c echo.Context) error {
	id := c.Param("id")
	url := "https://api.coingecko.com/api/v3/coins/" + id + "?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false"
	responseByte := helpers.GetData(url)
	crypto := &structs.FiatCurPrice{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, struct{ Msg string }{"Unable to unmarshall response."})
	}
	if crypto.ID == "" {
		return c.JSON(http.StatusOK, struct{ Msg string }{"Given crypto doesn't exist... yet"})
	}
	return c.JSON(http.StatusOK, crypto)
}

func (ch *CryptoHandler) Trending24h(c echo.Context) error {
	url := "https://api.coingecko.com/api/v3/search/trending"
	responseByte := helpers.GetData(url)
	crypto := &structs.Trending24h{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, struct{ Msg string }{"Unable to unmarshall response."})
	}
	return c.JSON(http.StatusOK, crypto)
}

func (ch *CryptoHandler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Msg string }{"API LIVE!"})
}
