package types

type (
	// Location represents a location in LibreNMS.
	Location struct {
		ID               int     `json:"id,omitempty"`
		FixedCoordinates Bool    `json:"fixed_coordinates,omitempty"`
		Latitude         Float64 `json:"lat,omitempty"`
		Longitude        Float64 `json:"lng,omitempty"`
		Name             string  `json:"location,omitempty"`
		Timestamp        string  `json:"timestamp,omitempty"`
	}

	// LocationCreateRequest represents the request payload for creating a location.
	LocationCreateRequest struct {
		Name             string  `json:"location,omitempty"`
		FixedCoordinates Bool    `json:"fixed_coordinates,omitempty"`
		Latitude         float64 `json:"lat,omitempty"`
		Longitude        float64 `json:"lng,omitempty"`
	}

	// LocationUpdateRequest represents the request payload for updating a location.
	//
	// Only set the field(s) you want to update. Trying to patch fields that have not changed will
	// result in an HTTP 500 error.
	LocationUpdateRequest struct {
		Name             string
		FixedCoordinates *bool
		Latitude         *float64
		Longitude        *float64
	}

	// LocationResponse represents a response containing a single location from the LibreNMS API.
	LocationResponse struct {
		Status   string   `json:"status,omitempty"`
		Location Location `json:"get_location,omitempty"`
	}

	// LocationsResponse represents a response containing a list of locations from the LibreNMS API.
	LocationsResponse struct {
		BaseResponse
		Locations []Location `json:"locations"`
	}
)

// NewLocationUpdateRequest creates a new, empty LocationUpdateRequest.
func NewLocationUpdateRequest() *LocationUpdateRequest {
	return &LocationUpdateRequest{}
}

// SetFixedCoordinates sets whether the location has fixed coordinates in the LocationUpdateRequest.
func (r *LocationUpdateRequest) SetFixedCoordinates(fixed bool) *LocationUpdateRequest {
	r.FixedCoordinates = &fixed
	return r
}

// SetLatitude sets the latitude of the location in the LocationUpdateRequest.
func (r *LocationUpdateRequest) SetLatitude(lat float64) *LocationUpdateRequest {
	r.Latitude = &lat
	return r
}

// SetLongitude sets the longitude of the location in the LocationUpdateRequest.
func (r *LocationUpdateRequest) SetLongitude(lng float64) *LocationUpdateRequest {
	r.Longitude = &lng
	return r
}

// SetName sets the name of the location in the LocationUpdateRequest.
func (r *LocationUpdateRequest) SetName(name string) *LocationUpdateRequest {
	r.Name = name
	return r
}

// payload converts the LocationUpdateRequest to a map for the API request.
func (r *LocationUpdateRequest) Payload() map[string]interface{} {
	payload := make(map[string]interface{})
	if r.Name != "" {
		payload["location"] = r.Name
	}
	if r.FixedCoordinates != nil {
		payload["fixed_coordinates"] = *r.FixedCoordinates
	}
	if r.Latitude != nil {
		payload["lat"] = *r.Latitude
	}
	if r.Longitude != nil {
		payload["lng"] = *r.Longitude
	}
	return payload
}
