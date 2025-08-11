package librenms

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	inventoryEndpoint = "inventory"
)

type (
	InventoryItem struct {
		EntPhysicalID           string `json:"entPhysical_id"`
		DeviceID                string `json:"device_id"`
		EntPhysicalIndex        string `json:"entPhysicalIndex"`
		EntPhysicalDescr        string `json:"entPhysicalDescr"`
		EntPhysicalClass        string `json:"entPhysicalClass"`
		EntPhysicalName         string `json:"entPhysicalName"`
		EntPhysicalHardwareRev  string `json:"entPhysicalHardwareRev"`
		EntPhysicalFirmwareRev  string `json:"entPhysicalFirmwareRev"`
		EntPhysicalSoftwareRev  string `json:"entPhysicalSoftwareRev"`
		EntPhysicalSerialNum    string `json:"entPhysicalSerialNum"`
		EntPhysicalModelName    string `json:"entPhysicalModelName"`
		EntPhysicalMfgName      string `json:"entPhysicalMfgName"`
		EntPhysicalIsFRU        Bool   `json:"entPhysicalIsFRU"`
		EntPhysicalAlias        string `json:"entPhysicalAlias"`
		EntPhysicalAssetID      string `json:"entPhysicalAssetID"`
		EntPhysicalContainedIn  string `json:"entPhysicalContainedIn"`
		EntPhysicalParentRelPos string `json:"entPhysicalParentRelPos"`
		EntPhysicalMfgDate      string `json:"entPhysicalMfgDate"`
		EntPhysicalUris         string `json:"entPhysicalUris"`
		EntPhysicalVendorType   string `json:"entPhysicalVendorType"`
		IfIndex                 string `json:"ifIndex"`
		Deleted                 Bool   `json:"deleted"`
	}

	InventoryResponse struct {
		BaseResponse
		Count     int             `json:"count"`
		Inventory []InventoryItem `json:"inventory"`
	}

	InventoryParams struct {
		EntPhysicalClass       *string `url:"entPhysicalClass,omitempty"`
		EntPhysicalContainedIn *string `url:"entPhysicalContainedIn,omitempty"`
	}
)

// GetInventory retrieves the inventory for a device with optional filtering
// This enables recursive lookup by specifying entPhysicalContainedIn parameter
//
// Documentation: https://docs.librenms.org/API/Inventory/#get_inventory
// Route: /api/v0/inventory/:hostname
func (i *InventoryAPI) GetInventory(hostname string, params *InventoryParams) (*InventoryResponse, error) {
	path := fmt.Sprintf("%s/%s", inventoryEndpoint, hostname)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.EntPhysicalClass != nil {
			query.Set("entPhysicalClass", *params.EntPhysicalClass)
		}
		if params.EntPhysicalContainedIn != nil {
			query.Set("entPhysicalContainedIn", *params.EntPhysicalContainedIn)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp InventoryResponse
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
func (i *InventoryAPI) GetInventoryForDevice(hostname string) (*InventoryResponse, error) {
	path := fmt.Sprintf("%s/%s/all", inventoryEndpoint, hostname)
	var resp InventoryResponse
	httpReq, err := i.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	err = i.client.do(httpReq, &resp)
	return &resp, err
}
