FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 8080
EXPOSE 8090
CMD ["/go/bin/app"]
