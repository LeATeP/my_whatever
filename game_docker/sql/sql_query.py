#!/bin/env python3

create_inv = """create table inv (
id serial primary key,
iron_chunk int,
copper_chunk int)"""
create_building_list = ""

class psql:
    def create_table(conn):
        cur = conn.cursor()
        cur.execute(create_inv)
        cur.close()
    
        conn.commit()	
        
    def get_value(connection, command):
        cur = connection.cursor()
        cur.execute(command)
        x = cur.fetchone()
        print(x, type(x))
        
        cur.close()



