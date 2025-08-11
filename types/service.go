package types

type (
	// Service represents a service in LibreNMS.
	Service struct {
		ID          int    `json:"service_id"`
		Changed     int64  `json:"service_changed"`
		Description string `json:"service_desc"`
		DeviceID    int    `json:"device_id"`
		DS          string `json:"service_ds"`
		Ignore      Bool   `json:"service_ignore"`
		IP          string `json:"service_ip"`
		Message     string `json:"service_message"`
		Name        string `json:"service_name"`
		Param       string `json:"service_param"`
		Status      int    `json:"service_status"` // assuming this follows Nagios conventions, 0=ok, 1=warning, 2=critical, 3=unknown
		TemplateID  int    `json:"service_template_id"`
		Type        string `json:"service_type"`
	}

	// ServiceCreateRequest represents the request payload for creating a service.
	ServiceCreateRequest struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"desc,omitempty"`
		IP          string `json:"ip,omitempty"`
		Ignore      Bool   `json:"ignore,omitempty"`
		Param       string `json:"param,omitempty"`
		Type        string `json:"type"`
	}

	// ServiceUpdateRequest represents the request payload for updating a service.
	//
	// Only set the field(s) you want to update. Trying to patch fields that have not changed will
	// result in an HTTP 500 error.
	ServiceUpdateRequest struct {
		Name        *string
		Description *string
		IP          *string
		Ignore      *bool
		Param       *string
		Type        *string
	}

	// serviceResponse is the internal response structure for services.
	//
	// The raw response is returned as a list of service lists, but it seems that
	// all services are always returned in the first list. This also causes the count
	// to always reflect 1 in the response. So we're going to collapse this into
	// a 1-dimensional slice and update count for easier client handling.
	ServiceResponseMulti struct {
		BaseResponse
		Services [][]Service `json:"services"`
	}

	// ServiceResponse is the response structure for services.
	ServiceResponse struct {
		BaseResponse
		Services []Service `json:"services"`
	}
)

// getServices flattens the slice of slices into a single slice.
func (s *ServiceResponseMulti) GetServices() []Service {
	flatServices := make([]Service, 0)
	for _, serviceList := range s.Services {
		flatServices = append(flatServices, serviceList...)
	}
	return flatServices
}

// NewServiceUpdateRequest creates a new ServiceUpdateRequest instance.
func NewServiceUpdateRequest() *ServiceUpdateRequest {
	return &ServiceUpdateRequest{}
}

// SetDescription sets the description of the service in the update request.
func (r *ServiceUpdateRequest) SetDescription(description string) *ServiceUpdateRequest {
	r.Description = &description
	return r
}

// SetName sets the name of the service in the update request.
func (r *ServiceUpdateRequest) SetName(name string) *ServiceUpdateRequest {
	r.Name = &name
	return r
}

// SetIP sets the IP address of the service in the update request.
func (r *ServiceUpdateRequest) SetIP(ip string) *ServiceUpdateRequest {
	r.IP = &ip
	return r
}

// SetIgnore sets the ignore status of the service in the update request.
func (r *ServiceUpdateRequest) SetIgnore(ignore bool) *ServiceUpdateRequest {
	r.Ignore = &ignore
	return r
}

// SetParam sets the parameter of the service in the update request.
func (r *ServiceUpdateRequest) SetParam(param string) *ServiceUpdateRequest {
	r.Param = &param
	return r
}

// SetType sets the type of the service in the update request.
func (r *ServiceUpdateRequest) SetType(serviceType string) *ServiceUpdateRequest {
	r.Type = &serviceType
	return r
}

// payload generates the actual update payload for the request, only including fields that are not nil.
// This will allow us to send a partial list of fields and still be able to send 'empty' values (avoids
// `omitempty` issues with the JSON encoder).
func (r *ServiceUpdateRequest) Payload() map[string]interface{} {
	payload := make(map[string]interface{})
	if r.Name != nil {
		payload["service_name"] = *r.Name
	}
	if r.Description != nil {
		payload["service_desc"] = *r.Description
	}
	if r.IP != nil {
		payload["service_ip"] = *r.IP
	}
	if r.Ignore != nil {
		payload["service_ignore"] = func() int {
			if *r.Ignore {
				return 1
			}
			return 0
		}()
	}

	if r.Param != nil {
		payload["service_param"] = *r.Param
	}
	if r.Type != nil {
		payload["service_type"] = *r.Type
	}
	return payload
}
