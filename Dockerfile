FROM golang:1.22.2

WORKDIR /app

ADD . /app/

RUN go build -o ./goltpb .
EXPOSE 8080

ENTRYPOINT ["./goltpb"]
