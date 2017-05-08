package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("`curl http://localhost:8080` --> Hello from server")
	fmt.Println("`curl http://localhost:8080/weather/CITY_NAME` --> Weather now in CITY_NAME")
	fmt.Println("Ctrl+c to stop the server")

	http.HandleFunc("/", helloFunc) //Call to root path calls to helloFunc
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		queryData, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		var clientData clientWeatherData
		clientData.Name = city
		clientData.Main.Kelvin = queryData.Main.Kelvin
		clientData.Main.Degrees = clientData.Main.Kelvin - 273.15
		totalTime := time.Since(begin).String()
		clientData.QueryTime = totalTime

		json.NewEncoder(w).Encode(clientData)
	})

	http.ListenAndServe(":8080", nil)
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You asked for nothing! Here you have a Hello World so. :D\n"))
}

func query(city string) (openWeatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=254502a503e64cd51c607d2107a912d8&q=" + city)
	if err != nil {
		return openWeatherData{}, err
	}

	defer resp.Body.Close()

	var d openWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return openWeatherData{}, err
	}

	return d, nil
}

type openWeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type clientWeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin  float64 `json:"kelvin"`
		Degrees float64 `json:"degrees"`
	} `json:"main"`
	QueryTime string `json:"query_time"`
}
