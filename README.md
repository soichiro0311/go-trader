# go-trader

## Overview
- trading simulate cli application. `quote-application` publish QuoteInfo(BTC/USD). `order-application` order latest quote.

## Getting Started
```
cd order/resource/docker
docker-compose up -d
cd ../../application
go run main.go <- order-application
cd ../../quote/application
go run main.go <- quote-application
```

## APIs
### quote
- `/quote/latest/BTC/USD` : get latest quote infomation(BTC/USD).This infomation is always displayed console.
- `/quote/history/BTC/USD` : get all quote infomation(BTC/USD).

### order
- `/order/register` : register order. you can make a POST request curl command. Example is 
```
curl -X POST -H "Content-Type: application/json" -d '{"currencycode1":"BTC", "currencycode2":"USD"}' localhost:8090/order/register
```

- `/order/all` : show all your orders. 