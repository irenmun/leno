package tests

import (
	"testing"

	"github.com/irenmun/wind-flight-predictor/internal/predictor"
)

func TestHeadwindPrediction(t *testing.T) {

	input := predictor.FlightInput{
		CruiseSpeed: 60,
		WindSpeed:   20,
		Direction:   predictor.Headwind,
		Payload:     2,
		MaxRange:    50,
	}

	result := predictor.Calculate(input)

	if result.GroundSpeed != 40 {
		t.Errorf("expected ground speed 40, got %.1f", result.GroundSpeed)
	}

	if result.EstimatedRange <= 0 {
		t.Error("estimated range should be greater than zero")
	}
}

func TestTailwindPrediction(t *testing.T) {

	input := predictor.FlightInput{
		CruiseSpeed: 60,
		WindSpeed:   10,
		Direction:   predictor.Tailwind,
		Payload:     1,
		MaxRange:    45,
	}

	result := predictor.Calculate(input)

	if result.GroundSpeed <= input.CruiseSpeed {
		t.Error("tailwind should increase ground speed")
	}
}

func TestValidation(t *testing.T) {

	input := predictor.FlightInput{
		CruiseSpeed: -5,
		WindSpeed:   10,
		Direction:   predictor.Headwind,
		Payload:     2,
		MaxRange:    50,
	}

	if predictor.Validate(input) == nil {
		t.Error("expected validation error")
	}
}

func TestHighRisk(t *testing.T) {

	input := predictor.FlightInput{
		CruiseSpeed: 55,
		WindSpeed:   35,
		Direction:   predictor.Headwind,
		Payload:     1,
		MaxRange:    40,
	}

	result := predictor.Calculate(input)

	if result.Risk != "HIGH" {
		t.Errorf("expected HIGH risk, got %s", result.Risk)
	}
}
