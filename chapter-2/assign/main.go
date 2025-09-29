package main

var (
	maxSpeed         int     = 600
	maxSpinSpeed     int     = 400
	maxDarkThreshold int     = 500
	isActive         bool    = true
	isCalibrated     bool    = false
	maxRatio         float32 = 1.34e22
)

func main() {

	println("Max Speed:", maxSpeed)
	println("Max Spin Speed:", maxSpinSpeed)
	println("Max Dark Threshold:", maxDarkThreshold)
	println("Is Active:", isActive)
	println("Is Calibrated:", isCalibrated)
	println("Max Ratio:", maxRatio)
}
