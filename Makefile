.PHONY: build-docker-compose, run-docker-compose

all: build-docker-compose run-docker-compose

build-docker-compose:
	docker-compose build

run-docker-compose:
	docker-compose up --remove-orphans