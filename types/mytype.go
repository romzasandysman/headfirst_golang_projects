package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel.ToLiters(), busFuel.ToGallons())
	fmt.Println(carFuel.ToMilliliters())
}

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L\n", l)
}

func (l Liters) ToGallons() Gallons{
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons{
	return Gallons(m * 0.264)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL\n", m)
}

func (g Gallons) ToLiters() Liters{
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters{
	return Milliliters(g * 3785.41)
}

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal\n", g)
}