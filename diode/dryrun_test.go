package diode

import (
	"context"
	"os"
	"path/filepath"
	"strings"
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
	first := loaded[0].ConvertToProtoEntity()
	require.Equal(t, "dev1", first.GetDevice().GetName())
	require.Equal(t, "Test device", first.GetDevice().GetDescription())

	second := loaded[1].ConvertToProtoEntity()
	require.Equal(t, "dev2", second.GetDevice().GetName())
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

func TestLoadDryRunEntitiesFixture(t *testing.T) {
	// Test loading from fixture
	fixtureFile := filepath.Join("testdata", "dryrun_fixture.json")

	entities, err := LoadDryRunEntities(fixtureFile)
	require.NoError(t, err)
	require.Len(t, entities, 94)

	// Verify first entity (ASN)
	first := entities[0].ConvertToProtoEntity()
	require.Equal(t, int64(555), first.GetAsn().GetAsn())
	// Verify entity at index 33 (IP Address)
	ipAddr := entities[33].ConvertToProtoEntity()
	require.Equal(t, "192.168.100.1/24", ipAddr.GetIpAddress().GetAddress())
	require.NotNil(t, ipAddr.GetIpAddress().GetAssignedObjectInterface())
	require.Equal(t, "GigabitEthernet1/0/1", ipAddr.GetIpAddress().GetAssignedObjectInterface().GetName())
	// Verify last entity (Wireless Link)
	last := entities[93].ConvertToProtoEntity()
	require.Equal(t, "P2P-Link-1", last.GetWirelessLink().GetSsid())

	// Test dry run client with output file
	tmpDir := t.TempDir()
	outputFile := filepath.Join(tmpDir, "out.json")

	c, err := NewDryRunClient(outputFile)
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	_, err = drc.Ingest(context.Background(), entities)
	require.NoError(t, err)

	// Verify output file exists and starts with "["
	content, err := os.ReadFile(outputFile)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(string(content), "["))

	// Load entities from output file and verify
	reloadedEntities, err := LoadDryRunEntities(outputFile)
	require.NoError(t, err)
	require.Len(t, reloadedEntities, 94)

	// Verify first entity (ASN)
	first = entities[0].ConvertToProtoEntity()
	require.Equal(t, int64(555), first.GetAsn().GetAsn())
	// Verify entity at index 33 (IP Address)
	ipAddr = entities[33].ConvertToProtoEntity()
	require.Equal(t, "192.168.100.1/24", ipAddr.GetIpAddress().GetAddress())
	require.NotNil(t, ipAddr.GetIpAddress().GetAssignedObjectInterface())
	require.Equal(t, "GigabitEthernet1/0/1", ipAddr.GetIpAddress().GetAssignedObjectInterface().GetName())
	// Verify last entity (Wireless Link)
	last = entities[93].ConvertToProtoEntity()
	require.Equal(t, "P2P-Link-1", last.GetWirelessLink().GetSsid())
}

func TestLoadDryRunEntitiesError(t *testing.T) {
	_, err := LoadDryRunEntities("not-exist.json")
	require.Error(t, err)
}
