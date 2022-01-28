#!/bin/python3

import time

x = 0
start = time.time()
while x < 1000000:
    x += 1
    time.sleep(0.001)
    
finish = time.time()
print(finish - start)