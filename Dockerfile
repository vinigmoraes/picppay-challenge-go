FROM golang:1.23 AS build

WORKDIR /app

COPY . /app

RUN go build -o api /app/cmd/api/main.go

EXPOSE 8080

CMD ["/app/api"]

