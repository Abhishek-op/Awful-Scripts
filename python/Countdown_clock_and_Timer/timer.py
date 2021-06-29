import os
import time

os.system('clear')

# using ctime() to show present time
times = time.ctime()
print("\nCurrent Time: ", times)

print(
    "\n     Welcome to Awfulscripts\n"
)
hours = int(input("Enter hours? "))
minutes = int(input("Enter minutes? "))
seconds = int(input("Enter seconds? "))
hrsToSec = (hours * 60) * 60
mnsToSec = (minutes * 60)
seconds = seconds
seconds = hrsToSec + mnsToSec + seconds
print("\n Timer has been set for " + str(seconds) + " seconds.")
for i in range(seconds, -1, -1):
    displayHours = int(seconds / 3600)
    displayMinutes = int(seconds / 60)
    if displayMinutes >= 60:
        displayMinutes = displayMinutes - (displayHours * 60)
    else:
        displayMinutes = displayMinutes
    displaySeconds = int(seconds % 60)
    print("\n     Your time remaining is: {}:{}:{}".format(
        str(displayHours).zfill(2),
        str(displayMinutes).zfill(2),
        str(displaySeconds).zfill(2)))
    seconds -= 1
    time.sleep(1)

print("\n Time is ended.")
