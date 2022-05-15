# Build the Go Binary.
FROM golang:1.17-alpine AS builder

ENV CGO_ENABLED 0
ARG BUILD_REF

ARG GOARCH

# Copy the source code into the container.
COPY . /cloud-front-test

# Build the service binary.
WORKDIR /cloud-front-test
RUN GOARCH=${GOARCH}  go build -ldflags "-X main.build=${BUILD_REF}"

# Run the Go Binary in Alpine.
FROM alpine:latest

ARG BUILD_DATE
ARG BUILD_REF

COPY --from=builder /cloud-front-test/mybuildings /cloud-front-test/mybuildings
COPY --from=builder /cloud-front-test/data /cloud-front-test/data
WORKDIR /cloud-front-test

CMD ["./cloud-front-test"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="mybuildings" \
      org.opencontainers.image.vendor="Circutor"
