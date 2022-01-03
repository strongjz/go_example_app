OUTPUT ?= golang_example
REGISTRY ?= strongjz
IMAGE ?= golang_example
VERSION ?= 0.0.7

include .env
export

.PHONY: test clean install

test:
	go test ./...

clean:
	rm -f $(OUTPUT)

install:
	go get .

build: install
	go build -o $(OUTPUT) main.go

run: install
	go run main.go

compose_up:
	docker-compose up

docker_build:
	docker build -t $(REGISTRY)/$(IMAGE):$(VERSION) .

docker_run:
	docker run --env-file=.env -it --rm -p $(PORT):$(PORT) -p $(AD_PORT):$(AD_PORT) $(REGISTRY)/$(IMAGE):$(VERSION)

docker_push: docker_build
	docker push $(REGISTRY)/$(IMAGE):$(VERSION); \
	git tag -a $(VERSION); \
	git push origin --tags
