#!/bin/bash

sudo apt install vim && \
    sudo apt install git 

mkdir code
(cd code && git clone https://github.com/LeATeP/my_whatever)

cp /home/code/my_whatever/.bashrc ~/
cp /home/code/my_whatever/.bash-aliases ~/
cp /home/code/my_whatever/.vimrc ~/
