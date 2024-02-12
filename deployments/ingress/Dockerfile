FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ingress cmd/ingress/main.go

CMD ["./ingress"]

EXPOSE 8080
