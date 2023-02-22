CMD ?= sh
ENV_FILE ?= .env
IMAGE := dockerUsername/demobot

# Need to hold off on this target until we get a dockerfile made (and probably docker-compose set up since we'll have a handful of containers, including possibly mongodb for storage)
# build:
	# -docker build -t ${IMAGE} .

# TODO: Add protobuf commands. Need to get this installed on my local machine (windows), but am trying to set this up so it will run on mac
# proto:
	# -

# This is a reference - it will not be used. Well, at least not with its current flags. We will probably use a `make run` command eventually.
# run:
	# -docker run --rm --env-file=${ENV_FILE} -v ${HOME}/.aws:/root/.aws -it ${IMAGE} ${CMD}

.PHONY: build run