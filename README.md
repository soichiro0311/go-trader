# go-trader

## Overview
- trading simulate cli application. `board-application` publish boardInfo(BTC/USD). `order-application` order latest board.

## Getting Started
```
cd order/resource/docker
docker-compose up -d
cd ../../application
go run main.go <- order-application
cd ../../board/application
go run main.go <- board-application
```

## APIs
### board
- `/board/latest/BTC/USD` : get latest board infomation(BTC/USD).This infomation is always displayed console.
- `/board/history/BTC/USD` : get all board infomation(BTC/USD).

### order
- `/order/register` : register order. you can make a POST request curl command. Example is 
```
curl -X POST -H "Content-Type: application/json" -d '{"currencycode1":"BTC", "currencycode2":"USD"}' localhost:8090/order/register
```

- `/order/all` : show all your orders. 