FROM golang:1.22 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
COPY vendor vendor/
COPY main.go main.go
COPY server server/
COPY fibonacci fibonacci/

RUN go build -mod=vendor -o serverapp main.go

FROM fedora:latest

WORKDIR /
COPY --from=builder /workspace/serverapp .
USER 65532:65532
EXPOSE 80
ENTRYPOINT ["/serverapp"]
