# GoClimaWear

GoClimaWear is a simple Go application that helps you decide what clothes to wear based on the current temperature and humidity for private family purposes (very handy with a ssh shortcut automation on your phone). Content from other projects can be piped.

## Project Description

This project takes in two parameters - temperature and humidity, and based on these parameters, it suggests the type of clothes that would be suitable to wear. This can be particularly useful for planning your day or for packing for a trip.

## How to Use
*Option 1*
1. Input the current temperature and humidity. e.g. `go run main.go -temperature 25 -humidity 10`

*Option 2*
1. content also can be piped from my other c# [project](https://github.com/phpanhey/ClimaCoreFindorff): `echo "22 55" | go run main.go` => `./ClimaCoreFindorff | ./GoClimaWear`

*The application will output a suggestion for what to wear.*


## Development

This project is developed using Go to evaluate the language features. 

## License

[MIT](https://choosealicense.com/licenses/mit/)
