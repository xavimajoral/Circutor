# Build the Go Binary.
FROM golang:1.17-alpine AS builder

RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache ca-certificates \
    && apk add --update gcc musl-dev

ENV CGO_ENABLED 0
ARG BUILD_REF

ARG GOARCH

# Copy the source code into the container.
COPY . /cloud-front-test

# Build the service binary.
WORKDIR /cloud-front-test
RUN GOARCH=${GOARCH}  go build -ldflags "-w -s -X main.build=${BUILD_REF}"

# Run the Go Binary in Alpine.
FROM alpine:latest

ARG BUILD_DATE
ARG BUILD_REF

COPY --from=builder /cloud-front-test /cloud-front-test
#COPY --from=builder /cloud-front-test/data /cloud-front-test/data
WORKDIR /cloud-front-test

CMD ["./cloud-front-test"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="mybuildings" \
      org.opencontainers.image.vendor="Circutor"
