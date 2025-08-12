package librenms

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/javen-yan/librenms-go/types"
)

const (
	inventoryEndpoint = "inventory"
)

// GetInventory retrieves the inventory for a device with optional filtering
// This enables recursive lookup by specifying entPhysicalContainedIn parameter
//
// Documentation: https://docs.librenms.org/API/Inventory/#get_inventory
// Route: /api/v0/inventory/:hostname
func (i *InventoryAPI) GetInventory(hostname string, params *types.InventoryParams) (*types.InventoryResponse, error) {
	path := fmt.Sprintf("%s/%s", inventoryEndpoint, hostname)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.EntPhysicalClass != "" {
			query.Set("entPhysicalClass", params.EntPhysicalClass)
		}
		if params.EntPhysicalContainedIn != "" {
			query.Set("entPhysicalContainedIn", params.EntPhysicalContainedIn)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.InventoryResponse
	httpReq, err := i.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = i.client.do(httpReq, &resp)
	return &resp, err
}

// GetInventoryForDevice retrieves the flattened inventory for a device
// This retrieves all inventory items regardless of their structure
//
// Documentation: https://docs.librenms.org/API/Inventory/#get_inventory_for_device
// Route: /api/v0/inventory/:hostname/all
func (i *InventoryAPI) GetInventoryForDevice(hostname string) (*types.InventoryResponse, error) {
	path := fmt.Sprintf("%s/%s/all", inventoryEndpoint, hostname)
	var resp types.InventoryResponse
	httpReq, err := i.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	err = i.client.do(httpReq, &resp)
	return &resp, err
}
