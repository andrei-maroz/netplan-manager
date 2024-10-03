package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/juju/juju/network/netplan"
)

const netplanDir = "/etc/netplan"
const version = 2 // Netplan support only v2.

type NetplanManager struct {
	netplanConfig *netplan.Netplan
}

func NewNetplanManager(dir string) (*NetplanManager, error) {
	config, err := netplan.ReadDirectory(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading Netplan directory: %w", err)
	}
	return &NetplanManager{netplanConfig: &config}, nil
}

func (nm *NetplanManager) Apply() error {
	// Apply Netplan configuration.
	cmd := exec.Command("netplan", "apply")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error applying Netplan configuration: %w", err)
	}

	fmt.Println("Netplan configuration applied successfully.")
	return nil
}

func (nm *NetplanManager) EnableDhcp4(interfaceName string) {

	if ethernet, ok := nm.netplanConfig.Network.Ethernets[interfaceName]; ok {
		*ethernet.Interface.DHCP4 = true
		ethernet.Interface.Addresses = []string{}

		// Update netplanConfig.
		nm.netplanConfig.Network.Ethernets[interfaceName] = ethernet
	} else {
		log.Fatalf("Interface %s not found", interfaceName)
	}
}

func (nm *NetplanManager) SetIp4(interfaceName string, ip string) {

	if ethernet, ok := nm.netplanConfig.Network.Ethernets[interfaceName]; ok {
		*ethernet.Interface.DHCP4 = false
		ethernet.Interface.Addresses = []string{ip}

		// Update netplanConfig.
		nm.netplanConfig.Network.Ethernets[interfaceName] = ethernet
	} else {
		log.Fatalf("Interface %s not found", interfaceName)
	}
}

func (nm *NetplanManager) PrintConfig() {
	// Convert to json.
	formattedConfig, err := json.MarshalIndent(nm.netplanConfig, "", "  ")
	if err != nil {
		log.Fatalf("error formatting Netplan config: %v", err)
	}
	fmt.Println(string(formattedConfig))
}

func (nm *NetplanManager) WriteInterfaceConfig(interfaceName string) {

	interfaceConfig, exists := nm.netplanConfig.Network.Ethernets[interfaceName]
	if !exists {
		log.Fatalf("interface %s does not exist in the current configuration", interfaceName)
	}

	// Create netplan config only with necessory interface.
	tempNetplanConfig := &netplan.Netplan{
		Network: netplan.Network{
			Ethernets: map[string]netplan.Ethernet{
				interfaceName: interfaceConfig,
			},
			Version: version,
		},
	}

	filePath := fmt.Sprintf("%s/%s.yaml", netplanDir, interfaceName)

	if _, err := tempNetplanConfig.Write(filePath); err != nil {
		log.Fatalf("error writing interface config: %v", err)
	}
	fmt.Printf("Configuration for interface %s written to %s successfully.\n", interfaceName, filePath)
}

func main() {
	netplanManager, err := NewNetplanManager(netplanDir)
	if err != nil {
		log.Fatalf("failed to create NetplanManager: %v", err)
	}

	netplanManager.EnableDhcp4("enp0s8")
	netplanManager.WriteInterfaceConfig("enp0s8")

	netplanManager.SetIp4("enp0s3", "192.168.0.199/24")
	netplanManager.WriteInterfaceConfig("enp0s3")

	if err := netplanManager.Apply(); err != nil {
		log.Fatalf("failed to apply changes: %v", err)
	}
	netplanManager.PrintConfig()
}
