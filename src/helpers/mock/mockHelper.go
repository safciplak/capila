package mock

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
)

type jsonResponse struct {
	Data interface{} `json:"data"`
}

//nolint:govet // Skip small optimization check
type hotel struct {
	GUID              string `json:"guid"`
	SystemName        string `json:"systemName"`
	Code              string `json:"code"`
	Currency          string `json:"currency"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Image             string `json:"image"`
	Website           string `json:"website"`
	CheckInTime       string `json:"checkInTime"`
	PrepaidPercentage int    `json:"prepaidPercentage"`
	Stars             int    `json:"stars"`
	Address           struct {
		City                 string  `json:"city"`
		Street               string  `json:"street"`
		StreetNumber         string  `json:"streetNumber"`
		StreetNumberAddition string  `json:"streetNumberAddition"`
		PostalCode           string  `json:"postalCode"`
		CountryCode          string  `json:"countryCode"`
		Gps                  string  `json:"gps"`
		Latitude             float64 `json:"latitude"`
		Longitude            float64 `json:"longitude"`
	} `json:"address"`
}

// FullPath creates the fullPath with some defaults for the StructFromMockJSON func.
func FullPath(path, version, folder, targetFile string) string {
	return filepath.Join(path, "mock", version, folder, targetFile)
}

// StructFromMockJSON attempts to marshall a mock file for use in testing.
func StructFromMockJSON(fullPath string, targetStruct interface{}) error {
	if reflect.ValueOf(targetStruct).Kind() != reflect.Ptr {
		return errors.New("targetStruct should be a pointer")
	}

	var (
		content  []byte
		rawData  []byte
		err      error
		response = jsonResponse{}
	)

	content, err = os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &response)
	if err != nil {
		return err
	}

	rawData, err = json.Marshal(response.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawData, &targetStruct)
}
