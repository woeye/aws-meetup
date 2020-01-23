tag = $(shell date +"%Y%m%d%H%M")

build:
	docker build -t awsmeetup/calc-service:latest -t awsmeetup/calc-service:$(tag) .
	docker tag awsmeetup/calc-service:$(tag) liprotectregistry.azurecr.io/awsmeetup/calc-service:$(tag)
	docker push liprotectregistry.azurecr.io/awsmeetup/calc-service:$(tag)

gen_proto:
	protoc -I ./protos ./protos/calc_service.proto --go_out=plugins=grpc:calc_service