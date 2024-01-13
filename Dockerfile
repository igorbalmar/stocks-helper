FROM golang:bookworm

WORKDIR /usr/src/stocks-helper

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY stocks ./stocks
COPY main.go .
COPY tickers.env /usr/local/etc/stocks-helper/
COPY check-tickers.sh /usr/local/bin/

RUN echo "TZ='America/Sao_Paulo'; export TZ" >> ~/.profile && \
    . ~/.profile && \
    go build -v -o /usr/local/bin/stocks-helper

CMD ["/usr/local/bin/stocks-helper"]
