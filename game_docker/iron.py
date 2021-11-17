#!/bin/env python3

from time import sleep

iron_ore = 0
while True:
    sleep(1)
    iron_ore += 1
    print(f"You got {iron_ore}")
