FROM golang:1.20-alpine as build-stage
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /dist

# Tests
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /dist /dist
EXPOSE 8005
CMD [ "/dist" ]