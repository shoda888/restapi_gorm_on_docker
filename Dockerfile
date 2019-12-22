FROM golang:1.11.2-alpine3.8

WORKDIR /go
ADD . /go

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/jinzhu/gorm

    # go get github.com/labstack/echo/middleware && \

CMD ["go", "run", "main.go"]