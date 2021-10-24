#!/bin/env python3

from sys import argv
from random import randint



def f(arg):
    y = arg[1]
    x = str(y) + '-' + str(randint(1000,9999))
    return x
    


x = f(argv)
print(x)