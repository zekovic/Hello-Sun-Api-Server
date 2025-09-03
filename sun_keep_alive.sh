if ! ps aux | grep "[s]un_api_server" >/dev/null; then
	./sun_api_server &
	echo "$(date +'%Y-%m-%d %H:%M:%S') Restarting sun_api..." >> sun.log
fi
