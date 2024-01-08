FROM golang:1.21-alpine

WORKDIR /app

COPY . .

# TODO: To include command to install packages

RUN go build -o ./bin/authapp ./

EXPOSE 8080

CMD ["./bin/authapp"]
