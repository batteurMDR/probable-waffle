# syntax=docker/dockerfile:1
FROM golang:1.18 AS builder
WORKDIR /go/src/github.com/estiam/probable-waffle
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/estiam/probable-waffle/config.yml ./
COPY --from=builder /go/src/github.com/estiam/probable-waffle/server ./
RUN mkdir upload
CMD ["./server"]