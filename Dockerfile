FROM golang:1.22 as build

WORKDIR /app

COPY . /app

RUN go get -u && \
  CGO_ENABLED=0 go build -ldflags "-s -w" -o ecr-login .

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/ecr-login /app

ENTRYPOINT ["/app/ecr-login"]
