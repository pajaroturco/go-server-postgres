# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.20.0-alpine

ENV GO111MODULE=on

RUN apk update  && \
    apk add git \
    make openssh-client \
    curl

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air -c .air.toml