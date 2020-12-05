FROM golang:1.15.3-alpine
ENV GO111MODULE=auto
WORKDIR /app
COPY /backend /app
EXPOSE 8080
CMD ["go","run","main.go"]
