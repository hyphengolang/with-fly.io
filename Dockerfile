# syntax=docker/dockerfile:1

FROM golang:latest as builder
WORKDIR /home/hyphengolang
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o gofly ./cmd/gofly

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/hyphengolang/gofly .
CMD ["./gofly"]