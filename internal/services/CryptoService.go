package services

import (
	"context"
	"encoding/json"
	"go-redis-crypto/internal/helpers"
	"go-redis-crypto/internal/structs"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type CryptoService struct {
	RedisConn *redis.Client
}

func (cs *CryptoService) Top10() (crytop10 *structs.Cryptop10, err error) {
	crypto := &structs.Cryptop10{}
	val, err := cs.RedisConn.Get(context.Background(), "top10").Result()
	if err == redis.Nil {
		/* Non cached */
		url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"
		responseByte := helpers.GetData(url)
		err := json.Unmarshal(responseByte, &crypto)
		if err != nil {
			return nil, err
		}
		b, err := json.Marshal(crypto)
		if err != nil {
			return nil, err
		}
		val := string(b)
		_, err = cs.RedisConn.Set(context.Background(), "top10", val, 30*time.Second).Result()
		log.Println("API FETCH")
		return crypto, err
	} else if err != nil {
		return nil, err
	} else {
		/* Cached */
		err = json.Unmarshal([]byte(val), &crypto)
		if err != nil {
			return nil, err
		}
		log.Println("CACHE HIT")
	}
	return crypto, nil
}
