#!/bin/env bash
# right way to use iter in dict
y=( ["3"]="pling" ["5"]="plang" 
    ["7"]="plong" ["1"]="plingplong")
for i in "${!y[@]}"; do
    echo "$i ${y[$i]}";
    done
