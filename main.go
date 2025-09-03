package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var token string = ""

var serverPort int = 8070
var brojac int = 0
var wg sync.WaitGroup

func main() {
	
	data, err := os.ReadFile("key")
	if err != nil {
		fmt.Printf("Err while getting API key...  [%v] \n", err)
		return
	}
	token = string(data)
	token = strings.TrimSpace(token)
	if token == "" {
		fmt.Printf("No API key found... \n")
		return
	}
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	//mux.HandleFunc("GET /stop", handleStop)
	mux.HandleFunc("GET /aqi_info", handleAqiInfo)
	mux.HandleFunc("GET /list_locations", handleListLocations)
	
	serverStr := fmt.Sprintf(":%v", serverPort)
	fmt.Printf("Hello Sun , listening %v", serverStr)
	
	wg.Add(1)
	go http.ListenAndServe(serverStr, mux)
	wg.Wait()
	// http.ListenAndServe(server_str, mux)
}

func handleAqiInfo(w http.ResponseWriter, r *http.Request) {
	
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		fmt.Printf("Err parsing latitude: %v", err)
	}
	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		fmt.Printf("Err parsing longitude: %v", err)
	}
	uid, err := strconv.ParseInt(r.URL.Query().Get("uid"), 10, 64)
	if err != nil {
		fmt.Printf("Err parsing UID: %v", err)
	}
	response := getAndParse(lat, lon, int(uid))
	airResponse := AirResponse{}
	airResponse.Status.Error = ""
	airResponse.Status.Token = "asdfg"
	
	if len(response.Data.City.Geo) == 2 {
		airResponse.City.Lat = response.Data.City.Geo[0]
		airResponse.City.Lon = response.Data.City.Geo[1]
	}
	airResponse.City.Name = response.Data.City.Name
	airResponse.City.URL = response.Data.City.URL
	airResponse.City.Location = response.Data.City.Location
	airResponse.City.Idx = response.Data.Idx
	
	airResponse.Values.PM10 = response.Data.IAQI.PM10.V
	airResponse.Values.PM25 = response.Data.IAQI.PM25.V
	
	w.Header().Set("Content-Type", "application/json")
	
	data, err := json.Marshal(airResponse)
	if err != nil {
		fmt.Println("Error while making config json. ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleListLocations(w http.ResponseWriter, r *http.Request) {
	// https://api.waqi.info/v2/map/bounds?latlng=39.379436,116.091230,40.235643,116.784382&networks=all&token=demo
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		fmt.Printf("Err parsing latitude: %v", err)
	}
	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		fmt.Printf("Err parsing longitude: %v", err)
	}
	
	//fmt.Fprint(w, getBounds(lat, lon))
	w.Header().Set("Content-Type", "application/json")
	w.Write(getBounds(lat, lon))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Sun [%v]", r.RemoteAddr)
	addrArr := strings.Split(r.RemoteAddr, ":")
	if len(addrArr) > 0 {
		fmt.Fprintf(w, "IP: %v", addrArr[0])
	}
}


