#!/bin/env sh

net="p_net"

sudo docker run -it \
    --network "$net" \
    -v /home/leatep/code/Vaava:/Vaava \
    -w /Vaava \
    leatep/workspace
