FROM golang:latest AS builder
LABEL maintainer "mats@viacard.com"

WORKDIR /go/src
COPY main.go /go/src/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o /go/bin/hookr .

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/hookr /go/bin/hookr

USER nobody
EXPOSE 8080
VOLUME ["/data"]

ENTRYPOINT ["/go/bin/hookr"]
CMD ["/data/"]
