FROM golang:1.21-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o green-api-app .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/green-api-app .
COPY static ./static

EXPOSE 8080

ENV PORT=8080 \
    ENV=production

CMD ["./green-api-app"]