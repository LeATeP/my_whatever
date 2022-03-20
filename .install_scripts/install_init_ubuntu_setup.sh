#!/bin/bash

# 1. installing go, python/pip, vscode, google-chrome, data grip, lolcat, tweaks
# 1.1 virtualbox, kind/k8s
# 2. changing power-saving settings 
# 3. set up bashrc, aliases, scripts
# 4. setting up git, github ssh keys, 
# 5. git, pulling main repos
# 6. docker.. getting docker backup's, pull images: [ubuntu, python, golang]

# installing go
curl -Lo ./go.tar.gz https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz
go version
# installing kind
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.11.1/kind-linux-amd64
chmod +x ./kind
mv ./kind /home/leatep/go/bin/kind