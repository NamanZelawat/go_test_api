app:
	docker-compose --verbose -f deployments/docker-compose.yml up -d --build api ingress

api_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d --build api

ingress_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d ingress

mysql_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d mysql

proto_files:
	buf generate