CMD ?= sh
ENV_FILE ?= .env
IMAGE := dockerUsername/demobot

build:
	-docker build -t ${IMAGE} .

# TODO: Add protobuf commands. Need to get this installed on my local machine (windows), but am trying to set this up so it will run on mac
# proto:
	# -

# run:
	# -docker run --rm --env-file=${ENV_FILE} -v ${HOME}/.aws:/root/.aws -it ${IMAGE} ${CMD}

.PHONY: build run