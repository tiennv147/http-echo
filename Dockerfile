#build stage
FROM golang:alpine AS builder
ADD . /go/src/github.com/tiennv147/http-echo
WORKDIR /go/src/github.com/tiennv147/http-echo
RUN apk add --no-cache git
RUN go mod vendor
RUN go build -o ./main ./main.go

#final stage
FROM alpine:latest
ARG env=dev
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/tiennv147/http-echo/main /app/server
COPY --from=builder /go/src/github.com/tiennv147/http-echo/config.yaml /app/config.yaml

EXPOSE 8080

CMD ["./app/server", "-c", "/app/config.yaml"]