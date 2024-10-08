up:
	docker-compose up -d --build

stop:
	docker-compose stop

down:
	docker-compose down

unit:
	go test ./...

test:
	go test -tags integrations ./...
