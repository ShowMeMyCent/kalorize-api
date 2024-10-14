    FROM golang:1.19-alpine
    LABEL version="1.0"
    LABEL maintainer="Glorious Satria Dhamang Aji"
    WORKDIR /app
    COPY . . 
    RUN go mod tidy
    RUN GOOS=linux GOARCH=amd64 go build server.go
    EXPOSE 8080
    ENTRYPOINT ["./server"]