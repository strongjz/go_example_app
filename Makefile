OUTPUT = app
REGISTRY = strongjz
IMAGE = golang_example
VERSION = 0.0.1

.PHONY: test clean install

test:
	go test ./...

clean:
	rm -f $(OUTPUT)

install:
	go get .

build:
	go build -o $(OUTPUT) main.go

run:
	go run main.go

compose_up:
	source .env && docker-compose up

docker_build:
	docker build -t $(REGISTRY)/$(IMAGE):$(VERSION) .

docker_run:
	docker run --env-file=.env -it --rm -p 8080:8080 $(REGISTRY)/$(IMAGE):$(VERSION)

docker_push:
	docker push $(REGISTRY)/$(IMAGE):$(VERSION)
