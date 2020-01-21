tag = $(shell date +"%Y%m%d%H%M")

build:
	docker build -t awsmeetup/service-1:latest -t awsmeetup/service-1:$(tag) .

push: build
	docker tag awsmeetup/service-1:$(tag) 192.168.178.70:30000/awsmeetup/service-1:$(tag)
	docker push 192.168.178.70:30000/awsmeetup/service-1:$(tag)
