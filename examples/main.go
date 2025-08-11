package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/javen-yan/librenms-go"
)

func main() {
	// Create a new LibreNMS client
	// Using the new NewClient function
	client, err := librenms.NewClient(
		"http://localhost:8000",
		"AUTH_TOKEN",
		librenms.WithLogLevel(slog.LevelDebug),
	)
	if err != nil {
		log.Fatalf("Failed to create LibreNMS client: %v", err)
	}

	// Get system information
	fmt.Println("\n=== System Information ===")
	systemInfo, err := client.System.Get()
	if err != nil {
		log.Printf("Failed to get system information: %v", err)
	} else {
		if len(systemInfo.System) > 0 {
			sys := systemInfo.System[0]
			fmt.Printf("LibreNMS Version: %s\n", sys.LocalVer)
			fmt.Printf("Database Version: %s\n", sys.DatabaseVer)
			fmt.Printf("PHP Version: %s\n", sys.PHPVer)
			fmt.Printf("Python Version: %s\n", sys.PythonVer)
			fmt.Printf("RRDTool Version: %s\n", sys.RRDToolVer)
			fmt.Printf("Build Date: %s\n", sys.LocalDate)
		}
	}

	// Use Device API
	fmt.Println("\n=== Device Management ===")

	// List all devices
	devices, err := client.Device.List(nil)
	if err != nil {
		log.Printf("Failed to get device list: %v", err)
	} else {
		fmt.Printf("Found %d devices\n", devices.Count)
		for _, device := range devices.Devices {
			display := ""
			if device.Display != nil {
				display = *device.Display
			}
			fmt.Printf("- %s (%s) - %s\n", device.Hostname, display, device.OS)
		}
	}

	// Get specific device
	if len(devices.Devices) > 0 {
		deviceID := devices.Devices[0].DeviceID
		device, err := client.Device.Get(fmt.Sprintf("%d", deviceID))
		if err != nil {
			log.Printf("Failed to get device details: %v", err)
		} else {
			fmt.Printf("\nDevice Details:\n")
			fmt.Printf("ID: %d\n", device.Devices[0].DeviceID)
			fmt.Printf("Hostname: %s\n", device.Devices[0].Hostname)
			fmt.Printf("Operating System: %s\n", device.Devices[0].OS)
			fmt.Printf("Status: %v\n", device.Devices[0].Status)
		}
	}

	// Use Alert API
	fmt.Println("\n=== Alert Management ===")

	// List all alerts
	alerts, err := client.Alert.List(nil)
	if err != nil {
		log.Printf("Failed to get alert list: %v", err)
	} else {
		fmt.Printf("Found %d alerts\n", alerts.Count)
		for _, alert := range alerts.Alerts {
			fmt.Printf("- %s (%s) - %s\n", alert.Name, alert.Hostname, alert.Severity)
		}
	}

	// Use Location API
	fmt.Println("\n=== Location Management ===")
	locations, err := client.Location.List()
	if err != nil {
		log.Printf("Failed to get location list: %v", err)
	} else {
		fmt.Printf("Found %d locations\n", locations.Count)
		for _, location := range locations.Locations {
			fmt.Printf("- %s\n", location.Name)
		}
	}

	// Use Service API
	fmt.Println("\n=== Service Management ===")
	services, err := client.Service.List()
	if err != nil {
		log.Printf("Failed to get service list: %v", err)
	} else {
		fmt.Printf("Found %d services\n", services.Count)
		for _, service := range services.Services {
			fmt.Printf("- %s - %s\n", service.Name, service.Type)
		}
	}

	fmt.Println("\n=== SDK Usage Example Completed ===")
}
