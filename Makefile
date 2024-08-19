up:
	docker-compose -f ./deployment/docker-compose.yaml up -d --build

stop:
	docker-compose -f ./deployment/docker-compose.yaml stop

down:
	docker-compose -f ./deployment/docker-compose.yaml down