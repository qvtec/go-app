FROM golang:1.20-alpine

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@v1.44.0

WORKDIR /app

COPY ../../go.* ./
RUN go mod download && go mod verify

COPY ../../  ./
RUN go build -o main ./cmd/main.go

EXPOSE 8088

CMD ["air", "-c", ".air.toml"]