FROM golang:1.16 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app

FROM alpine

WORKDIR /

COPY --from=build /app/main /server

EXPOSE 8080

CMD /server