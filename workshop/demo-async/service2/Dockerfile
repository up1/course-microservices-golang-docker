FROM golang:1.15.6-alpine3.12 as build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app /src/cmd/main.go

FROM alpine:3.13.0  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV TZ=Asia/Bangkok
COPY --from=build /src/app ./app
CMD ["./app"] 
