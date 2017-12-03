#!/usr/bin/env bash

set -eu

docker login -u $DOCKER_USER -p $DOCKER_PASS
docker push imagespy/imagespy:master
