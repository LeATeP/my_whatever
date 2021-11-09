#!/usr/bin/env bash



for i in $(seq 30 37); do
	printf "\e[${i}m%s\n" "hello";  
	done			
# '\e' stand for color, '[' for value 'm' close

	


# q1() {
# for i in *; do 
#   echo "$i"; done
#}
#
#q1 | nl | lolcat
