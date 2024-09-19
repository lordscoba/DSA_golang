package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)
// car = Drive(car)
// => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
func Drive(car Car) Car {

	newBattery := car.battery - car.batteryDrain

	if newBattery < 0 {
		return car

	} else {
		car.battery = newBattery
		car.distance = car.distance + car.speed

		return car
	}
}

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)

// distance := 100
// track := NewTrack(distance)

// CanFinish(car, track)
// => true
// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

	return car.battery >= car.batteryDrain*track.distance/car.speed
}
