package models

// Star represents a stellar object with basic physical properties
type Star struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Mass        float64 `json:"mass"`   // in solar masses
	Radius      float64 `json:"radius"` // in solar radii
	Temperature float64 `json:"temperature"` // in Kelvin
}

// OrbitParams describes the orbital characteristics of a companion star
type OrbitParams struct {
	SemimajorAxis float64 `json:"semimajor_axis"` // in AU
	Eccentricity  float64 `json:"eccentricity"`
	Inclination   float64 `json:"inclination"`    // in degrees
	ArgumentOfPeriapsis float64 `json:"argument_of_periapsis"` // in degrees
	LongitudeOfAscendingNode float64 `json:"longitude_of_ascending_node"` // in degrees
	Period        float64 `json:"period"`         // in years
}

// Observation represents a single astronomical observation of stellar activity
type Observation struct {
	Timestamp   int64   `json:"timestamp"`      // Unix timestamp
	StarID      string  `json:"star_id"`
	ActivityLevel float64 `json:"activity_level"` // normalized measure of atmospheric disturbance
	WakePattern []float64 `json:"wake_pattern"`   // array of positional measurements
}