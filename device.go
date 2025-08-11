package librenms

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/javen-yan/librenms-go/types"
)

const (
	// deviceEndpoint is the API endpoint for devices.
	deviceEndpoint = "devices"
)

// Create creates a device by hostname/IP.
//
// Documentation: https://docs.librenms.org/API/Devices/#add_device
func (d *DeviceAPI) Create(payload *types.DeviceCreateRequest) (*types.DeviceResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/", deviceEndpoint), payload, nil)
	if err != nil {
		return nil, err
	}
	deviceResp := new(types.DeviceResponse)
	return deviceResp, c.do(req, deviceResp)
}

// Delete deletes a device by its ID or hostname from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Devices/#del_device
func (d *DeviceAPI) Delete(identifier string) (*types.DeviceResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/%s", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	deviceResp := new(types.DeviceResponse)
	return deviceResp, c.do(req, deviceResp)
}

// Get retrieves a device by its ID or hostname from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_device
func (d *DeviceAPI) Get(identifier string) (*types.DeviceResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	deviceResp := new(types.DeviceResponse)
	return deviceResp, c.do(req, deviceResp)
}

// List retrieves a list of devices from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Devices/#list_devices
func (d *DeviceAPI) List(query *types.DevicesQuery) (*types.DeviceResponse, error) {
	c := d.client
	params, err := parseParams(query)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodGet, deviceEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	deviceResp := new(types.DeviceResponse)
	return deviceResp, c.do(req, deviceResp)
}

// Update updates a device by its ID or hostname.
//
// Documentation: https://docs.librenms.org/API/Devices/#update_device_field
func (d *DeviceAPI) Update(identifier string, payload *types.DeviceUpdateRequest) (*types.BaseResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("%s/%s", deviceEndpoint, identifier), payload, nil)
	if err != nil {
		return nil, err
	}
	patchResp := new(types.BaseResponse)
	return patchResp, c.do(req, patchResp)
}

// Discover triggers a discovery of the given device.
//
// Documentation: https://docs.librenms.org/API/Devices/#discover_device
func (d *DeviceAPI) Discover(identifier string) (*types.BaseResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/discover", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// GetAvailability retrieves calculated availabilities of the given device.
//
// Documentation: https://docs.librenms.org/API/Devices/#availability
func (d *DeviceAPI) GetAvailability(identifier string) (*types.DeviceAvailabilityResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/availability", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceAvailabilityResponse)
	return resp, c.do(req, resp)
}

// GetOutages retrieves detected outages of the given device.
//
// Documentation: https://docs.librenms.org/API/Devices/#outages
func (d *DeviceAPI) GetOutages(identifier string) (*types.DeviceOutagesResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/outages", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceOutagesResponse)
	return resp, c.do(req, resp)
}

// GetGraphs retrieves a list of available graphs for a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_graphs
func (d *DeviceAPI) GetGraphs(identifier string) (*types.DeviceGraphsResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/graphs", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceGraphsResponse)
	return resp, c.do(req, resp)
}

// GetHealthGraphs retrieves health graphs for a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#list_available_health_graphs
func (d *DeviceAPI) GetHealthGraphs(identifier string, graphType string, sensorID string) (*types.DeviceGraphsResponse, error) {
	c := d.client
	var endpoint string
	if sensorID != "" {
		endpoint = fmt.Sprintf("%s/%s/health/%s/%s", deviceEndpoint, identifier, graphType, sensorID)
	} else if graphType != "" {
		endpoint = fmt.Sprintf("%s/%s/health/%s", deviceEndpoint, identifier, graphType)
	} else {
		endpoint = fmt.Sprintf("%s/%s/health", deviceEndpoint, identifier)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceGraphsResponse)
	return resp, c.do(req, resp)
}

// GetWirelessGraphs retrieves wireless graphs for a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#list_available_wireless_graphs
func (d *DeviceAPI) GetWirelessGraphs(identifier string, graphType string, sensorID string) (*types.DeviceGraphsResponse, error) {
	c := d.client
	var endpoint string
	if sensorID != "" {
		endpoint = fmt.Sprintf("%s/%s/wireless/%s/%s", deviceEndpoint, identifier, graphType, sensorID)
	} else if graphType != "" {
		endpoint = fmt.Sprintf("%s/%s/wireless/%s", deviceEndpoint, identifier, graphType)
	} else {
		endpoint = fmt.Sprintf("%s/%s/wireless", deviceEndpoint, identifier)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceGraphsResponse)
	return resp, c.do(req, resp)
}

// GetPorts retrieves a list of ports for a particular device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_port_graphs
func (d *DeviceAPI) GetPorts(identifier string, columns string) (*types.DevicePortsResponse, error) {
	c := d.client
	endpoint := fmt.Sprintf("%s/%s/ports", deviceEndpoint, identifier)

	var params *url.Values
	if columns != "" {
		params = &url.Values{}
		params.Set("columns", columns)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, params)
	if err != nil {
		return nil, err
	}
	resp := new(types.DevicePortsResponse)
	return resp, c.do(req, resp)
}

// GetDeviceFDB retrieves FDB entries associated with a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_device_fdb
func (d *DeviceAPI) GetDeviceFDB(identifier string) (*types.DeviceFDBResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/fdb", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceFDBResponse)
	return resp, c.do(req, resp)
}

// GetDeviceNAC retrieves NAC entries associated with a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_device_nac
func (d *DeviceAPI) GetDeviceNAC(identifier string) (*types.DeviceNACResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/nac", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceNACResponse)
	return resp, c.do(req, resp)
}

// GetDeviceIPAddresses retrieves IP addresses associated with a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_device_ip_addresses
func (d *DeviceAPI) GetDeviceIPAddresses(identifier string) (*types.DeviceIPAddressesResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/ip", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceIPAddressesResponse)
	return resp, c.do(req, resp)
}

// GetPortStack retrieves port mappings for a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_port_stack
func (d *DeviceAPI) GetPortStack(identifier string, validMappings bool) (*types.DevicePortStackResponse, error) {
	c := d.client
	endpoint := fmt.Sprintf("%s/%s/port_stack", deviceEndpoint, identifier)

	var params *url.Values
	if validMappings {
		params = &url.Values{}
		params.Set("valid_mappings", "")
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, params)
	if err != nil {
		return nil, err
	}
	resp := new(types.DevicePortStackResponse)
	return resp, c.do(req, resp)
}

// GetDeviceTransceivers retrieves transceivers associated with a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_device_transceivers
func (d *DeviceAPI) GetDeviceTransceivers(identifier string) (*types.DeviceTransceiversResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/transceivers", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceTransceiversResponse)
	return resp, c.do(req, resp)
}

// GetComponents retrieves components for a particular device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_components
func (d *DeviceAPI) GetComponents(identifier string, query *types.ComponentsQuery) (*types.DeviceComponentsResponse, error) {
	c := d.client
	endpoint := fmt.Sprintf("%s/%s/components", deviceEndpoint, identifier)

	params, err := parseParams(query)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest(http.MethodGet, endpoint, nil, params)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceComponentsResponse)
	return resp, c.do(req, resp)
}

// AddComponent creates a new component of a type on a particular device.
//
// Documentation: https://docs.librenms.org/API/Devices/#add_components
func (d *DeviceAPI) AddComponent(identifier string, componentType string) (*types.DeviceComponentsResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/%s/components/%s", deviceEndpoint, identifier, componentType), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceComponentsResponse)
	return resp, c.do(req, resp)
}

// EditComponents edits existing components on a particular device.
//
// Documentation: https://docs.librenms.org/API/Devices/#edit_components
func (d *DeviceAPI) EditComponents(identifier string, payload map[string]interface{}) (*types.BaseResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("%s/%s/components", deviceEndpoint, identifier), payload, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// DeleteComponent deletes an existing component on a particular device.
//
// Documentation: https://docs.librenms.org/API/Devices/#delete_components
func (d *DeviceAPI) DeleteComponent(identifier string, componentID string) (*types.BaseResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/%s/components/%s", deviceEndpoint, identifier, componentID), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// GetPortStats retrieves information about a particular port for a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_port_stats_by_port_hostname
func (d *DeviceAPI) GetPortStats(identifier string, ifName string, columns string) (*types.DevicePortStatsResponse, error) {
	c := d.client
	endpoint := fmt.Sprintf("%s/%s/ports/%s", deviceEndpoint, identifier, ifName)

	var params *url.Values
	if columns != "" {
		params = &url.Values{}
		params.Set("columns", columns)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, params)
	if err != nil {
		return nil, err
	}
	resp := new(types.DevicePortStatsResponse)
	return resp, c.do(req, resp)
}

// GetDeviceMaintenance retrieves the current maintenance status of a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#device_under_maintenance
func (d *DeviceAPI) GetDeviceMaintenance(identifier string) (*types.DeviceMaintenanceResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/maintenance", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceMaintenanceResponse)
	return resp, c.do(req, resp)
}

// SetDeviceMaintenance sets a device into maintenance mode.
//
// Documentation: https://docs.librenms.org/API/Devices/#maintenance_device
func (d *DeviceAPI) SetDeviceMaintenance(identifier string, payload *types.DeviceMaintenanceRequest) (*types.BaseResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/%s/maintenance", deviceEndpoint, identifier), payload, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// RenameDevice renames a device.
//
// Documentation: https://docs.librenms.org/API/Devices/#rename_device
func (d *DeviceAPI) RenameDevice(identifier string, newHostname string) (*types.BaseResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("%s/%s/rename/%s", deviceEndpoint, identifier, newHostname), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// GetDeviceGroups lists the device groups that a device is matched on.
//
// Documentation: https://docs.librenms.org/API/Devices/#get_device_groups
func (d *DeviceAPI) GetDeviceGroups(identifier string) (*types.DeviceGroupsResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s/groups", deviceEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.DeviceGroupsResponse)
	return resp, c.do(req, resp)
}

// UpdateDevicePortNotes updates a device port notes field.
//
// Documentation: https://docs.librenms.org/API/Devices/#update_device_port_notes
func (d *DeviceAPI) UpdateDevicePortNotes(identifier string, portID int, notes string) (*types.BaseResponse, error) {
	c := d.client
	payload := map[string]string{"notes": notes}
	req, err := c.newRequest(http.MethodPatch, fmt.Sprintf("%s/%s/port/%d", deviceEndpoint, identifier, portID), payload, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}
