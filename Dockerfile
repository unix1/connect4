FROM golang:1.15.2

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build .

CMD ["./connect4"]
