FROM golang:latest AS builder
WORKDIR /go/src
COPY main.go /go/src/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/hookr .
RUN find /go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/hookr .
EXPOSE 8080
CMD ["./main", "/tmp"] 
