FROM golang:alpine
LABEL maintainer="deemanth1995@gmail.com"

WORKDIR /build

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY src/main.go /build/

EXPOSE 7000

CMD ["go run /build/main.go"]