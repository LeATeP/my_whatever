#!/bin/env python3
from psycopg2 import connect, DatabaseError
from os import environ 

env = environ.get

def conn():
    conn = None
    
    try:
        print('Connecting to postgres server')
        conn = connect(host=str(env("POSTGRES_HOST")), 
                       database=str(env("POSTGRES_DB")), 
                       user=str(env("POSTGRES_USER")), 
                       password=str(env("POSTGRES_PASSWORD")))
        
        
        cur = conn.cursor()
        
        print('Postgres db version:')
        cur.execute('SELECT version()')
        
        db_version = cur.fetchone()
        print(db_version)
        
        cur.close()
    except (Exception, DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()
            print('Database connection closed.')
    

conn()
        






