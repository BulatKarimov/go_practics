FROM golang:1.13-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/practics cmd/app/main.go
ENTRYPOINT ["bin/practics"]