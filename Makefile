APP := docker-compose exec -T app

up:
	docker-compose up -d

down:
	docker-compose down

exec:
	$(APP) go run ./cmd/main.go ${target}
