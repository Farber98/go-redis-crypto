# go-redis-crypto

Playing with Redis to reduce the amount of fetches done to Coingecko API.

### 1. Run Redis with Docker
```
$ docker pull redis
$ docker run --name redis-instance -p 6379:6379 -d redis
```
