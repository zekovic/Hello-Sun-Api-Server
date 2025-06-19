package main

type AQIResponse struct {
	Status string `json:"status"`
	Data   struct {
		AQI         int    `json:"aqi"`
		Idx       int32    `json:"idx"`
		Attributions []struct {
			URL  string `json:"url"`
			Name string `json:"name"`
			Logo string `json:"logo,omitempty"`
		} `json:"attributions"`
		City struct {
			Geo      []float64 `json:"geo"`
			Name     string    `json:"name"`
			URL      string    `json:"url"`
			Location string    `json:"location"`
		} `json:"city"`
		Dominentpol string `json:"dominentpol"`
		IAQI        struct {
			CO struct {
				V float64    `json:"v"`
			}    `json:"co"`
			DEW struct {
				V float64    `json:"v"`
			}    `json:"dew"`
			H  struct {
				V float64    `json:"v"`
			}    `json:"h"`
			NO2 struct {
				V float64    `json:"v"`
			}    `json:"no2"`
			O3 struct {
				V float64    `json:"v"`
			}    `json:"o3"`
			P  struct {
				V float64    `json:"v"`
			}    `json:"p"`
			PM10 struct {
				V float64    `json:"v"`
			}    `json:"pm10"`
			PM25 struct {
				V float64    `json:"v"`
			}    `json:"pm25"`
			SO2 struct {
				V float64    `json:"v"`
			}    `json:"so2"`
			T  struct {
				V float64    `json:"v"`
			}    `json:"t"`
			W  struct {
				V float64    `json:"v"`
			}    `json:"w"`
			WG struct {
				V float64    `json:"v"`
			}    `json:"wg"`
		} `json:"iaqi"`
		Time struct {
			S   string `json:"s"`
			TZ  string `json:"tz"`
			V   int64  `json:"v"`
			ISO string `json:"iso"`
		} `json:"time"`
		// Forecast struct {
		// 	Daily map[string][]struct {
		// 		Avg int    `json:"avg"`
		// 		Day string `json:"day"`
		// 		Max int    `json:"max"`
		// 		Min int    `json:"min"`
		// 	} `json:"daily"`
		// } `json:"forecast"`
		// Debug struct {
		// 	Sync string `json:"sync"`
		// } `json:"debug"`
	} `json:"data"`
}


type BoundsResponse struct {
	Status string   `json:"status"`
	Data []struct {
		Lat float64   `json:"lat"`
		Lon float64   `json:"lon"`
		UID int32     `json:"uid"`
		AQI string    `json:"aqi"`
		Station struct {
			Name string   `json:"name"`
			Time string   `json:"time"`
		}   `json:"station"`
	}   `json:"data"`
}

type AirResponse struct {
	Status struct {
		Token string   `json:"token"`
		Error string   `json:"error"`
	}   `json:"status"`
	City struct {
		Lat      float64   `json:"lat"`
		Lon     float64   `json:"long"`
		Name     string    `json:"name"`
		URL      string    `json:"url"`
		Location string    `json:"location"`
		Idx       int32    `json:"idx"`
	} `json:"city"`
	Values struct {
		PM10   float64   `json:"pm10"`
		PM25   float64   `json:"pm25"`
	}   `json:"values"`
}

type LocationResponse []struct {
	
}




