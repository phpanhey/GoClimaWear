package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	temperature, humidity := getTemperatureAndHumidity()

	fmt.Println("Temperature:", temperature)
	fmt.Println("Humidity:", humidity)

}

func getTemperatureAndHumidity() (float64, float64) {
	tempStr := flag.String("temperature", "20", "The current temperature")
	humidityStr := flag.String("humidity", "50", "The current humidity")
	flag.Parse()
	temp, err := strconv.ParseFloat(*tempStr, 64)
	if err != nil {
		fmt.Println("Error parsing temperature:", err)

	}
	humidity, err := strconv.ParseFloat(*humidityStr, 64)
	if err != nil {
		fmt.Println("Error parsing humidity:", err)

	}
	return temp, humidity
}
