PROJECT_NAME: BetelgeuseCompanionTracker

# BetelgeuseCompanionTracker

A Go-based astronomical analysis tool that simulates and visualizes the orbital dynamics of Betelgeuse's hidden companion star, Siwarha, based on observed atmospheric disturbances.

## Description

This project addresses the challenge of analyzing complex stellar interactions by creating a simulation framework that models how a companion star affects its giant star host's atmospheric structure. Inspired by the discovery that Betelgeuse's mysterious dimming and surface disturbances were caused by a hidden companion creating observable gas wakes, this tool helps astronomers understand the orbital mechanics and interaction patterns between binary star systems.

The application processes observational data to detect and visualize the characteristic "wake" patterns that reveal the presence and orbit of invisible companions, providing insights into stellar evolution and binary system dynamics.

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/BetelgeuseCompanionTracker.git

# Navigate to project directory
cd BetelgeuseCompanionTracker

# Install dependencies
go mod tidy

# Build the project
go build -o betelgeuse-tracker main.go
```

## Usage

```bash
# Run the main application with sample data
./betelgeuse-tracker

# Analyze specific observational data files
./betelgeuse-tracker -input data/betelgeuse_observation.csv -output results/simulation.json

# Visualize the orbital wake pattern
./betelgeuse-tracker -visualize -orbit 3.84 -distance 1.2

# Run simulation with custom parameters
./betelgeuse-tracker -mass 0.8 -radius 1200 -period 2.5 -angle 45
```

## Features

- **Orbital Dynamics Simulation**: Models the gravitational interaction between binary stars
- **Atmospheric Wake Detection**: Identifies characteristic gas trail patterns in stellar atmospheres
- **Data Visualization**: Generates plots showing companion star trajectories and wake formation
- **Observational Analysis**: Processes real astronomical data to detect hidden companions
- **Parameter Estimation**: Calculates orbital parameters from observed atmospheric disturbances

## Example Output

The tool generates JSON reports showing:
- Companion star orbit parameters (semi-major axis, eccentricity)
- Wake pattern characteristics
- Atmospheric density distributions
- Time-series predictions of stellar brightness variations

## Requirements

- Go 1.19 or higher
- Standard library only (no external dependencies required)

## License

MIT License - see LICENSE file for details.