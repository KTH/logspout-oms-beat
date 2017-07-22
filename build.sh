#!/bin/bash
#
# Compiles statically linked binary and builds docker file.
# Options are passed as options to the docker build command.
#
set -ex
rm beat
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o beat .

docker build $* .
