Generate proto files

cd api

protoc --go_out=..//internal --go_opt=paths=source_relative --go-grpc_out=../internal --go-grpc_opt=paths=source_relative proto/im
age/test.proto
