FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /hellod

FROM alpine:3.18
WORKDIR /
COPY --from=builder /hellod /hellod
EXPOSE 8080
RUN adduser -D hellod
USER hellod
ENTRYPOINT ["/hellod"]
