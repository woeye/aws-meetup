tag = $(shell date +"%Y%m%d%H%M")

build:
	docker build -t awsmeetup/service-1:latest -t awsmeetup/service-1:$(tag) .
	docker tag awsmeetup/service-1:$(tag) registry.k3s:80/awsmeetup/service-1:$(tag)
	docker push registry.k3s:80/awsmeetup/service-1:$(tag)
