package diode

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDryRunClient(t *testing.T) {
	client, err := NewDryRunClient("")
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestDryRunClientIngestAndLoad(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "out.json")

	c, err := NewDryRunClient(file)
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	entities := []Entity{&Device{Name: String("dev1"), Description: String("Test device")}}

	_, err = drc.Ingest(context.Background(), entities)
	require.NoError(t, err)
	_, err = drc.Ingest(context.Background(), []Entity{&Device{Name: String("dev2")}})
	require.NoError(t, err)

	require.NoError(t, drc.Close())

	data, err := os.ReadFile(file)
	require.NoError(t, err)
	require.NotEmpty(t, data)

	loaded, err := LoadDryRunEntities(file)
	require.NoError(t, err)
	require.Len(t, loaded, 2)
	require.Equal(t, "dev1", loaded[0].GetDevice().GetName())
	require.Equal(t, "Test device", loaded[0].GetDevice().GetDescription())
	require.Equal(t, "dev2", loaded[1].GetDevice().GetName())
}

func TestNewDryRunClientEnvVar(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "out.json")

	err := os.Setenv(DiodeDryRunFileEnvVarName, file)
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv(DiodeDryRunFileEnvVarName)
		require.NoError(t, err)
	}()

	c, err := NewDryRunClient("")
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	_, err = drc.Ingest(context.Background(), []Entity{&Device{Name: String("dev1")}})
	require.NoError(t, err)
	require.NoError(t, drc.Close())

	_, err = os.Stat(file)
	require.NoError(t, err)
}

// func TestLoadDryRunEntitiesFixture(t *testing.T) {
// 	// Test loading from fixture
// 	fixtureFile := filepath.Join("testdata", "dryrun_fixture.json")

// 	entities, err := LoadDryRunEntities(fixtureFile)
// 	require.NoError(t, err)
// 	require.Len(t, entities, 94)

// 	// Verify first entity (ASN)
// 	asn, ok := entities[0].Entity.(*diodepb.Entity_Asn)
// 	require.True(t, ok, "Entity at index 0 should be ASN")
// 	require.NotNil(t, asn.Asn)
// 	assert.Equal(t, int32(555), asn.Asn.Asn)

// 	// Verify entity at index 33 (IP Address)
// 	ipAddr, ok := entities[33].Entity.(*diodepb.Entity_IpAddress)
// 	require.True(t, ok, "Entity at index 33 should be IP Address")
// 	require.NotNil(t, ipAddr.IpAddress)
// 	assert.Equal(t, "192.168.100.1/24", ipAddr.IpAddress.Address)
// 	require.NotNil(t, ipAddr.IpAddress.AssignedObjectInterface)
// 	assert.Equal(t, "GigabitEthernet1/0/1", ipAddr.IpAddress.AssignedObjectInterface.Name)

// 	// Verify last entity (Wireless Link)
// 	wirelessLink, ok := entities[93].Entity.(*diodepb.Entity_WirelessLink)
// 	require.True(t, ok, "Entity at index 93 should be Wireless Link")
// 	require.NotNil(t, wirelessLink.WirelessLink)
// 	assert.Equal(t, "P2P-Link-1", wirelessLink.WirelessLink.Ssid)

// 	// Test dry run client with output file
// 	tmpDir := t.TempDir()
// 	outputFile := filepath.Join(tmpDir, "out.json")

// 	c, err := NewDryRunClient(outputFile)
// 	require.NoError(t, err)
// 	drc := c.(*DryRunClient)

// 	_, err = drc.Ingest(context.Background(), entities)
// 	require.NoError(t, err)

// 	// Verify output file exists and starts with "["
// 	content, err := os.ReadFile(outputFile)
// 	require.NoError(t, err)
// 	assert.True(t, strings.HasPrefix(string(content), "["))

// 	// Load entities from output file and verify
// 	reloadedEntities, err := LoadDryRunEntities(outputFile)
// 	require.NoError(t, err)
// 	require.Len(t, reloadedEntities, 94)

// 	// Verify reloaded entities have same properties
// 	asn2, ok := reloadedEntities[0].Entity.(*diodepb.Entity_Asn)
// 	require.True(t, ok, "Reloaded entity at index 0 should be ASN")
// 	require.NotNil(t, asn2.Asn)
// 	assert.Equal(t, int32(555), asn2.Asn.Asn)

// 	ipAddr2, ok := reloadedEntities[33].Entity.(*diodepb.Entity_IpAddress)
// 	require.True(t, ok, "Reloaded entity at index 33 should be IP Address")
// 	require.NotNil(t, ipAddr2.IpAddress)
// 	assert.Equal(t, "192.168.100.1/24", ipAddr2.IpAddress.Address)
// 	require.NotNil(t, ipAddr2.IpAddress.AssignedObjectInterface)
// 	assert.Equal(t, "GigabitEthernet1/0/1", ipAddr2.IpAddress.AssignedObjectInterface.Name)

// 	wirelessLink2, ok := reloadedEntities[93].Entity.(*diodepb.Entity_WirelessLink)
// 	require.True(t, ok, "Reloaded entity at index 93 should be Wireless Link")
// 	require.NotNil(t, wirelessLink2.WirelessLink)
// 	assert.Equal(t, "P2P-Link-1", wirelessLink2.WirelessLink.Ssid)
// }

func TestLoadDryRunEntitiesError(t *testing.T) {
	_, err := LoadDryRunEntities("not-exist.json")
	require.Error(t, err)
}
