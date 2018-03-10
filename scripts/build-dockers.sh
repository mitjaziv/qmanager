#!/bin/bash

docker build --rm -t mitjaziv/qmanager -f ./scripts/docker/Dockerfile.qmanager --build-arg GITVERSION=${CI_COMMIT_ID} --build-arg GITTAG=${CI_TAG_AUTO} .
#docker build --rm -t mitjaziv/qmworker -f ./scripts/docker/Dockerfile.qmworker --build-arg GITVERSION=${CI_COMMIT_ID} --build-arg GITTAG=${CI_TAG_AUTO} .
docker build --rm -t mitjaziv/qmsimulator -f ./scripts/docker/Dockerfile.qmsimulator --build-arg GITVERSION=${CI_COMMIT_ID} --build-arg GITTAG=${CI_TAG_AUTO} .
#docker build --rm -t mitjaziv/qmclient -f ./scripts/docker/Dockerfile.qmclient --build-arg GITVERSION=${CI_COMMIT_ID} --build-arg GITTAG=${CI_TAG_AUTO} .
