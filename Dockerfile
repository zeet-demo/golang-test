FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV PORT 8080
ENV GIN_MODE=release
EXPOSE 8080

RUN go build

CMD ["./golang-test"]