case $1 in
-h)
	echo "Made by Abhishek"
	echo "Usage: weather [options]"
	echo "-h              Show the help screen"
	echo "-l [location]   Enter your location"
  echo "Example: ./weather.sh -l India"
	;;
-l)
	curl https://wttr.in/$2
	;;
*)
	echo("go to help[./weather.sh -h]")
	;;
esac