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

# Python 2.x
import urllib2

# Python 3.x 
#import urllib.request

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
        move = getMoveFromServer()
        ms = 0x00;
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
        if move != 0x00:
			win32api.keybd_event(ms, 0)
			time.sleep(0.05)
			win32api.keybd_event(ms, 0 , 0x0002 , 0)
        time.sleep(5)

# Gets a random move for testing
def getMove():
    return random.choice(['up','down','left','right','up','down','left','right','up','down','left','right','start','select','a','b'])

# Gets a random move from the server - TODO
def getMoveFromServer():
	return urllib2.urlopen("http://54.213.250.201/move").read() # Python 2.x
    # return urllib.request.urlopen("http://54.213.250.201/move").read() # Python 3.x

if __name__ == '__main__':
    main()
