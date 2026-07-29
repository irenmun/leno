package predictor

import (
	"errors"
)

func Validate(input FlightInput) error {

	if input.CruiseSpeed <= 0 {
		return errors.New("invalid cruise speed")
	}

	if input.MaxRange <= 0 {
		return errors.New("invalid maximum range")
	}

	if input.WindSpeed < 0 {
		return errors.New("invalid wind speed")
	}

	if input.Payload < 0 {
		return errors.New("invalid payload weight")
	}

	if input.Direction < Headwind || input.Direction > Crosswind {
		return errors.New("invalid wind direction")
	}

	return nil
}
