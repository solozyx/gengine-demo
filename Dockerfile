FROM golang:1.13-alpine as builder
WORKDIR /app/
COPY . /app/
#RUN export GO111MODULE=on && \
#    export GOPROXY=https://mirrors.aliyun.com/goproxy/,https://goproxy.cn,https://goproxy.io,direct && \
#    echo 'nameserver 223.5.5.5' > /etc/resolv.conf && \
#    go build  -o ./gengine  ./main.go

#RUN export GO111MODULE=on && \
#  export GOPROXY=https://goproxy.cn,direct && \
#  go build  -o ./gengine  ./main.go

RUN export GO111MODULE=on && \
    export GOPROXY=https://goproxy.cn,direct && \
    echo 'nameserver 223.5.5.5' > /etc/resolv.conf && \
    go build  -o ./gengine  ./main.go

FROM alpine:latest as prod
ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone 
#RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=0 /app/config.yml .
COPY --from=0 /app/gengine .
# 文件上传下载
RUN mkdir file
CMD ["./gengine"]
