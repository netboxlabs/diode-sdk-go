package diode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

func TestDeviceMethods(t *testing.T) {
	tests := []struct {
		name     string
		device   *Device
		expected interface{}
		method   func(*Device) interface{}
	}{
		{
			name:     "GetName",
			device:   &Device{Name: String("device-1")},
			expected: String("device-1"),
			method: func(d *Device) interface{} {
				return d.GetName()
			},
		},
		{
			name:     "GetDeviceType",
			device:   &Device{DeviceType: &DeviceType{Model: String("model-1")}},
			expected: &diodepb.DeviceType{Model: "model-1"},
			method: func(d *Device) interface{} {
				return d.GetDeviceType()
			},
		},
		{
			name:     "GetRole",
			device:   &Device{Role: &DeviceRole{Name: String("role-1")}},
			expected: &diodepb.DeviceRole{Name: "role-1"},
			method: func(d *Device) interface{} {
				return d.GetRole()
			},
		},
		{
			name:     "GetPlatform",
			device:   &Device{Platform: &Platform{Name: String("platform-1")}},
			expected: &diodepb.Platform{Name: "platform-1"},
			method: func(d *Device) interface{} {
				return d.GetPlatform()
			},
		},
		{
			name:     "GetSerial",
			device:   &Device{Serial: String("serial-1")},
			expected: String("serial-1"),
			method: func(d *Device) interface{} {
				return d.GetSerial()
			},
		},
		{
			name:     "GetSite",
			device:   &Device{Site: &Site{Name: String("site-1")}},
			expected: &diodepb.Site{Name: "site-1"},
			method: func(d *Device) interface{} {
				return d.GetSite()
			},
		},
		{
			name:     "GetAssetTag",
			device:   &Device{AssetTag: String("asset-1")},
			expected: String("asset-1"),
			method: func(d *Device) interface{} {
				return d.GetAssetTag()
			},
		},
		{
			name:     "GetStatus",
			device:   &Device{Status: String("active")},
			expected: String("active"),
			method: func(d *Device) interface{} {
				return d.GetStatus()
			},
		},
		{
			name:     "GetDescription",
			device:   &Device{Description: String("description")},
			expected: String("description"),
			method: func(d *Device) interface{} {
				return d.GetDescription()
			},
		},
		{
			name:     "GetComments",
			device:   &Device{Comments: String("comments")},
			expected: String("comments"),
			method: func(d *Device) interface{} {
				return d.GetComments()
			},
		},
		{
			name:   "GetTags",
			device: &Device{Tags: []*Tag{{Name: String("tag-1")}}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(d *Device) interface{} {
				return d.GetTags()
			},
		},
		{
			name:   "GetPrimaryIp4",
			device: &Device{PrimaryIp4: &IPAddress{Address: String("192.168.1.1")}},
			expected: &diodepb.IPAddress{
				Address: "192.168.1.1",
			},
			method: func(d *Device) interface{} {
				return d.GetPrimaryIp4()
			},
		},
		{
			name:   "GetPrimaryIp6",
			device: &Device{PrimaryIp6: &IPAddress{Address: String("::1")}},
			expected: &diodepb.IPAddress{
				Address: "::1",
			},
			method: func(d *Device) interface{} {
				return d.GetPrimaryIp6()
			},
		},
		{
			name:   "ConvertToProtoMessage",
			device: &Device{Name: String("device-1")},
			expected: &diodepb.Device{
				Name: String("device-1"),
			},
			method: func(d *Device) interface{} {
				return d.ConvertToProtoMessage()
			},
		},
		{
			name:   "ConvertToProtoEntity",
			device: &Device{Name: String("device-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Device{
					Device: &diodepb.Device{
						Name: String("device-1"),
					},
				},
			},
			method: func(d *Device) interface{} {
				return d.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.device))
		})
	}
}

func TestDeviceTypeMethods(t *testing.T) {
	tests := []struct {
		name       string
		deviceType *DeviceType
		expected   interface{}
		method     func(*DeviceType) interface{}
	}{
		{
			name:       "GetModel",
			deviceType: &DeviceType{Model: String("model-1")},
			expected:   "model-1",
			method: func(d *DeviceType) interface{} {
				return d.GetModel()
			},
		},
		{
			name: "GetManufacturer",
			deviceType: &DeviceType{Manufacturer: &Manufacturer{
				Name: String("manufacturer-1"),
			}},
			expected: &diodepb.Manufacturer{
				Name: "manufacturer-1",
			},
			method: func(d *DeviceType) interface{} {
				return d.GetManufacturer()
			},
		},
		{
			name:       "GetPartNumber",
			deviceType: &DeviceType{PartNumber: String("part-1")},
			expected:   String("part-1"),
			method: func(d *DeviceType) interface{} {
				return d.GetPartNumber()
			},
		},
		{
			name:       "GetDescription",
			deviceType: &DeviceType{Description: String("description")},
			expected:   String("description"),
			method: func(d *DeviceType) interface{} {
				return d.GetDescription()
			},
		},
		{
			name:       "GetComments",
			deviceType: &DeviceType{Comments: String("comments")},
			expected:   String("comments"),
			method: func(d *DeviceType) interface{} {
				return d.GetComments()
			},
		},
		{
			name:       "GetTags",
			deviceType: &DeviceType{Tags: []*Tag{{Name: String("tag-1")}}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(d *DeviceType) interface{} {
				return d.GetTags()
			},
		},
		{
			name:       "ConvertToProtoMessage",
			deviceType: &DeviceType{Model: String("model-1")},
			expected: &diodepb.DeviceType{
				Model: "model-1",
			},
			method: func(d *DeviceType) interface{} {
				return d.ConvertToProtoMessage()
			},
		},
		{
			name:       "ConvertToProtoEntity",
			deviceType: &DeviceType{Model: String("model-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_DeviceType{
					DeviceType: &diodepb.DeviceType{
						Model: "model-1",
					},
				},
			},
			method: func(d *DeviceType) interface{} {
				return d.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.deviceType))
		})
	}
}

func TestIPAddressMethods(t *testing.T) {
	tests := []struct {
		name      string
		ipAddress *IPAddress
		expected  interface{}
		method    func(*IPAddress) interface{}
	}{
		{
			name:      "ConvertToProtoMessage",
			ipAddress: &IPAddress{Address: String("192.168.1.1")},
			expected: &diodepb.IPAddress{
				Address: "192.168.1.1",
			},
			method: func(ip *IPAddress) interface{} {
				return ip.ConvertToProtoMessage()
			},
		},
		{
			name:      "GetAddress",
			ipAddress: &IPAddress{Address: String("192.168.1.1")},
			expected:  "192.168.1.1",
			method: func(ip *IPAddress) interface{} {
				return ip.GetAddress()
			},
		},
		{
			name:      "GetAssignedObject",
			ipAddress: &IPAddress{AssignedObject: &Interface{}},
			expected:  &diodepb.IPAddress_AssignedObjectInterface{AssignedObjectInterface: &diodepb.Interface{}},
			method: func(ip *IPAddress) interface{} {
				return ip.GetAssignedObject()
			},
		},
		{
			name:      "GetStatus",
			ipAddress: &IPAddress{Status: String("active")},
			expected:  String("active"),
			method: func(ip *IPAddress) interface{} {
				return ip.GetStatus()
			},
		},
		{
			name:      "GetRole",
			ipAddress: &IPAddress{Role: String("admin")},
			expected:  String("admin"),
			method: func(ip *IPAddress) interface{} {
				return ip.GetRole()
			},
		},
		{
			name:      "GetDnsName",
			ipAddress: &IPAddress{DnsName: String("example.com")},
			expected:  String("example.com"),
			method: func(ip *IPAddress) interface{} {
				return ip.GetDnsName()
			},
		},
		{
			name:      "GetDescription",
			ipAddress: &IPAddress{Description: String("Test description")},
			expected:  String("Test description"),
			method: func(ip *IPAddress) interface{} {
				return ip.GetDescription()
			},
		},
		{
			name:      "GetComments",
			ipAddress: &IPAddress{Comments: String("Test comments")},
			expected:  String("Test comments"),
			method: func(ip *IPAddress) interface{} {
				return ip.GetComments()
			},
		},
		{
			name:      "GetTags",
			ipAddress: &IPAddress{Tags: []*Tag{{Name: String("tag-1")}}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(ip *IPAddress) interface{} {
				return ip.GetTags()
			},
		},
		{
			name:      "ConvertToProtoEntity",
			ipAddress: &IPAddress{Address: String("192.168.1.1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_IpAddress{
					IpAddress: &diodepb.IPAddress{
						Address: "192.168.1.1",
					},
				},
			},
			method: func(ip *IPAddress) interface{} {
				return ip.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.ipAddress))
		})
	}
}

func TestInterfaceMethods(t *testing.T) {
	tests := []struct {
		name     string
		iface    *Interface
		expected interface{}
		method   func(*Interface) interface{}
	}{
		{
			name:     "GetDevice",
			iface:    &Interface{Device: &Device{Name: String("device-1")}},
			expected: &diodepb.Device{Name: String("device-1")},
			method: func(i *Interface) interface{} {
				return i.GetDevice()
			},
		},
		{
			name:     "GetName",
			iface:    &Interface{Name: String("eth0")},
			expected: "eth0",
			method: func(i *Interface) interface{} {
				return i.GetName()
			},
		},
		{
			name:     "GetLabel",
			iface:    &Interface{Label: String("label-1")},
			expected: String("label-1"),
			method: func(i *Interface) interface{} {
				return i.GetLabel()
			},
		},
		{
			name:     "GetType",
			iface:    &Interface{Type: String("type-1")},
			expected: "type-1",
			method: func(i *Interface) interface{} {
				return i.GetType()
			},
		},
		{
			name:     "GetEnabled",
			iface:    &Interface{Enabled: Bool(true)},
			expected: Bool(true),
			method: func(i *Interface) interface{} {
				return i.GetEnabled()
			},
		},
		{
			name:     "GetMtu",
			iface:    &Interface{Mtu: Int64(1500)},
			expected: Int64(1500),
			method: func(i *Interface) interface{} {
				return i.GetMtu()
			},
		},
		{
			name:     "GetPrimaryMacAddress",
			iface:    &Interface{PrimaryMacAddress: &MACAddress{MacAddress: String("00:1A:2B:3C:4D:5E")}},
			expected: &diodepb.MACAddress{MacAddress: "00:1A:2B:3C:4D:5E"},
			method: func(i *Interface) interface{} {
				return i.GetPrimaryMacAddress()
			},
		},
		{
			name:     "GetSpeed",
			iface:    &Interface{Speed: Int64(1000)},
			expected: Int64(1000),
			method: func(i *Interface) interface{} {
				return i.GetSpeed()
			},
		},
		{
			name:     "GetWwn",
			iface:    &Interface{Wwn: String("wwn-1")},
			expected: String("wwn-1"),
			method: func(i *Interface) interface{} {
				return i.GetWwn()
			},
		},
		{
			name:     "GetMgmtOnly",
			iface:    &Interface{MgmtOnly: Bool(true)},
			expected: Bool(true),
			method: func(i *Interface) interface{} {
				return i.GetMgmtOnly()
			},
		},
		{
			name:     "GetDescription",
			iface:    &Interface{Description: String("Test description")},
			expected: String("Test description"),
			method: func(i *Interface) interface{} {
				return i.GetDescription()
			},
		},
		{
			name:     "GetMarkConnected",
			iface:    &Interface{MarkConnected: Bool(true)},
			expected: Bool(true),
			method: func(i *Interface) interface{} {
				return i.GetMarkConnected()
			},
		},
		{
			name:     "GetMode",
			iface:    &Interface{Mode: String("mode-1")},
			expected: String("mode-1"),
			method: func(i *Interface) interface{} {
				return i.GetMode()
			},
		},
		{
			name:     "GetTags",
			iface:    &Interface{Tags: []*Tag{{Name: String("tag-1")}}},
			expected: []*diodepb.Tag{{Name: "tag-1"}},
			method: func(i *Interface) interface{} {
				return i.GetTags()
			},
		},
		{
			name:     "ConvertToProtoMessage",
			iface:    &Interface{Name: String("eth0")},
			expected: &diodepb.Interface{Name: "eth0"},
			method: func(i *Interface) interface{} {
				return i.ConvertToProtoMessage()
			},
		},
		{
			name:  "ConvertToProtoEntity",
			iface: &Interface{Name: String("eth0")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Interface{
					Interface: &diodepb.Interface{Name: "eth0"},
				},
			},
			method: func(i *Interface) interface{} {
				return i.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.iface))
		})
	}
}

func TestManufacturerMethods(t *testing.T) {
	tests := []struct {
		name         string
		manufacturer *Manufacturer
		expected     interface{}
		method       func(*Manufacturer) interface{}
	}{
		{
			name:         "GetName",
			manufacturer: &Manufacturer{Name: String("manufacturer-1")},
			expected:     "manufacturer-1",
			method: func(m *Manufacturer) interface{} {
				return m.GetName()
			},
		},
		{
			name:         "GetSlug",
			manufacturer: &Manufacturer{Slug: String("slug-1")},
			expected:     "slug-1",
			method: func(m *Manufacturer) interface{} {
				return m.GetSlug()
			},
		},
		{
			name:         "GetDescription",
			manufacturer: &Manufacturer{Description: String("Test description")},
			expected:     String("Test description"),
			method: func(m *Manufacturer) interface{} {
				return m.GetDescription()
			},
		},
		{
			name: "GetTags",
			manufacturer: &Manufacturer{Tags: []*Tag{
				{Name: String("tag-1")},
			}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(m *Manufacturer) interface{} {
				return m.GetTags()
			},
		},
		{
			name:         "ConvertToProtoMessage",
			manufacturer: &Manufacturer{Name: String("manufacturer-1")},
			expected: &diodepb.Manufacturer{
				Name: "manufacturer-1",
			},
			method: func(m *Manufacturer) interface{} {
				return m.ConvertToProtoMessage()
			},
		},
		{
			name:         "ConvertToProtoEntity",
			manufacturer: &Manufacturer{Name: String("manufacturer-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Manufacturer{
					Manufacturer: &diodepb.Manufacturer{
						Name: "manufacturer-1",
					},
				},
			},
			method: func(m *Manufacturer) interface{} {
				return m.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.manufacturer))
		})
	}
}

func TestPlatformMethods(t *testing.T) {
	tests := []struct {
		name     string
		platform *Platform
		expected interface{}
		method   func(*Platform) interface{}
	}{
		{
			name:     "GetName",
			platform: &Platform{Name: String("platform-1")},
			expected: "platform-1",
			method: func(p *Platform) interface{} {
				return p.GetName()
			},
		},
		{
			name:     "GetSlug",
			platform: &Platform{Slug: String("slug-1")},
			expected: "slug-1",
			method: func(p *Platform) interface{} {
				return p.GetSlug()
			},
		},
		{
			name:     "GetManufacturer",
			platform: &Platform{Manufacturer: &Manufacturer{Name: String("manufacturer-1")}},
			expected: &diodepb.Manufacturer{Name: "manufacturer-1"},
			method: func(p *Platform) interface{} {
				return p.GetManufacturer()
			},
		},
		{
			name:     "GetDescription",
			platform: &Platform{Description: String("Test description")},
			expected: String("Test description"),
			method: func(p *Platform) interface{} {
				return p.GetDescription()
			},
		},
		{
			name: "GetTags",
			platform: &Platform{Tags: []*Tag{
				{Name: String("tag-1")},
			}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(p *Platform) interface{} {
				return p.GetTags()
			},
		},
		{
			name:     "ConvertToProtoMessage",
			platform: &Platform{Name: String("platform-1")},
			expected: &diodepb.Platform{
				Name: "platform-1",
			},
			method: func(p *Platform) interface{} {
				return p.ConvertToProtoMessage()
			},
		},
		{
			name:     "ConvertToProtoEntity",
			platform: &Platform{Name: String("platform-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Platform{
					Platform: &diodepb.Platform{Name: "platform-1"},
				},
			},
			method: func(p *Platform) interface{} {
				return p.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.platform))
		})
	}
}

func TestPrefixMethods(t *testing.T) {
	tests := []struct {
		name     string
		prefix   *Prefix
		expected interface{}
		method   func(*Prefix) interface{}
	}{
		{
			name:     "GetPrefix",
			prefix:   &Prefix{Prefix: String("prefix-1")},
			expected: "prefix-1",
			method: func(p *Prefix) interface{} {
				return p.GetPrefix()
			},
		},
		{
			name:     "GetScope",
			prefix:   &Prefix{Scope: &Site{Name: String("site-1")}},
			expected: &diodepb.Prefix_ScopeSite{ScopeSite: &diodepb.Site{Name: "site-1"}},
			method: func(p *Prefix) interface{} {
				return p.GetScope()
			},
		},
		{
			name:     "GetStatus",
			prefix:   &Prefix{Status: String("active")},
			expected: String("active"),
			method: func(p *Prefix) interface{} {
				return p.GetStatus()
			},
		},
		{
			name:     "GetIsPool",
			prefix:   &Prefix{IsPool: Bool(true)},
			expected: Bool(true),
			method: func(p *Prefix) interface{} {
				return p.GetIsPool()
			},
		},
		{
			name:     "GetMarkUtilized",
			prefix:   &Prefix{MarkUtilized: Bool(true)},
			expected: Bool(true),
			method: func(p *Prefix) interface{} {
				return p.GetMarkUtilized()
			},
		},
		{
			name:     "GetDescription",
			prefix:   &Prefix{Description: String("Test description")},
			expected: String("Test description"),
			method: func(p *Prefix) interface{} {
				return p.GetDescription()
			},
		},
		{
			name:     "GetComments",
			prefix:   &Prefix{Comments: String("Test comments")},
			expected: String("Test comments"),
			method: func(p *Prefix) interface{} {
				return p.GetComments()
			},
		},
		{
			name: "GetTags",
			prefix: &Prefix{Tags: []*Tag{
				{Name: String("tag-1")},
			}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(p *Prefix) interface{} {
				return p.GetTags()
			},
		},
		{
			name:     "ConvertToProtoMessage",
			prefix:   &Prefix{Prefix: String("prefix-1")},
			expected: &diodepb.Prefix{Prefix: "prefix-1"},
			method: func(p *Prefix) interface{} {
				return p.ConvertToProtoMessage()
			},
		},
		{
			name:   "ConvertToProtoEntity",
			prefix: &Prefix{Prefix: String("prefix-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Prefix{
					Prefix: &diodepb.Prefix{Prefix: "prefix-1"},
				},
			},
			method: func(p *Prefix) interface{} {
				return p.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.prefix))
		})
	}
}

func TestDeviceRoleMethods(t *testing.T) {
	tests := []struct {
		name     string
		role     *DeviceRole
		expected interface{}
		method   func(*DeviceRole) interface{}
	}{
		{
			name:     "GetName",
			role:     &DeviceRole{Name: String("role-1")},
			expected: "role-1",
			method: func(r *DeviceRole) interface{} {
				return r.GetName()
			},
		},
		{
			name:     "GetSlug",
			role:     &DeviceRole{Slug: String("slug-1")},
			expected: "slug-1",
			method: func(r *DeviceRole) interface{} {
				return r.GetSlug()
			},
		},
		{
			name:     "GetColor",
			role:     &DeviceRole{Color: String("color-1")},
			expected: String("color-1"),
			method: func(r *DeviceRole) interface{} {
				return r.GetColor()
			},
		},
		{
			name:     "GetDescription",
			role:     &DeviceRole{Description: String("Test description")},
			expected: String("Test description"),
			method: func(r *DeviceRole) interface{} {
				return r.GetDescription()
			},
		},
		{
			name: "GetTags",
			role: &DeviceRole{Tags: []*Tag{
				{Name: String("tag-1")},
			}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(r *DeviceRole) interface{} {
				return r.GetTags()
			},
		},
		{
			name:     "ConvertToProtoMessage",
			role:     &DeviceRole{Name: String("role-1")},
			expected: &diodepb.DeviceRole{Name: "role-1"},
			method: func(r *DeviceRole) interface{} {
				return r.ConvertToProtoMessage()
			},
		},
		{
			name: "ConvertToProtoEntity",
			role: &DeviceRole{Name: String("role-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_DeviceRole{
					DeviceRole: &diodepb.DeviceRole{Name: "role-1"},
				},
			},
			method: func(r *DeviceRole) interface{} {
				return r.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.role))
		})
	}
}

func TestSiteMethods(t *testing.T) {
	tests := []struct {
		name     string
		site     *Site
		expected interface{}
		method   func(*Site) interface{}
	}{
		{
			name:     "GetName",
			site:     &Site{Name: String("site-1")},
			expected: "site-1",
			method: func(s *Site) interface{} {
				return s.GetName()
			},
		},
		{
			name:     "GetSlug",
			site:     &Site{Slug: String("slug-1")},
			expected: "slug-1",
			method: func(s *Site) interface{} {
				return s.GetSlug()
			},
		},
		{
			name:     "GetStatus",
			site:     &Site{Status: String("active")},
			expected: String("active"),
			method: func(s *Site) interface{} {
				return s.GetStatus()
			},
		},
		{
			name:     "GetFacility",
			site:     &Site{Facility: String("facility-1")},
			expected: String("facility-1"),
			method: func(s *Site) interface{} {
				return s.GetFacility()
			},
		},
		{
			name:     "GetTimeZone",
			site:     &Site{TimeZone: String("UTC")},
			expected: String("UTC"),
			method: func(s *Site) interface{} {
				return s.GetTimeZone()
			},
		},
		{
			name:     "GetDescription",
			site:     &Site{Description: String("Test description")},
			expected: String("Test description"),
			method: func(s *Site) interface{} {
				return s.GetDescription()
			},
		},
		{
			name:     "GetComments",
			site:     &Site{Comments: String("Test comments")},
			expected: String("Test comments"),
			method: func(s *Site) interface{} {
				return s.GetComments()
			},
		},
		{
			name: "GetTags",
			site: &Site{Tags: []*Tag{
				{Name: String("tag-1")},
			}},
			expected: []*diodepb.Tag{
				{Name: "tag-1"},
			},
			method: func(s *Site) interface{} {
				return s.GetTags()
			},
		},
		{
			name:     "ConvertToProtoMessage",
			site:     &Site{Name: String("site-1")},
			expected: &diodepb.Site{Name: "site-1"},
			method: func(s *Site) interface{} {
				return s.ConvertToProtoMessage()
			},
		},
		{
			name: "ConvertToProtoEntity",
			site: &Site{Name: String("site-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Site{
					Site: &diodepb.Site{Name: "site-1"},
				},
			},
			method: func(s *Site) interface{} {
				return s.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.site))
		})
	}
}

func TestTagMethods(t *testing.T) {
	tests := []struct {
		name     string
		tag      *Tag
		expected interface{}
		method   func(*Tag) interface{}
	}{
		{
			name:     "GetName",
			tag:      &Tag{Name: String("tag-1")},
			expected: "tag-1",
			method: func(t *Tag) interface{} {
				return t.GetName()
			},
		},
		{
			name:     "GetSlug",
			tag:      &Tag{Slug: String("slug-1")},
			expected: "slug-1",
			method: func(t *Tag) interface{} {
				return t.GetSlug()
			},
		},
		{
			name:     "GetColor",
			tag:      &Tag{Color: String("color-1")},
			expected: String("color-1"),
			method: func(t *Tag) interface{} {
				return t.GetColor()
			},
		},
		{
			name:     "ConvertToProtoMessage",
			tag:      &Tag{Name: String("tag-1")},
			expected: &diodepb.Tag{Name: "tag-1"},
			method: func(t *Tag) interface{} {
				return t.ConvertToProtoMessage()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.tag))
		})
	}
}

func TestClusterMethods(t *testing.T) {
	tests := []struct {
		name     string
		cluster  *Cluster
		expected interface{}
		method   func(*Cluster) interface{}
	}{
		{
			name:     "GetName",
			cluster:  &Cluster{Name: String("cluster-1")},
			expected: "cluster-1",
			method: func(c *Cluster) interface{} {
				return c.GetName()
			},
		},
		{
			name:     "GetType",
			cluster:  &Cluster{Type: &ClusterType{}},
			expected: &diodepb.ClusterType{},
			method: func(c *Cluster) interface{} {
				return c.GetType()
			},
		},
		{
			name:     "GetGroup",
			cluster:  &Cluster{Group: &ClusterGroup{}},
			expected: &diodepb.ClusterGroup{},
			method: func(c *Cluster) interface{} {
				return c.GetGroup()
			},
		},
		{
			name:    "GetScope",
			cluster: &Cluster{Scope: &Site{}},
			expected: &diodepb.Cluster_ScopeSite{
				ScopeSite: &diodepb.Site{},
			},
			method: func(c *Cluster) interface{} {
				return c.GetScope()
			},
		},
		{
			name:     "GetStatus",
			cluster:  &Cluster{Status: String("active")},
			expected: String("active"),
			method: func(c *Cluster) interface{} {
				return c.GetStatus()
			},
		},
		{
			name:     "GetDescription",
			cluster:  &Cluster{Description: String("Test description")},
			expected: String("Test description"),
			method: func(c *Cluster) interface{} {
				return c.GetDescription()
			},
		},
		{
			name:     "GetTags",
			cluster:  &Cluster{Tags: []*Tag{{Name: String("tag-1")}}},
			expected: []*diodepb.Tag{{Name: "tag-1"}},
			method: func(c *Cluster) interface{} {
				return c.GetTags()
			},
		},
		{
			name:     "ConvertToProtoMessage",
			cluster:  &Cluster{Name: String("cluster-1")},
			expected: &diodepb.Cluster{Name: "cluster-1"},
			method: func(c *Cluster) interface{} {
				return c.ConvertToProtoMessage()
			},
		},
		{
			name:    "ConvertToProtoEntity",
			cluster: &Cluster{Name: String("cluster-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Cluster{
					Cluster: &diodepb.Cluster{Name: "cluster-1"},
				},
			},
			method: func(c *Cluster) interface{} {
				return c.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.cluster))
		})
	}
}

func TestClusterGroupMethods(t *testing.T) {
	tests := []struct {
		name         string
		clusterGroup *ClusterGroup
		expected     interface{}
		method       func(*ClusterGroup) interface{}
	}{
		{
			name:         "GetName",
			clusterGroup: &ClusterGroup{Name: String("group-1")},
			expected:     "group-1",
			method: func(cg *ClusterGroup) interface{} {
				return cg.GetName()
			},
		},
		{
			name:         "GetSlug",
			clusterGroup: &ClusterGroup{Slug: String("group-slug")},
			expected:     "group-slug",
			method: func(cg *ClusterGroup) interface{} {
				return cg.GetSlug()
			},
		},
		{
			name:         "GetDescription",
			clusterGroup: &ClusterGroup{Description: String("Test description")},
			expected:     String("Test description"),
			method: func(cg *ClusterGroup) interface{} {
				return cg.GetDescription()
			},
		},
		{
			name:         "GetTags",
			clusterGroup: &ClusterGroup{Tags: []*Tag{{Name: String("tag-1")}}},
			expected:     []*diodepb.Tag{{Name: "tag-1"}},
			method: func(cg *ClusterGroup) interface{} {
				return cg.GetTags()
			},
		},
		{
			name:         "ConvertToProtoMessage",
			clusterGroup: &ClusterGroup{Name: String("group-1")},
			expected:     &diodepb.ClusterGroup{Name: "group-1"},
			method: func(cg *ClusterGroup) interface{} {
				return cg.ConvertToProtoMessage()
			},
		},
		{
			name:         "ConvertToProtoEntity",
			clusterGroup: &ClusterGroup{Name: String("group-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_ClusterGroup{
					ClusterGroup: &diodepb.ClusterGroup{Name: "group-1"},
				},
			},
			method: func(cg *ClusterGroup) interface{} {
				return cg.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.clusterGroup))
		})
	}
}

func TestClusterTypeMethods(t *testing.T) {
	tests := []struct {
		name        string
		clusterType *ClusterType
		expected    interface{}
		method      func(*ClusterType) interface{}
	}{
		{
			name:        "GetName",
			clusterType: &ClusterType{Name: String("type-1")},
			expected:    "type-1",
			method: func(ct *ClusterType) interface{} {
				return ct.GetName()
			},
		},
		{
			name:        "GetSlug",
			clusterType: &ClusterType{Slug: String("type-slug")},
			expected:    "type-slug",
			method: func(ct *ClusterType) interface{} {
				return ct.GetSlug()
			},
		},
		{
			name:        "GetDescription",
			clusterType: &ClusterType{Description: String("Test description")},
			expected:    String("Test description"),
			method: func(ct *ClusterType) interface{} {
				return ct.GetDescription()
			},
		},
		{
			name:        "GetTags",
			clusterType: &ClusterType{Tags: []*Tag{{Name: String("tag-1")}}},
			expected:    []*diodepb.Tag{{Name: "tag-1"}},
			method: func(ct *ClusterType) interface{} {
				return ct.GetTags()
			},
		},
		{
			name:        "ConvertToProtoMessage",
			clusterType: &ClusterType{Name: String("type-1")},
			expected:    &diodepb.ClusterType{Name: "type-1"},
			method: func(ct *ClusterType) interface{} {
				return ct.ConvertToProtoMessage()
			},
		},
		{
			name:        "ConvertToProtoEntity",
			clusterType: &ClusterType{Name: String("type-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_ClusterType{
					ClusterType: &diodepb.ClusterType{Name: "type-1"},
				},
			},
			method: func(ct *ClusterType) interface{} {
				return ct.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.clusterType))
		})
	}
}

func TestVirtualMachineMethods(t *testing.T) {
	tests := []struct {
		name           string
		virtualMachine *VirtualMachine
		expected       interface{}
		method         func(*VirtualMachine) interface{}
	}{
		{
			name:           "GetName",
			virtualMachine: &VirtualMachine{Name: String("vm-1")},
			expected:       "vm-1",
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetName()
			},
		},
		{
			name:           "GetStatus",
			virtualMachine: &VirtualMachine{Status: String("running")},
			expected:       String("running"),
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetStatus()
			},
		},
		{
			name:           "GetSite",
			virtualMachine: &VirtualMachine{Site: &Site{Name: String("site-1")}},
			expected:       &diodepb.Site{Name: "site-1"},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetSite()
			},
		},
		{
			name:           "GetCluster",
			virtualMachine: &VirtualMachine{Cluster: &Cluster{Name: String("cluster-1")}},
			expected:       &diodepb.Cluster{Name: "cluster-1"},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetCluster()
			},
		},
		{
			name:           "GetRole",
			virtualMachine: &VirtualMachine{Role: &DeviceRole{Name: String("role-1")}},
			expected:       &diodepb.DeviceRole{Name: "role-1"},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetRole()
			},
		},
		{
			name:           "GetDevice",
			virtualMachine: &VirtualMachine{Device: &Device{Name: String("device-1")}},
			expected:       &diodepb.Device{Name: String("device-1")},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetDevice()
			},
		},
		{
			name:           "GetPlatform",
			virtualMachine: &VirtualMachine{Platform: &Platform{Name: String("platform-1")}},
			expected:       &diodepb.Platform{Name: "platform-1"},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetPlatform()
			},
		},
		{
			name:           "GetPrimaryIp4",
			virtualMachine: &VirtualMachine{PrimaryIp4: &IPAddress{Address: String("192.168.1.1")}},
			expected:       &diodepb.IPAddress{Address: "192.168.1.1"},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetPrimaryIp4()
			},
		},
		{
			name:           "GetPrimaryIp6",
			virtualMachine: &VirtualMachine{PrimaryIp6: &IPAddress{Address: String("::1")}},
			expected:       &diodepb.IPAddress{Address: "::1"},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetPrimaryIp6()
			},
		},
		{
			name:           "GetVcpus",
			virtualMachine: &VirtualMachine{Vcpus: Float64(4)},
			expected:       Float64(4),
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetVcpus()
			},
		},
		{
			name:           "GetMemory",
			virtualMachine: &VirtualMachine{Memory: Int64(8192)},
			expected:       Int64(8192),
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetMemory()
			},
		},
		{
			name:           "GetDisk",
			virtualMachine: &VirtualMachine{Disk: Int64(100)},
			expected:       Int64(100),
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetDisk()
			},
		},
		{
			name:           "GetDescription",
			virtualMachine: &VirtualMachine{Description: String("Test VM")},
			expected:       String("Test VM"),
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetDescription()
			},
		},
		{
			name:           "GetComments",
			virtualMachine: &VirtualMachine{Comments: String("No comments")},
			expected:       String("No comments"),
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetComments()
			},
		},
		{
			name:           "GetTags",
			virtualMachine: &VirtualMachine{Tags: []*Tag{{Name: String("tag-1")}}},
			expected:       []*diodepb.Tag{{Name: "tag-1"}},
			method: func(vm *VirtualMachine) interface{} {
				return vm.GetTags()
			},
		},
		{
			name:           "ConvertToProtoMessage",
			virtualMachine: &VirtualMachine{Name: String("vm-1")},
			expected:       &diodepb.VirtualMachine{Name: "vm-1"},
			method: func(vm *VirtualMachine) interface{} {
				return vm.ConvertToProtoMessage()
			},
		},
		{
			name:           "ConvertToProtoEntity",
			virtualMachine: &VirtualMachine{Name: String("vm-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_VirtualMachine{
					VirtualMachine: &diodepb.VirtualMachine{Name: "vm-1"},
				},
			},
			method: func(vm *VirtualMachine) interface{} {
				return vm.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.virtualMachine))
		})
	}
}

func TestVMInterfaceMethods(t *testing.T) {
	tests := []struct {
		name        string
		vmInterface *VMInterface
		expected    interface{}
		method      func(*VMInterface) interface{}
	}{
		{
			name:        "GetName",
			vmInterface: &VMInterface{Name: String("vminterface-1")},
			expected:    "vminterface-1",
			method: func(vmi *VMInterface) interface{} {
				return vmi.GetName()
			},
		},
		{
			name:        "GetPrimaryMacAddress",
			vmInterface: &VMInterface{PrimaryMacAddress: &MACAddress{MacAddress: String("00:1A:2B:3C:4D:5E")}},
			expected:    &diodepb.MACAddress{MacAddress: "00:1A:2B:3C:4D:5E"},
			method: func(vmi *VMInterface) interface{} {
				return vmi.GetPrimaryMacAddress()
			},
		},
		{
			name:        "GetDescription",
			vmInterface: &VMInterface{Description: String("Test description")},
			expected:    String("Test description"),
			method: func(vmi *VMInterface) interface{} {
				return vmi.GetDescription()
			},
		},
		{
			name:        "GetTags",
			vmInterface: &VMInterface{Tags: []*Tag{{Name: String("tag-1")}}},
			expected:    []*diodepb.Tag{{Name: "tag-1"}},
			method: func(vmi *VMInterface) interface{} {
				return vmi.GetTags()
			},
		},
		{
			name:        "ConvertToProtoMessage",
			vmInterface: &VMInterface{Name: String("vminterface-1")},
			expected:    &diodepb.VMInterface{Name: "vminterface-1"},
			method: func(vmi *VMInterface) interface{} {
				return vmi.ConvertToProtoMessage()
			},
		},
		{
			name:        "ConvertToProtoEntity",
			vmInterface: &VMInterface{Name: String("vminterface-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_VmInterface{
					VmInterface: &diodepb.VMInterface{Name: "vminterface-1"},
				},
			},
			method: func(vmi *VMInterface) interface{} {
				return vmi.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.vmInterface))
		})
	}
}

func TestVirtualDiskMethods(t *testing.T) {
	tests := []struct {
		name        string
		virtualDisk *VirtualDisk
		expected    interface{}
		method      func(*VirtualDisk) interface{}
	}{
		{
			name:        "GetVirtualMachine",
			virtualDisk: &VirtualDisk{VirtualMachine: &VirtualMachine{Name: String("vm-1")}},
			expected:    &diodepb.VirtualMachine{Name: "vm-1"},
			method: func(vd *VirtualDisk) interface{} {
				return vd.GetVirtualMachine()
			},
		},
		{
			name:        "GetName",
			virtualDisk: &VirtualDisk{Name: String("disk-1")},
			expected:    "disk-1",
			method: func(vd *VirtualDisk) interface{} {
				return vd.GetName()
			},
		},
		{
			name:        "GetSize",
			virtualDisk: &VirtualDisk{Size: Int64(1024)},
			expected:    *Int64(1024),
			method: func(vd *VirtualDisk) interface{} {
				return vd.GetSize()
			},
		},
		{
			name:        "GetDescription",
			virtualDisk: &VirtualDisk{Description: String("Test disk")},
			expected:    String("Test disk"),
			method: func(vd *VirtualDisk) interface{} {
				return vd.GetDescription()
			},
		},
		{
			name:        "GetTags",
			virtualDisk: &VirtualDisk{Tags: []*Tag{{Name: String("tag-1")}}},
			expected:    []*diodepb.Tag{{Name: "tag-1"}},
			method: func(vd *VirtualDisk) interface{} {
				return vd.GetTags()
			},
		},
		{
			name:        "ConvertToProtoMessage",
			virtualDisk: &VirtualDisk{Name: String("disk-1")},
			expected:    &diodepb.VirtualDisk{Name: "disk-1"},
			method: func(vd *VirtualDisk) interface{} {
				return vd.ConvertToProtoMessage()
			},
		},
		{
			name:        "ConvertToProtoEntity",
			virtualDisk: &VirtualDisk{Name: String("disk-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_VirtualDisk{
					VirtualDisk: &diodepb.VirtualDisk{Name: "disk-1"},
				},
			},
			method: func(vd *VirtualDisk) interface{} {
				return vd.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.virtualDisk))
		})
	}
}

func TestMetadataConversion(t *testing.T) {
	tests := []struct {
		name            string
		entity          Entity
		metadata        Metadata
		getMetadataFunc func(*diodepb.Entity) *structpb.Struct
	}{
		{
			name: "ASN with metadata",
			entity: &ASN{
				Asn: Int64(64512),
				Metadata: Metadata{
					"source":      "import",
					"imported_at": "2024-01-01T00:00:00Z",
					"priority":    1,
				},
			},
			metadata: Metadata{
				"source":      "import",
				"imported_at": "2024-01-01T00:00:00Z",
				"priority":    1,
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetAsn().GetMetadata()
			},
		},
		{
			name: "Device with metadata",
			entity: &Device{
				Name: String("device-1"),
				Metadata: Metadata{
					"environment": "production",
					"owner":       "team-a",
					"cost_center": 12345,
				},
			},
			metadata: Metadata{
				"environment": "production",
				"owner":       "team-a",
				"cost_center": 12345,
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetDevice().GetMetadata()
			},
		},
		{
			name: "IPAddress with metadata",
			entity: &IPAddress{
				Address: String("192.168.1.1/24"),
				Metadata: Metadata{
					"last_seen":  "2024-01-01T00:00:00Z",
					"discovered": true,
					"scan_id":    "scan-123",
				},
			},
			metadata: Metadata{
				"last_seen":  "2024-01-01T00:00:00Z",
				"discovered": true,
				"scan_id":    "scan-123",
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetIpAddress().GetMetadata()
			},
		},
		{
			name: "Site with metadata",
			entity: &Site{
				Name: String("site-1"),
				Metadata: Metadata{
					"region":    "us-west",
					"capacity":  500,
					"is_active": true,
				},
			},
			metadata: Metadata{
				"region":    "us-west",
				"capacity":  500,
				"is_active": true,
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetSite().GetMetadata()
			},
		},
		{
			name: "Interface with metadata",
			entity: &Interface{
				Name: String("eth0"),
				Metadata: Metadata{
					"monitored":    true,
					"vlan_id":      100,
					"uplink_speed": "10G",
				},
			},
			metadata: Metadata{
				"monitored":    true,
				"vlan_id":      100,
				"uplink_speed": "10G",
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetInterface().GetMetadata()
			},
		},
		{
			name: "VirtualMachine with metadata",
			entity: &VirtualMachine{
				Name: String("vm-1"),
				Metadata: Metadata{
					"hypervisor": "vmware",
					"cluster_id": "cluster-001",
					"tier":       2,
				},
			},
			metadata: Metadata{
				"hypervisor": "vmware",
				"cluster_id": "cluster-001",
				"tier":       2,
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetVirtualMachine().GetMetadata()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protoEntity := tt.entity.ConvertToProtoEntity()
			metadata := tt.getMetadataFunc(protoEntity)

			require.NotNil(t, metadata)

			for key, expectedValue := range tt.metadata {
				actualValue := metadata.Fields[key]
				require.NotNil(t, actualValue, "metadata key %s should be present", key)

				switch v := expectedValue.(type) {
				case string:
					require.Equal(t, v, actualValue.GetStringValue())
				case int:
					require.Equal(t, float64(v), actualValue.GetNumberValue())
				case bool:
					require.Equal(t, v, actualValue.GetBoolValue())
				}
			}
		})
	}
}

func TestMetadataConversionNil(t *testing.T) {
	tests := []struct {
		name            string
		entity          Entity
		getMetadataFunc func(*diodepb.Entity) *structpb.Struct
	}{
		{
			name:   "ASN without metadata",
			entity: &ASN{Asn: Int64(64512)},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetAsn().GetMetadata()
			},
		},
		{
			name:   "Device without metadata",
			entity: &Device{Name: String("device-1")},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetDevice().GetMetadata()
			},
		},
		{
			name:   "IPAddress without metadata",
			entity: &IPAddress{Address: String("192.168.1.1/24")},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetIpAddress().GetMetadata()
			},
		},
		{
			name:   "Site without metadata",
			entity: &Site{Name: String("site-1")},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetSite().GetMetadata()
			},
		},
		{
			name:   "Interface without metadata",
			entity: &Interface{Name: String("eth0")},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetInterface().GetMetadata()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protoEntity := tt.entity.ConvertToProtoEntity()

			metadata := tt.getMetadataFunc(protoEntity)
			require.Nil(t, metadata)
		})
	}
}

func TestMetadataConversionEmpty(t *testing.T) {
	tests := []struct {
		name            string
		entity          Entity
		getMetadataFunc func(*diodepb.Entity) *structpb.Struct
	}{
		{
			name: "ASN with empty metadata",
			entity: &ASN{
				Asn:      Int64(64512),
				Metadata: Metadata{},
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetAsn().GetMetadata()
			},
		},
		{
			name: "Device with empty metadata",
			entity: &Device{
				Name:     String("device-1"),
				Metadata: Metadata{},
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetDevice().GetMetadata()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protoEntity := tt.entity.ConvertToProtoEntity()
			metadata := tt.getMetadataFunc(protoEntity)

			require.NotNil(t, metadata)
			require.Empty(t, metadata.Fields)
		})
	}
}

func TestMetadataConversionComplexTypes(t *testing.T) {
	tests := []struct {
		name            string
		entity          Entity
		metadata        Metadata
		getMetadataFunc func(*diodepb.Entity) *structpb.Struct
	}{
		{
			name: "Device with nested metadata",
			entity: &Device{
				Name: String("device-1"),
				Metadata: Metadata{
					"config": map[string]interface{}{
						"enabled": true,
						"timeout": 30,
					},
					"tags": []interface{}{"production", "critical"},
				},
			},
			metadata: Metadata{
				"config": map[string]interface{}{
					"enabled": true,
					"timeout": 30,
				},
				"tags": []interface{}{"production", "critical"},
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetDevice().GetMetadata()
			},
		},
		{
			name: "IPAddress with mixed types",
			entity: &IPAddress{
				Address: String("10.0.0.1/24"),
				Metadata: Metadata{
					"null_value":   nil,
					"string_value": "test",
					"number_value": 42.5,
					"bool_value":   false,
				},
			},
			metadata: Metadata{
				"null_value":   nil,
				"string_value": "test",
				"number_value": 42.5,
				"bool_value":   false,
			},
			getMetadataFunc: func(e *diodepb.Entity) *structpb.Struct {
				return e.GetIpAddress().GetMetadata()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protoEntity := tt.entity.ConvertToProtoEntity()
			metadata := tt.getMetadataFunc(protoEntity)

			require.NotNil(t, metadata)

			for key := range tt.metadata {
				_, exists := metadata.Fields[key]
				require.True(t, exists, "metadata key %s should be present", key)
			}
		})
	}
}

func TestMetadataConversionNestedEntities(t *testing.T) {
	tests := []struct {
		name             string
		entity           Entity
		parentMetadata   Metadata
		nestedMetadata   Metadata
		verifyNestedFunc func(*testing.T, *diodepb.Entity)
	}{
		{
			name: "Device with Site both having metadata",
			entity: &Device{
				Name: String("device-1"),
				Site: &Site{
					Name: String("site-1"),
					Metadata: Metadata{
						"site_region":   "us-west",
						"site_capacity": 1000,
					},
				},
				Metadata: Metadata{
					"device_owner": "team-a",
					"device_tier":  1,
				},
			},
			parentMetadata: Metadata{
				"device_owner": "team-a",
				"device_tier":  1,
			},
			nestedMetadata: Metadata{
				"site_region":   "us-west",
				"site_capacity": 1000,
			},
			verifyNestedFunc: func(t *testing.T, protoEntity *diodepb.Entity) {
				deviceEntity := protoEntity.GetDevice()
				require.NotNil(t, deviceEntity)

				site := deviceEntity.GetSite()
				require.NotNil(t, site)
				require.Equal(t, "site-1", site.GetName())
			},
		},
		{
			name: "IPAddress with Interface both having metadata",
			entity: &IPAddress{
				Address: String("192.168.1.1/24"),
				AssignedObject: &Interface{
					Name: String("eth0"),
					Metadata: Metadata{
						"interface_speed": "10G",
						"interface_vlan":  100,
					},
				},
				Metadata: Metadata{
					"ip_discovered": true,
					"ip_scan_id":    "scan-123",
				},
			},
			parentMetadata: Metadata{
				"ip_discovered": true,
				"ip_scan_id":    "scan-123",
			},
			nestedMetadata: Metadata{
				"interface_speed": "10G",
				"interface_vlan":  100,
			},
			verifyNestedFunc: func(t *testing.T, protoEntity *diodepb.Entity) {
				ipEntity := protoEntity.GetIpAddress()
				require.NotNil(t, ipEntity)

				assignedObj := ipEntity.GetAssignedObjectInterface()
				require.NotNil(t, assignedObj)
				require.Equal(t, "eth0", assignedObj.GetName())
			},
		},
		{
			name: "VirtualMachine with Cluster both having metadata",
			entity: &VirtualMachine{
				Name: String("vm-1"),
				Cluster: &Cluster{
					Name: String("cluster-1"),
					Metadata: Metadata{
						"cluster_size": 5,
						"cluster_type": "production",
					},
				},
				Metadata: Metadata{
					"vm_hypervisor": "vmware",
					"vm_template":   "ubuntu-20.04",
				},
			},
			parentMetadata: Metadata{
				"vm_hypervisor": "vmware",
				"vm_template":   "ubuntu-20.04",
			},
			nestedMetadata: Metadata{
				"cluster_size": 5,
				"cluster_type": "production",
			},
			verifyNestedFunc: func(t *testing.T, protoEntity *diodepb.Entity) {
				vmEntity := protoEntity.GetVirtualMachine()
				require.NotNil(t, vmEntity)

				cluster := vmEntity.GetCluster()
				require.NotNil(t, cluster)
				require.Equal(t, "cluster-1", cluster.GetName())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert to proto entity
			protoEntity := tt.entity.ConvertToProtoEntity()

			// Get the parent entity's metadata from the specific entity type
			var parentMetadata *structpb.Struct
			switch protoEntity.GetEntity().(type) {
			case *diodepb.Entity_Device:
				parentMetadata = protoEntity.GetDevice().GetMetadata()
			case *diodepb.Entity_IpAddress:
				parentMetadata = protoEntity.GetIpAddress().GetMetadata()
			case *diodepb.Entity_VirtualMachine:
				parentMetadata = protoEntity.GetVirtualMachine().GetMetadata()
			}

			// Verify parent entity metadata is present
			require.NotNil(t, parentMetadata, "parent entity should have metadata")

			// Verify parent metadata fields
			for key, expectedValue := range tt.parentMetadata {
				actualValue := parentMetadata.Fields[key]
				require.NotNil(t, actualValue, "parent metadata key %s should be present", key)

				switch v := expectedValue.(type) {
				case string:
					require.Equal(t, v, actualValue.GetStringValue())
				case int:
					require.Equal(t, float64(v), actualValue.GetNumberValue())
				case bool:
					require.Equal(t, v, actualValue.GetBoolValue())
				}
			}

			// Verify nested entity structure
			tt.verifyNestedFunc(t, protoEntity)
		})
	}
}

func TestOwnerGroupMethods(t *testing.T) {
	tests := []struct {
		name       string
		ownerGroup *OwnerGroup
		expected   interface{}
		method     func(*OwnerGroup) interface{}
	}{
		{
			name:       "GetName",
			ownerGroup: &OwnerGroup{Name: String("owner-group-1")},
			expected:   "owner-group-1",
			method: func(o *OwnerGroup) interface{} {
				return o.GetName()
			},
		},
		{
			name:       "GetDescription",
			ownerGroup: &OwnerGroup{Description: String("Test description")},
			expected:   String("Test description"),
			method: func(o *OwnerGroup) interface{} {
				return o.GetDescription()
			},
		},
		{
			name:       "ConvertToProtoMessage",
			ownerGroup: &OwnerGroup{Name: String("owner-group-1")},
			expected: &diodepb.OwnerGroup{
				Name: "owner-group-1",
			},
			method: func(o *OwnerGroup) interface{} {
				return o.ConvertToProtoMessage()
			},
		},
		{
			name:       "ConvertToProtoEntity",
			ownerGroup: &OwnerGroup{Name: String("owner-group-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_OwnerGroup{
					OwnerGroup: &diodepb.OwnerGroup{
						Name: "owner-group-1",
					},
				},
			},
			method: func(o *OwnerGroup) interface{} {
				return o.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.ownerGroup))
		})
	}
}

func TestOwnerMethods(t *testing.T) {
	tests := []struct {
		name     string
		owner    *Owner
		expected interface{}
		method   func(*Owner) interface{}
	}{
		{
			name:     "GetName",
			owner:    &Owner{Name: String("owner-1")},
			expected: "owner-1",
			method: func(o *Owner) interface{} {
				return o.GetName()
			},
		},
		{
			name:     "GetGroup",
			owner:    &Owner{Group: &OwnerGroup{Name: String("owner-group-1")}},
			expected: &diodepb.OwnerGroup{Name: "owner-group-1"},
			method: func(o *Owner) interface{} {
				return o.GetGroup()
			},
		},
		{
			name:     "GetDescription",
			owner:    &Owner{Description: String("Test description")},
			expected: String("Test description"),
			method: func(o *Owner) interface{} {
				return o.GetDescription()
			},
		},
		{
			name:  "ConvertToProtoMessage",
			owner: &Owner{Name: String("owner-1")},
			expected: &diodepb.Owner{
				Name: "owner-1",
			},
			method: func(o *Owner) interface{} {
				return o.ConvertToProtoMessage()
			},
		},
		{
			name:  "ConvertToProtoEntity",
			owner: &Owner{Name: String("owner-1")},
			expected: &diodepb.Entity{
				Entity: &diodepb.Entity_Owner{
					Owner: &diodepb.Owner{
						Name: "owner-1",
					},
				},
			},
			method: func(o *Owner) interface{} {
				return o.ConvertToProtoEntity()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.method(tt.owner))
		})
	}
}
