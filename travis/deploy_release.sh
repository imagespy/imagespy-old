#!/usr/bin/env bash

set -eu

docker login -u $DOCKER_USER -p $DOCKER_PASS
docker tag imagespy/imagespy:master imagespy/imagespy:latest
docker tag imagespy/imagespy:master imagespy/imagespy:$TRAVIS_TAG
docker push imagespy/imagespy:latest
docker push imagespy/imagespy:$TRAVIS_TAG
