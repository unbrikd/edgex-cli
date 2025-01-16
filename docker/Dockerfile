# This dockerfile is intended to be used for building the application and not to run as a container.
# The final image is a scratch image with the binary copied to it that will be copied to the host machine using the 
# provided Makefile.

FROM golang:1.22.2-bullseye AS builder

ARG GOOS
ARG GOARCH
ARG GO_APP
ARG GO_LDFLAGS_EXTRA
ARG GO_BUILDFLAGS_EXTRA
ENV CGO_ENABLED=1

WORKDIR /build

# Install system dependencies
RUN apt-get update && apt-get install -y \
gcc \
libc6-dev \
libzmq3-dev \
pkg-config \
&& rm -rf /var/lib/apt/lists/*

# Clear module cache and ensure clean download
RUN go clean -modcache

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w ${GO_LDFLAGS_EXTRA}" ${GO_BUILDFLAGS_EXTRA} -o /app/${GO_APP} .

FROM scratch

ARG GO_APP

WORKDIR /app
COPY --from=builder /app/${GO_APP} /app/${GO_APP}

ENTRYPOINT ["/app/${GO_APP}"]