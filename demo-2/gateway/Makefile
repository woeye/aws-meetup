tag = $(shell date +"%Y%m%d%H%M")

build:
	docker build -t awsmeetup/gateway:latest -t awsmeetup/gateway:$(tag) .
	docker tag awsmeetup/gateway:$(tag) liprotectregistry.azurecr.io/awsmeetup/gateway:$(tag)
	docker push liprotectregistry.azurecr.io/awsmeetup/gateway:$(tag)

gen_proto:
	protoc -I ./protos ./protos/calc_service.proto --go_out=plugins=grpc:calc_service
