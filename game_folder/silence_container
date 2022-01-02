#!/bin/env sh

net="p_net"

sudo docker run -d --rm \
    --network "$net" \
    -v /home/leatep/scripts/game_docker:/game_docker \
    -e GAME="/game_docker" \
    -w /game_docker \
    silent_python3 "$@"
