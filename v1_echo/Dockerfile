FROM golang:1.20.4-alpine3.17

WORKDIR /app

COPY . .

RUN go mod tidy
RUN GO111MODULE=on go build -o main .

CMD ["./main"]