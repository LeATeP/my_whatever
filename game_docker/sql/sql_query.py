#!/bin/env python3

create_inv = """create table inv (
id serial primary key,
iron_chunk int,
copper_chunk int)"""
create_building_list = ""

item_list = ['Iron_chunk', 'Copper_chunk', 'Gold_coin']


class psql:
    def __init__(self, connection) -> None:
        self.conn = connection
    
    def create_table(self):
        cursor = self.conn.cursor()
        cursor.execute(create_inv)
        cursor.close()
    
        self.conn.commit()	
        
    def get_value(self, command) -> dict:
        cursor = self.conn.cursor()
        cursor.execute(command)
        
        data = {};
        for n in list(cursor.fetchall()):    
            data[n[0]] = {n[1]: n[2]} # "1 {'iron_chunk': 0}"
            
        cursor.close()
        return data
        
    def update_value(self, command):
        pass


