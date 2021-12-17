#!/bin/env python3
from os import environ
from sys import path, argv, flags
from random import randint
from time import sleep

env = environ.get
path.append(env("GAME")) # my_whatever/game_docker
from sql.sql_query import psql
from colors import fg_rgb

class unit:
    def __init__(self, cmd) -> None:
        unit_data = cmd('''select id from worker limit 1;''')
        print(unit_data)

    
class battle:
    def __init__(self) -> None:
        self.cmd = psql().cmd
    
    def start_battle(self):
        unit1 = unit(self.cmd)



if __name__ == "__main__":
    battle().start_battle(argv)
        






