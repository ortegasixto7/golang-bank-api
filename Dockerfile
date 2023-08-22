FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /dist
EXPOSE 8005
CMD [ "/dist" ]