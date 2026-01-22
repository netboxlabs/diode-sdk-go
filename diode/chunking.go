package diode

import (
	"google.golang.org/protobuf/proto"

	pb "github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	// DefaultMaxChunkSizeMB is the default maximum chunk size in megabytes.
	// This provides a safe margin below the gRPC 4 MB message size limit,
	// accounting for protobuf serialization overhead.
	DefaultMaxChunkSizeMB = 3.0
)

// CreateMessageChunks creates size-aware chunks from entities using greedy bin-packing.
//
// This function chunks entities to ensure each chunk stays under the specified
// size limit. It uses a greedy bin-packing algorithm that accumulates entities
// until adding the next entity would exceed the limit, then starts a new chunk.
//
// The default chunk size of 3.0 MB provides a safe margin below the gRPC 4 MB
// message size limit, accounting for protobuf serialization overhead.
//
// Parameters:
//   - entities: Slice of Entity protobuf messages to chunk
//   - maxChunkSizeMB: Maximum chunk size in MB (use 0 for default of 3.0 MB)
//
// Returns:
//   - Slice of entity chunks, each under maxChunkSizeMB. Returns at least
//     one chunk even if the input is empty.
//
// Examples:
//
//	entities := []*diodepb.Entity{entity1, entity2, entity3, ...}
//	chunks := CreateMessageChunks(entities, 0) // Use default 3.0 MB
//	for _, chunk := range chunks {
//	    client.IngestProto(ctx, chunk)
//	}
//
//	// Use a custom chunk size
//	chunks = CreateMessageChunks(entities, 3.5)
func CreateMessageChunks(entities []*pb.Entity, maxChunkSizeMB float64) [][]*pb.Entity {
	// Use default if not specified
	if maxChunkSizeMB <= 0 {
		maxChunkSizeMB = DefaultMaxChunkSizeMB
	}

	// Handle empty input
	if len(entities) == 0 {
		return [][]*pb.Entity{entities}
	}

	// Convert MB to bytes
	maxChunkSizeBytes := int(maxChunkSizeMB * 1024 * 1024)

	// Quick check: if all entities fit in one chunk, return early
	totalSize := EstimateMessageSize(entities)
	if totalSize <= maxChunkSizeBytes {
		return [][]*pb.Entity{entities}
	}

	// Greedy bin-packing: accumulate entities until limit reached
	baseOverhead := proto.Size(&pb.IngestRequest{})
	chunks := make([][]*pb.Entity, 0)
	currentChunk := make([]*pb.Entity, 0)
	currentChunkSize := baseOverhead // Start with overhead for the chunk

	for _, entity := range entities {
		entitySize := proto.Size(entity)
		projectedSize := currentChunkSize + entitySize

		// Check if adding this entity would exceed limit
		if len(currentChunk) > 0 && projectedSize > maxChunkSizeBytes {
			// Finalize current chunk and start new one
			chunks = append(chunks, currentChunk)
			currentChunk = []*pb.Entity{entity}
			currentChunkSize = baseOverhead + entitySize
		} else {
			// Add entity to current chunk
			currentChunk = append(currentChunk, entity)
			currentChunkSize = projectedSize
		}
	}

	// Add final chunk if not empty
	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}

	// Return chunks or original entities if no chunks were created
	if len(chunks) == 0 {
		return [][]*pb.Entity{entities}
	}

	return chunks
}

// EstimateMessageSize estimates the serialized size of entities in bytes.
//
// Calculates the total size by summing individual entity sizes plus the
// IngestRequest protobuf overhead.
//
// Parameters:
//   - entities: Slice of Entity protobuf messages
//
// Returns:
//   - Estimated size in bytes including IngestRequest overhead
//
// Examples:
//
//	entities := []*diodepb.Entity{entity1, entity2, entity3}
//	sizeBytes := EstimateMessageSize(entities)
//	sizeMB := float64(sizeBytes) / (1024 * 1024)
//	fmt.Printf("Estimated size: %.2f MB\n", sizeMB)
func EstimateMessageSize(entities []*pb.Entity) int {
	baseOverhead := proto.Size(&pb.IngestRequest{})
	entitySizesSum := 0
	for _, entity := range entities {
		entitySizesSum += proto.Size(entity)
	}
	return baseOverhead + entitySizesSum
}
