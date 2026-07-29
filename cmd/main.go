package main

import (
	"fmt"

	"github.com/irenmun/wind-flight-predictor/internal/predictor"
)

func main() {

	input := predictor.FlightInput{
		CruiseSpeed: 60,
		WindSpeed:   18,
		Direction:   predictor.Headwind,
		Payload:     2.4,
		MaxRange:    42,
	}

	result := predictor.Calculate(input)

	fmt.Println("========== Wind Flight Predictor ==========")
	fmt.Printf("Ground Speed      : %.1f km/h\n", result.GroundSpeed)
	fmt.Printf("Estimated Range   : %.1f km\n", result.EstimatedRange)
	fmt.Printf("Estimated Time    : %.1f minutes\n", result.EstimatedMinutes)
	fmt.Printf("Risk Level        : %s\n", result.Risk)
}
