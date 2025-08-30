package main

import (
	"fmt"
	"math"
)

func main() {
	// Typed variable
	var gravityFormula = "Gmm/r^2"
	var gravityConstant float64 = 6.67430e-11 // in m^3 kg^-1 s^-2
	const massOfEarth float64 = 5.972e24      // in kg
	const radiusOfEarth float64 = 6.371e6     // in meters
	const massOfMars float64 = 6.39e23        // in kg
	const radiusOfMars float64 = 3.3895e6     // in meters
	const massOfMercury float64 = 3.3011e23   // in kg
	const radiusOfMercury float64 = 2.4397e6  // in meters
	const massOfVenus float64 = 4.8675e24     // in kg
	const radiusOfVenus float64 = 6.0518e6    // in meters
	const massOfJupiter float64 = 1.898e27    // in kg
	const radiusOfJupiter float64 = 6.9911e7  // in meters
	const massOfSaturn float64 = 5.683e26     // in kg
	const radiusOfSaturn float64 = 5.8232e7   // in meters
	const massOfUranus float64 = 8.681e25     // in kg
	const radiusOfUranus float64 = 2.5362e7   // in meters
	const massOfNeptune float64 = 1.024e26    // in kg
	const radiusOfNeptune float64 = 2.4622e7  // in meters
	const massOfPluto float64 = 1.303e22      // in kg
	const radiusOfPluto float64 = 1.1883e6    // in meters
	var tenthRadiusEarthGravity float64
	var tenthRadiusMarsGravity float64
	var planetNames [8]string

	// Untyped variable
	const speedOfLight = 299792458                   // in meters per second
	const distanceFromSunToEarth = 149600000000      // in meters
	const distanceFromMarsToEarth = 225000000000     // in meters
	const distanceFromMercuryToEarth = 91700000000   // in meters
	const distanceFromVenusToEarth = 41400000000     // in meters
	const distanceFromJupiterToEarth = 628730000000  // in meters
	const distanceFromSaturnToEarth = 1275000000000  // in meters
	const distanceFromUranusToEarth = 2723950000000  // in meters
	const distanceFromNeptuneToEarth = 4351400000000 // in meters
	const distanceFromPlutoToEarth = 5906380000000   // in meters

	// Undeclared variable
	relativeDistanceMercuryMars := distanceFromMarsToEarth / distanceFromMercuryToEarth

	fmt.Println("Gravity Formula:", gravityFormula)
	fmt.Println("Gravity Constant:", gravityConstant)
	fmt.Println("Speed of Light:", speedOfLight)

	fmt.Println("Distance from Sun to Earth:", distanceFromSunToEarth)
	fmt.Println("Distance from Mars to Earth:", distanceFromMarsToEarth)
	fmt.Println("Distance from Mercury to Earth:", distanceFromMercuryToEarth)
	fmt.Println("Distance from Venus to Earth:", distanceFromVenusToEarth)
	fmt.Println("Distance from Jupiter to Earth:", distanceFromJupiterToEarth)
	fmt.Println("Distance from Saturn to Earth:", distanceFromSaturnToEarth)
	fmt.Println("Distance from Uranus to Earth:", distanceFromUranusToEarth)
	fmt.Println("Distance from Neptune to Earth:", distanceFromNeptuneToEarth)
	fmt.Println("Distance from Pluto to Earth:", distanceFromPlutoToEarth)
	fmt.Println("Relative Distance from Earth to Mercury and Mars:", relativeDistanceMercuryMars)
	fmt.Println("Sin of Distance from Sun to Earth:", math.Sin(float64(distanceFromSunToEarth)))

	// Interface Types - We can pass functions as variables / Lambda Functions
	gravityFunction := func(planetMass float64, planetRadius float64) float64 {
		return (gravityConstant * planetMass) / (planetRadius * planetRadius)
	}

	// For Loops
	for i := 1; i <= 10; i++ {
		radius := float64(i) * radiusOfEarth
		if i == 1 {
			fmt.Println("\nVarious Gravity Figures of Earth")
		}
		fmt.Println(i, " times the radius of earth: ", float64(i)*radiusOfEarth, " meters")
		gravity := gravityFunction(massOfEarth, radius)
		fmt.Println("Gravity on earth at radius", radius, "gravity:", gravity, "m/s^2")
		if i == 10 {
			tenthRadiusEarthGravity = gravity
		}
	}

	// If and Else Statements
	for i := range 11 {
		radius := float64(i) * radiusOfMars
		if i == 1 {
			fmt.Println("\nVarious Gravity Figures of Mars")
		} else if i == 0 {
			continue
		}
		fmt.Println(i, " times the radius of Mars: ", float64(i)*radiusOfMars, " meters")
		gravity := gravityFunction(massOfMars, radius)
		fmt.Println("Gravity on mars at radius", radius, "gravity:", gravity, "m/s^2")
		if i == 10 {
			tenthRadiusMarsGravity = gravity
		}
	}

	// Switch
	switch math.Ceil(tenthRadiusEarthGravity + tenthRadiusMarsGravity) {
	case 1:
		fmt.Println("\nThe sum of gravity on 10th radius of Earth and Mars is approximately 1 m/s^2")
	case 2:
		fmt.Println("\nThe sum of gravity on 10th radius of Earth and Mars is approximately 2 m/s^2")
	default:
		fmt.Println("\nThe sum of gravity on 10th radius of Earth and Mars is neither 1 nor 2 m/s^2")
	}

	// Arrays

	planetNames = [8]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	var planetGravities [8]float64
	for i, planet := range planetNames {
		fmt.Println(i+1, ":", planet)
		var gravity float64
		switch planet {
		case "Mercury":
			gravity = gravityFunction(massOfMercury, radiusOfMercury)
		case "Venus":
			gravity = gravityFunction(massOfVenus, radiusOfVenus)
		case "Earth":
			gravity = gravityFunction(massOfEarth, radiusOfEarth)
		case "Mars":
			gravity = gravityFunction(massOfMars, radiusOfMars)
		case "Jupiter":
			gravity = gravityFunction(massOfJupiter, radiusOfJupiter)
		case "Saturn":
			gravity = gravityFunction(massOfSaturn, radiusOfSaturn)
		case "Uranus":
			gravity = gravityFunction(massOfUranus, radiusOfUranus)
		case "Neptune":
			gravity = gravityFunction(massOfNeptune, radiusOfNeptune)
		case "Pluto":
			gravity = gravityFunction(massOfPluto, radiusOfPluto)
		}
		planetGravities[i] = gravity
		fmt.Println("Gravity on", planet, "is:", gravity, "m/s^2")
	}
	fmt.Println("\nPlanet Gravities Array:", planetGravities)
}
