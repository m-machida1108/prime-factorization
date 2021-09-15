APP := docker-compose exec -T app

up:
	docker-compose up -d

down:
	docker-compose down

exec:
	$(APP) go run ./cmd/main.go ${target}

image-build:
	docker-compose down
	docker-compose build --no-cache