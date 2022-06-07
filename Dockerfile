FROM golang:bullseye
WORKDIR /data
COPY . .
RUN GOPROXY=https://proxy.golang.com.cn,direct go build -o server

FROM debian:bullseye-slim
MAINTAINER bipy notbipy@gmail.com

WORKDIR /app

COPY --from=0 /data/server server
COPY --from=0 /data/entrypoint.sh entrypoint.sh
COPY --from=0 /data/*.env ./

ENV DEBIAN_FRONTEND=noninteractive
RUN echo "deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free" > /etc/apt/sources.list
RUN apt update; apt install -y ffmpeg

EXPOSE 80

ENTRYPOINT ["/app/entrypoint.sh"]
