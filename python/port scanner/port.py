__auther__ ="Abhishek"

import socket
import time
import threading
from queue import Queue


socket.setdefaulttimeout(0.25)
#avoide multiple modification
print_lock = threading.Lock()

print("""
  _                        
 |_) _  ._ _|_             
 |  (_) |   |_             
                           
  __                       
 (_   _  _. ._  ._   _  ._ 
 __) (_ (_| | | | | (/_ |  
                           
          """)
          
target = input('Enter the host to be scanned: ')
#Hostname to ip convertion
IP = socket.gethostbyname(target)
print ('Starting scan on host: ', IP)

def portscan(port):
   s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
   try:
      con = s.connect((IP, port))
      
      name = socket.getservbyport(port, "tcp") #port name
      with print_lock:
         print('PortNumber:', port, name, 'is open')
      con.close()
   except:
      pass

def threader():
   while True:
      worker = q.get()
      portscan(worker)
      q.task_done()
      
q = Queue()
startTime = time.time()
   
for x in range(100):
   t = threading.Thread(target = threader)
   t.daemon = True
   t.start()
   
for worker in range(1, 500):
   q.put(worker)
   
q.join()
print('Time taken:', time.time() - startTime)