# Architecture

## Overview

Wind Flight Predictor follows a layered architecture to keep prediction logic isolated from presentation and reporting.

```
cmd/
 └── main.go
        │
        ▼
 predictor
 ├── validator.go
 ├── calculator.go
 └── model.go
        │
        ├────────► weather
        │            └── wind.go
        │
        └────────► report
                     └── report.go
```

## Modules

### predictor

Contains the core business logic responsible for validating user input and calculating the predicted flight characteristics.

### weather

Provides helper functions for classifying wind direction and severity.

### report

Creates human-readable flight reports that can be stored or shared.

### configs

Stores configurable application parameters.

### tests

Contains unit tests covering the prediction engine.

## Design Goals

- Simple architecture
- Easily testable
- Minimal dependencies
- Clear separation of responsibilities
- Extensible prediction engine
