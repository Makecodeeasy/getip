FROM golang:1.18rc1-alpine as builder

COPY getip /workspace/release

WORKDIR /workspace/release
ENV GOPROXY https://goproxy.cn/
RUN go build -o getip

# FROM {{ RUNTIME_IMAGE }}
FROM alpine

COPY --from=builder /workspace/release/getip /getip
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8080

CMD ["/getip"]
