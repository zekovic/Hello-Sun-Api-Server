API server for retrieving air quality data from https://api.waqi.info and returning it to [Hello Sun](https://github.com/zekovic/Hello-Sun) app. 

To run the server, file named "key" should be created, with API token inside of it. 

Air quality API token can be acquired at: https://aqicn.org/data-platform/token/

#### List locations of air measuring stations, in area around coordinates:
```
http://<your-server>:8070/list_locations?lat=<latitude>&lon=<longitude>
```

#### Get air quality info of selected measuring station by it's UID:
```
http://<your-server>:8070/aqi_info?uid=<station_uid>
```

#### Get air quality info of selected measuring station if UID is missing:
```
http://<your-server>:8070/aqi_info?lat=<latitude>&lon=<longitude>
```
