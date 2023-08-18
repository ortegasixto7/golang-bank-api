FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

WORKDIR /app/src

RUN go build -o /go-bank-api

EXPOSE 8005

CMD [ "/go-bank-api" ]