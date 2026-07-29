# Wind Flight Predictor

Wind Flight Predictor is a command-line application written in Go that estimates the impact of wind conditions on the expected flight distance of a small unmanned aerial vehicle (UAV).

The program combines wind speed, wind direction, payload weight and aircraft cruise speed to produce a simplified flight prediction.

---

## Features

- Wind influence calculation
- Ground speed estimation
- Flight time estimation
- Distance prediction
- Payload penalty
- Flight risk assessment
- JSON configuration support
- Terminal-friendly output

---

## Project Structure

```
cmd/
internal/
configs/
docs/
tests/
```

---

## Input Parameters

- Cruise speed
- Wind speed
- Wind direction
- Payload weight
- Maximum range

---

## Example

Input

Cruise Speed: 60 km/h
Wind Speed: 18 km/h
Direction: Headwind
Payload: 2.4 kg
Maximum Range: 42 km

Output

```
Ground Speed : 42 km/h

Estimated Range : 29.4 km

Estimated Flight Time : 42 minutes

Risk Level : Medium
```

---

## Build

```bash
go build ./cmd
```

Run

```bash
go run ./cmd
```

---

## Future Improvements

- Live weather API integration

- Terrain profile correction

- Battery degradation model

- Crosswind compensation

- Interactive terminal dashboard

---

## License

MIT
