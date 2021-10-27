#!/usr/bin/env bash

q1() {
for i in *; do 
   echo "$i"; done
}

q1 | nl | lolcat