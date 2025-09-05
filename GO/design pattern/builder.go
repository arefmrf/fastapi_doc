package main

import (
	"fmt"
)

type Car struct {
	Brand  string
	Engine string
	Seats  int
	Color  string
}

// Display prints the car details.
func (c *Car) Display() {
	fmt.Printf("Car: %s, Engine: %s, Seats: %d, Color: %s\n", c.Brand, c.Engine, c.Seats, c.Color)
}

type CarBuilder struct {
	brand  string
	engine string
	seats  int
	color  string
}

// NewCarBuilder initializes a new CarBuilder.
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

// SetBrand sets the car brand.
func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
	cb.brand = brand
	return cb
}

// SetEngine sets the car engine type.
func (cb *CarBuilder) SetEngine(engine string) *CarBuilder {
	cb.engine = engine
	return cb
}

// SetSeats sets the number of seats.
func (cb *CarBuilder) SetSeats(seats int) *CarBuilder {
	cb.seats = seats
	return cb
}

// SetColor sets the car color.
func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.color = color
	return cb
}

// Build constructs the final Car object.
func (cb *CarBuilder) Build() *Car {
	return &Car{
		Brand:  cb.brand,
		Engine: cb.engine,
		Seats:  cb.seats,
		Color:  cb.color,
	}
}

func CarBuild() {
	// Using the builder to create a Car object
	car := NewCarBuilder().
		SetBrand("Tesla").
		SetEngine("Electric").
		SetSeats(4).
		SetColor("Red").
		Build()

	car.Display()
}

// second method using director

type Director struct {
	builder *CarBuilder
}

// NewDirector initializes a Director with a builder
func NewDirector(b *CarBuilder) *Director {
	return &Director{builder: b}
}

// BuildSportsCar creates a sports car with default values
func (d *Director) BuildSportsCar() *Car {
	return d.builder.SetBrand("Ferrari").
		SetEngine("V8").
		SetSeats(2).
		SetColor("Red").
		Build()
}

// BuildSUV creates an SUV with default values
func (d *Director) BuildSUV() *Car {
	return d.builder.SetBrand("Range Rover").
		SetEngine("V6").
		SetSeats(5).
		SetColor("Black").
		Build()
}

func DirectorCarBuild() {
	// Using the Director to create preset cars
	builder := NewCarBuilder()
	director := NewDirector(builder)

	sportsCar := director.BuildSportsCar()
	sportsCar.Display()

	suv := director.BuildSUV()
	suv.Display()
}
