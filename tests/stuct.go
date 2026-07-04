package main

import "fmt"

type Vehicle interface {
	move()
	refuel(amount int, bank *FuelBank)
	status()
}

type FuelBank struct {
	overflow int
}

type Car struct {
	fuel int
}

type Plane struct {
	fuel int
}

type Ship struct {
	fuel int
}

//
// CAR
//
func (c *Car) move() {
	fmt.Println("Car is driving!")
	c.fuel -= 20
	if c.fuel < 0 {
		c.fuel = 0
	}
}

func (c *Car) refuel(amount int, bank *FuelBank) {
	maxFuel := 60

	if c.fuel+amount > maxFuel {
		overflow := c.fuel + amount - maxFuel
		c.fuel = maxFuel
		bank.overflow += overflow
		fmt.Println("Car is full! overflow -> bank =", overflow)
	} else {
		c.fuel += amount
	}
}

func (c *Car) status() {
	fmt.Println("Car fuel:", c.fuel)
}

//
// PLANE
//
func (p *Plane) move() {
	fmt.Println("Plane is flying!")
	p.fuel -= 50
	if p.fuel < 0 {
		p.fuel = 0
	}
}

func (p *Plane) refuel(amount int, bank *FuelBank) {
	maxFuel := 100

	if p.fuel+amount > maxFuel {
		overflow := p.fuel + amount - maxFuel
		p.fuel = maxFuel
		bank.overflow += overflow
		fmt.Println("Plane is full! overflow -> bank =", overflow)
	} else {
		p.fuel += amount
	}
}

func (p *Plane) status() {
	fmt.Println("Plane fuel:", p.fuel)
}

//
// SHIP
//
func (s *Ship) move() {
	fmt.Println("Ship is sailing!")
	s.fuel -= 50
	if s.fuel < 0 {
		s.fuel = 0
	}
}

func (s *Ship) refuel(amount int, bank *FuelBank) {
	maxFuel := 100

	if s.fuel+amount > maxFuel {
		overflow := s.fuel + amount - maxFuel
		s.fuel = maxFuel
		bank.overflow += overflow
		fmt.Println("Ship is full! overflow -> bank =", overflow)
	} else {
		s.fuel += amount
	}
}

func (s *Ship) status() {
	fmt.Println("Ship fuel:", s.fuel)
}

//
// MAIN
//
func main() {
	bank := &FuelBank{}

	car := &Car{fuel: 30}
	plane := &Plane{fuel: 20}
	ship := &Ship{fuel: 50}

	vehicles := []Vehicle{car, plane, ship}

	fmt.Println("=== CYCLE 1 ===")
	for _, v := range vehicles {
		v.move()
		v.refuel(50, bank)
		v.status()
	}

	fmt.Println("\nBank overflow:", bank.overflow)
}