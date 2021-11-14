#!/usr/bin/env bash


# 256 colors
for i in $(seq 0 255); do
	printf "\e[1;38;5;${i}m%s" "hello $i"
done	# '\e[' initiate.. 1 for font style... 
		# '38' for color mode.. 
		# '5' for 255 / 2 for R;G;B
		# '%s' for display.. + '\n'
		

# q1() {
# for i in *; do 
#   echo "$i"; done
#}
#
#q1 | nl | lolcat
