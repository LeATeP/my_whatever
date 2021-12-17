#!/bin/env python3
from os import environ
from sys import path
from random import randint
from time import sleep

env = environ.get
path.append(env("GAME")) # my_whatever/game_docker
from sql.sql_query import psql
from colors import fg_rgb

class mining:
    def __init__(self) -> None:
        self.worker_ids = []
        # self.user_config = {}
        # self.tick_time = 1 if self.user_config['-t'] is None else self.user_config['-t']
    
    def start_work(self):
        try:
            cmd = psql().cmd
            data = cmd('''select * from worker
                       where working = false limit 1;''')
            
            # if no worker are available then break
            if not data:
                raise Exception('no worker are Free')
            
            # change worker status on True
            for worker in data.values():
                id = worker['id']
                cmd(f'''update worker set working = true
                        where id = {id} returning working;''')
                self.worker_ids.append(id) # add worker to pull
                
            
            good = True    
            while good:
                sleep(0.1)
                drop = self.gen_drop()
                for item in drop:
                    amount = cmd(f'''update items
                    set amount = amount + {drop[item]}
                    where name = '{item}' returning amount;''')[0]['amount']
                    s = f"{item}: {amount}"
                    print(fg_rgb(s))

                
                
                
        except (Exception) as error:
            print(error.args)
        finally:
            for id in self.worker_ids:
                result = cmd(f'''update worker
                        set working = false
                        where id = {id}
                        returning working;''')
        
    def gen_drop(self):
        drop =  {'Rock': randint(1,2) if randint(1, 10) > 5 else None,
                'Energy rock': 1 if randint(1, 10) > 8 else None,
                'Iron ore': 1 if randint(1, 100) > 95 else None}

        drop = {n: drop[n] for n in drop if drop[n] is not None}
        return drop
        


if __name__ == "__main__":
    mining().start_work()
        






