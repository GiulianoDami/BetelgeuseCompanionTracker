package simulation

import (
	"errors"
	"math"
)

// OrbitalSimulator simulates the orbital dynamics and atmospheric wake patterns
// of a binary star system involving Betelgeuse and its companion Siwarha
type OrbitalSimulator struct {
	// Simulation parameters
	OrbitalPeriod float64 // in Earth days
	SemimajorAxis float64 // in astronomical units
	Eccentricity  float64 // orbital eccentricity
	Inclination   float64 // orbital inclination in degrees
	ArgumentOfPeriapsis float64 // in degrees
	LongitudeOfAscendingNode float64 // in degrees
	
	// Time parameters
	StartTime     float64 // in days
	EndTime       float64 // in days
	TimeStep      float64 // in days
	
	// Atmospheric parameters
	WakeAmplitude float64 // amplitude of atmospheric disturbance
	WakeWidth     float64 // width of the wake pattern
}

// Run executes the orbital simulation and returns any errors encountered
func (os *OrbitalSimulator) Run() error {
	if os.OrbitalPeriod <= 0 {
		return errors.New("orbital period must be positive")
	}
	
	if os.SemimajorAxis <= 0 {
		return errors.New("semimajor axis must be positive")
	}
	
	if os.Eccentricity < 0 || os.Eccentricity >= 1 {
		return errors.New("eccentricity must be between 0 and 1")
	}
	
	if os.TimeStep <= 0 {
		return errors.New("time step must be positive")
	}
	
	if os.EndTime <= os.StartTime {
		return errors.New("end time must be greater than start time")
	}
	
	// Simulate orbital positions over time
	currentTime := os.StartTime
	for currentTime <= os.EndTime {
		// Calculate orbital position using Keplerian elements
		_, _ = os.calculateOrbitalPosition(currentTime)
		
		// Calculate atmospheric wake pattern
		_ = os.calculateWakePattern(currentTime)
		
		currentTime += os.TimeStep
	}
	
	return nil
}

// calculateOrbitalPosition computes the position of the companion star at a given time
func (os *OrbitalSimulator) calculateOrbitalPosition(time float64) (float64, float64) {
	// Simplified calculation using mean anomaly and eccentric anomaly
	meanAnomaly := 2 * math.Pi * time / os.OrbitalPeriod
	
	// Solve Kepler's equation using Newton-Raphson method (simplified)
	eccentricAnomaly := meanAnomaly
	for i := 0; i < 10; i++ {
		eccentricAnomaly = eccentricAnomaly - (eccentricAnomaly - os.Eccentricity*math.Sin(eccentricAnomaly) - meanAnomaly) / (1 - os.Eccentricity*math.Cos(eccentricAnomaly))
	}
	
	// Calculate true anomaly
	trueAnomaly := 2 * math.Atan(math.Sqrt((1+os.Eccentricity)/(1-os.Eccentricity)) * math.Tan(eccentricAnomaly/2))
	
	// Calculate distance from primary
	distance := os.SemimajorAxis * (1 - os.Eccentricity*os.Eccentricity) / (1 + os.Eccentricity*math.Cos(trueAnomaly))
	
	// Convert to Cartesian coordinates
	x := distance * math.Cos(trueAnomaly)
	y := distance * math.Sin(trueAnomaly)
	
	return x, y
}

// calculateWakePattern computes the atmospheric wake pattern at a given time
func (os *OrbitalSimulator) calculateWakePattern(time float64) float64 {
	// Get companion position
	x, y := os.calculateOrbitalPosition(time)
	
	// Calculate distance from primary (Betelgeuse)
	distance := math.Sqrt(x*x + y*y)
	
	// Simple model for wake amplitude based on distance and time
	wake := os.WakeAmplitude * math.Exp(-distance/os.WakeWidth)
	
	return wake
}