FROM golang:1.15.8-alpine as builder

RUN apk --update upgrade \
    && apk --no-cache --no-progress add make \
    && rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/jokly/go-click

COPY go.mod go.sum ./
RUN GO111MODULE=on go mod download

COPY . .

RUN make binary

FROM alpine:3.13

COPY --from=builder /go/src/github.com/jokly/go-click/go-click /

EXPOSE 8080

ENTRYPOINT [ "/go-click" ]