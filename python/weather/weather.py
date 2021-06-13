import requests


def main():
	print('[--]Enter city name:')
	a = input()
	hurl="https://api.openweathermap.org/data/2.5/weather?q="
	b = "&appid="
	token = "f9485f4a8083d0333d16c46db5c31e32"
	url = hurl+a+b+token
	response = requests.get(url)
	x=response.json()
	if x["cod"]!="404":
	           u=x["coord"]
	           lon=u["lon"]
	           lat=u["lat"]
	           
	           y=x["main"]
	           current_temp = y["temp"]
	           current_hum = y["humidity"]
	           pressure = y["pressure"]
	           sea_level =y["sea_level"]
	           
	           z = x["weather"]
	           weather_description = z[0]["description"]
	           
	           w = x['wind']
	           speed = w['speed']
	           degree = w['deg']
	           
	           s = x['sys']
	           con=s['country']
	           sunrise = s["sunrise"]
	           sunset = s["sunset"]
	           print(
	           '',
        '-----------------------------------------------------------------------',
	           '%-12s => %24s' % ("[-]Country", con),
                '%-12s => %24s' % ("[-]lon", lon),
        '%-12s => %24s' % ("[-]Lat", lat),
        '%-12s => %24s' % ("[-]Sunrise", sunrise),
	           '%-12s => %24s' % ("[-]Sunset", sunset),
                
        '%-12s => %24s' % ("[-]Temperature(Kelvin)", current_temp),
                '%-12s => %24s' % ("[-]Humidity", current_hum),
        '%-12s => %24s' % ("[-]Wind Speed(km/hour)", speed),
        '%-12s => %24s' % ("[-]Degree", degree),
        '%-12s => %24s' % ("[-]Pressure", pressure),
        '%-12s => %24s' % ("[-]Sea level", sea_level),
        '%-12s => %24s' % ("[-]Weather Discription", weather_description),
        
               sep="\n"
               )
            
if __name__ == '__main__':
    main()            		