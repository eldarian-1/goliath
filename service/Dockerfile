FROM golang:1.25.5

RUN apt update && apt install -y ffmpeg

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE $GOLIATH_PORT
CMD ["go", "run", "main.go"]
