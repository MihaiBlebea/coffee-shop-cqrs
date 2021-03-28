setup: build up

build:
	docker-compose build

up:
	docker-compose up -d

setup-test: env-file build-test up-test cover-html

env-file: 
	cp ./.env.example ./.env

stop: 
	docker stop coffee-shop-cqrs

build-test:
	docker build \
		--no-cache \
		--file ./Dockerfile.test \
		-t serbanblebea/coffee-shop-cqrs:test \
		.

up-test:
	docker run \
		-v ${PWD}:/app \
		--rm \
		--name coffee-shop-cqrs-test \
		--env-file ./.env \
		serbanblebea/coffee-shop-cqrs:test

cover-html:
	go tool \
		cover -html=cover.out \
		-o cover.html \
		&& open cover.html

go-build:
	go build -o=./coffee-shop . && \
	./coffee-shop start

go-test:
	go test -v ./...
