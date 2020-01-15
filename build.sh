#!/bin/bash

TAG=quay.io/mvala/dummylogger:latest

docker build . -t "${TAG}"
docker push "${TAG}"
