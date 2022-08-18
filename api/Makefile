ARCH :=  $(shell ./get-go-arch.sh)
OS := darwin
API_IMAGE_NAME := cloud-front-test
VERSION := 1.1

tidy:
	go mod tidy

run:
	go run main.go

build:
	GOARCH=$(ARCH) GOOS=$(OS) go build

build-docker:
	docker build \
		-f cloud-front-test.dockerfile \
		-t $(API_IMAGE_NAME):$(VERSION) \
		--build-arg GOARCH=$(ARCH) \
		.

build-docker-multiplatform:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		-f cloud-front-test.dockerfile \
		-t circutor/$(API_IMAGE_NAME):$(VERSION) \
		--push \
		.


docs:
	~/go/bin/swag init -g main.go --parseVendor --parseDependency --parseInternal