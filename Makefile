.PHONY: init build build-server build-storage up run start down stop restart

init:
	@if [ ! -f '.env' ]; then \
		echo 'Copying .env file...'; \
		cp .env.example .env; \
	fi;
	@if [ ! -f 'src/.env' ]; then \
		echo 'Copying src/.env file...'; \
		cp src/.env.example src/.env; \
	fi;
	@if [ ! -f 'docker-compose.yml' ]; then \
		echo 'Copying docker-compose.yml file...'; \
		cp docker-compose.example.yml docker-compose.yml; \
	fi;
	@echo ''
	@echo 'NOTE: Please check your configuration in ".env" before run.'
	@echo 'NOTE: Please check your configuration in "docker-compose.yml" before run.'
	@echo ''

build: build-server build-storage

build-server:
	docker-compose build server

build-storage:
	docker-compose build storage

start: run
up: run
run:
	docker-compose up -d --force-recreate

down: stop
stop:
	docker-compose down

restart: stop run
