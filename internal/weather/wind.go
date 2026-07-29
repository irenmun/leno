package weather

import "github.com/irenmun/wind-flight-predictor/internal/predictor"

func Describe(direction predictor.WindDirection) string {

	switch direction {

	case predictor.Headwind:
		return "Headwind"

	case predictor.Tailwind:
		return "Tailwind"

	case predictor.Crosswind:
		return "Crosswind"

	default:
		return "Unknown"
	}
}

func Severity(speed float64) string {

	switch {

	case speed < 8:
		return "Light"

	case speed < 18:
		return "Moderate"

	case speed < 28:
		return "Strong"

	default:
		return "Extreme"
	}
}
