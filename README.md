# ton-api
grpc proxy into tonlib-go

TODO:
 + Fix Dockerfile to use libtonlibjson.so from tonlib-go module 
 + Add LISTEN_PORT ENV var instead of port hardcode
 + Remove liteclient stuff from config
 + Remove GetBetMethodID from config
 + Move config params to ENV vars

## build
```docker build -t ton-api .```

## run (development)
```docker run -e CONTRACT_ADDR='kQBgMySA9l-X185xzrEm20M0INDbi0AavYOog9p0133yq-ZJ' --name ton-api ton-api```

## ENV VARS
TONLIB_CFG_PATH - path to tonlib.config.json. Default value is: `/usr/local/bin/app/tonlib.config.json.example`, the path to embedded version of config.
LISTEN_PORT - default value is `5400`.
CONTRACT_ADDR - required variable, no default value.

