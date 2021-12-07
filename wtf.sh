#!/usr/bin/env bash

# for q in $(seq 0 15 255); do
#     for w in $(seq 0 15 255); do
#         for e in $(seq 0 15 255); do
#             printf "\e[38;2;${q};${w};${e}m%s $1"  
#             sleep .005  
#         done
#     done
# done


# 256 colors
# for i in $(seq 0 255); do
	#  printf "\e[1;38;5;${i}m%s\n" "hello $i"
# done	# '\e[' initiate.. 1 for font style... 
		# '38' for color mode.. 
		# '5' for 255 / 2 for R;G;B
		# '%s' for display.. + '\n'

		# terminal theme color '[30m .. [37m' /// to clean screen '\e[H\e[2J
    	# to reset colors '\e[0m'

	

# q1() {
# for i in *; do 
#   echo "$i"; done
#}
#
#q1 | nl | lolcat
