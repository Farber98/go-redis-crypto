# go-redis-crypto

Using Redis as distributed cache to reduce the amount of fetches and response times to Coingecko API.

## How to play:
#### 1. Clone repo
#### 2. Run where docker-compose.yaml is situated. This will run Go API [:1525] and Redis [:6379] 
```
$ docker-compose up --build
```
#### 3. Run swagger.yaml to understand the endpoints. 
```
[VSCODE]: ALT+SHIFT+P inside .yaml || Right click + Preview Swagger on .yaml
```
#### 4. Compare API FETCH vs CACHE HIT
![CacheVsFetch](/img/cachevsfetch.png)
