# build phase
FROM golang:1.20

WORKDIR /app
COPY . /app
ENV CGO_ENABLED=0

RUN go build -o order-tracker .

# execution phase
FROM alpine:latest

WORKDIR /
COPY --from=0 /app/order-tracker ./
EXPOSE 8080

CMD ["./order-tracker"]

