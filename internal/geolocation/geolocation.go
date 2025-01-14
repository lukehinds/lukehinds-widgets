package geolocation

type Location struct {
	City    string
	Country string
	Lat     float64
	Lon     float64
}

// TODO: Implement IP geolocation service
func GetLocationByIP(ip string) (*Location, error) {
	// Placeholder for IP geolocation implementation
	return &Location{
		City:    "London",
		Country: "UK",
		Lat:     51.5074,
		Lon:     -0.1278,
	}, nil
}