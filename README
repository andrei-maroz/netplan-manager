# NetplanManager

This is a test `NetplanManager` for managing network configurations using Netplan in Ubuntu 24. It provides functionalities to enable DHCP, set static IP addresses, write interface configurations to separate config files, and apply changes to the network configuration.

## Functions

- **NewNetplanManager(dir string) (*NetplanManager, error)**: Initializes a new `NetplanManager` instance by reading the Netplan configuration from the specified directory.

- **Apply() error**: Applies the current Netplan configuration using the `netplan apply` command.

- **EnableDhcp4(interfaceName string)**: Enables DHCP for the specified network interface.

- **SetIp4(interfaceName string, ip string)**: Sets a static IPv4 address for the specified network interface.

- **PrintConfig()**: Prints the current Netplan configuration in JSON format.

- **WriteInterfaceConfig(interfaceName string)**: Writes the configuration of the specified network interface to the Netplan directory.

## Main Function

The `main` function demonstrates the usage of `NetplanManager` by:
1. Creating a new instance of `NetplanManager`.
2. Enabling DHCP for the interface `enp0s8`.
3. Writing the configuration for `enp0s8`.
4. Setting a static IP address (`192.168.0.199/24`) for the interface `enp0s3`.
5. Writing the configuration for `enp0s3`.
6. Applying the changes to the network configuration.
7. Printing the current Netplan configuration.
