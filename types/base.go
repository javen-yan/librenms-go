package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type (
	// Bool represents a boolean value, used for JSON marshaling. The API
	// returns some fields as 0/1 instead of true/false, so we use this custom type.
	Bool bool

	// Float64 represents a float64 value, used for JSON marshaling. The API
	// returns some fields as strings instead of numbers, so we use this custom type.
	Float64 float64

	// BaseResponse is the base structure for API responses.
	BaseResponse struct {
		// Status indicates the success or failure of the API call.
		Status string `json:"status"`
		// Message contains additional information about the API call.
		Message string `json:"message"`
		Count   int    `json:"count"`
	}
)

// MarshalJSON implements the JSON marshaling for the Bool type.
func (b *Bool) MarshalJSON() ([]byte, error) {
	if *b {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

// UnmarshalJSON implements the JSON unmarshalling for the Bool type.
func (b *Bool) UnmarshalJSON(data []byte) error {
	// attempt to unmarshal as a boolean first
	var valueBool bool
	if err := json.Unmarshal(data, &valueBool); err == nil {
		*b = Bool(valueBool)
		return nil
	}

	// if that fails, try to unmarshal as an integer (0 or 1)
	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("failed to unmarshal Bool: %w", err)
	}
	*b = value != 0
	return nil
}

// MarshalJSON implements the JSON marshaling for the Float64 type.
func (f *Float64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(*f), 'f', -1, 64)), nil
}

// UnmarshalJSON implements the JSON unmarshalling for the Float64 type.
func (f *Float64) UnmarshalJSON(data []byte) error {
	// attempt to unmarshal as a float64 first
	var valueFloat64 float64
	if err := json.Unmarshal(data, &valueFloat64); err == nil {
		*f = Float64(valueFloat64)
		return nil
	}

	// if that fails, try to unmarshal and parse as a string
	var valueString string
	if err := json.Unmarshal(data, &valueString); err != nil {
		return fmt.Errorf("failed to unmarshal Float64: %w", err)
	}

	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Float64 from string: %w", err)
	}

	*f = Float64(value)
	return nil
}
