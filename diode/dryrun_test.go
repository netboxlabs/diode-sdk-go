package diode

import (
	"bytes"
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

	entities := []Entity{&Device{Name: String("dev1")}}

	_, err = drc.Ingest(context.Background(), entities)
	require.NoError(t, err)
	require.NoError(t, drc.Close())

	data, err := os.ReadFile(file)
	require.NoError(t, err)
	require.True(t, bytes.HasSuffix(data, []byte("\n")))

	loaded, err := LoadDryRunEntities(file)
	require.NoError(t, err)
	require.Len(t, loaded, 1)
	require.Equal(t, "dev1", loaded[0].GetDevice().GetName())
}

func TestNewDryRunClientEnvVar(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "out.json")

	_ = os.Setenv(DiodeDryRunFileEnvVarName, file)
	defer os.Unsetenv(DiodeDryRunFileEnvVarName)

	c, err := NewDryRunClient("")
	require.NoError(t, err)
	drc := c.(*DryRunClient)

	_, err = drc.Ingest(context.Background(), []Entity{&Device{Name: String("dev1")}})
	require.NoError(t, err)
	require.NoError(t, drc.Close())

	_, err = os.Stat(file)
	require.NoError(t, err)
}

func TestLoadDryRunEntitiesError(t *testing.T) {
	_, err := LoadDryRunEntities("not-exist.json")
	require.Error(t, err)
}
