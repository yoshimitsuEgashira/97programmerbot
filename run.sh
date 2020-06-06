#!/bin/sh
set -eux
sudo docker build -t 97programmer:latest .
sudo docker run --env-file=.env 97programmer:latest
sudo docker container prune -f
