FROM golang:1.22-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY backend/ .

RUN go mod tidy && go build -o server .

EXPOSE 8080
CMD ["./server"]
