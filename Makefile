

run:
	docker-compose up --build
db:
	docker-compose exec mysql bash
test:
	go test ./...

