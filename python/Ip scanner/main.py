import requests


def findip():
    print("[--]Enter Hostname")
    IP = input()
    try:
        hurl = 'http://ip-api.com/json/'
        url = hurl+IP
        response = requests.get(url)
        return response.json()
    except Exception as e:
        print("[-]Error: check your net connection")
        return {}


def main():
    
    data = findip()
    print('[-]Fetching details about hostname')
    print(
        '%-12s => %24s' % ("[-]GLOBAL IP", data.get("query")),
                '%-12s => %24s' % ("[-]ISP", data.get("isp")),
        '%-12s => %24s' % ("[-]ORG", data.get("org")),
        '%-12s => %24s' % ("[-]CITY", data.get("city")),
        '%-12s => %24s' % ("[-]Zip code", data.get("zip")),
        '%-12s => %24s' % ("[-]COUNTRY", data.get("country")),
        '%-12s => %24s' % ("[-]TIMEZONE", data.get("timezone")),
        '%-12s => %24s' % ("[-]LATITUDE", data.get("lat")),
        '%-12s => %24s' % ("[-]LONGITUDE", data.get("lon")),
        '%-12s => %24s' % ("[-]TIMEZONE", data.get("timezone")),
        '%-12s => %24s' % ("[-]ORG", data.get("org")),
        '%-12s => %24s' % ("[-]As", data.get("as")),

        sep="\n"
    )
    


if __name__ == '__main__':
    main()