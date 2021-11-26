#!/bin/env python3
from psycopg2 import connect, DatabaseError
from os import environ 
from sql.sql_query import psql
from time import sleep
from sys import argv

env = environ.get


def conn():
    conn = None
    
    try:
        print('Connecting to postgres server')
        connection = connect(host=str(env("POSTGRES_HOST")), 
                       database=str(env("POSTGRES_DB")), 
                       user=str(env("POSTGRES_USER")), 
                       password=str(env("POSTGRES_PASSWORD")))
        
        sql = psql(connection)
        get, update = sql.get_value, sql.update_value
        
        while True:
            update(column="amount", 
                   item_name=f"{argv[1]}", 
                   operator="+", amount="1")
            sleep(5)
        
        
    except (Exception, DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()
            print('Database connection closed.')
    

if __name__ == "__main__":
    conn()
        






