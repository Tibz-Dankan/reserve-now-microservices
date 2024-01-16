FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o ./bin/roomapp ./

EXPOSE 8080

CMD ["./bin/roomapp"]
