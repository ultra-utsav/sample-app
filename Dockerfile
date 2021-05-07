FROM golang:alpine AS builder

ADD . /go/src/sample-app

WORKDIR /go/src/sample-app

RUN go get ./...

RUN go build -o app

RUN go install

FROM alpine

RUN mkdir sample-app
WORKDIR /sample-app

COPY --from=builder /go/src/sample-app/app .

RUN mkdir static
COPY --from=builder /go/src/sample-app/static ./static

ENTRYPOINT ["./app"]