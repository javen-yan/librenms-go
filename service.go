package librenms

import (
	"fmt"
	"net/http"

	"github.com/javen-yan/librenms-go/types"
)

const (
	serviceEndpoint = "services"
)

// Create creates a service for the specified device id or hostname.
//
// Documentation: https://docs.librenms.org/API/Services/#add_service_for_host
func (s *ServiceAPI) Create(deviceIdentifier string, service *types.ServiceCreateRequest) (*types.ServiceResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/%s", serviceEndpoint, deviceIdentifier), service, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.ServiceResponse)
	return resp, c.do(req, resp)
}

// Delete deletes a service by its ID.
//
// Documentation: https://docs.librenms.org/API/Services/#delete_service_from_host
func (s *ServiceAPI) Delete(serviceID int) (*types.BaseResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/%d", serviceEndpoint, serviceID), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// Get retrieves a service by ID from the LibreNMS API.
//
// Similar to GetDeviceGroup, this uses the same endpoint as GetServices, but it returns a
// modified payload with the single host (if a match is found).
// This is primarily a convenience function for the Terraform provider.
func (s *ServiceAPI) Get(serviceID int) (*types.ServiceResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodGet, serviceEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	internalResp := new(types.ServiceResponseMulti)
	if err = c.do(req, internalResp); err != nil {
		return nil, err
	}

	resp := &types.ServiceResponse{
		BaseResponse: internalResp.BaseResponse,
		Services:     internalResp.GetServices(),
	}

	if len(resp.Services) == 0 {
		return resp, nil
	}

	// look for a matching service by ID
	singleServiceResp := &types.ServiceResponse{
		Services: make([]types.Service, 0),
	}
	singleServiceResp.Message = resp.Message
	singleServiceResp.Status = resp.Status

	for _, service := range resp.Services {
		if service.ID == serviceID {
			singleServiceResp.Services = append(singleServiceResp.Services, service)
			singleServiceResp.Count = 1
			break
		}
	}

	return singleServiceResp, err
}

// List retrieves all services from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Services/#list_services
func (s *ServiceAPI) List() (*types.ServiceResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodGet, serviceEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	internalResp := new(types.ServiceResponseMulti)
	if err = c.do(req, internalResp); err != nil {
		return nil, err
	}

	services := internalResp.GetServices()
	return &types.ServiceResponse{
		BaseResponse: types.BaseResponse{
			Status:  internalResp.Status,
			Message: internalResp.Message,
			Count:   len(services),
		},
		Services: services,
	}, err
}

// GetForHost retrieves all services for a specific host by ID or name from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Services/#get_service_for_host
func (s *ServiceAPI) GetForHost(deviceIdentifier string) (*types.ServiceResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceEndpoint, deviceIdentifier), nil, nil)
	if err != nil {
		return nil, err
	}

	internalResp := new(types.ServiceResponseMulti)
	if err = c.do(req, internalResp); err != nil {
		return nil, err
	}

	services := internalResp.GetServices()
	return &types.ServiceResponse{
		BaseResponse: types.BaseResponse{
			Status:  internalResp.Status,
			Message: internalResp.Message,
			Count:   len(services),
		},
		Services: services,
	}, err
}

// Update updates a service for the specified service ID.
//
// Documentation: https://docs.librenms.org/API/Services/#edit_service_from_host
func (s *ServiceAPI) Update(serviceID int, service *types.ServiceUpdateRequest) (*types.ServiceResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("%s/%d", serviceEndpoint, serviceID), service.Payload(), nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.ServiceResponse)
	return resp, c.do(req, resp)
}
