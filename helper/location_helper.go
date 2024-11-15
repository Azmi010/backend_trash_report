package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GeoCodeResponse struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetAddressFromCoordinates(lat, lng string) (string, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", lat, lng, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var geoCodeResponse GeoCodeResponse
	if err := json.Unmarshal(body, &geoCodeResponse); err != nil {
		return "", err
	}

	if geoCodeResponse.Status != "OK" {
		return "", fmt.Errorf("failed to fetch address: %s", geoCodeResponse.Status)
	}

	if len(geoCodeResponse.Results) == 0 {
		return "", fmt.Errorf("no results found for the provided coordinates")
	}

	return geoCodeResponse.Results[0].FormattedAddress, nil
}
