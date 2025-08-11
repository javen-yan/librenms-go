package types

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
