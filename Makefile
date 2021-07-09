.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot


build-image:
	docker docker build -t pocketbot:v0.1 .

start-container:
	docker run --name pocketbot -p 80:80 --env-file .env pocketbot:v0.1