# docker-hookr

Capture webhook -requests from GitHub to be processed by separate process.

The server listens to incoming http requests at port 8080 and dumps both the header 
and body into a file in the /data/ directory. 

Each request gets dumped into a separate file named YYYY-MM-DD_HH.MM.SS.nnnnnnnnn.hook

The /data/ directory must be mounted as a volume in the host and must be writable by user 65534 (nobody).

Please note that no provisons are made to support HTTPS since it it meant to be run together with a reverse proxy like Traefik that handles the TLS termination before proxying the request to hookr.

## Using
If used separately it could be used like this:
```
docker pull viacard/hookr
docker run -d -p 8080:8080 -v /tmp:/data hookr
curl http://localhost:8080
```
You will now have a file containing the request in the /tmp directory on the host.

More realisticly it will be run from docker-compose of something similar.
```
  hookr:
    image: hookr
    container_name: hookr
    restart: always
    volumes:
      - ./data/hookrfiles/:/data
    labels:
      - "traefik.docker.network=proxy"
      - "traefik.enable=true"
      - "traefik.basic.frontend.rule=Host:hook.example.com"
    networks:
      - proxy
```

## Repo

All files can be found at https://github.com/viacard/docker-hookr

## Dockerfile
```
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
```

## Release history

- v1.0 July 7 2019 - Initial release
