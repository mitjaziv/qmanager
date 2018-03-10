.PHONY: all build clean build-in-docker build-dockers run-docker run-simulator-docker clean-docker

# build executable.
all:
	@echo 'Usage: make <build|clean|build-in-docker|build-dockers|run-docker|run-simulator-docker|clean-docker>'

# build executable.
build:
	./scripts/build.sh

# cleans up any binaries
clean:
	@rm -rf build

# build in docker
build-in-docker:
	./scripts/build-in-docker.sh

# build dockers
build-dockers:
	./scripts/build-dockers.sh

# run docker
run-docker:
	./scripts/run-docker.sh

# run simulator docker
run-simulator-docker: build-in-docker build-dockers
	./scripts/run-simulator-docker.sh

# clean docker
clean-docker:
	./scripts/clean-docker.sh
