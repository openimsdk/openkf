#!/bin/sh

./dbmigration -c ./config.docker.yaml
./server -c ./config.docker.yaml