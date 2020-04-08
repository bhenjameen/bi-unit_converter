package main

import (
	"fmt"
	"math"
)

//PI Formula
const PI = math.Pi

//Converter is a Bi-Directional unit converter
type Converter struct {
}

//Feet is a unit type
type Feet float64

//Centimeter is a unit type
type Centimeter float64

//Minute is a unit type
type Minute float64

//Second is a unit type
type Second float64

//Millisecond is a unit type
type Millisecond float64

//Celcius is a unit type
type Celcius float64

//Fahrenheit is a unit type
type Fahrenheit float64

//Radian is a unit type
type Radian float64

//Degree is a unit type
type Degree float64

//Kilogram is a unit type
type Kilogram float64

//Pound is a unit type
type Pound float64

//CentimetertoFeet conversion
func (cvr Converter) CentimetertoFeet(c Centimeter) Feet {
	return Feet(c / 30.48)
}

//FeettoCentimeter conversion
func (cvr Converter) FeettoCentimeter(f Feet) Centimeter {
	return Centimeter(f * 30.48)
}

//MinutetoSecond conversion
func (cvr Converter) MinutetoSecond(m Minute) Second {
	return Second(m * 60)
}

//SecondtoMinute conversion
func (cvr Converter) SecondtoMinute(s Second) Minute {
	return Minute(s / 60)
}

//SecondtoMillisecond conversion
func (cvr Converter) SecondtoMillisecond(s Second) Millisecond {
	return Millisecond(s * 1000)
}

//MillisecondtoSecond conversion
func (cvr Converter) MillisecondtoSecond(m Millisecond) Second {
	return Second(m / 1000)
}

//CelciustoFahrenheit conversion
func (cvr Converter) CelciustoFahrenheit(c Celcius) Fahrenheit { //(0°C × 9/5) + 32
	return Fahrenheit((c * (9.0 / 5)) + 32)
}

//FahrenheittoCelcius conversion
func (cvr Converter) FahrenheittoCelcius(f Fahrenheit) Celcius { //(0°F − 32) × 5/9
	return Celcius((f - 32) * (5.0 / 9))
}

//RadiantoDegree conversion
func (cvr Converter) RadiantoDegree(r Radian) Degree { //1rad × 180/π
	return Degree(r * (180 / PI))
}

//DegreetoRadian conversion
func (cvr Converter) DegreetoRadian(d Degree) Radian { //1° × π/180
	return Radian(d * (PI / 180))
}

//KilogramtoPound conversion
func (cvr Converter) KilogramtoPound(k Kilogram) Pound {
	return Pound(k * 2.205)
}

//PoundtoKilogram conversion
func (cvr Converter) PoundtoKilogram(p Pound) Kilogram {
	return Kilogram(p / 2.205)
}

func main() {
	converter := new(Converter)

	fmt.Printf("13.50 Centimeters to Feet is: ")
	fmt.Println(converter.CentimetertoFeet(13.50))

	fmt.Printf("10 Feet to Centimeter is: ")
	fmt.Println(converter.FeettoCentimeter(10))

	fmt.Printf("5 Minutes to Second is: ")
	fmt.Println(converter.MinutetoSecond(5))

	fmt.Printf("1000 Seconds to Minute is: ")
	fmt.Println(converter.SecondtoMinute(1000))

	fmt.Printf("300 Seconds to Millisecond is: ")
	fmt.Println(converter.SecondtoMillisecond(300))

	fmt.Printf("600000 Millisecond to Second is: ")
	fmt.Println(converter.MillisecondtoSecond(600000))

	fmt.Printf("45 Celcius to Fahrenheit is: ")
	fmt.Println(converter.CelciustoFahrenheit(45))

	fmt.Printf("90 Fahrenheit to Celcius is: ")
	fmt.Println(converter.FahrenheittoCelcius(90))

	fmt.Printf("4 Radians to Degree is: ")
	fmt.Println(converter.RadiantoDegree(4))

	fmt.Printf("90 Degrees to Radian is: ")
	fmt.Println(converter.DegreetoRadian(90))

	fmt.Printf("1000 Kilograms to Pound is: ")
	fmt.Println(converter.KilogramtoPound(1000))

	fmt.Printf("500 Pounds to Kilogram is: ")
	fmt.Println(converter.PoundtoKilogram(500))

}
