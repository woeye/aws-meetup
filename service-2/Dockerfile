# Builder
FROM golang:1.13-alpine as builder
RUN mkdir /build
WORKDIR /build
COPY go.mod go.sum *.go /build/
RUN go get -v
RUN go build -o hello-app-2

# Final image
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/hello-app-2 .
EXPOSE 8000
CMD ["/app/hello-app-2"]
