# Diode Go SDK - Entity Examples

Source: NetBox v4.6.0
Generated: 2026-05-14 20:30:15Z

## Prerequisites

- Go 1.25 or later
- Diode SDK for Go

## Installation

```bash
go get github.com/netboxlabs/diode-sdk-go@latest
```

## Configuration

Each example uses constants for configuration. You can modify these in the example code:

```go
target     = "grpc://localhost:8080/diode"
appName    = "example-app"
appVersion = "1.0.0"
```

## Quick Start

Each entity example is in its own directory with a complete Go module. To run an example:

```bash
cd examples/device
go run main.go
```

## Example Patterns

Each example includes three patterns:

- **Minimal**: Only required fields
- **Extended**: Required fields plus common optional fields
- **Explicit**: Fully nested objects with all common fields

Switch between patterns by uncommenting the desired function call in `main()`.

## Available Entity Examples

### Circuits

- [Circuit](examples/circuit/)
- [CircuitGroup](examples/circuit_group/)
- [CircuitGroupAssignment](examples/circuit_group_assignment/)
- [CircuitTermination](examples/circuit_termination/)
- [CircuitType](examples/circuit_type/)
- [Provider](examples/provider/)
- [ProviderAccount](examples/provider_account/)
- [ProviderNetwork](examples/provider_network/)
- [VirtualCircuit](examples/virtual_circuit/)
- [VirtualCircuitTermination](examples/virtual_circuit_termination/)
- [VirtualCircuitType](examples/virtual_circuit_type/)

### DCIM

- [Cable](examples/cable/)
- [CablePath](examples/cable_path/)
- [CableTermination](examples/cable_termination/)
- [ConsolePort](examples/console_port/)
- [ConsoleServerPort](examples/console_server_port/)
- [Device](examples/device/)
- [DeviceBay](examples/device_bay/)
- [DeviceRole](examples/device_role/)
- [DeviceType](examples/device_type/)
- [FrontPort](examples/front_port/)
- [Interface](examples/interface_entity/)
- [InventoryItem](examples/inventory_item/)
- [InventoryItemRole](examples/inventory_item_role/)
- [Location](examples/location/)
- [Manufacturer](examples/manufacturer/)
- [Module](examples/module/)
- [ModuleBay](examples/module_bay/)
- [ModuleType](examples/module_type/)
- [ModuleTypeProfile](examples/module_type_profile/)
- [Platform](examples/platform/)
- [PowerFeed](examples/power_feed/)
- [PowerOutlet](examples/power_outlet/)
- [PowerPanel](examples/power_panel/)
- [PowerPort](examples/power_port/)
- [Rack](examples/rack/)
- [RackReservation](examples/rack_reservation/)
- [RackRole](examples/rack_role/)
- [RackType](examples/rack_type/)
- [RearPort](examples/rear_port/)
- [Region](examples/region/)
- [Site](examples/site/)
- [SiteGroup](examples/site_group/)
- [VirtualChassis](examples/virtual_chassis/)
- [VirtualDeviceContext](examples/virtual_device_context/)

### Extras

- [CustomField](examples/custom_field/)
- [CustomFieldChoiceSet](examples/custom_field_choice_set/)
- [CustomLink](examples/custom_link/)
- [JournalEntry](examples/journal_entry/)
- [Tag](examples/tag/)

### IPAM

- [ASN](examples/asn/)
- [ASNRange](examples/asn_range/)
- [Aggregate](examples/aggregate/)
- [FHRPGroup](examples/fhrp_group/)
- [FHRPGroupAssignment](examples/fhrp_group_assignment/)
- [IPAddress](examples/ip_address/)
- [IPRange](examples/ip_range/)
- [MACAddress](examples/mac_address/)
- [Prefix](examples/prefix/)
- [RIR](examples/rir/)
- [Role](examples/role/)
- [RouteTarget](examples/route_target/)
- [Service](examples/service/)
- [VLAN](examples/vlan/)
- [VLANGroup](examples/vlan_group/)
- [VLANTranslationPolicy](examples/vlan_translation_policy/)
- [VLANTranslationRule](examples/vlan_translation_rule/)
- [VRF](examples/vrf/)

### Other

- [CableBundle](examples/cable_bundle/)
- [DeviceConfig](examples/device_config/)
- [RackGroup](examples/rack_group/)
- [ScriptModule](examples/script_module/)
- [VirtualMachineType](examples/virtual_machine_type/)

### Tenancy

- [Contact](examples/contact/)
- [ContactAssignment](examples/contact_assignment/)
- [ContactGroup](examples/contact_group/)
- [ContactRole](examples/contact_role/)
- [Owner](examples/owner/)
- [OwnerGroup](examples/owner_group/)
- [Tenant](examples/tenant/)
- [TenantGroup](examples/tenant_group/)

### VPN

- [IKEPolicy](examples/ike_policy/)
- [IKEProposal](examples/ike_proposal/)
- [IPSecPolicy](examples/ip_sec_policy/)
- [IPSecProfile](examples/ip_sec_profile/)
- [IPSecProposal](examples/ip_sec_proposal/)
- [L2VPN](examples/l2vpn/)
- [L2VPNTermination](examples/l2vpn_termination/)
- [Tunnel](examples/tunnel/)
- [TunnelGroup](examples/tunnel_group/)
- [TunnelTermination](examples/tunnel_termination/)

### Virtualization

- [Cluster](examples/cluster/)
- [ClusterGroup](examples/cluster_group/)
- [ClusterType](examples/cluster_type/)
- [VMInterface](examples/vm_interface/)
- [VirtualDisk](examples/virtual_disk/)
- [VirtualMachine](examples/virtual_machine/)

### Wireless

- [WirelessLAN](examples/wireless_lan/)
- [WirelessLANGroup](examples/wireless_lan_group/)
- [WirelessLink](examples/wireless_link/)
