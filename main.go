package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	temperature, humidity := getTemperatureAndHumidity()
	clothingType := PredictClothingType(temperature, humidity)
	fmt.Println("Temperature:", temperature, "°C")
	fmt.Print(clothingType)
}

func getTemperatureAndHumidity() (float64, float64) {
	tempStr := flag.String("temperature", "", "The current temperature")
	humidityStr := flag.String("humidity", "", "The current humidity")
	flag.Parse()

	// If flags are empty, try reading from stdin
	if *tempStr == "" || *humidityStr == "" {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			if len(parts) >= 2 {
				tempStr = &parts[0]
				humidityStr = &parts[1]
			}
		}
	}

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

func PredictClothingType(temperature, humidity float64) string {
	switch {
	case temperature > 25:
		return recommendForHotWeather(humidity)
	case temperature > 18:
		return recommendForWarmWeather(humidity)
	case temperature > 10:
		return recommendForCoolWeather(humidity)
	default:
		return recommendForColdWeather(humidity)
	}
}

func recommendForHotWeather(humidity float64) string {
	if humidity > 60 {
		return "🥵 Very light, breathable clothing due to high humidity"
	}
	return "☀️ Lightweight clothing"
}

func recommendForWarmWeather(humidity float64) string {
	if humidity > 70 {
		return "😅 Comfortable, moisture-wicking clothing due to moderate temperature and high humidity"
	}
	return "🌤 Comfortable clothing with an option to layer"
}

func recommendForCoolWeather(humidity float64) string {
	return "🍂 Layered clothing with a jacket"
}

func recommendForColdWeather(humidity float64) string {
	if humidity > 80 {
		return "☔️ Warm, waterproof clothing due to cold and high humidity"
	}
	return "❄️ Warm clothing with coat, hat, and gloves"
}
