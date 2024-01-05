FROM golang:bookworm

WORKDIR /usr/src/stocks-helper

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY stocks ./stocks
COPY main.go .
COPY tickers.env /usr/local/etc/stocks-helper/
COPY check-tickers.sh /usr/local/bin/

RUN go build -v -o /usr/local/bin/stocks-helper ./ && \
    chmod 0004  /usr/local/etc/stocks-helper/tickers.env && \
    chmod 0005 /usr/local/bin/stocks-helper /usr/local/bin/check-tickers.sh 

CMD ["bash", "/usr/local/bin/check-tickers.sh"]
