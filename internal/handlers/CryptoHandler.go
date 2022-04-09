package handlers

import (
	"go-redis-crypto/internal/infraestructure"
	"go-redis-crypto/internal/services"
	"net/http"

	"github.com/labstack/echo"
)

type CryptoHandler struct {
	Redis         *infraestructure.Redis
	CryptoService *services.CryptoService
}

func (ch *CryptoHandler) Top10(c echo.Context) error {
	crypto, err := ch.CryptoService.Top10()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, struct{ Msg string }{"ERR: " + err.Error()})
	}
	return c.JSON(http.StatusOK, crypto)
}

func (ch *CryptoHandler) FiatCurPrice(c echo.Context) error {
	crypto, err := ch.CryptoService.FiatCurPrice(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, struct{ Msg string }{"ERR: " + err.Error()})
	}
	if crypto.ID == "" {
		return c.JSON(http.StatusOK, struct{ Msg string }{"Given crypto doesn't exist... yet"})
	}
	return c.JSON(http.StatusOK, crypto)
}

/* func (ch *CryptoHandler) Trending24h(c echo.Context) error {
	url := "https://api.coingecko.com/api/v3/search/trending"
	responseByte := helpers.GetData(url)
	crypto := &structs.Trending24h{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, struct{ Msg string }{"Unable to unmarshall response."})
	}
	return c.JSON(http.StatusOK, crypto)
} */
func (ch *CryptoHandler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Msg string }{"API LIVE!"})
}
