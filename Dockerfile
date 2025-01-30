FROM golang:1.23.5-alpine3.21

WORKDIR /app

COPY go.mode .

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

CMD [ "./main" ]