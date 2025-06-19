package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getFromApi(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	client := &http.Client{Timeout: time.Second * 10}
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getBounds(lat, lon float64) []byte {
	distance := 0.25
	
	url := fmt.Sprintf("https://api.waqi.info/v2/map/bounds?latlng=%v,%v,%v,%v&networks=all&token=%v",
		lat - distance, lon - distance, lat + distance, lon + distance,
		token)
	
	data, err := getFromApi(url)
	if err != nil {
		fmt.Printf("Err on request: %v", err)
		return nil
	}
	return data
}

func getAndParse(lat, lon float64, uid int) AQIResponse {
	
	url := fmt.Sprintf("https://api.waqi.info/feed/geo:%v;%v/?token=%v", lat, lon, token)
	if uid != 0 {
		url = fmt.Sprintf("https://api.waqi.info/feed/@%v/?token=%v", uid, token)
	}
	
	data, err := getFromApi(url)
	if err != nil {
		fmt.Printf("Err on request: %v", err)
	}
	//fmt.Printf("Result: %v \n", string(data) )
	response := AQIResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		fmt.Printf("Err on AQI response %v", err)
		return response
	}
	
	/**wr, err = getWeatherFromJson(data)
	if err != nil {
		fmt.Printf("Err on JSON: %v", err)
	}*/
	return response
}

