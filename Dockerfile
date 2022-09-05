FROM golang:alpine

WORKDIR /app

COPY . .

EXPOSE 3000

RUN go build

ENTRYPOINT ["./sso"]

