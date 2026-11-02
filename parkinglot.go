package main

type VehicleSize int

type Vehicle interface { //abstraction
	Size() VehicleSize
	LicensePlate() string
}

type spot struct {
	ID string
	Size VehicleSize
	Vehicle Vehicle
}

