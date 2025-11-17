package diode

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

func TestNewDryRunClient(t *testing.T) {
	client, err := NewDryRunClient("", "")
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestDryRunClientIngestAndLoad(t *testing.T) {
	dir := t.TempDir()

	c, err := NewDryRunClient("app", dir)
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	entities := []Entity{&Device{Name: String("dev1"), Description: String("Test device")}}

	_, err = drc.Ingest(context.Background(), entities)
	require.NoError(t, err)
	_, err = drc.Ingest(context.Background(), []Entity{&Device{Name: String("dev2")}})
	require.NoError(t, err)

	require.NoError(t, drc.Close())

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 2)

	for i, entry := range entries {
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		require.NoError(t, err)
		require.NotEmpty(t, data)

		loaded, err := LoadDryRunEntities(filepath.Join(dir, entry.Name()))
		require.NoError(t, err)
		require.Len(t, loaded, 1)
		entity := loaded[0]
		expected := fmt.Sprintf("dev%d", i+1)
		require.Equal(t, expected, entity.GetDevice().GetName())
	}
}

func TestNewDryRunClientEnvVar(t *testing.T) {
	dir := t.TempDir()

	err := os.Setenv(DiodeDryRunOutputDirEnvVarName, dir)
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv(DiodeDryRunOutputDirEnvVarName)
		require.NoError(t, err)
	}()

	c, err := NewDryRunClient("envapp", "")
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	_, err = drc.Ingest(context.Background(), []Entity{&Device{Name: String("dev1")}})
	require.NoError(t, err)
	require.NoError(t, drc.Close())

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 1)
}

func TestLoadDryRunEntitiesFixture(t *testing.T) {
	// Test loading from fixture
	fixtureFile := filepath.Join("testdata", "dryrun_fixture.json")

	entities, err := LoadDryRunEntities(fixtureFile)
	require.NoError(t, err)
	require.Len(t, entities, 94)

	// Verify first entity (ASN)
	first := entities[0]
	require.Equal(t, int64(555), first.GetAsn().GetAsn())
	// Verify entity at index 33 (IP Address)
	ipAddr := entities[33]
	require.Equal(t, "192.168.100.1/24", ipAddr.GetIpAddress().GetAddress())
	require.NotNil(t, ipAddr.GetIpAddress().GetAssignedObjectInterface())
	require.Equal(t, "GigabitEthernet1/0/1", ipAddr.GetIpAddress().GetAssignedObjectInterface().GetName())
	// Verify last entity (Wireless Link)
	last := entities[93]
	require.Equal(t, "P2P-Link-1", last.GetWirelessLink().GetSsid())

	// Test dry run client with output file
	tmpDir := t.TempDir()

	c, err := NewDryRunClient("app", tmpDir)
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	_, err = drc.IngestProto(context.Background(), entities)
	require.NoError(t, err)

	// Verify output file exists and starts with "{"
	filenames, err := os.ReadDir(tmpDir)
	require.NoError(t, err)
	fullPath := filepath.Join(tmpDir, filenames[0].Name())
	content, err := os.ReadFile(fullPath)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(string(content), "{"))

	// Load entities from output file and verify
	reloadedEntities, err := LoadDryRunEntities(fullPath)
	require.NoError(t, err)
	require.Len(t, reloadedEntities, 94)

	// Verify first entity (ASN)
	first = reloadedEntities[0]
	require.Equal(t, int64(555), first.GetAsn().GetAsn())
	// Verify entity at index 33 (IP Address)
	ipAddr = reloadedEntities[33]
	require.Equal(t, "192.168.100.1/24", ipAddr.GetIpAddress().GetAddress())
	require.NotNil(t, ipAddr.GetIpAddress().GetAssignedObjectInterface())
	require.Equal(t, "GigabitEthernet1/0/1", ipAddr.GetIpAddress().GetAssignedObjectInterface().GetName())
	// Verify last entity (Wireless Link)
	last = reloadedEntities[93]
	require.Equal(t, "P2P-Link-1", last.GetWirelessLink().GetSsid())
}

func TestLoadDryRunEntitiesError(t *testing.T) {
	_, err := LoadDryRunEntities("not-exist.json")
	require.Error(t, err)
}

func TestDryRunClientSDKVersionCaching(t *testing.T) {
	client, err := NewDryRunClient("test-app", "")
	require.NoError(t, err)
	require.NotNil(t, client)

	dryRunClient := client.(*DryRunClient)

	// Verify that SDK name and version are cached during initialization
	require.Equal(t, SDKName, dryRunClient.sdkName)
	require.NotEmpty(t, dryRunClient.sdkVersion)

	// The cached version should be consistent
	cachedVersion := dryRunClient.sdkVersion
	cachedName := dryRunClient.sdkName

	// Create another client and verify it gets the same version
	client2, err := NewDryRunClient("test-app-2", "")
	require.NoError(t, err)
	defer func() {
		err := client2.Close()
		require.NoError(t, err)
	}()

	dryRunClient2 := client2.(*DryRunClient)
	require.Equal(t, cachedVersion, dryRunClient2.sdkVersion)
	require.Equal(t, cachedName, dryRunClient2.sdkName)
}

func TestDryRunClientIngestWithMetadata(t *testing.T) {
	dir := t.TempDir()

	c, err := NewDryRunClient("app", dir)
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	entities := []Entity{
		&Device{Name: String("dev1"), Description: String("Test device")},
	}

	// Test with metadata
	metadata := Metadata{
		"batch_id": "batch-789",
		"source":   "auto_discovery",
		"priority": 2,
		"verified": false,
	}

	_, err = drc.Ingest(context.Background(), entities, WithIngestMetadata(metadata))
	require.NoError(t, err)

	// Test without metadata (backward compatibility)
	_, err = drc.Ingest(context.Background(), entities)
	require.NoError(t, err)

	require.NoError(t, drc.Close())

	// Verify two files were created
	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 2)

	// Load the first file (with metadata) and verify metadata is present
	loaded, err := LoadDryRunEntities(filepath.Join(dir, entries[0].Name()))
	require.NoError(t, err)
	require.Len(t, loaded, 1)

	// Load the IngestRequest to verify metadata is included
	data, err := os.ReadFile(filepath.Join(dir, entries[0].Name()))
	require.NoError(t, err)

	var req diodepb.IngestRequest
	err = protojson.Unmarshal(data, &req)
	require.NoError(t, err)
	require.NotNil(t, req.Metadata)

	metadataMap := req.Metadata.AsMap()
	require.Equal(t, "batch-789", metadataMap["batch_id"])
	require.Equal(t, "auto_discovery", metadataMap["source"])
	require.Equal(t, float64(2), metadataMap["priority"])
	require.Equal(t, false, metadataMap["verified"])
}
