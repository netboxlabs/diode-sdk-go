package diode

import (
	pb "github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
	"google.golang.org/protobuf/proto"
)

// ProtoEntity wraps a protobuf Entity and implements the Entity interface.
type ProtoEntity struct {
	PB *pb.Entity
}

// ConvertToProtoMessage returns the underlying protobuf message.
func (p ProtoEntity) ConvertToProtoMessage() proto.Message {
	if p.PB == nil {
		return nil
	}
	return p.PB
}

// ConvertToProtoEntity returns the underlying protobuf Entity.
func (p ProtoEntity) ConvertToProtoEntity() *pb.Entity {
	return p.PB
}
