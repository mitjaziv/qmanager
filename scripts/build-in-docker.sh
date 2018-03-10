#!/bin/bash

docker build --rm -t qmanager-build -f ./scripts/docker/Dockerfile.build .
docker run --rm -it -v $PWD:/go/src/github.com/mitjaziv/qmanager qmanager-build ./scripts/build.sh
docker rmi qmanager-build
