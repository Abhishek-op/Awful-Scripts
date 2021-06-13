from requests import get

from bs4 import BeautifulSoup
def main():
	print("\033[1;95;40m \033")
	print("""
	 _     _       _                                  
| |   | |     | |                                 
| |   | | ____| |                                 
| |   | |/ ___) |                                 
| |___| | |   | |                                 
 \______|_|   |_|                                 
                                                  
    _    _                                        
   | |  | |                _                      
    \ \ | | _   ___   ____| |_  ____   ____  ____ 
     \ \| || \ / _ \ / ___)  _)|  _ \ / _  )/ ___)
 _____) ) | | | |_| | |   | |__| | | ( (/ /| |    
(______/|_| |_|\___/|_|    \___)_| |_|\____)_|    
                                                  """)
	url='https://tinyurl.com/api-create.php?url='
	print('[---]Made by Abhi')
	print("")
	
	print('[-]Enter url:')
	print("")
	print("")
	a=input()
	print("")
	print("")
	main_url=url+a
	res = get(main_url)
	print('[-]Main URL')
	print("")
	print(a)
	print("")
	print("")
	print('[-]Shorted URL')
	print("\033[1;32;40m \033")
	print(res.text)
if __name__ == '__main__':
    main()