package librenms_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/javen-yan/librenms-go/types"
	"github.com/stretchr/testify/require"
)

const (
	testEndpointInventory = "/api/v0/inventory"
)

// TestClient_GetInventory tests the GetInventory functionality
func TestClient_GetInventory(t *testing.T) {
	r := require.New(t)

	r.NotNil(testAPIClient, "Global testAPIClient should be initialized")

	inventoryResp, err := testAPIClient.Inventory.GetInventory("test-device", nil)

	r.NoError(err, "GetInventory returned an error")
	r.NotNil(inventoryResp, "GetInventory response is nil")

	r.Equal("ok", inventoryResp.Status, "Expected status 'ok'")
	r.Equal(1, inventoryResp.Count, "Expected count 1")
	r.Len(inventoryResp.Inventory, 1, "Expected 1 inventory item")

	item := inventoryResp.Inventory[0]
	r.Equal("1", item.EntPhysicalID, "Expected Inventory ID 1")
	r.Equal("1", item.DeviceID, "Expected DeviceID 1")
	r.Equal("1", item.EntPhysicalIndex, "Expected EntPhysicalIndex 1")
	r.Equal("Cisco IOS Software, C3560 Software (C3560-IPBASEK9-M), Version 12.2(53)SEY4, RELEASE SOFTWARE (fc1)", item.EntPhysicalDescr, "Expected EntPhysicalDescr")
	r.Equal("chassis", item.EntPhysicalClass, "Expected EntPhysicalClass 'chassis'")
	r.Equal("C3560-24PS-S", item.EntPhysicalName, "Expected EntPhysicalName 'C3560-24PS-S'")
	r.Equal("V02", item.EntPhysicalHardwareRev, "Expected EntPhysicalHardwareRev 'V02'")
	r.Equal("12.2(53)SEY4", item.EntPhysicalFirmwareRev, "Expected EntPhysicalFirmwareRev '12.2(53)SEY4'")
	r.Equal("12.2(53)SEY4", item.EntPhysicalSoftwareRev, "Expected EntPhysicalSoftwareRev '12.2(53)SEY4'")
	r.Equal("FOC1234X0YX", item.EntPhysicalSerialNum, "Expected EntPhysicalSerialNum 'FOC1234X0YX'")
	r.Equal("WS-C3560-24PS-S", item.EntPhysicalModelName, "Expected EntPhysicalModelName 'WS-C3560-24PS-S'")
	r.Equal("Cisco Systems, Inc.", item.EntPhysicalMfgName, "Expected EntPhysicalMfgName 'Cisco Systems, Inc.'")
	r.Equal(types.Bool(true), item.EntPhysicalIsFRU, "Expected EntPhysicalIsFRU true")
	r.Equal("Core Switch", item.EntPhysicalAlias, "Expected EntPhysicalAlias 'Core Switch'")
	r.Equal("ASSET001", item.EntPhysicalAssetID, "Expected EntPhysicalAssetID 'ASSET001'")
	r.Equal("0", item.EntPhysicalContainedIn, "Expected EntPhysicalContainedIn 0")
	r.Equal("-1", item.EntPhysicalParentRelPos, "Expected EntPhysicalParentRelPos -1")
	r.Equal("2023-01-15", item.EntPhysicalMfgDate, "Expected EntPhysicalMfgDate '2023-01-15'")
	r.Equal("http://www.cisco.com/go/c3560", item.EntPhysicalUris, "Expected EntPhysicalUris")
	r.Equal(types.Bool(false), item.Deleted, "Expected Deleted false")
}

// This init function will register handlers for inventory-related API endpoints.
// This is used when running the full test suite.
func init() {
	handleEndpoint(testEndpointInventory, mockResponses{
		http.MethodGet: loadMockResponse("get_inventory_200.json"),
	})

	// Register handler for specific device inventory endpoint
	mux.HandleFunc("/api/v0/inventory/test-device", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			_, err := w.Write(loadMockResponse("get_inventory_200.json"))
			if err != nil {
				log.Printf("Error writing response: %v", err)
				http.Error(w, "Failed to write response", http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
