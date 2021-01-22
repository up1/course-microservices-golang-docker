FROM golang:1.15.7-alpine3.13 AS builder
WORKDIR /src  
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/main.go

FROM alpine:3.13  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /src/app .
CMD ["./app"] 