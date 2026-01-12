package visualization

import (
	"fmt"
	"os"
	"path/filepath"

	"../simulation"
)

// PlotOrbit plots the orbital path of a celestial body
func PlotOrbit(orbit simulation.OrbitParams) error {
	// Create output directory if it doesn't exist
	outputDir := "plots"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate plot filename
	filename := filepath.Join(outputDir, fmt.Sprintf("orbit_%s.png", orbit.Name))

	// In a real implementation, this would use a plotting library like gonum/plot
	// For now, we'll just create a placeholder file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create plot file: %w", err)
	}
	defer file.Close()

	// Placeholder content - in reality this would generate actual plot data
	fmt.Fprintf(file, "Plot of orbit for %s\n", orbit.Name)
	fmt.Fprintf(file, "Semi-major axis: %.2f AU\n", orbit.SemiMajorAxis)
	fmt.Fprintf(file, "Eccentricity: %.2f\n", orbit.Eccentricity)
	fmt.Fprintf(file, "Period: %.2f days\n", orbit.Period)

	return nil
}

// PlotWake plots the wake pattern data
func PlotWake(wakeData []float64) error {
	// Create output directory if it doesn't exist
	outputDir := "plots"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate plot filename
	filename := filepath.Join(outputDir, "wake_pattern.png")

	// In a real implementation, this would use a plotting library like gonum/plot
	// For now, we'll just create a placeholder file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create plot file: %w", err)
	}
	defer file.Close()

	// Placeholder content - in reality this would generate actual plot data
	fmt.Fprintf(file, "Wake pattern analysis\n")
	fmt.Fprintf(file, "Data points: %d\n", len(wakeData))
	fmt.Fprintf(file, "Max amplitude: %.2f\n", max(wakeData))
	fmt.Fprintf(file, "Min amplitude: %.2f\n", min(wakeData))

	return nil
}

// Helper function to find maximum value
func max(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// Helper function to find minimum value
func min(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	minVal := data[0]
	for _, v := range data {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}
