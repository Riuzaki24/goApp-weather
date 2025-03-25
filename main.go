package main

import (
	"flag"
	"fmt"

	"github.com/riuzaki24/app-5/geo"
	"github.com/riuzaki24/app-5/weather"
)

func main() {
	fmt.Println("Новый проект")

	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погода")

	flag.Parse()
	
	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
