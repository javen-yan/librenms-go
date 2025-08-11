package librenms

import (
	"fmt"
	"net/http"

	"github.com/javen-yan/librenms-go/types"
)

// Create creates a new location in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Locations/#add_location
func (l *LocationAPI) Create(location *types.LocationCreateRequest) (*types.BaseResponse, error) {
	c := l.client
	req, err := c.newRequest(http.MethodPost, "locations", location, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// Delete deletes a location by its ID in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Locations/#delete_location
func (l *LocationAPI) Delete(locationID int) (*types.BaseResponse, error) {
	c := l.client
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("locations/%d", locationID), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// Get retrieves a location by its ID from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Locations/#get_location
func (l *LocationAPI) Get(locationID int) (*types.LocationResponse, error) {
	c := l.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("location/%d", locationID), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.LocationResponse)
	return resp, c.do(req, resp)
}

// List retrieves a list of locations from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Locations/#list_locations
func (l *LocationAPI) List() (*types.LocationsResponse, error) {
	c := l.client
	req, err := c.newRequest(http.MethodGet, "resources/locations", nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.LocationsResponse)
	return resp, c.do(req, resp)
}

// Update updates a location by its ID in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Locations/#edit_location
func (l *LocationAPI) Update(locationID int, location *types.LocationUpdateRequest) (*types.BaseResponse, error) {
	c := l.client
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("locations/%d", locationID), location.Payload(), nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}
