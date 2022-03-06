FROM golang:1.16-alpine

RUN apk add --no-cache git

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o server .

EXPOSE 9000
CMD [ "/app/server" ]