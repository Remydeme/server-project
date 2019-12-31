FROM golang:alpine

MAINTAINER Rémy DEME <demeremy@gmail.com>

RUN apk update && apk add --no-cache git  && \
        apk add --no-cache make

RUN mkdir -p /go/src/github.com/Remydeme/esme-devops-project

WORKDIR /go/src/github.com/Remydeme/esme-devops-project

COPY . .

WORKDIR /go/src/github.com/Remydeme/esme-devops-project

EXPOSE 8080

CMD ["make", "build"]
