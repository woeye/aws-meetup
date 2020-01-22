# Builder
FROM arm32v7/golang:1.13-alpine as builder
RUN mkdir /build
WORKDIR /build
COPY go.mod go.sum *.go /build/
RUN go get -v
RUN go build -o hello-app

# Final image
FROM arm32v7/alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/hello-app .
EXPOSE 8000
CMD ["/app/hello-app"]
