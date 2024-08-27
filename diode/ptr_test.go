package diode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	val := true
	ptr := Bool(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestString(t *testing.T) {
	val := "test"
	ptr := String(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestInt(t *testing.T) {
	val := 42
	ptr := Int(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestInt32(t *testing.T) {
	val := int32(42)
	ptr := Int32(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestInt64(t *testing.T) {
	val := int64(42)
	ptr := Int64(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestUint(t *testing.T) {
	val := uint(42)
	ptr := Uint(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestUint32(t *testing.T) {
	val := uint32(42)
	ptr := Uint32(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestUint64(t *testing.T) {
	val := uint64(42)
	ptr := Uint64(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestFloat32(t *testing.T) {
	val := float32(42.0)
	ptr := Float32(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}

func TestFloat64(t *testing.T) {
	val := float64(42.0)
	ptr := Float64(val)
	require.NotNil(t, ptr)
	require.Equal(t, val, *ptr)
}
