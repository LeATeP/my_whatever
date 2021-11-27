# REQUIREMENTS 

- Linux
- Docker
- Python/bash

game_docker set up:
"docker build -t game ."
"run pdb create" # creating postgres db
"run create_hub" # main menu

there is no auto add tables or stuff... 
items.py display table "items" that include 'id', 'name', 'amount'

mining.py main worker container that edit/add into db