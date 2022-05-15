ARCH :=  $(shell ./get-go-arch.sh)
OS := darwin
API_IMAGE_NAME := cloud-front-test
VERSION := latest

tidy:
	go mod tidy

run:
	go run server.go

build:
	GOARCH=$(ARCH) GOOS=$(OS)  go build -ldflags

build-docker:
	docker build \
		-f cloud-front-test.dockerfile \
		-t $(API_IMAGE_NAME):$(VERSION) \
		--build-arg GOARCH=$(ARCH) \
		.

docs:
	~/go/bin/swag init -g server.go --parseVendor --parseDependency --parseInternal