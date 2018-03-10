#!/bin/bash

docker rm -f qmanager
docker rm -f qmworker
docker rm -f qmclient
docker rm -f qmsimulator

docker network rm qmanager
docker network create qmanager
