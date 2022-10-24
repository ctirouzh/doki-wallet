FROM docker.io/library/golang:buster AS builder
ENV GOPROXY=https://proxy.golang.org,direct
WORKDIR /src

# Install vendor dependencies
COPY go.* ./
RUN go mod download

# Build project
COPY . .
RUN go build -o app ./cmd/server

# --------------------------------------------------

FROM docker.io/library/debian:buster-slim

# Fix certificate issues
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    apt-get clean

COPY .docker/docker-entrypoint.sh /
COPY --from=builder /src/app /usr/local/bin/

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD []