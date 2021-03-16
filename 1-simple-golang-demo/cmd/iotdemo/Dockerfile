FROM golang:1.13.4-alpine3.10

WORKDIR $GOPATH/src/iotdemo

COPY  ./pkg/ ./pkg/

COPY ./cmd/iotdemo ./cmd/iotdemo

COPY ./vendor/ ./vendor/

COPY ./module/ ./module/

COPY ./docs/ ./docs/

COPY ./test/ ./test/

COPY ./go.mod ./go.mod

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-w -s' -o iotdemo ./cmd/iotdemo

FROM alpine:3.10.3

WORKDIR /iotdemo/app

COPY --from=0 /go/src/iotdemo/iotdemo ./

COPY --from=0 /go/src/iotdemo/docs ./

WORKDIR /iotdemo

ENTRYPOINT ["./app/iotdemo"]

