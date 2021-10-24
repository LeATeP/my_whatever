#!/usr/bin/env bash
name=$(docker_name_gen.py workshop)
sudo docker run -it \
    -v "/home/leatep/Desktop/code/docker/$name:/home" \
    -h "${name}" \
    --name "${name}" \
    rwxrob/workspace