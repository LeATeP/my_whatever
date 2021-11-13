#!/usr/bin/env bash
name=$(./docker_name_gen.py workshop)
sudo docker run -it \
    -v "/home/leatep/Desktop/code/docker/$name:/home" \
    -v "/var/run/docker.sock:/var/run/docker.sock" \
    -h "${name}" \
    --name "${name}" \
    rwxrob/workspace