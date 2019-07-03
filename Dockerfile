FROM golang:1.12-alpine
LABEL maintainer="Luca Bruzzone"

WORKDIR $GOPATH/src/github.com/unsign3d/smart-thermometer-server

COPY . .

RUN apk update
RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep


RUN .ci/bootstrap.sh

RUN go build -o out

RUN chmod +x out

EXPOSE 8080

ENTRYPOINT ["./out"]
