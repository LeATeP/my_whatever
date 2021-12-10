#!/bin/env python3

class psql:
    def __init__(self, cursor, connection) -> None:
        self.cursor = cursor
        self.conn = connection
    
    def cmd(self, cmd):
        try:
            self.cursor.execute(cmd); self.conn.commit()

            fetch = self.cursor.fetchall()
            row_name = [n[0] for n in self.cursor.description]


            data = {}
            for row in fetch:
                data[row[0]] = {f"{row_name[0]}": row[0]}
                for i in range(len(row_name)):
                    data[row[0]].update({row_name[i]: row[i]})
            return data
        except Exception:
            return False
