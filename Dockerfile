FROM golang:1.21 AS builder

WORKDIR /build

COPY . /build

RUN go mod tidy && go build -o executor .

# TODO: Change to scratch
FROM ubuntu:22.04
RUN useradd -m chain
USER chain
COPY --from=builder /build/executor /usr/bin/
COPY --from=ghcr.io/grpc-ecosystem/grpc-health-probe:v0.4.19 /ko-app/grpc-health-probe /usr/bin/
CMD ["executor"]
