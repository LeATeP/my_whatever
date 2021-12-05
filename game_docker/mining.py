#!/bin/env python3
from os import environ 
from time import sleep

env = environ.get
from hub.wrap import init_conn


def conn(cmd):
    
   while True:
       cmd['update'](item_name="Rock",
                     amount="1")
       sleep(5)
    
if __name__ == "__main__":
    init_conn(conn)
        






