package diode

import (
	"testing"

	pb "github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
	"google.golang.org/protobuf/proto"
)

// createTestEntity creates a test entity with a device name
func createTestEntity(name string) *pb.Entity {
	entity := &pb.Entity{
		Entity: &pb.Entity_Device{
			Device: &pb.Device{
				Name: String(name),
			},
		},
	}
	return entity
}

// TestCreateMessageChunksEmptyList tests CreateMessageChunks with an empty entity list
func TestCreateMessageChunksEmptyList(t *testing.T) {
	entities := []*pb.Entity{}
	chunks := CreateMessageChunks(entities, 0)

	if len(chunks) != 1 {
		t.Errorf("Expected 1 chunk, got %d", len(chunks))
	}

	if len(chunks[0]) != 0 {
		t.Errorf("Expected empty chunk, got %d entities", len(chunks[0]))
	}
}

// TestCreateMessageChunksSingleChunk tests CreateMessageChunks when entities fit in a single chunk
func TestCreateMessageChunksSingleChunk(t *testing.T) {
	// Create small entities that will fit in one chunk
	entities := make([]*pb.Entity, 5)
	for i := 0; i < 5; i++ {
		entities[i] = createTestEntity("test_device")
	}

	chunks := CreateMessageChunks(entities, 3.0)

	if len(chunks) != 1 {
		t.Errorf("Expected 1 chunk, got %d", len(chunks))
	}

	if len(chunks[0]) != 5 {
		t.Errorf("Expected 5 entities in chunk, got %d", len(chunks[0]))
	}

	// Verify same entities
	for i, entity := range chunks[0] {
		if entity != entities[i] {
			t.Errorf("Entity at position %d doesn't match", i)
		}
	}
}

// TestCreateMessageChunksMultipleChunks tests CreateMessageChunks when entities need to be split
func TestCreateMessageChunksMultipleChunks(t *testing.T) {
	// Create entities with large data to force multiple chunks
	entities := make([]*pb.Entity, 100)
	for i := 0; i < 100; i++ {
		// Create an entity with a large description to increase size
		largeString := make([]byte, 50000) // ~50KB per entity
		for j := range largeString {
			largeString[j] = 'A'
		}
		entity := &pb.Entity{
			Entity: &pb.Entity_Device{
				Device: &pb.Device{
					Name:        String("large_device"),
					Description: String(string(largeString)),
				},
			},
		}
		entities[i] = entity
	}

	// Use a smaller chunk size to force multiple chunks
	chunks := CreateMessageChunks(entities, 1.0) // 1 MB chunks

	// Should have multiple chunks
	if len(chunks) <= 1 {
		t.Errorf("Expected multiple chunks, got %d", len(chunks))
	}

	// All entities should be present across chunks
	totalEntities := 0
	for _, chunk := range chunks {
		totalEntities += len(chunk)
	}

	if totalEntities != 100 {
		t.Errorf("Expected 100 total entities, got %d", totalEntities)
	}

	// Each chunk should have at least 1 entity
	for i, chunk := range chunks {
		if len(chunk) < 1 {
			t.Errorf("Chunk %d has no entities", i)
		}
	}
}

// TestCreateMessageChunksOneEntityPerChunk tests when each entity needs its own chunk
func TestCreateMessageChunksOneEntityPerChunk(t *testing.T) {
	// Create very large entities
	entities := make([]*pb.Entity, 3)
	largeString := make([]byte, 2*1024*1024) // 2 MB per entity
	for j := range largeString {
		largeString[j] = 'B'
	}

	for i := 0; i < 3; i++ {
		entity := &pb.Entity{
			Entity: &pb.Entity_Device{
				Device: &pb.Device{
					Name:        String("huge_device"),
					Description: String(string(largeString)),
				},
			},
		}
		entities[i] = entity
	}

	// Use default 3 MB chunk size - each entity is ~2 MB so should get one per chunk
	chunks := CreateMessageChunks(entities, 3.0)

	// Should have 3 chunks with 1 entity each
	if len(chunks) != 3 {
		t.Errorf("Expected 3 chunks, got %d", len(chunks))
	}

	for i, chunk := range chunks {
		if len(chunk) != 1 {
			t.Errorf("Chunk %d: expected 1 entity, got %d", i, len(chunk))
		}
	}
}

// TestEstimateMessageSize tests the EstimateMessageSize function
func TestEstimateMessageSize(t *testing.T) {
	// Create test entities
	entities := make([]*pb.Entity, 3)
	for i := 0; i < 3; i++ {
		entities[i] = createTestEntity("test_device")
	}

	size := EstimateMessageSize(entities)

	// Should return a positive integer
	if size <= 0 {
		t.Errorf("Expected positive size, got %d", size)
	}

	// Verify calculation: should be base overhead + sum of entity sizes
	expectedSize := proto.Size(&pb.IngestRequest{})
	for _, entity := range entities {
		expectedSize += proto.Size(entity)
	}

	if size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}
}

// TestEstimateMessageSizeEmptyList tests EstimateMessageSize with an empty entity list
func TestEstimateMessageSizeEmptyList(t *testing.T) {
	entities := []*pb.Entity{}
	size := EstimateMessageSize(entities)

	// Should return base overhead (positive value for protobuf header)
	if size < 0 {
		t.Errorf("Expected non-negative size, got %d", size)
	}

	// Should equal the base IngestRequest size
	expectedSize := proto.Size(&pb.IngestRequest{})
	if size != expectedSize {
		t.Errorf("Expected size %d (base overhead), got %d", expectedSize, size)
	}
}

// TestCreateMessageChunksCustomChunkSize tests CreateMessageChunks with a custom chunk size
func TestCreateMessageChunksCustomChunkSize(t *testing.T) {
	// Create entities with medium-sized data
	entities := make([]*pb.Entity, 50)
	for i := 0; i < 50; i++ {
		mediumString := make([]byte, 100000) // ~100KB per entity
		for j := range mediumString {
			mediumString[j] = 'C'
		}
		entity := &pb.Entity{
			Entity: &pb.Entity_Device{
				Device: &pb.Device{
					Name:        String("test_device"),
					Description: String(string(mediumString)),
				},
			},
		}
		entities[i] = entity
	}

	// Use 3.5 MB chunk size (like orb-discovery in Python SDK)
	chunks := CreateMessageChunks(entities, 3.5)

	// Should have multiple chunks due to size
	if len(chunks) <= 1 {
		t.Errorf("Expected multiple chunks, got %d", len(chunks))
	}

	// All entities should be present
	totalEntities := 0
	for _, chunk := range chunks {
		totalEntities += len(chunk)
	}

	if totalEntities != 50 {
		t.Errorf("Expected 50 total entities, got %d", totalEntities)
	}
}

// TestCreateMessageChunksPreservesOrder tests that CreateMessageChunks preserves entity order
func TestCreateMessageChunksPreservesOrder(t *testing.T) {
	// Create entities with identifiable names
	entities := make([]*pb.Entity, 20)
	for i := 0; i < 20; i++ {
		largeString := make([]byte, 200000) // ~200KB per entity
		for j := range largeString {
			largeString[j] = 'D'
		}
		entity := &pb.Entity{
			Entity: &pb.Entity_Device{
				Device: &pb.Device{
					Name:        String(formatDeviceName(i)),
					Description: String(string(largeString)),
				},
			},
		}
		entities[i] = entity
	}

	// Force multiple chunks with smaller size
	chunks := CreateMessageChunks(entities, 2.0)

	// Flatten chunks and verify order
	flattened := make([]*pb.Entity, 0)
	for _, chunk := range chunks {
		flattened = append(flattened, chunk...)
	}

	if len(flattened) != 20 {
		t.Errorf("Expected 20 entities after flattening, got %d", len(flattened))
	}

	for i, entity := range flattened {
		expectedName := formatDeviceName(i)
		actualName := entity.GetDevice().GetName()
		if actualName != expectedName {
			t.Errorf("Entity at position %d: expected name %s, got %s", i, expectedName, actualName)
		}
	}
}

// TestCreateMessageChunksSingleLargeEntity tests with a single entity that exceeds chunk size
func TestCreateMessageChunksSingleLargeEntity(t *testing.T) {
	// Create a very large entity (5 MB) that exceeds 3 MB limit
	largeString := make([]byte, 5*1024*1024) // 5 MB
	for j := range largeString {
		largeString[j] = 'E'
	}
	entity := &pb.Entity{
		Entity: &pb.Entity_Device{
			Device: &pb.Device{
				Name:        String("huge_device"),
				Description: String(string(largeString)),
			},
		},
	}
	entities := []*pb.Entity{entity}

	chunks := CreateMessageChunks(entities, 3.0)

	// Should still return one chunk with the single entity
	if len(chunks) != 1 {
		t.Errorf("Expected 1 chunk, got %d", len(chunks))
	}

	if len(chunks[0]) != 1 {
		t.Errorf("Expected 1 entity in chunk, got %d", len(chunks[0]))
	}

	if chunks[0][0].GetDevice().GetName() != "huge_device" {
		t.Errorf("Entity name mismatch")
	}
}

// TestCreateMessageChunksDefaultChunkSize tests that default chunk size is used when 0 is passed
func TestCreateMessageChunksDefaultChunkSize(t *testing.T) {
	entities := make([]*pb.Entity, 5)
	for i := 0; i < 5; i++ {
		entities[i] = createTestEntity("test_device")
	}

	// Pass 0 to use default
	chunksDefault := CreateMessageChunks(entities, 0)

	// Pass explicit default
	chunksExplicit := CreateMessageChunks(entities, DefaultMaxChunkSizeMB)

	// Should produce same result
	if len(chunksDefault) != len(chunksExplicit) {
		t.Errorf("Default behavior mismatch: got %d chunks with 0, %d with explicit default",
			len(chunksDefault), len(chunksExplicit))
	}
}

// Helper function to format device names with leading zeros
func formatDeviceName(index int) string {
	if index < 10 {
		return "device_00" + string(rune('0'+index))
	} else if index < 100 {
		tens := index / 10
		ones := index % 10
		return "device_0" + string(rune('0'+tens)) + string(rune('0'+ones))
	}
	return "device_" + string(rune('0'+index))
}
