#!/bin/bash


MAIN_DOCKERFILE=$(cat <<EOF
FROM golang:alpine as builder

WORKDIR /go
COPY main.go .
RUN CGO_ENABLED=0 go build -o main main.go


FROM alpine
RUN adduser -D app
USER app
WORKDIR /
ENV PORT 8000
COPY --from=builder --chown=app:app /go/main /main

ENTRYPOINT ["/main"]
EOF
)

for i in 8080 3000 ; do
    echo "Building for port $i"
    echo "$MAIN_DOCKERFILE" | sed "s/8000/$i/" > Dockerfile
    docker build -t "regmicmahesh/echo-server:$i" .
    docker push "regmicmahesh/echo-server:$i"
done

docker build -t regmicmahesh/echo-server:latest .
docker push regmicmahesh/echo-server:latest
