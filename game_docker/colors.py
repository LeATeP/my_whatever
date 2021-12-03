#!/bin/env python3
from sys import argv
from os import environ


def fg_rgb(text: str, r: int = None, g: int = None, b: int = None, mode: str = None): 
    colors = {
    "black": "\u001b[30m",
    "red": f"\u001b[31m{text}\u001b[0m",
    "green": "\u001b[32m",
    "yellow": "\u001b[33m",
    "blue": "\u001b[34m",
    "magenta": "\u001b[35m",
    "cyan": "\u001b[36m",
    "white": "\u001b[37m",
    }
    if r or g or b:
        return f"\u001b[38;2;{r};{g};{b}m{text}\033[0m"
    if mode:
        return colors[mode]
    return None

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


