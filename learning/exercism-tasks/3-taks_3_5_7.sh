#!/bin/env bash








# most efficient way to write division combination's like 3-5-7
# divide () {
#     x=""
#     (($1 % 3)) || x+="Pling"
#     (($1 % 5)) || x+="Plang"
#     (($1 % 7)) || x+="Plong"

#     if test "$x" != ""; then
#         echo "$x"; else 
#         echo "$1";
#     fi

# }
# divide "$@"


# # right way to use iter in dict
# y=( ["3"]="pling" ["5"]="plang" ["7"]="plong" ["1"]="plingplong")
# for i in "${!y[@]}"; do
#     echo "$i ${y[$i]}";
#     done
