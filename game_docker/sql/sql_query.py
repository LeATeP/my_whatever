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

    def insert_new(self, table: str, row: str, race_name: str):
        command = f'''
        insert into {table} ({row})
        values ('{race_name}')
        returning id, race, grade, xp, level;'''
        
        cursor = self.conn.cursor()
        cursor.execute(command)
        result = cursor.fetchall()
        
        data = {}
        rows = [n[0] for n in cursor.description]
        for i in range(len(row) + 1):
            data[rows[i]] = result[0][i]
        
        cursor.close() 
        
        self.conn.commit()
        return data
        
    def get_value(self, table: str) -> dict:
        

        command = f'''
        select * from {table};'''
            
        cursor = self.conn.cursor()
        cursor.execute(command)
        
        data = {}
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


