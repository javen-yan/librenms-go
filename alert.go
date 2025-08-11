package librenms

import (
	"fmt"
	"net/http"

	"github.com/javen-yan/librenms-go/types"
)

const (
	alertEndpoint = "alerts"
)

// Ack acknowledges an alert by its ID in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#ack_alert
func (a *AlertAPI) Ack(alertID int, payload *types.AlertAckRequest) (*types.BaseResponse, error) {
	c := a.client
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("%s/%d", alertEndpoint, alertID), payload, nil)
	if err != nil {
		return nil, err
	}

	alertsResp := new(types.BaseResponse)
	return alertsResp, c.do(req, alertsResp)
}

// Get retrieves a specific alert by its ID from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#get_alert
func (a *AlertAPI) Get(alertID int) (*types.AlertsResponse, error) {
	c := a.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%d", alertEndpoint, alertID), nil, nil)
	if err != nil {
		return nil, err
	}

	alertsResp := new(types.AlertsResponse)
	return alertsResp, c.do(req, alertsResp)
}

// List retrieves a list of alerts from the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#list_alerts
func (a *AlertAPI) List(query *types.AlertsQuery) (*types.AlertsResponse, error) {
	c := a.client
	if query == nil {
		query = types.NewAlertsQuery()
	}
	req, err := c.newRequest(http.MethodGet, alertEndpoint, nil, query.Values())
	if err != nil {
		return nil, err
	}

	alertsResp := new(types.AlertsResponse)
	return alertsResp, c.do(req, alertsResp)
}

// UnmuteAlert unmutes an alert by its ID in the LibreNMS API.
//
// Documentation: https://docs.librenms.org/API/Alerts/#unmute_alert
func (c *Client) UnmuteAlert(alertID int) (*types.BaseResponse, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("%s/unmute/%d", alertEndpoint, alertID), nil, nil)
	if err != nil {
		return nil, err
	}

	alertsResp := new(types.BaseResponse)
	return alertsResp, c.do(req, alertsResp)
}
