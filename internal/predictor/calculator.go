package predictor

import "math"

const (
	payloadPenaltyFactor = 0.05
	crosswindPenalty     = 0.10
)

func Calculate(input FlightInput) FlightResult {

	groundSpeed := calculateGroundSpeed(input)

	payloadFactor := 1.0 - (input.Payload * payloadPenaltyFactor)

	if payloadFactor < 0.55 {
		payloadFactor = 0.55
	}

	estimatedRange := input.MaxRange *
		(groundSpeed / input.CruiseSpeed) *
		payloadFactor

	if estimatedRange < 0 {
		estimatedRange = 0
	}

	var minutes float64

	if groundSpeed > 0 {
		minutes = estimatedRange / groundSpeed * 60
	}

	return FlightResult{
		GroundSpeed:      round(groundSpeed),
		EstimatedRange:   round(estimatedRange),
		EstimatedMinutes: round(minutes),
		Risk:             determineRisk(input, groundSpeed),
	}
}

func calculateGroundSpeed(input FlightInput) float64 {

	speed := input.CruiseSpeed

	switch input.Direction {

	case Headwind:
		speed -= input.WindSpeed

	case Tailwind:
		speed += input.WindSpeed * 0.70

	case Crosswind:
		speed -= input.WindSpeed * crosswindPenalty
	}

	if speed < 5 {
		speed = 5
	}

	return speed
}

func determineRisk(input FlightInput, ground float64) string {

	if input.WindSpeed >= 30 {
		return "HIGH"
	}

	if input.Payload >= 4 {
		return "HIGH"
	}

	if ground < input.CruiseSpeed*0.60 {
		return "MEDIUM"
	}

	if input.Direction == Crosswind && input.WindSpeed > 20 {
		return "MEDIUM"
	}

	return "LOW"
}

func round(v float64) float64 {
	return math.Round(v*10) / 10
}
