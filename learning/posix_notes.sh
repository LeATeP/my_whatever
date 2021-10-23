#!/bin/env sh
set -x



# third  10.23.2021

test_test() {
    test "$1" &&
    echo "good"
}

test_test "$@"



# second  10.21.2021

# if_statement() {
#     if [ "$#" = "1" ]; then   # or if test "$#" -e "1"; then
#         echo "$# good $1"
#         exit 1
#     fi
#     echo "no good"
# }
# if_statement "$@"



# First  10.20.2021

# function
# say_hello() {
#     echo "$#"
#     for i in "$@";
#     do
#         echo "$i";
#     done
# }
# say_hello "good"   # "$@"








