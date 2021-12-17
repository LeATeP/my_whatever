#!/bin/env python3
from sys import argv
from os import environ
from random import choice, randint


def fg_rgb(text: str): 
    return f"\u001b[38;2;{randint(0, 255)};{randint(0, 255)};{randint(0,255)}m{text}\u001b[0m"


def bg_rgb(text: str, r: int = None, g: int = None, b: int = None, mode: str = None):
    colors = {
    "black": "\u001b[40m" + text + "\u001b[0m",
    "red": f"\u001b[41m{text}\u001b[0m",
    "green": "\u001b[42m",
    "yellow": "\u001b[43m",
    "blue": "\u001b[44m",
    "magenta": "\u001b[45m",
    "cyan": "\u001b[46m",
    "white": "\u001b[47m",
    }
    if r and g and b:
        return f"\u001b[48;2;{r};{g};{b}m{text}\033[0m"
    if mode:
        return colors[mode]
    return None
    
# x = "Hello world!"
# y = fg_rgb(x, mode="red")
# print(y)


