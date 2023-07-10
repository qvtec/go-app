FROM golang:1.20

# "golang-migrate"のインストール
RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY ../../go.* ./
RUN go mod download && go mod verify

COPY ../../  ./
RUN go build -o main ./cmd/main.go

CMD ["./main"]