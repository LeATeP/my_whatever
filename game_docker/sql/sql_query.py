#!/bin/env python3

create_inv = """create table inv (
id serial primary key,
iron_chunk int,
copper_chunk int)"""
create_building_list = ""

item_list = ['Iron_chunk', 'Copper_chunk', 'Gold_coin']


class psql:
    def __init__(self, cursor, connection) -> None:
        self.cursor = cursor
        self.conn = connection
    
    def create_table(self):   
        self.cursor.execute(create_inv); self.conn.commit()
    

    def insert_new(self, table: str, row: str, name: str, returning: str):
        command = f'''
        insert into {table} ({row})
        values ('{name}')
        returning {returning};'''
        
        
        self.cursor.execute(command); self.conn.commit()
        result = self.cursor.fetchall()
        
        
        data = {}
        rows = [n[0] for n in self.cursor.description]
        for i in range(len(row) + 1):
            data[rows[i]] = result[0][i]
        
        return data
        
    def get_value(self, table: str, select: str = "*", values: str = None) -> dict:
        if values is not None:
            command = f'''
            select {select} from {table}
            values ({values});'''
        else:
            command = f'''
            select {select} from {table};'''
        
        self.cursor.execute(command)
        
        data = {}
        for n in list(self.cursor.fetchall()):    
            data[n[1]] = n[2] # "iron_chunk: 0"
            
        return data
        
    def update_value(self,
                    item_name: str,
                    table: str = "items",
                    column: str = "amount",
                    operator: str = '+', 
                    amount: str = '0') -> None:
        command = f'''
        update {table}
        set {column} = {column} {operator} {amount}
        where name = '{item_name}';
        '''
 
        self.cursor.execute(command); self.conn.commit()
        
        return True


