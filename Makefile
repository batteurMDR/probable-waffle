build-run:
	go mod tidy
	go mod vendor
	docker build -t pw-server .
	docker-compose up -d

build:
	docker build -t pw-server .