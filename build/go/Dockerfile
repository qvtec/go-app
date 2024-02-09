FROM golang:1.20-alpine

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@v1.44.0
RUN go install github.com/vektra/mockery/v2@v2.32.0
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY ../../go.* ./
RUN go mod download && go mod verify

COPY ../../  ./
RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]