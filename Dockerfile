FROM golang:1.26.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY api/ ./api/ 
COPY cmd/ ./cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

EXPOSE 8080
CMD ["./main"]
