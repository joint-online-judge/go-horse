FROM golang:1.20 AS build

WORKDIR /go/src/fiber

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/fiber/app .

ENV HOST="localhost" \
    PORT=34764 \
    JWT_SECRET="secret"

EXPOSE $PORT

CMD ["./app"]
