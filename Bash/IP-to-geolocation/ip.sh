blue="\e[0;94m"
green="\e[0;92m"
bold="\e[1m"
if ! [ -x "$(command -v jq)" ]; then
  echo 'Error: jq ,https://stedolan.github.io/jq/download/'
  exit 1
fi

if [[ $# -ne 1 ]]; then
	echo 'Provide I.P as command line parameter. Usage:  ' $0 ' 15.45.0.1 '
	exit 1
fi
link=$(echo "http://ip-api.com/json/"$1)
data=$(curl $link -s) # -s for slient output

status=$(echo $data | jq '.status' -r)

if [[ $status == "success" ]]; then
	
	city=$(echo $data | jq '.city' -r)
	region=$(echo $data | jq '.regionName' -r)
	country=$( echo $data | jq '.country' -r)
  lat=$(echo $data | jq '.lat' -r)
  lon=$(echo $data | jq '.lon' -r)
  zip=$(echo $data | jq '.zip' -r)
  code=$(echo $data | jq '.countryCode' -r)
  timezone=$(echo $data | jq '.timezone' -r)
  isp=$(echo $data | jq '.isp' -r)
  org=$(echo $data | jq '.org' -r)
  as=$(echo $data | jq '.as' -r)
  echo -e "${blue}${bold}[x]Made by Abhishek-op"
  echo -e "${blue}${bold}[x]Starting scan for ip: "$1
  echo ""
	echo -e  "${blue}${bold}[^]IPinfo        Data"
  echo -e "${green}[-]City          "$city
  echo -e "${green}[-]zip           "$zip
  echo -e "${green}[-]Region        "$region
  echo -e "${green}[-]Country       "$country
  echo -e "${green}[-]Countrycode   "$cod
  echo -e "${green}[-]lon           "$lon
  echo -e "${green}[-]lag           "$lat
  echo -e "${green}[-]Timezone      "$timezone
  echo -e "${green}[-]ISP           "$isp
  echo -e "${green}[-]ORG           "$org
  echo -e "${green}[-]As            "$as
fi