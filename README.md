# ton-api
grpc proxy into tonlib-go

TODO:
 - Fix Dockerfile because ./lib/ was deleted from tonlib-go
 - Add LISTEN_PORT ENV var instead of port hardcode
 - Remove liteclient stuff from config
 - Remove GetBetMethodID from config
 - Move config params to ENV vars

## build
```docker build -t ton-api .```

## run
```docker run --name ton-api --network dice-network -d ton-api```
