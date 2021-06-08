bold="\e[1m"
blue="\e[0;94m"
purple="\e[0;92m"
case $1 in
-h)
	echo -e "${blue}${purple}Made by Abhishek-op"
	echo -e "${blue}${bold}Usage:           [options]"
	echo -e "${blue}${bold}-h            Show the help screen"
	echo -e "${blue}${bold}-f              Enter filename"
  echo -e "${blue}${bold}Example: ./share.sh -f image.jpg"
	;;
-f)
  echo -e "${purple}${bold}Welcome to file-sharing script. made by Abhishek-op"
  echo ""
  echo -e "${purple}${bold}Here is your link${blue}"
	curl -F "file=@$2" https://file.io -s | sed 's/,/\n/g' | awk '/link/{gsub(/"/," ",$0); print $NF}'
	;;
*)
	echo -e "${blue}go to [./share.sh -h] help"
	;;
esac