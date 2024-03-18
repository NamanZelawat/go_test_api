app:
	docker-compose --verbose -f deployments/docker-compose.yml up -d --build api ingress minio kafka model

api_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d --build api

ingress_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d ingress

mysql_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d mysql

model_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d model

kafka_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d kafka

proto_files:
	buf generate