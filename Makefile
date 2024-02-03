api:
	docker-compose --verbose -f deployments/docker-compose.yml up -d --build api mysql

mysql:
	docker-compose --verbose -f deployments/docker-compose.yml up -d mysql