tag = $(shell date +"%Y%m%d%H%M")

build:
	docker build -t awsmeetup/hello-app-2:latest -t awsmeetup/hello-app-2:$(tag) .
	docker tag awsmeetup/hello-app-2:$(tag) liprotectregistry.azurecr.io/awsmeetup/hello-app-2:$(tag)
	docker push liprotectregistry.azurecr.io/awsmeetup/hello-app-2:$(tag)
