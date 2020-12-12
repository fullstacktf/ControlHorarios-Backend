FROM golang:1.15.3-alpine as builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go build main.go
FROM alpine
COPY --from=builder /build/main /app/
EXPOSE 8080
WORKDIR /app
CMD ["./main"]
