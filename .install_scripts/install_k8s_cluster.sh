#!/bin/bash

## issue 1 network
# look more at file vmNetworkSetup.yaml

## issue 2 swap
# if http://localhost:10248/healthz' failed with error: 
# 1. sudo swapoff -a
#  remove anything with swap from /etc/fstab or make a cmd like -- ## sudo sed -i '/ swap / s/^/#/' /etc/fstab
# 2. add line to daemon
# sudo mkdir /etc/docker
# cat <<EOF | sudo tee /etc/docker/daemon.json
# {
# "exec-opts": ["native.cgroupdriver=systemd"],
# }
# EOF
# 3. reboot
# 4. if anything, check errors from `journalctl -xeu kubelet` 
# sudo systemctl enable docker && sudo systemctl daemon-reload && \
    # sudo systemctl restart docker && sudo systemctl status kubelet



sudo apt install vim && \
    sudo apt install git 
sudo apt-get update && \
    sudo apt-get install -y apt-transport-https ca-certificates curl
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update && \
    sudo apt-get install -y kubelet kubeadm kubectl docker.io
