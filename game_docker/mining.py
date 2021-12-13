#!/bin/env python3
from os import environ
from sys import path
from random import randint
from time import sleep

env = environ.get
path.append(env("GAME")) # my_whatever/game_docker
from hub.wrap import init_conn
from colors import fg_rgb
    
class mining:
    def __init__(self) -> None:
        pass
    
    def start_work(self, cmd):
        try:
            data = cmd('''select * from worker
                       where working = false limit 1;''')
            if data:
                key = self.ret_key(data)
                result = cmd(f'''update worker
                    set working = true
                    where id = {key}
                    returning working;''')
                result_confirmation = self.ret_key(result)
                
                ###
                
                try:
                    if result_confirmation:
                        while True:
                            sleep(1)
                            drop = self.gen_drop()
                            drop = {n: drop[n] for n in drop if drop[n] is not None}
                            if not drop:
                                continue
                            for n in drop:
                                command = f'''update items
                                set amount = amount + {drop[n]}
                                where name = '{n}'
                                returning amount;'''
                                x = cmd(command)
                        
                except Exception:
                    pass
                
                ###
            
        except Exception:
            pass
        finally:
            print('finally', str(key))
            result = cmd(f'''update worker
                    set working = false
                    where id = {key}
                    returning working;''')
        
    def gen_drop(self):
        drop =  {'Rock': randint(1,2) if randint(1, 10) > 5 else None,
                'Energy rock': 1 if randint(1, 10) > 8 else None,
                'Iron ore': 1 if randint(1, 100) > 95 else None}

        return drop
        
    
    def ret_key(self, data):
        key = [n for n in data][0]
        return key


if __name__ == "__main__":
    init_conn(mining().start_work)
        






