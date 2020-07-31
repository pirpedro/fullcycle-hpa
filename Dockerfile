FROM golang:1.14.4-alpine3.12 as builder
COPY src/hello/main.go .
RUN CGO_ENABLED=0 go build -o hello main.go
#RUN apk add musl-utils
#RUN ldd hello

FROM scratch
COPY --from=builder /go/hello /hello
EXPOSE 8000
CMD ["/hello"]