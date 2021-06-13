import requests

def main():
  print("Welcome to awfulscripts")
  print("")
  a=input('Enter a complete url of site: ')
  try:
    res = requests.get(a+'/robots.txt')
    if res.status_code == 200:
      print("here is robote.txt file")
      print(res.text)
      print("")
      print("robot.txt file is saved")
      
      file=open("robot.txt", "w")
      file.write(res.text)
      file.close()
    else:
      print("Enter a valid url")
  except requests.exceptions.MissingSchema:
    print("Enter a valid url")
    print("Example:https://www.google.com")
		
if __name__ == '__main__':
	main()  		