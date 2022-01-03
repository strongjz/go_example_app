FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM scratch
ARG PORT
ARG AD_PORT

COPY --from=builder /go/bin/app /go/bin/app

EXPOSE $PORT
EXPOSE $AD_PORT
CMD ["/go/bin/app"]
