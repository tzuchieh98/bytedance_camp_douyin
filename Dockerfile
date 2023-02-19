FROM golang:1.19.5 as builder

WORKDIR /go/src/github.com/linzijie1998/bytedance_camp_douyin
COPY . .

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED 0

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN  go build -o server .

FROM alpine:latest

LABEL MAINTAINER="linzijie1998@126.com"

WORKDIR /go/src/github.com/linzijie1998/bytedance_camp_douyin

COPY --from=0 /go/src/github.com/linzijie1998/bytedance_camp_douyin/server ./
COPY --from=0 /go/src/github.com/linzijie1998/bytedance_camp_douyin/config.yaml ./

EXPOSE 8888
ENTRYPOINT ./server