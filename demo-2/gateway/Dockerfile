# Builder
FROM golang:1.13-alpine as builder
RUN apk add --no-cache protoc git
RUN mkdir /build
WORKDIR /build
RUN go get -u github.com/golang/protobuf/protoc-gen-go
COPY go.mod go.sum *.go /build/
COPY calc_service /build/calc_service
COPY protos /build/protos
RUN go get -v
RUN protoc -I ./protos ./protos/calc_service.proto --go_out=plugins=grpc:calc_service
RUN go build -o gateway

# Final image
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/gateway .
EXPOSE 8000
CMD ["/app/gateway"]
