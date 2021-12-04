FROM golang:alpine AS BUILD

WORKDIR /go/src/app

COPY . .

RUN go mod download \
  && go mod verify \
  && go mod tidy

CMD go run main.go
