FROM golang:1.15.3-alpine
ENV GO111MODULE=auto
WORKDIR /app
COPY . /app
RUN go run main.go
EXPOSE 8080
