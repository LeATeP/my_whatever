#!/bin/bash

sudo apt-get update && apt-get install -y vim git

sudo apt-get update && apt-get install -y docker.io

sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update && apt-get install -y kubelet kubeadm kubectl

mkdir code
(cd code && git clone https://github.com/LeATeP/my_whatever)

cp /home/code/my_whatever/.bashrc ~/
cp /home/code/my_whatever/.bash-aliases ~/
cp /home/code/my_whatever/.vimrc ~/