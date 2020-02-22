FROM ubuntu:18.04 as builder
RUN apt-get update \
 && apt-get install -y curl cmake clang ninja-build pkg-config libssl-dev zlib1g-dev libreadline-dev libmicrohttpd-dev gperf git wget \
 && rm -rf /var/lib/apt/lists/* \
 && curl -LSs https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz -o go.tar.gz \
 && tar -xf go.tar.gz \
 && rm -v go.tar.gz \
 && mv go /usr/local/
ENV PATH=${PATH}:/usr/local/go/bin
WORKDIR /go/src/build
ADD . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o ton-api ./cmd

FROM frolvlad/alpine-glibc
COPY --from=builder /go/src/build/ton-api /usr/local/bin/app/
COPY --from=builder /go/src/build/config.yml /usr/local/bin/app/
COPY --from=builder /go/src/build/tonlib.config.json.example /usr/local/bin/app/
COPY --from=builder /go/src/build/lib/linux /usr/lib/
RUN apk add --no-cache libstdc++
WORKDIR /usr/local/bin/app/
RUN mkdir test.keys
CMD ["./ton-api"]