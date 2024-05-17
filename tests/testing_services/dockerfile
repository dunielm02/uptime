FROM golang:1.22.2 as go_build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./test ./services/http_server/main.go

FROM ubuntu:latest as prod

COPY --from=go_build /app/test /app/test

RUN apt-get update && apt-get install -y openssh-server supervisor
RUN mkdir -p /var/run/sshd /var/log/supervisor

RUN useradd -m -s /bin/bash test
RUN echo "test:test" | chpasswd

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 22 8000
CMD ["/usr/bin/supervisord"]