#!/bin/bash

docker rm -f qmanager
docker rm -f qmsimulator

docker network rm qmanager
docker network create qmanager

docker run --name qmanager --net=qmanager --restart=always -p 8080:8080 -p 8090:8090 -d mitjaziv/qmanager -http qmanager:8080 -rpc qmanager:8090
docker run --name qmsimulator --net=qmanager --restart=always -d mitjaziv/qmsimulator -rpc qmanager:8090
