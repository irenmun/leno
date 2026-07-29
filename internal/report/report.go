package report

import (
	"fmt"
	"os"

	"github.com/irenmun/wind-flight-predictor/internal/predictor"
)

func Save(path string, input predictor.FlightInput, result predictor.FlightResult) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	fmt.Fprintln(file, "Wind Flight Predictor Report")
	fmt.Fprintln(file, "============================")
	fmt.Fprintln(file)

	fmt.Fprintf(file, "Cruise Speed : %.1f km/h\n", input.CruiseSpeed)
	fmt.Fprintf(file, "Wind Speed   : %.1f km/h\n", input.WindSpeed)
	fmt.Fprintf(file, "Payload      : %.1f kg\n", input.Payload)
	fmt.Fprintf(file, "Max Range    : %.1f km\n", input.MaxRange)

	fmt.Fprintln(file)

	fmt.Fprintf(file, "Ground Speed      : %.1f km/h\n", result.GroundSpeed)
	fmt.Fprintf(file, "Estimated Range   : %.1f km\n", result.EstimatedRange)
	fmt.Fprintf(file, "Estimated Minutes : %.1f\n", result.EstimatedMinutes)
	fmt.Fprintf(file, "Risk Level        : %s\n", result.Risk)

	return nil
}
