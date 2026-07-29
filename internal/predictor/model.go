package predictor

type WindDirection int

const (
	Headwind WindDirection = iota
	Tailwind
	Crosswind
)

type FlightInput struct {
	CruiseSpeed float64
	WindSpeed   float64
	Direction   WindDirection
	Payload     float64
	MaxRange    float64
}

type FlightResult struct {
	GroundSpeed     float64
	EstimatedRange  float64
	EstimatedMinutes float64
	Risk            string
}
