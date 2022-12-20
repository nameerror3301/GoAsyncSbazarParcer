FROM golang:1.19 AS builder

WORKDIR /service

COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /go-async-sbazar ./app/cmd/main.go


FROM alpine:3.10
COPY --from=builder /go-async-sbazar /bin

CMD ["sh", "-c", "/bin/go-async-sbazar"]