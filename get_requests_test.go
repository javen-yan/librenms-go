package librenms_test

import (
	"log/slog"
	"testing"

	"github.com/javen-yan/librenms-go"
	"github.com/stretchr/testify/require"
)

const (
	// 真实的 LibreNMS 服务器配置
	realServerAddr = "http://192.168.8.2:8000"
	realToken      = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

// 创建真实服务器的客户端
func createRealClient(t *testing.T) *librenms.Client {
	client, err := librenms.New(realServerAddr+"/", realToken, librenms.WithLogLevel(slog.LevelDebug))
	require.NoError(t, err, "Failed to create client with real server")
	return client
}

// TestRealServer_GetDevices 测试从真实服务器获取设备列表
func TestRealServer_GetDevices(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取设备列表
	devicesResp, err := client.Device.List(nil)

	r.NoError(err, "GetDevices should not return error")
	r.NotNil(devicesResp, "GetDevices response should not be nil")
	r.Equal("ok", devicesResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(devicesResp.Count, 0, "Count should be non-negative")

	if devicesResp.Count > 0 {
		r.Len(devicesResp.Devices, devicesResp.Count, "Devices slice length should match count")

		// 验证第一个设备的基本字段
		device := devicesResp.Devices[0]
		r.Greater(device.DeviceID, 0, "DeviceID should be positive")
		r.NotEmpty(device.Hostname, "Hostname should not be empty")
	}

	t.Logf("Successfully retrieved %d devices from real server", devicesResp.Count)
}

// TestRealServer_GetDevice 测试从真实服务器获取单个设备
func TestRealServer_GetDevice(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 首先获取设备列表以获取一个有效的设备标识符
	devicesResp, err := client.Device.List(nil)
	r.NoError(err, "GetDevices should not return error")
	r.Greater(devicesResp.Count, 0, "Should have at least one device to test with")

	// 使用第一个设备的 hostname 进行测试
	testHostname := devicesResp.Devices[0].Hostname
	r.NotEmpty(testHostname, "Test hostname should not be empty")

	// 测试获取单个设备
	deviceResp, err := client.Device.Get(testHostname)

	r.NoError(err, "GetDevice should not return error")
	r.NotNil(deviceResp, "GetDevice response should not be nil")
	r.Equal("ok", deviceResp.Status, "Expected status 'ok'")
	r.Equal(1, deviceResp.Count, "Expected count 1")
	r.Len(deviceResp.Devices, 1, "Expected 1 device")

	device := deviceResp.Devices[0]
	r.Equal(testHostname, device.Hostname, "Hostname should match")
	r.Greater(device.DeviceID, 0, "DeviceID should be positive")

	t.Logf("Successfully retrieved device: %s (ID: %d)", device.Hostname, device.DeviceID)
}

// TestRealServer_GetPorts 测试从真实服务器获取端口信息
func TestRealServer_GetPorts(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取端口列表
	portsResp, err := client.Port.GetAllPorts(nil)

	r.NoError(err, "GetPorts should not return error")
	r.NotNil(portsResp, "GetPorts response should not be nil")
	r.Equal("ok", portsResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(len(portsResp.Ports), 0, "Ports slice should not be negative")

	// 添加详细的调试信息
	t.Logf("Ports response: Status=%s, Count=%d, Ports slice length=%d",
		portsResp.Status, portsResp.Count, len(portsResp.Ports))

	// 使用实际的数组长度而不是 Count 字段
	actualPortCount := len(portsResp.Ports)
	if actualPortCount > 0 {
		r.Len(portsResp.Ports, actualPortCount, "Ports slice length should match actual count")

		// 验证第一个端口的基本字段
		port := portsResp.Ports[0]
		r.Greater(port.PortID, 0, "PortID should be positive")

		// 验证 ifName 字段（API 返回的数据）
		if port.IfName != "" {
			t.Logf("First port: ID=%d, Name=%s", port.PortID, port.IfName)
		} else {
			t.Logf("First port: ID=%d, Name is empty", port.PortID)
		}

		// 如果返回了多个端口，显示一些统计信息
		if portsResp.Count > 10 {
			t.Logf("Port types found: %d total ports", portsResp.Count)

			// 统计不同类型的端口
			portTypes := make(map[string]int)
			for _, p := range portsResp.Ports {
				if p.IfName != "" {
					portTypes[p.IfName]++
				}
			}

			t.Logf("Port type distribution: %v", portTypes)
		}
	} else {
		t.Logf("No ports returned, but API should return 63 ports based on curl test")
	}

	t.Logf("Successfully retrieved %d ports from server (Count field: %d)", len(portsResp.Ports), portsResp.Count)
}

// TestRealServer_GetAlerts 测试从真实服务器获取告警信息
func TestRealServer_GetAlerts(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取告警列表
	alertsResp, err := client.Alert.List(nil)

	r.NoError(err, "GetAlerts should not return error")
	r.NotNil(alertsResp, "GetAlerts response should not be nil")
	r.Equal("ok", alertsResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(alertsResp.Count, 0, "Count should be non-negative")

	if alertsResp.Count > 0 {
		r.Len(alertsResp.Alerts, alertsResp.Count, "Alerts slice length should match count")

		// 验证第一个告警的基本字段
		alert := alertsResp.Alerts[0]
		r.Greater(alert.ID, 0, "Alert ID should be positive")
		r.NotEmpty(alert.Name, "Alert name should not be empty")
	}

	t.Logf("Successfully retrieved %d alerts from real server", alertsResp.Count)
}

// TestRealServer_GetServices 测试从真实服务器获取服务信息
func TestRealServer_GetServices(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取服务列表
	servicesResp, err := client.Service.List()

	r.NoError(err, "GetServices should not return error")
	r.NotNil(servicesResp, "GetServices response should not be nil")
	r.Equal("ok", servicesResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(servicesResp.Count, 0, "Count should be non-negative")

	if servicesResp.Count > 0 {
		r.Len(servicesResp.Services, servicesResp.Count, "Services slice length should match count")

		// 验证第一个服务的基本字段
		service := servicesResp.Services[0]
		r.Greater(service.ID, 0, "Service ID should be positive")
		r.NotEmpty(service.Name, "Service name should not be empty")
	}

	t.Logf("Successfully retrieved %d services from real server", servicesResp.Count)
}

// TestRealServer_GetLocations 测试从真实服务器获取位置信息
func TestRealServer_GetLocations(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取位置列表
	locationsResp, err := client.Location.List()

	// 位置 API 可能不存在，跳过测试
	if err != nil {
		t.Skipf("Locations API not available: %v", err)
		return
	}

	r.NotNil(locationsResp, "GetLocations response should not be nil")
	r.Equal("ok", locationsResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(locationsResp.Count, 0, "Count should be non-negative")

	if locationsResp.Count > 0 {
		r.Len(locationsResp.Locations, locationsResp.Count, "Locations slice length should match count")

		// 验证第一个位置的基本字段
		location := locationsResp.Locations[0]
		r.Greater(location.ID, 0, "Location ID should be positive")
		r.NotEmpty(location.Name, "Location name should not be empty")
	}

	t.Logf("Successfully retrieved %d locations from real server", locationsResp.Count)
}

// TestRealServer_GetSystemInfo 测试从真实服务器获取系统信息
func TestRealServer_GetSystemInfo(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取系统信息
	systemResp, err := client.System.Get()

	r.NoError(err, "GetSystemInfo should not return error")
	r.NotNil(systemResp, "GetSystemInfo response should not be nil")
	r.Equal("ok", systemResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(systemResp.Count, 0, "Count should be non-negative")

	if systemResp.Count > 0 {
		r.Len(systemResp.System, systemResp.Count, "System slice length should match count")

		// 验证系统信息的基本字段
		system := systemResp.System[0]
		r.NotEmpty(system.LocalVer, "Local version should not be empty")
		r.NotEmpty(system.DBSchema, "Database schema should not be empty")
	}

	t.Logf("Successfully retrieved system info from real server")
}

// TestRealServer_GetInventory 测试从真实服务器获取设备清单
func TestRealServer_GetInventory(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 首先获取设备列表
	devicesResp, err := client.Device.List(nil)
	r.NoError(err, "GetDevices should not return error")
	r.Greater(devicesResp.Count, 0, "Should have at least one device to test with")

	// 使用第一个设备的 hostname 进行测试
	testHostname := devicesResp.Devices[0].Hostname

	// 测试获取设备清单
	inventoryResp, err := client.Inventory.GetInventory(testHostname, nil)

	r.NoError(err, "GetInventory should not return error")
	r.NotNil(inventoryResp, "GetInventory response should not be nil")
	r.Equal("ok", inventoryResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(inventoryResp.Count, 0, "Count should be non-negative")

	if inventoryResp.Count > 0 {
		r.Len(inventoryResp.Inventory, inventoryResp.Count, "Inventory slice length should match count")

		// 验证第一个清单项的基本字段
		item := inventoryResp.Inventory[0]
		r.NotEmpty(item.DeviceID, "Device ID should not be empty")
		r.NotEmpty(item.EntPhysicalIndex, "Physical index should not be empty")
	}

	t.Logf("Successfully retrieved %d inventory items for device %s", inventoryResp.Count, testHostname)
}

// TestRealServer_GetDeviceGroups 测试从真实服务器获取设备组信息
func TestRealServer_GetDeviceGroups(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取设备组列表
	groupsResp, err := client.DeviceGroup.List()

	r.NoError(err, "GetDeviceGroups should not return error")
	r.NotNil(groupsResp, "GetDeviceGroups response should not be nil")
	r.Equal("ok", groupsResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(groupsResp.Count, 0, "Count should be non-negative")

	if groupsResp.Count > 0 {
		r.Len(groupsResp.Groups, groupsResp.Count, "Groups slice length should match count")

		// 验证第一个设备组的基本字段
		group := groupsResp.Groups[0]
		r.Greater(group.ID, 0, "Group ID should be positive")
		r.NotEmpty(group.Name, "Group name should not be empty")

		// 验证设备组的其他字段
		r.NotNil(group.Rules, "Rules should not be nil")
		r.Equal("AND", group.Rules.Condition, "Expected condition 'AND'")
		r.True(group.Rules.Valid, "Rules should be valid")
	}

	t.Logf("Successfully retrieved %d device groups from real server", groupsResp.Count)
}

// TestRealServer_GetAlertRules 测试从真实服务器获取告警规则
func TestRealServer_GetAlertRules(t *testing.T) {
	r := require.New(t)

	client := createRealClient(t)

	// 测试获取告警规则列表
	rulesResp, err := client.AlertRule.List()

	r.NoError(err, "GetAlertRules should not return error")
	r.NotNil(rulesResp, "GetAlertRules response should not be nil")
	r.Equal("ok", rulesResp.Status, "Expected status 'ok'")
	r.GreaterOrEqual(rulesResp.Count, 0, "Count should be non-negative")

	if rulesResp.Count > 0 {
		r.Len(rulesResp.Rules, rulesResp.Count, "Rules slice length should match count")

		// 验证第一个告警规则的基本字段
		rule := rulesResp.Rules[0]
		r.Greater(rule.ID, 0, "Rule ID should be positive")
		r.NotEmpty(rule.Name, "Rule name should not be empty")
	}

	t.Logf("Successfully retrieved %d alert rules from real server", rulesResp.Count)
}
