package librenms

import (
	"fmt"
	"net/http"

	"github.com/javen-yan/librenms-go/types"
)

const (
	alertRuleEndpoint = "rules"
)

// Create creates a specific alert rule in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#add_rule
func (a *AlertRuleAPI) Create(payload *types.AlertRuleCreateRequest) (*types.BaseResponse, error) {
	c := a.client
	// as a convenience/hack, add a -1 to Devices if Devices is empty
	if len(payload.Devices) == 0 {
		payload.Devices = []int{-1}
	}

	req, err := c.newRequest(http.MethodPost, alertRuleEndpoint, payload, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// DeleteAlertRule deletes a specific alert rule by its ID from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#delete_rule
func (a *AlertRuleAPI) Delete(id int) (*types.BaseResponse, error) {
	c := a.client
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("%s/%d", alertRuleEndpoint, id), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// GetAlertRule retrieves a specific alert rule by its ID from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#get_alert_rule
func (a *AlertRuleAPI) Get(id int) (*types.AlertRuleResponse, error) {
	c := a.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%d", alertRuleEndpoint, id), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.AlertRuleResponse)
	return resp, c.do(req, resp)
}

// GetAlertRules retrieves all alert rules from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#list_alert_rules
func (a *AlertRuleAPI) List() (*types.AlertRuleResponse, error) {
	c := a.client
	req, err := c.newRequest(http.MethodGet, alertRuleEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.AlertRuleResponse)
	return resp, c.do(req, resp)
}

// UpdateAlertRule updates a specific alert rule in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#edit_rule
func (a *AlertRuleAPI) Update(payload *types.AlertRuleUpdateRequest) (*types.BaseResponse, error) {
	c := a.client
	if payload.ID < 1 {
		return nil, fmt.Errorf("rule ID is required for updating an alert rule")
	}

	// as a convenience/hack, add a -1 to Devices if Devices is empty
	if len(payload.Devices) == 0 {
		payload.Devices = []int{-1}
	}

	req, err := c.newRequest(http.MethodPut, alertRuleEndpoint, payload, nil)
	if err != nil {
		return nil, err
	}
	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}
