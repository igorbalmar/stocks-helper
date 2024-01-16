ARG GO_VERSION=1.21

FROM golang:${GO_VERSION}-alpine as build

WORKDIR $GOPATH/build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod .
RUN go mod download && go mod verify

COPY . .
COPY main.go .

RUN echo "TZ='America/Sao_Paulo'; export TZ" >> ~/.profile && \
    . ~/.profile && \
    go build -v -o stocks-helper

FROM alpine

WORKDIR /usr/local/bin/

COPY --from=build  /go/build/stocks-helper /usr/local/bin/

CMD [ "./stocks-helper" ]
