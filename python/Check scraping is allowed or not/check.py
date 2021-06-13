import requests
site = input("Enter website url: ")
response = requests.get(site)
print(response.text)
if (response.status_code == 200):
    response = requests.get(site + '/admin.php')
    if (response.status_code == 200):
        print("Vulnerable")
    else:
        print("Non-vulnerable")
else:
    print("Failure")