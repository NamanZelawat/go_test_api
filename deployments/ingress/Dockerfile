FROM golang:alpine as build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/ingress/main.go

FROM alpine:latest as final

COPY --from=build /app/main .

CMD ["./main"]