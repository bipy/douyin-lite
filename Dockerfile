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

EXPOSE 80

ENTRYPOINT ["/app/entrypoint.sh"]
