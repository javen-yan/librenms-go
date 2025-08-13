package types

type (
	InventoryItem struct {
		EntPhysicalID           int    `json:"entPhysical_id,omitempty"`
		DeviceID                int    `json:"device_id,omitempty"`
		EntPhysicalIndex        int    `json:"entPhysicalIndex,omitempty"`
		EntPhysicalDescr        string `json:"entPhysicalDescr,omitempty"`
		EntPhysicalClass        string `json:"entPhysicalClass,omitempty"`
		EntPhysicalName         string `json:"entPhysicalName,omitempty"`
		EntPhysicalHardwareRev  string `json:"entPhysicalHardwareRev,omitempty"`
		EntPhysicalFirmwareRev  string `json:"entPhysicalFirmwareRev,omitempty"`
		EntPhysicalSoftwareRev  string `json:"entPhysicalSoftwareRev,omitempty"`
		EntPhysicalSerialNum    string `json:"entPhysicalSerialNum,omitempty"`
		EntPhysicalModelName    string `json:"entPhysicalModelName,omitempty"`
		EntPhysicalMfgName      string `json:"entPhysicalMfgName,omitempty"`
		EntPhysicalIsFRU        Bool   `json:"entPhysicalIsFRU,omitempty"`
		EntPhysicalAlias        string `json:"entPhysicalAlias,omitempty"`
		EntPhysicalAssetID      string `json:"entPhysicalAssetID,omitempty"`
		EntPhysicalContainedIn  int    `json:"entPhysicalContainedIn,omitempty"`
		EntPhysicalParentRelPos int    `json:"entPhysicalParentRelPos,omitempty"`
		EntPhysicalMfgDate      string `json:"entPhysicalMfgDate,omitempty"`
		EntPhysicalUris         string `json:"entPhysicalUris,omitempty"`
		EntPhysicalVendorType   string `json:"entPhysicalVendorType,omitempty"`
		IfIndex                 string `json:"ifIndex,omitempty"`
		Deleted                 Bool   `json:"deleted,omitempty"`
	}

	InventoryResponse struct {
		BaseResponse
		Count     int             `json:"count,omitempty"`
		Inventory []InventoryItem `json:"inventory"`
	}

	InventoryParams struct {
		EntPhysicalClass       string `form:"entPhysicalClass,omitempty"`
		EntPhysicalContainedIn string `form:"entPhysicalContainedIn,omitempty"`
	}
)
