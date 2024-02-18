## Build
FROM golang:alpine as builder
RUN apk update
RUN apk add git
RUN mkdir /build

COPY . /build

WORKDIR /build

RUN go get -u ./...
RUN go build -o books-api .

## Deploy
FROM alpine:latest
RUN addgroup -g 1000 noroot
RUN adduser -u 1000 -G noroot -h /home/noroot -D noroot
RUN mkdir /home/noroot/app
WORKDIR /home/noroot/app

COPY --from=builder /build/books-api /home/noroot/app/
COPY --from=builder /build/app.* /home/noroot/app/

EXPOSE $API_PORT

CMD ["./books-api"]