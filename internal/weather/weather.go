package weather

type Weather struct {
	Temperature float64
	Condition   string
}

// TODO: Implement weather service
func GetWeather(lat, lon float64) (*Weather, error) {
	// Placeholder for weather API implementation
	return &Weather{
		Temperature: 20.5,
		Condition:   "Sunny",
	}, nil
}