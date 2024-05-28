FROM golang:1.20.5

WORKDIR /app

ADD . /app/

RUN go build -o ./goltpb .
EXPOSE 8080

ENTRYPOINT ["./goltpb"]
