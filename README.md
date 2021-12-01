# REQUIREMENTS 

- Linux
- Docker
- Python/bash

1. set up image "docker build -t game ."

2. create db container "run my_whatever/game_docker/pdb create" # creating postgres db

3. change project location in "create_hub" / "silence_container" to your project clone location and run docker container script "my_whatever/game_docker/create_hub"

4. to test that everything worked you can run "items.py" or "my_whatever/game_docker/hub/items.py"


# but it will display nothing cuz you have nothing in db


items.py display table "items" that include 'id', 'name', 'amount'

mining.py main worker file and container "silence container" work together that edit/add into db
