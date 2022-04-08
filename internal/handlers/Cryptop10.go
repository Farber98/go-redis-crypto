package handlers

import (
	"encoding/json"
	"go-redis-crypto/internal/helpers"
	"go-redis-crypto/internal/structs"
	"net/http"

	"github.com/labstack/echo"
)

func Top10(c echo.Context) error {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"
	responseByte := helpers.GetData(url)
	crypto := &structs.Cryptop10{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, struct{ Msg string }{"Unable to unmarshall response."})
	}
	return c.JSON(http.StatusOK, crypto)
}
