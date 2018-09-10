.PHONY : init clean package test tag push

GIT_HASH := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
APP_NAME := "webhealth"
IMAGE_NAME := "dronedeploy/$(APP_NAME)"
GCR_IMAGE := "gcr.io/dronedeploy-code-delivery-0/$(APP_NAME)"

package:
	docker build --build-arg GIT_HASH=$(GIT_HASH) -t $(IMAGE_NAME):$(GIT_HASH) .

test:
	@#no tests

tag:
	docker tag $(IMAGE_NAME):$(GIT_HASH) $(IMAGE_NAME):$(GIT_BRANCH)
	docker tag $(IMAGE_NAME):$(GIT_HASH) $(GCR_IMAGE):$(GIT_BRANCH)
	docker tag $(IMAGE_NAME):$(GIT_HASH) $(GCR_IMAGE):$(GIT_HASH)

push: tag
	time docker push $(IMAGE_NAME):$(GIT_HASH)
	time docker push $(IMAGE_NAME):$(GIT_BRANCH)
	docker push $(GCR_IMAGE):$(GIT_HASH)
	docker push $(GCR_IMAGE):$(GIT_BRANCH)
