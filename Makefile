app:
	docker-compose --verbose -f deployments/docker-compose.yml up -d --build api ingress minio model message

api_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d --build api

ingress_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d ingress

mysql_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d mysql

model_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d model

message_cont:
	docker-compose --verbose -f deployments/docker-compose.yml up -d message

build_images:
	docker-compose --verbose -f deployments/docker-compose.yml build api ingress

push_images:
	docker-compose --verbose -f deployments/docker-compose.yml push api ingress

proto_files:
	buf generate