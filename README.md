# go-redis-crypto

Using Redis to reduce the amount of fetches done to Coingecko API.

## How to play:
### 1. Clone
<!--### 2. Run Swagger to understand the endpoints. 
```
ALT+SHIFT+P inside .yaml || Right click + Preview Swagger on .yaml
``` -->
### 2. Run where docker-compose.yml is situated 
```
$ docker-compose up --build
```
### 3. Compare API FETCH vs CACHE HIT
![CacheVsFetch](/img/cachevsfetch.png)
