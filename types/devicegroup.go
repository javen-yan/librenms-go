package types

import "encoding/json"

type (
	// DeviceGroup represents a device group in LibreNMS.
	DeviceGroup struct {
		ID          int                      `json:"id"`
		Name        string                   `json:"name"`
		Description *string                  `json:"desc"`
		Pattern     *string                  `json:"pattern"`
		Rules       DeviceGroupRuleContainer `json:"rules"`
		Type        string                   `json:"type"`
	}

	// DeviceGroupRuleContainer represents the top-level container for device group rules.
	DeviceGroupRuleContainer struct {
		Condition string            `json:"condition"`
		Joins     [][]string        `json:"joins"`
		Rules     []DeviceGroupRule `json:"rules"`
		Valid     bool              `json:"valid"`
	}

	// DeviceGroupRule represents a rule within a device group. This is a recursive structure.
	// It can contain nested rules, allowing for complex conditions.
	//
	// A terminal section defines id, field, type, input, operator, and value.
	// A non-terminal section defines condition and a list of rules.
	DeviceGroupRule struct {
		ID        string            `json:"id,omitempty"`
		Condition string            `json:"condition,omitempty"`
		Field     string            `json:"field,omitempty"`
		Input     string            `json:"input,omitempty"`
		Operator  string            `json:"operator,omitempty"`
		Rules     []DeviceGroupRule `json:"rules,omitempty"`
		Type      string            `json:"type,omitempty"`
		Value     string            `json:"value,omitempty"`
	}

	// DeviceGroupCreateRequest represents the request payload for creating a device group.
	//
	// The rules should be a serialized JSON string that matches the DeviceGroupRuleContainer
	// structure. Define your rules using the DeviceGroupRuleContainer struct and then
	// serialize it using its JSON() method.
	DeviceGroupCreateRequest struct {
		Name        string  `json:"name"`
		Description *string `json:"desc,omitempty"`
		Devices     []int   `json:"devices,omitempty"`
		Rules       *string `json:"rules,omitempty"`
		Type        string  `json:"type"`
	}

	// DeviceGroupUpdateRequest represents the request payload for updating a device group.
	//
	// The rules should be a serialized JSON string that matches the DeviceGroupRuleContainer
	// structure. Define your rules using the DeviceGroupRuleContainer struct and then
	// serialize it using its JSON() method.
	DeviceGroupUpdateRequest struct {
		Name        string  `json:"name,omitempty"`
		Description *string `json:"desc,omitempty"`
		Devices     []int   `json:"devices,omitempty"`
		Rules       *string `json:"rules,omitempty"`
		Type        string  `json:"type,omitempty"`
	}

	// DeviceGroupResponse represents a response containing a list of device groups from the LibreNMS API.
	DeviceGroupResponse struct {
		BaseResponse
		Groups []DeviceGroup `json:"groups"`
	}

	// DeviceGroupMember represents a member of a device group.
	DeviceGroupMember struct {
		ID int `json:"device_id"`
	}

	// DeviceGroupMembersResponse represents a response containing the members of a device group.
	DeviceGroupMembersResponse struct {
		BaseResponse
		Devices []DeviceGroupMember `json:"devices"`
	}

	// DeviceGroupCreateResponse represents a creation response.
	DeviceGroupCreateResponse struct {
		BaseResponse
		ID int `json:"id"`
	}
)

// JSON is a helper function that serializes the DeviceGroupRuleContainer to JSON format.
func (g *DeviceGroupRuleContainer) JSON() (string, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// MustJSON is a helper function that serializes the DeviceGroupRuleContainer to JSON format.
// It returns an empty string if the marshalling fails.
func (g *DeviceGroupRuleContainer) MustJSON() string {
	data, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(data)
}
