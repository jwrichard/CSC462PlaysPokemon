#-------------------------------------------------------------------------------
# Name:        CSCPlaysPokemon
# Purpose:     Reads in moves from the server and plays the game
# Author:      Justin Richard
# Created:     31/07/2016
#-------------------------------------------------------------------------------
# IP: 54.200.67.158

import random
import time
import win32api
import win32con
import win32com

# Up - W
# Down - S
# Left - A
# Right - D
# A - Z
# B - X
# Select - C
# Start V
def main():
    # Activate pokemon window
    while True:
        move = getMove()
        if move == 'up':
            ms = 0x57
        if move == 'down':
            ms = 0x53
        if move == 'left':
            ms = 0x41
        if move == 'right':
            ms = 0x44
        if move == 'start':
            ms = 0x56
        if move == 'select':
            ms = 0x43
        if move == 'a':
            ms = 0x5A
        if move == 'b':
            ms = 0x58
        win32api.keybd_event(ms, 0)
        time.sleep(0.05)
        win32api.keybd_event(ms, 0 , 0x0002 , 0)

        time.sleep(1)

# Gets a random move for testing
def getMove():
    return random.choice(['up','down','left','right','up','down','left','right','up','down','left','right','start','select','a','b'])

# Gets a random move from the server - TODO
def getMoveFromServer():
    return "up"

if __name__ == '__main__':
    main()
