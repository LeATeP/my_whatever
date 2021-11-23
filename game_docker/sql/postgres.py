#!/bin/env python3
from psycopg2 import connect, DatabaseError
from os import environ 
from sql_query import psql

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
        get = sql.get_value
        inventory = get(command="select * from inv;")        
        [print(n, inventory[n]) for n in inventory] # "1 {'iron_chunk': 0}"
        
        
        # cur = conn.cursor()
        
        # print('Postgres db version:')
        # cur.execute('SELECT version()')
        
        # db_version = cur.fetchone()
        # print(db_version)
        
        # cur.close()
    except (Exception, DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()
            print('Database connection closed.')
    

if __name__ == "__main__":
    conn()
        






