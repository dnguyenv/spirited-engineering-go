# Build from golang-centos7 base image
# dnguyenv@us.ibm.com

FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
