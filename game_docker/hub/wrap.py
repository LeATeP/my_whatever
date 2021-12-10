#!/bin/env python3

from psycopg2 import connect, DatabaseError
from os import environ
from sys import path

env = environ.get


path.append(env("GAME")) # my_whatever/game_docker
from sql.sql_query import psql

class init_conn():
    def __init__(self, fn) -> None:     
        try: 
            connection = connect(host=str(env("POSTGRES_HOST")), 
                               database=str(env("POSTGRES_DB")), 
                               user=str(env("POSTGRES_USER")), 
                               password=str(env("POSTGRES_PASSWORD")))
            cursor = connection.cursor()
            sql = psql(cursor, connection)
            
            
            # #####
            fn(sql.cmd)
            # #####
            
            
            
            
            connection.commit()
            cursor.close()
            
        except (Exception, DatabaseError) as error:
            print(error)
        finally:
            if connection:
                connection.close()
                
        