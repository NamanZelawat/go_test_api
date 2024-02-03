FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o api cmd/api/main.go

CMD ["./api"]

EXPOSE 8080
