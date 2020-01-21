#!/bin/bash

TAG=quay.io/mvala/logcollector:latest

docker build . -t "${TAG}"
docker push "${TAG}"
