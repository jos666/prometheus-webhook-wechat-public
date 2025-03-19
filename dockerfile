FROM debian:latest
COPY .build/linux-amd64/prom-webhook-wechat /root/prom-webhook-wechat
COPY config.yaml /root/config.yaml
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list.d/debian.sources && apt update && apt  install --reinstall ca-certificates -y && update-ca-certificates
ENV LANG=C.utf8
ENV TZ=Asia/Shanghai
EXPOSE 8060
CMD ["/root/prom-webhook-wechat", "-config.file=/root/config.yaml"]
