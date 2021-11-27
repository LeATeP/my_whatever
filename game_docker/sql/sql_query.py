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
        
    def get_value(self, selects: str, table: str, item_name: str = None) -> dict:
        
        if item_name:
            command = f'''
            select {selects} from {table}
            where name = {item_name};'''
        else:
            command = f'''
            select {selects} from {table};'''
            
        cursor = self.conn.cursor()
        cursor.execute(command)
        
        data = {};
        for n in list(cursor.fetchall()):    
            data[n[1]] = n[2] # "iron_chunk: 0"
            
        cursor.close()
        return data
        
    def update_value(self, 
                    column: str,
                    item_name: str, 
                    operator: str = '+', 
                    amount: str = '0') -> None:
        command = f'''
        update items
        set {column} = {column} {operator} {amount}
        where name = '{item_name}';
        '''
 
        cursor = self.conn.cursor()
        cursor.execute(command); self.conn.commit()
        cursor.close()


