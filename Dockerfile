# Build from golang-centos7 base image
# dnguyenv@us.ibm.com

FROM golang:latest

ADD . /app/
WORKDIR /app

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
