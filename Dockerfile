FROM golang:1.17 AS build

WORKDIR /go/src/app

COPY cmd/ cmd/
COPY go.mod go.sum ./

RUN go build -o app cmd/main.go

FROM alpine:latest

COPY --from=build /go/src/app/app /app

EXPOSE 8080

CMD ["/app"]
