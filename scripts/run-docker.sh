#!/bin/bash

docker rm -f qmanager
docker run --name qmanager --restart=always -p 8080:8080 -p 8090:8090 -d mitjaziv/qmanager -http 0.0.0.0:8080 -rpc 0.0.0.0:8090
