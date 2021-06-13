import requests


def getLocationInfo():
    try:
        response = requests.get('http://ip-api.com/json')
        return response.json()
    except Exception as e:
        print("[-]Error: check your net connection")
        return {}


def main():
    print('[-]Fetching your location details based on your ip address')
    data = getLocationInfo()
    print(
        '%-12s => %24s' % ("[-]GLOBAL IP", data.get("query")),
                '%-12s => %24s' % ("[-]ISP", data.get("isp")),
        '%-12s => %24s' % ("[-]ORG", data.get("org")),
        '%-12s => %24s' % ("[-]CITY", data.get("city")),
        '%-12s => %24s' % ("[-]COUNTRY", data.get("country")),
        '%-12s => %24s' % ("[-]TIMEZONE", data.get("timezone")),
        '%-12s => %24s' % ("[-]LATITUDE", data.get("lat")),
        '%-12s => %24s' % ("[-]LONGITUDE", data.get("lon")),
        '%-12s => %24s' % ("[-]TIMEZONE", data.get("timezone")),
        sep="\n"
    )
    


if __name__ == '__main__':
    main()
