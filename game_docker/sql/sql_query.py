#!/bin/env python3

from psycopg2 import connect, DatabaseError
from os import environ
from sys import path

env = environ.get

class psql:
    def __init__(self) -> None:
        self.conn = connect(host=str(env("POSTGRES_HOST")), 
                           database=str(env("POSTGRES_DB")), 
                           user=str(env("POSTGRES_USER")), 
                           password=str(env("POSTGRES_PASSWORD")))
        self.curs = self.conn.cursor()
        
    def cmd(self, cmd):
        try:
            self.curs.execute(cmd)
            self.conn.commit()

            fetch = self.curs.fetchall()
            row_name = [n[0] for n in self.curs.description]

            data = {}
            for i in range(len(fetch)):
                data[i] = dict(zip(row_name, fetch[i]))
            
            return data
           

        except (Exception, DatabaseError) as error:
            print(error)