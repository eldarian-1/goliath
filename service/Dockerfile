FROM golang:1.25.5-alpine

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE $GOLIATH_PORT
CMD ["go", "run", "server.go"]
