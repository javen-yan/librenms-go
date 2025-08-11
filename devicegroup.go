package librenms

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/javen-yan/librenms-go/types"
)

const (
	// deviceGroupEndpoint is the API endpoint for devices.
	deviceGroupEndpoint = "devicegroups"
)

// Create creates a device group in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/DeviceGroups/#add_devicegroup
func (d *DeviceGroupAPI) Create(group *types.DeviceGroupCreateRequest) (*types.DeviceGroupCreateResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodPost, deviceGroupEndpoint, group, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.DeviceGroupCreateResponse)
	return resp, c.do(req, resp)
}

// Delete deletes a group by its ID or hostname from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/DeviceGroups/#delete_devicegroup
func (d *DeviceGroupAPI) Delete(identifier string) (*types.BaseResponse, error) {
	c := d.client
	uri, err := url.Parse(fmt.Sprintf("%s/%s", deviceGroupEndpoint, identifier))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URI: %w", err)
	}

	req, err := c.newRequest(http.MethodDelete, uri.String(), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// Get uses the same endpoint as GetDeviceGroups, but it returns a
// modified payload with the single host (if a match is found).
// This is primarily a convenience function for the Terraform provider.
func (d *DeviceGroupAPI) Get(identifier string) (*types.DeviceGroupResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, deviceGroupEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.DeviceGroupResponse)
	if err = c.do(req, resp); err != nil {
		return resp, err
	}

	if len(resp.Groups) == 0 {
		return resp, nil
	}

	singleGroupResp := &types.DeviceGroupResponse{
		Groups: make([]types.DeviceGroup, 0),
	}
	singleGroupResp.Message = resp.Message
	singleGroupResp.Status = resp.Status

	for _, group := range resp.Groups {
		if group.Name == identifier || strconv.Itoa(group.ID) == identifier {
			singleGroupResp.Groups = append(singleGroupResp.Groups, group)
			singleGroupResp.Count = 1
			break
		}
	}

	return singleGroupResp, err
}

// List retrieves a list of device groups from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/DeviceGroups/#get_devicegroups
func (d *DeviceGroupAPI) List() (*types.DeviceGroupResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, deviceGroupEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.DeviceGroupResponse)
	return resp, c.do(req, resp)
}

// GetMembers retrieves a list of device group members from the LibreNMS API.
// The identifier can be either the group ID or the group name.
//
// Documentation: https://docs.librenms.org/API/DeviceGroups/#get_devices_by_group
func (d *DeviceGroupAPI) GetMembers(identifier string) (*types.DeviceGroupMembersResponse, error) {
	c := d.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s", deviceGroupEndpoint, identifier), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.DeviceGroupMembersResponse)
	return resp, c.do(req, resp)
}

// Update updates an existing device group in the LibreNMS API.
//
// The documentation states it uses name rather than ID to reference the group, but both seem to work (as of v25.5).
// Documentation: https://docs.librenms.org/API/DeviceGroups/#update_devicegroup
func (d *DeviceGroupAPI) Update(identifier string, payload *types.DeviceGroupUpdateRequest) (*types.BaseResponse, error) {
	c := d.client
	uri, err := url.Parse(fmt.Sprintf("%s/%s", deviceGroupEndpoint, identifier))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URI: %w", err)
	}

	req, err := c.newRequest(http.MethodPatch, uri.String(), payload, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}
