FROM golang:latest

LABEL maintainer="XML TEAM"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 9090

CMD ["./auth-service"]

