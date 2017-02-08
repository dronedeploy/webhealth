.PHONY : init clean package test tag push

GIT_HASH := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
IMAGE_NAME := "dronedeploy/webhealth"

package:
	docker build --build-arg GIT_HASH=$(GIT_HASH) -t $(IMAGE_NAME):$(GIT_HASH) .

test:
	@#no tests

tag:
	docker tag $(IMAGE_NAME):$(GIT_HASH) $(IMAGE_NAME):$(GIT_BRANCH)

push: tag
	time docker push $(IMAGE_NAME):$(GIT_HASH)
	time docker push $(IMAGE_NAME):$(GIT_BRANCH)
