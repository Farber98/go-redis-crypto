package handlers

import (
	"encoding/json"
	"go-redis-crypto/internal/helpers"
	"go-redis-crypto/internal/structs"
	"net/http"

	"github.com/labstack/echo"
)

func FiatCurPrice(c echo.Context) error {
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
