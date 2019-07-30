package main

import (
	"fmt"
	"github.com/jxy1996/study-go/structDemo"
)

func main() {
	// make some cars:
	ford := &structDemo.Car{Model: "Fiesta", BuildYear: 2008, Manufacturer: "Ford",}
	bmw := &structDemo.Car{Model: "XL 450", BuildYear: 2011, Manufacturer: "BMW"}
	merc := &structDemo.Car{Model: "D600", BuildYear: 2009, Manufacturer: "Mercedes"}
	bmw2 := &structDemo.Car{Model: "X 800", BuildYear: 2008, Manufacturer: "BMW"}

	cars := structDemo.Cars{ford, bmw, bmw2, merc}
	allBwmCar := cars.FindAll(func(car *structDemo.Car) bool {
		return car.BuildYear > 2010 && car.Manufacturer == "BMW"
	})

	fmt.Println("All bwm car:", allBwmCar)

}
