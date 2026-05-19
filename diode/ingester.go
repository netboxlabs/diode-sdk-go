// Generated code. DO NOT EDIT.
// Timestamp: 2026-05-14 20:30:14Z
//

package diode

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

// Entity is an interface that all entities must implement
type Entity interface {
	ConvertToProtoMessage() proto.Message
	ConvertToProtoEntity() *pb.Entity
}

// Metadata is a map of arbitrary key-value pairs that can be attached to entities
type Metadata map[string]any

// convertMetadata recursively converts Metadata to map[string]any
// so that nested Metadata values are properly handled by structpb.NewStruct
func convertMetadata(m Metadata) map[string]any {
	if m == nil {
		return nil
	}
	result := make(map[string]any, len(m))
	for k, v := range m {
		if nested, ok := v.(Metadata); ok {
			result[k] = convertMetadata(nested)
		} else {
			result[k] = v
		}
	}
	return result
}

type ASN struct {
	Asn          *int64
	Rir          *RIR
	Tenant       *Tenant
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Sites        []*Site
	Role         *Role
}

func (e *ASN) ConvertToProtoMessage() proto.Message {
	r := &pb.ASN{
		Asn:          e.GetAsn(),
		Rir:          e.GetRir(),
		Tenant:       e.GetTenant(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Sites:        e.GetSites(),
		Role:         e.GetRole(),
	}
	return r
}

func (e *ASN) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Asn{
			Asn: e.ConvertToProtoMessage().(*pb.ASN),
		},
	}
}

func (e *ASN) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ASN) GetAsn() int64 {
	var r int64
	if e != nil && e.Asn != nil {
		r = *e.Asn
	}
	return r
}

func (e *ASN) GetRir() *pb.RIR {
	var r *pb.RIR
	if e != nil && e.Rir != nil {
		r = e.Rir.ConvertToProtoMessage().(*pb.RIR)
	}
	return r
}

func (e *ASN) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *ASN) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ASN) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *ASN) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ASN) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ASN) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ASN) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *ASN) GetSites() []*pb.Site {
	var r []*pb.Site
	if e != nil && e.Sites != nil {
		for _, v := range e.Sites {
			r = append(r, v.ConvertToProtoMessage().(*pb.Site))
		}
	}
	return r
}

func (e *ASN) GetRole() *pb.Role {
	var r *pb.Role
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.Role)
	}
	return r
}

type ASNRange struct {
	Name         *string
	Slug         *string
	Rir          *RIR
	Start        *int64
	End          *int64
	Tenant       *Tenant
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *ASNRange) ConvertToProtoMessage() proto.Message {
	r := &pb.ASNRange{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Rir:          e.GetRir(),
		Start:        e.GetStart(),
		End:          e.GetEnd(),
		Tenant:       e.GetTenant(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *ASNRange) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_AsnRange{
			AsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
		},
	}
}

func (e *ASNRange) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ASNRange) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ASNRange) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *ASNRange) GetRir() *pb.RIR {
	var r *pb.RIR
	if e != nil && e.Rir != nil {
		r = e.Rir.ConvertToProtoMessage().(*pb.RIR)
	}
	return r
}

func (e *ASNRange) GetStart() int64 {
	var r int64
	if e != nil && e.Start != nil {
		r = *e.Start
	}
	return r
}

func (e *ASNRange) GetEnd() int64 {
	var r int64
	if e != nil && e.End != nil {
		r = *e.End
	}
	return r
}

func (e *ASNRange) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *ASNRange) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ASNRange) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ASNRange) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ASNRange) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ASNRange) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *ASNRange) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type Aggregate struct {
	Prefix       *string
	Rir          *RIR
	Tenant       *Tenant
	DateAdded    *time.Time
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *Aggregate) ConvertToProtoMessage() proto.Message {
	r := &pb.Aggregate{
		Prefix:       e.GetPrefix(),
		Rir:          e.GetRir(),
		Tenant:       e.GetTenant(),
		DateAdded:    e.GetDateAdded(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *Aggregate) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Aggregate{
			Aggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
		},
	}
}

func (e *Aggregate) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Aggregate) GetPrefix() string {
	var r string
	if e != nil && e.Prefix != nil {
		r = *e.Prefix
	}
	return r
}

func (e *Aggregate) GetRir() *pb.RIR {
	var r *pb.RIR
	if e != nil && e.Rir != nil {
		r = e.Rir.ConvertToProtoMessage().(*pb.RIR)
	}
	return r
}

func (e *Aggregate) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Aggregate) GetDateAdded() *timestamppb.Timestamp {
	var r *timestamppb.Timestamp
	if e != nil && e.DateAdded != nil {
		r = timestamppb.New(*e.DateAdded)
	}
	return r
}

func (e *Aggregate) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Aggregate) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Aggregate) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Aggregate) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Aggregate) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Aggregate) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Cable struct {
	Type          *string
	ATerminations []*GenericObject
	BTerminations []*GenericObject
	Status        *string
	Tenant        *Tenant
	Label         *string
	Color         *string
	Length        *float64
	LengthUnit    *string
	Description   *string
	Comments      *string
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Metadata      Metadata
	Profile       *string
	Owner         *Owner
	Bundle        *CableBundle
}

func (e *Cable) ConvertToProtoMessage() proto.Message {
	r := &pb.Cable{
		Type:          e.GetType(),
		ATerminations: e.GetATerminations(),
		BTerminations: e.GetBTerminations(),
		Status:        e.GetStatus(),
		Tenant:        e.GetTenant(),
		Label:         e.GetLabel(),
		Color:         e.GetColor(),
		Length:        e.GetLength(),
		LengthUnit:    e.GetLengthUnit(),
		Description:   e.GetDescription(),
		Comments:      e.GetComments(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Metadata:      e.GetMetadata(),
		Profile:       e.GetProfile(),
		Owner:         e.GetOwner(),
		Bundle:        e.GetBundle(),
	}
	return r
}

func (e *Cable) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Cable{
			Cable: e.ConvertToProtoMessage().(*pb.Cable),
		},
	}
}

func (e *Cable) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Cable) GetType() *string {
	var r *string
	if e != nil && e.Type != nil {
		r = e.Type
	}
	return r
}

func (e *Cable) GetATerminations() []*pb.GenericObject {
	var r []*pb.GenericObject
	if e != nil && e.ATerminations != nil {
		for _, v := range e.ATerminations {
			r = append(r, v.ConvertToProtoMessage().(*pb.GenericObject))
		}
	}
	return r
}

func (e *Cable) GetBTerminations() []*pb.GenericObject {
	var r []*pb.GenericObject
	if e != nil && e.BTerminations != nil {
		for _, v := range e.BTerminations {
			r = append(r, v.ConvertToProtoMessage().(*pb.GenericObject))
		}
	}
	return r
}

func (e *Cable) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Cable) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Cable) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *Cable) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *Cable) GetLength() *float64 {
	var r *float64
	if e != nil && e.Length != nil {
		r = e.Length
	}
	return r
}

func (e *Cable) GetLengthUnit() *string {
	var r *string
	if e != nil && e.LengthUnit != nil {
		r = e.LengthUnit
	}
	return r
}

func (e *Cable) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Cable) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Cable) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Cable) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Cable) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Cable) GetProfile() *string {
	var r *string
	if e != nil && e.Profile != nil {
		r = e.Profile
	}
	return r
}

func (e *Cable) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *Cable) GetBundle() *pb.CableBundle {
	var r *pb.CableBundle
	if e != nil && e.Bundle != nil {
		r = e.Bundle.ConvertToProtoMessage().(*pb.CableBundle)
	}
	return r
}

type CablePath struct {
	IsActive   *bool
	IsComplete *bool
	IsSplit    *bool
	Metadata   Metadata
}

func (e *CablePath) ConvertToProtoMessage() proto.Message {
	r := &pb.CablePath{
		IsActive:   e.GetIsActive(),
		IsComplete: e.GetIsComplete(),
		IsSplit:    e.GetIsSplit(),
		Metadata:   e.GetMetadata(),
	}
	return r
}

func (e *CablePath) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CablePath{
			CablePath: e.ConvertToProtoMessage().(*pb.CablePath),
		},
	}
}

func (e *CablePath) GetIsActive() *bool {
	var r *bool
	if e != nil && e.IsActive != nil {
		r = e.IsActive
	}
	return r
}

func (e *CablePath) GetIsComplete() *bool {
	var r *bool
	if e != nil && e.IsComplete != nil {
		r = e.IsComplete
	}
	return r
}

func (e *CablePath) GetIsSplit() *bool {
	var r *bool
	if e != nil && e.IsSplit != nil {
		r = e.IsSplit
	}
	return r
}

func (e *CablePath) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type CableTermination struct {
	Cable    *Cable
	CableEnd *string
	// Termination can be:
	//  - CircuitTermination
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - FrontPort
	//  - Interface
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPort
	//  - RearPort
	Termination anyCableTerminationTerminationValue
	Metadata    Metadata
}

func (e *CableTermination) ConvertToProtoMessage() proto.Message {
	r := &pb.CableTermination{
		Cable:    e.GetCable(),
		CableEnd: e.GetCableEnd(),
		Metadata: e.GetMetadata(),
	}
	if e.Termination != nil {
		e.Termination.anyCableTerminationTerminationValueApplyTo(r)
	}
	return r
}

func (e *CableTermination) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CableTermination{
			CableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
		},
	}
}

func (e *CableTermination) GetCable() *pb.Cable {
	var r *pb.Cable
	if e != nil && e.Cable != nil {
		r = e.Cable.ConvertToProtoMessage().(*pb.Cable)
	}
	return r
}

func (e *CableTermination) GetCableEnd() string {
	var r string
	if e != nil && e.CableEnd != nil {
		r = *e.CableEnd
	}
	return r
}

func (e *CableTermination) GetTermination() any {
	var r any
	if e != nil && e.Termination != nil {
		var tmp pb.CableTermination
		e.Termination.anyCableTerminationTerminationValueApplyTo(&tmp)
		r = tmp.Termination
	}
	return r
}

func (e *CableTermination) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type Circuit struct {
	Cid             *string
	Provider        *Provider
	ProviderAccount *ProviderAccount
	Type            *CircuitType
	Status          *string
	Tenant          *Tenant
	InstallDate     *time.Time
	TerminationDate *time.Time
	CommitRate      *int64
	Description     *string
	Distance        *float64
	DistanceUnit    *string
	Comments        *string
	Tags            []*Tag
	Assignments     []*CircuitGroupAssignment
	CustomFields    map[string]*CustomFieldValue
	Metadata        Metadata
	Owner           *Owner
}

func (e *Circuit) ConvertToProtoMessage() proto.Message {
	r := &pb.Circuit{
		Cid:             e.GetCid(),
		Provider:        e.GetProvider(),
		ProviderAccount: e.GetProviderAccount(),
		Type:            e.GetType(),
		Status:          e.GetStatus(),
		Tenant:          e.GetTenant(),
		InstallDate:     e.GetInstallDate(),
		TerminationDate: e.GetTerminationDate(),
		CommitRate:      e.GetCommitRate(),
		Description:     e.GetDescription(),
		Distance:        e.GetDistance(),
		DistanceUnit:    e.GetDistanceUnit(),
		Comments:        e.GetComments(),
		Tags:            e.GetTags(),
		Assignments:     e.GetAssignments(),
		CustomFields:    e.GetCustomFields(),
		Metadata:        e.GetMetadata(),
		Owner:           e.GetOwner(),
	}
	return r
}

func (e *Circuit) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Circuit{
			Circuit: e.ConvertToProtoMessage().(*pb.Circuit),
		},
	}
}

func (e *Circuit) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Circuit) GetCid() string {
	var r string
	if e != nil && e.Cid != nil {
		r = *e.Cid
	}
	return r
}

func (e *Circuit) GetProvider() *pb.Provider {
	var r *pb.Provider
	if e != nil && e.Provider != nil {
		r = e.Provider.ConvertToProtoMessage().(*pb.Provider)
	}
	return r
}

func (e *Circuit) GetProviderAccount() *pb.ProviderAccount {
	var r *pb.ProviderAccount
	if e != nil && e.ProviderAccount != nil {
		r = e.ProviderAccount.ConvertToProtoMessage().(*pb.ProviderAccount)
	}
	return r
}

func (e *Circuit) GetType() *pb.CircuitType {
	var r *pb.CircuitType
	if e != nil && e.Type != nil {
		r = e.Type.ConvertToProtoMessage().(*pb.CircuitType)
	}
	return r
}

func (e *Circuit) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Circuit) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Circuit) GetInstallDate() *timestamppb.Timestamp {
	var r *timestamppb.Timestamp
	if e != nil && e.InstallDate != nil {
		r = timestamppb.New(*e.InstallDate)
	}
	return r
}

func (e *Circuit) GetTerminationDate() *timestamppb.Timestamp {
	var r *timestamppb.Timestamp
	if e != nil && e.TerminationDate != nil {
		r = timestamppb.New(*e.TerminationDate)
	}
	return r
}

func (e *Circuit) GetCommitRate() *int64 {
	var r *int64
	if e != nil && e.CommitRate != nil {
		r = e.CommitRate
	}
	return r
}

func (e *Circuit) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Circuit) GetDistance() *float64 {
	var r *float64
	if e != nil && e.Distance != nil {
		r = e.Distance
	}
	return r
}

func (e *Circuit) GetDistanceUnit() *string {
	var r *string
	if e != nil && e.DistanceUnit != nil {
		r = e.DistanceUnit
	}
	return r
}

func (e *Circuit) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Circuit) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Circuit) GetAssignments() []*pb.CircuitGroupAssignment {
	var r []*pb.CircuitGroupAssignment
	if e != nil && e.Assignments != nil {
		for _, v := range e.Assignments {
			r = append(r, v.ConvertToProtoMessage().(*pb.CircuitGroupAssignment))
		}
	}
	return r
}

func (e *Circuit) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Circuit) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Circuit) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type CircuitGroup struct {
	Name         *string
	Slug         *string
	Description  *string
	Tenant       *Tenant
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *CircuitGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.CircuitGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Tenant:       e.GetTenant(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *CircuitGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CircuitGroup{
			CircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
		},
	}
}

func (e *CircuitGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *CircuitGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *CircuitGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *CircuitGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *CircuitGroup) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *CircuitGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *CircuitGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *CircuitGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *CircuitGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *CircuitGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type CircuitGroupAssignment struct {
	Group *CircuitGroup
	// Member can be:
	//  - Circuit
	//  - VirtualCircuit
	Member   anyCircuitGroupAssignmentMemberValue
	Priority *string
	Tags     []*Tag
	Metadata Metadata
}

func (e *CircuitGroupAssignment) ConvertToProtoMessage() proto.Message {
	r := &pb.CircuitGroupAssignment{
		Group:    e.GetGroup(),
		Priority: e.GetPriority(),
		Tags:     e.GetTags(),
		Metadata: e.GetMetadata(),
	}
	if e.Member != nil {
		e.Member.anyCircuitGroupAssignmentMemberValueApplyTo(r)
	}
	return r
}

func (e *CircuitGroupAssignment) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CircuitGroupAssignment{
			CircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
		},
	}
}

func (e *CircuitGroupAssignment) GetGroup() *pb.CircuitGroup {
	var r *pb.CircuitGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.CircuitGroup)
	}
	return r
}

func (e *CircuitGroupAssignment) GetMember() any {
	var r any
	if e != nil && e.Member != nil {
		var tmp pb.CircuitGroupAssignment
		e.Member.anyCircuitGroupAssignmentMemberValueApplyTo(&tmp)
		r = tmp.Member
	}
	return r
}

func (e *CircuitGroupAssignment) GetPriority() *string {
	var r *string
	if e != nil && e.Priority != nil {
		r = e.Priority
	}
	return r
}

func (e *CircuitGroupAssignment) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *CircuitGroupAssignment) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type CircuitTermination struct {
	Circuit  *Circuit
	TermSide *string
	// Termination can be:
	//  - Location
	//  - ProviderNetwork
	//  - Region
	//  - Site
	//  - SiteGroup
	Termination   anyCircuitTerminationTerminationValue
	PortSpeed     *int64
	UpstreamSpeed *int64
	XconnectId    *string
	PpInfo        *string
	Description   *string
	MarkConnected *bool
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Metadata      Metadata
}

func (e *CircuitTermination) ConvertToProtoMessage() proto.Message {
	r := &pb.CircuitTermination{
		Circuit:       e.GetCircuit(),
		TermSide:      e.GetTermSide(),
		PortSpeed:     e.GetPortSpeed(),
		UpstreamSpeed: e.GetUpstreamSpeed(),
		XconnectId:    e.GetXconnectId(),
		PpInfo:        e.GetPpInfo(),
		Description:   e.GetDescription(),
		MarkConnected: e.GetMarkConnected(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Metadata:      e.GetMetadata(),
	}
	if e.Termination != nil {
		e.Termination.anyCircuitTerminationTerminationValueApplyTo(r)
	}
	return r
}

func (e *CircuitTermination) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CircuitTermination{
			CircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
		},
	}
}

func (e *CircuitTermination) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *CircuitTermination) GetCircuit() *pb.Circuit {
	var r *pb.Circuit
	if e != nil && e.Circuit != nil {
		r = e.Circuit.ConvertToProtoMessage().(*pb.Circuit)
	}
	return r
}

func (e *CircuitTermination) GetTermSide() string {
	var r string
	if e != nil && e.TermSide != nil {
		r = *e.TermSide
	}
	return r
}

func (e *CircuitTermination) GetTermination() any {
	var r any
	if e != nil && e.Termination != nil {
		var tmp pb.CircuitTermination
		e.Termination.anyCircuitTerminationTerminationValueApplyTo(&tmp)
		r = tmp.Termination
	}
	return r
}

func (e *CircuitTermination) GetPortSpeed() *int64 {
	var r *int64
	if e != nil && e.PortSpeed != nil {
		r = e.PortSpeed
	}
	return r
}

func (e *CircuitTermination) GetUpstreamSpeed() *int64 {
	var r *int64
	if e != nil && e.UpstreamSpeed != nil {
		r = e.UpstreamSpeed
	}
	return r
}

func (e *CircuitTermination) GetXconnectId() *string {
	var r *string
	if e != nil && e.XconnectId != nil {
		r = e.XconnectId
	}
	return r
}

func (e *CircuitTermination) GetPpInfo() *string {
	var r *string
	if e != nil && e.PpInfo != nil {
		r = e.PpInfo
	}
	return r
}

func (e *CircuitTermination) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *CircuitTermination) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *CircuitTermination) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *CircuitTermination) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *CircuitTermination) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type CircuitType struct {
	Name         *string
	Slug         *string
	Color        *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *CircuitType) ConvertToProtoMessage() proto.Message {
	r := &pb.CircuitType{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Color:        e.GetColor(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *CircuitType) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CircuitType{
			CircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
		},
	}
}

func (e *CircuitType) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *CircuitType) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *CircuitType) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *CircuitType) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *CircuitType) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *CircuitType) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *CircuitType) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *CircuitType) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *CircuitType) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *CircuitType) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type Cluster struct {
	Name   *string
	Type   *ClusterType
	Group  *ClusterGroup
	Status *string
	Tenant *Tenant
	// Scope can be:
	//  - Location
	//  - Region
	//  - Site
	//  - SiteGroup
	Scope        anyClusterScopeValue
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *Cluster) ConvertToProtoMessage() proto.Message {
	r := &pb.Cluster{
		Name:         e.GetName(),
		Type:         e.GetType(),
		Group:        e.GetGroup(),
		Status:       e.GetStatus(),
		Tenant:       e.GetTenant(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	if e.Scope != nil {
		e.Scope.anyClusterScopeValueApplyTo(r)
	}
	return r
}

func (e *Cluster) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Cluster{
			Cluster: e.ConvertToProtoMessage().(*pb.Cluster),
		},
	}
}

func (e *Cluster) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Cluster) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Cluster) GetType() *pb.ClusterType {
	var r *pb.ClusterType
	if e != nil && e.Type != nil {
		r = e.Type.ConvertToProtoMessage().(*pb.ClusterType)
	}
	return r
}

func (e *Cluster) GetGroup() *pb.ClusterGroup {
	var r *pb.ClusterGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.ClusterGroup)
	}
	return r
}

func (e *Cluster) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Cluster) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Cluster) GetScope() any {
	var r any
	if e != nil && e.Scope != nil {
		var tmp pb.Cluster
		e.Scope.anyClusterScopeValueApplyTo(&tmp)
		r = tmp.Scope
	}
	return r
}

func (e *Cluster) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Cluster) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Cluster) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Cluster) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Cluster) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Cluster) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type ClusterGroup struct {
	Name         *string
	Slug         *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *ClusterGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.ClusterGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *ClusterGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ClusterGroup{
			ClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
		},
	}
}

func (e *ClusterGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ClusterGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ClusterGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *ClusterGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ClusterGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ClusterGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ClusterGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ClusterGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *ClusterGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type ClusterType struct {
	Name         *string
	Slug         *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *ClusterType) ConvertToProtoMessage() proto.Message {
	r := &pb.ClusterType{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *ClusterType) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ClusterType{
			ClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
		},
	}
}

func (e *ClusterType) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ClusterType) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ClusterType) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *ClusterType) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ClusterType) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ClusterType) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ClusterType) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ClusterType) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *ClusterType) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type ConsolePort struct {
	Device        *Device
	Module        *Module
	Name          *string
	Label         *string
	Type          *string
	Speed         *int64
	Description   *string
	MarkConnected *bool
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Metadata      Metadata
	Owner         *Owner
}

func (e *ConsolePort) ConvertToProtoMessage() proto.Message {
	r := &pb.ConsolePort{
		Device:        e.GetDevice(),
		Module:        e.GetModule(),
		Name:          e.GetName(),
		Label:         e.GetLabel(),
		Type:          e.GetType(),
		Speed:         e.GetSpeed(),
		Description:   e.GetDescription(),
		MarkConnected: e.GetMarkConnected(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *ConsolePort) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ConsolePort{
			ConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
		},
	}
}

func (e *ConsolePort) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ConsolePort) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *ConsolePort) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *ConsolePort) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ConsolePort) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *ConsolePort) GetType() *string {
	var r *string
	if e != nil && e.Type != nil {
		r = e.Type
	}
	return r
}

func (e *ConsolePort) GetSpeed() *int64 {
	var r *int64
	if e != nil && e.Speed != nil {
		r = e.Speed
	}
	return r
}

func (e *ConsolePort) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ConsolePort) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *ConsolePort) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ConsolePort) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ConsolePort) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ConsolePort) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type ConsoleServerPort struct {
	Device        *Device
	Module        *Module
	Name          *string
	Label         *string
	Type          *string
	Speed         *int64
	Description   *string
	MarkConnected *bool
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Metadata      Metadata
	Owner         *Owner
}

func (e *ConsoleServerPort) ConvertToProtoMessage() proto.Message {
	r := &pb.ConsoleServerPort{
		Device:        e.GetDevice(),
		Module:        e.GetModule(),
		Name:          e.GetName(),
		Label:         e.GetLabel(),
		Type:          e.GetType(),
		Speed:         e.GetSpeed(),
		Description:   e.GetDescription(),
		MarkConnected: e.GetMarkConnected(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *ConsoleServerPort) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ConsoleServerPort{
			ConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
		},
	}
}

func (e *ConsoleServerPort) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ConsoleServerPort) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *ConsoleServerPort) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *ConsoleServerPort) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ConsoleServerPort) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *ConsoleServerPort) GetType() *string {
	var r *string
	if e != nil && e.Type != nil {
		r = e.Type
	}
	return r
}

func (e *ConsoleServerPort) GetSpeed() *int64 {
	var r *int64
	if e != nil && e.Speed != nil {
		r = e.Speed
	}
	return r
}

func (e *ConsoleServerPort) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ConsoleServerPort) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *ConsoleServerPort) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ConsoleServerPort) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ConsoleServerPort) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ConsoleServerPort) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Contact struct {
	Group        *ContactGroup
	Name         *string
	Title        *string
	Phone        *string
	Email        *string
	Address      *string
	Link         *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Groups       []*ContactGroup
	Metadata     Metadata
	Owner        *Owner
}

func (e *Contact) ConvertToProtoMessage() proto.Message {
	r := &pb.Contact{
		Group:        e.GetGroup(),
		Name:         e.GetName(),
		Title:        e.GetTitle(),
		Phone:        e.GetPhone(),
		Email:        e.GetEmail(),
		Address:      e.GetAddress(),
		Link:         e.GetLink(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Groups:       e.GetGroups(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *Contact) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Contact{
			Contact: e.ConvertToProtoMessage().(*pb.Contact),
		},
	}
}

func (e *Contact) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Contact) GetGroup() *pb.ContactGroup {
	var r *pb.ContactGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.ContactGroup)
	}
	return r
}

func (e *Contact) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Contact) GetTitle() *string {
	var r *string
	if e != nil && e.Title != nil {
		r = e.Title
	}
	return r
}

func (e *Contact) GetPhone() *string {
	var r *string
	if e != nil && e.Phone != nil {
		r = e.Phone
	}
	return r
}

func (e *Contact) GetEmail() *string {
	var r *string
	if e != nil && e.Email != nil {
		r = e.Email
	}
	return r
}

func (e *Contact) GetAddress() *string {
	var r *string
	if e != nil && e.Address != nil {
		r = e.Address
	}
	return r
}

func (e *Contact) GetLink() *string {
	var r *string
	if e != nil && e.Link != nil {
		r = e.Link
	}
	return r
}

func (e *Contact) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Contact) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Contact) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Contact) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Contact) GetGroups() []*pb.ContactGroup {
	var r []*pb.ContactGroup
	if e != nil && e.Groups != nil {
		for _, v := range e.Groups {
			r = append(r, v.ConvertToProtoMessage().(*pb.ContactGroup))
		}
	}
	return r
}

func (e *Contact) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Contact) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type ContactAssignment struct {
	// Object can be:
	//  - ASN
	//  - ASNRange
	//  - Aggregate
	//  - Cable
	//  - CablePath
	//  - CableTermination
	//  - Circuit
	//  - CircuitGroup
	//  - CircuitGroupAssignment
	//  - CircuitTermination
	//  - CircuitType
	//  - Cluster
	//  - ClusterGroup
	//  - ClusterType
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - Contact
	//  - ContactAssignment
	//  - ContactGroup
	//  - ContactRole
	//  - Device
	//  - DeviceBay
	//  - DeviceRole
	//  - DeviceType
	//  - FHRPGroup
	//  - FHRPGroupAssignment
	//  - FrontPort
	//  - IKEPolicy
	//  - IKEProposal
	//  - IPAddress
	//  - IPRange
	//  - IPSecPolicy
	//  - IPSecProfile
	//  - IPSecProposal
	//  - Interface
	//  - InventoryItem
	//  - InventoryItemRole
	//  - L2VPN
	//  - L2VPNTermination
	//  - Location
	//  - MACAddress
	//  - Manufacturer
	//  - Module
	//  - ModuleBay
	//  - ModuleType
	//  - Platform
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPanel
	//  - PowerPort
	//  - Prefix
	//  - Provider
	//  - ProviderAccount
	//  - ProviderNetwork
	//  - RIR
	//  - Rack
	//  - RackReservation
	//  - RackRole
	//  - RackType
	//  - RearPort
	//  - Region
	//  - Role
	//  - RouteTarget
	//  - Service
	//  - Site
	//  - SiteGroup
	//  - Tag
	//  - Tenant
	//  - TenantGroup
	//  - Tunnel
	//  - TunnelGroup
	//  - TunnelTermination
	//  - VLAN
	//  - VLANGroup
	//  - VLANTranslationPolicy
	//  - VLANTranslationRule
	//  - VMInterface
	//  - VRF
	//  - VirtualChassis
	//  - VirtualCircuit
	//  - VirtualCircuitTermination
	//  - VirtualCircuitType
	//  - VirtualDeviceContext
	//  - VirtualDisk
	//  - VirtualMachine
	//  - WirelessLAN
	//  - WirelessLANGroup
	//  - WirelessLink
	//  - CustomField
	//  - CustomFieldChoiceSet
	//  - JournalEntry
	//  - ModuleTypeProfile
	//  - CustomLink
	//  - Owner
	//  - OwnerGroup
	//  - CableBundle
	//  - RackGroup
	//  - ScriptModule
	//  - VirtualMachineType
	Object       anyContactAssignmentObjectValue
	Contact      *Contact
	Role         *ContactRole
	Priority     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
}

func (e *ContactAssignment) ConvertToProtoMessage() proto.Message {
	r := &pb.ContactAssignment{
		Contact:      e.GetContact(),
		Role:         e.GetRole(),
		Priority:     e.GetPriority(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
	}
	if e.Object != nil {
		e.Object.anyContactAssignmentObjectValueApplyTo(r)
	}
	return r
}

func (e *ContactAssignment) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ContactAssignment{
			ContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
		},
	}
}

func (e *ContactAssignment) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ContactAssignment) GetObject() any {
	var r any
	if e != nil && e.Object != nil {
		var tmp pb.ContactAssignment
		e.Object.anyContactAssignmentObjectValueApplyTo(&tmp)
		r = tmp.Object
	}
	return r
}

func (e *ContactAssignment) GetContact() *pb.Contact {
	var r *pb.Contact
	if e != nil && e.Contact != nil {
		r = e.Contact.ConvertToProtoMessage().(*pb.Contact)
	}
	return r
}

func (e *ContactAssignment) GetRole() *pb.ContactRole {
	var r *pb.ContactRole
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.ContactRole)
	}
	return r
}

func (e *ContactAssignment) GetPriority() *string {
	var r *string
	if e != nil && e.Priority != nil {
		r = e.Priority
	}
	return r
}

func (e *ContactAssignment) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ContactAssignment) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ContactAssignment) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type ContactGroup struct {
	Name         *string
	Slug         *string
	Parent       *ContactGroup
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *ContactGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.ContactGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Parent:       e.GetParent(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *ContactGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ContactGroup{
			ContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
		},
	}
}

func (e *ContactGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ContactGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ContactGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *ContactGroup) GetParent() *pb.ContactGroup {
	var r *pb.ContactGroup
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.ContactGroup)
	}
	return r
}

func (e *ContactGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ContactGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ContactGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ContactGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *ContactGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ContactGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type ContactRole struct {
	Name         *string
	Slug         *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *ContactRole) ConvertToProtoMessage() proto.Message {
	r := &pb.ContactRole{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *ContactRole) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ContactRole{
			ContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
		},
	}
}

func (e *ContactRole) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ContactRole) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ContactRole) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *ContactRole) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ContactRole) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ContactRole) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ContactRole) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ContactRole) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *ContactRole) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type CustomFieldObjectReference struct {
	// Object can be:
	//  - ASN
	//  - ASNRange
	//  - Aggregate
	//  - Cable
	//  - CablePath
	//  - CableTermination
	//  - Circuit
	//  - CircuitGroup
	//  - CircuitGroupAssignment
	//  - CircuitTermination
	//  - CircuitType
	//  - Cluster
	//  - ClusterGroup
	//  - ClusterType
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - Contact
	//  - ContactAssignment
	//  - ContactGroup
	//  - ContactRole
	//  - Device
	//  - DeviceBay
	//  - DeviceRole
	//  - DeviceType
	//  - FHRPGroup
	//  - FHRPGroupAssignment
	//  - FrontPort
	//  - IKEPolicy
	//  - IKEProposal
	//  - IPAddress
	//  - IPRange
	//  - IPSecPolicy
	//  - IPSecProfile
	//  - IPSecProposal
	//  - Interface
	//  - InventoryItem
	//  - InventoryItemRole
	//  - L2VPN
	//  - L2VPNTermination
	//  - Location
	//  - MACAddress
	//  - Manufacturer
	//  - Module
	//  - ModuleBay
	//  - ModuleType
	//  - Platform
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPanel
	//  - PowerPort
	//  - Prefix
	//  - Provider
	//  - ProviderAccount
	//  - ProviderNetwork
	//  - RIR
	//  - Rack
	//  - RackReservation
	//  - RackRole
	//  - RackType
	//  - RearPort
	//  - Region
	//  - Role
	//  - RouteTarget
	//  - Service
	//  - Site
	//  - SiteGroup
	//  - Tag
	//  - Tenant
	//  - TenantGroup
	//  - Tunnel
	//  - TunnelGroup
	//  - TunnelTermination
	//  - VLAN
	//  - VLANGroup
	//  - VLANTranslationPolicy
	//  - VLANTranslationRule
	//  - VMInterface
	//  - VRF
	//  - VirtualChassis
	//  - VirtualCircuit
	//  - VirtualCircuitTermination
	//  - VirtualCircuitType
	//  - VirtualDeviceContext
	//  - VirtualDisk
	//  - VirtualMachine
	//  - WirelessLAN
	//  - WirelessLANGroup
	//  - WirelessLink
	//  - CustomField
	//  - CustomFieldChoiceSet
	//  - JournalEntry
	//  - ModuleTypeProfile
	//  - CustomLink
	//  - Owner
	//  - OwnerGroup
	//  - CableBundle
	//  - RackGroup
	//  - ScriptModule
	//  - VirtualMachineType
	Object anyCustomFieldObjectReferenceObjectValue
}

func (e *CustomFieldObjectReference) ConvertToProtoMessage() proto.Message {
	r := &pb.CustomFieldObjectReference{}
	if e.Object != nil {
		e.Object.anyCustomFieldObjectReferenceObjectValueApplyTo(r)
	}
	return r
}

func (e *CustomFieldObjectReference) GetObject() any {
	var r any
	if e != nil && e.Object != nil {
		var tmp pb.CustomFieldObjectReference
		e.Object.anyCustomFieldObjectReferenceObjectValueApplyTo(&tmp)
		r = tmp.Object
	}
	return r
}

type CustomFieldValue struct {
	MultipleSelection []string
	MultipleObjects   []*CustomFieldObjectReference
	// Value can be:
	//  - CustomFieldValueText (alias for string)
	//  - CustomFieldValueLongText (alias for string)
	//  - CustomFieldValueInteger (alias for int64)
	//  - CustomFieldValueDecimal (alias for float64)
	//  - CustomFieldValueBoolean (alias for bool)
	//  - CustomFieldValueDate (alias for time.Time)
	//  - CustomFieldValueDatetime (alias for time.Time)
	//  - CustomFieldValueUrl (alias for string)
	//  - CustomFieldValueJson (alias for string)
	//  - CustomFieldValueSelection (alias for string)
	//  - CustomFieldObjectReference
	Value anyCustomFieldValueValueValue
}

func (e *CustomFieldValue) ConvertToProtoMessage() proto.Message {
	r := &pb.CustomFieldValue{
		MultipleSelection: e.GetMultipleSelection(),
		MultipleObjects:   e.GetMultipleObjects(),
	}
	if e.Value != nil {
		e.Value.anyCustomFieldValueValueValueApplyTo(r)
	}
	return r
}

func (e *CustomFieldValue) GetMultipleSelection() []string {
	var r []string
	if e != nil && e.MultipleSelection != nil {
		for _, v := range e.MultipleSelection {
			r = append(r, v)
		}
	}
	return r
}

func (e *CustomFieldValue) GetMultipleObjects() []*pb.CustomFieldObjectReference {
	var r []*pb.CustomFieldObjectReference
	if e != nil && e.MultipleObjects != nil {
		for _, v := range e.MultipleObjects {
			r = append(r, v.ConvertToProtoMessage().(*pb.CustomFieldObjectReference))
		}
	}
	return r
}

func (e *CustomFieldValue) GetValue() any {
	var r any
	if e != nil && e.Value != nil {
		var tmp pb.CustomFieldValue
		e.Value.anyCustomFieldValueValueValueApplyTo(&tmp)
		r = tmp.Value
	}
	return r
}

type Device struct {
	Name           *string
	DeviceType     *DeviceType
	Role           *DeviceRole
	Tenant         *Tenant
	Platform       *Platform
	Serial         *string
	AssetTag       *string
	Site           *Site
	Location       *Location
	Rack           *Rack
	Position       *float64
	Face           *string
	Latitude       *float64
	Longitude      *float64
	Status         *string
	Airflow        *string
	PrimaryIp4     *IPAddress
	PrimaryIp6     *IPAddress
	OobIp          *IPAddress
	Cluster        *Cluster
	VirtualChassis *VirtualChassis
	VcPosition     *int64
	VcPriority     *int64
	Description    *string
	Comments       *string
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
	Owner          *Owner
	Config         *DeviceConfig
}

func (e *Device) ConvertToProtoMessage() proto.Message {
	r := &pb.Device{
		Name:           e.GetName(),
		DeviceType:     e.GetDeviceType(),
		Role:           e.GetRole(),
		Tenant:         e.GetTenant(),
		Platform:       e.GetPlatform(),
		Serial:         e.GetSerial(),
		AssetTag:       e.GetAssetTag(),
		Site:           e.GetSite(),
		Location:       e.GetLocation(),
		Rack:           e.GetRack(),
		Position:       e.GetPosition(),
		Face:           e.GetFace(),
		Latitude:       e.GetLatitude(),
		Longitude:      e.GetLongitude(),
		Status:         e.GetStatus(),
		Airflow:        e.GetAirflow(),
		PrimaryIp4:     e.GetPrimaryIp4(),
		PrimaryIp6:     e.GetPrimaryIp6(),
		OobIp:          e.GetOobIp(),
		Cluster:        e.GetCluster(),
		VirtualChassis: e.GetVirtualChassis(),
		VcPosition:     e.GetVcPosition(),
		VcPriority:     e.GetVcPriority(),
		Description:    e.GetDescription(),
		Comments:       e.GetComments(),
		Tags:           e.GetTags(),
		CustomFields:   e.GetCustomFields(),
		Metadata:       e.GetMetadata(),
		Owner:          e.GetOwner(),
		Config:         e.GetConfig(),
	}
	return r
}

func (e *Device) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Device{
			Device: e.ConvertToProtoMessage().(*pb.Device),
		},
	}
}

func (e *Device) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Device) GetName() *string {
	var r *string
	if e != nil && e.Name != nil {
		r = e.Name
	}
	return r
}

func (e *Device) GetDeviceType() *pb.DeviceType {
	var r *pb.DeviceType
	if e != nil && e.DeviceType != nil {
		r = e.DeviceType.ConvertToProtoMessage().(*pb.DeviceType)
	}
	return r
}

func (e *Device) GetRole() *pb.DeviceRole {
	var r *pb.DeviceRole
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.DeviceRole)
	}
	return r
}

func (e *Device) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Device) GetPlatform() *pb.Platform {
	var r *pb.Platform
	if e != nil && e.Platform != nil {
		r = e.Platform.ConvertToProtoMessage().(*pb.Platform)
	}
	return r
}

func (e *Device) GetSerial() *string {
	var r *string
	if e != nil && e.Serial != nil {
		r = e.Serial
	}
	return r
}

func (e *Device) GetAssetTag() *string {
	var r *string
	if e != nil && e.AssetTag != nil {
		r = e.AssetTag
	}
	return r
}

func (e *Device) GetSite() *pb.Site {
	var r *pb.Site
	if e != nil && e.Site != nil {
		r = e.Site.ConvertToProtoMessage().(*pb.Site)
	}
	return r
}

func (e *Device) GetLocation() *pb.Location {
	var r *pb.Location
	if e != nil && e.Location != nil {
		r = e.Location.ConvertToProtoMessage().(*pb.Location)
	}
	return r
}

func (e *Device) GetRack() *pb.Rack {
	var r *pb.Rack
	if e != nil && e.Rack != nil {
		r = e.Rack.ConvertToProtoMessage().(*pb.Rack)
	}
	return r
}

func (e *Device) GetPosition() *float64 {
	var r *float64
	if e != nil && e.Position != nil {
		r = e.Position
	}
	return r
}

func (e *Device) GetFace() *string {
	var r *string
	if e != nil && e.Face != nil {
		r = e.Face
	}
	return r
}

func (e *Device) GetLatitude() *float64 {
	var r *float64
	if e != nil && e.Latitude != nil {
		r = e.Latitude
	}
	return r
}

func (e *Device) GetLongitude() *float64 {
	var r *float64
	if e != nil && e.Longitude != nil {
		r = e.Longitude
	}
	return r
}

func (e *Device) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Device) GetAirflow() *string {
	var r *string
	if e != nil && e.Airflow != nil {
		r = e.Airflow
	}
	return r
}

func (e *Device) GetPrimaryIp4() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.PrimaryIp4 != nil {
		r = e.PrimaryIp4.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *Device) GetPrimaryIp6() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.PrimaryIp6 != nil {
		r = e.PrimaryIp6.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *Device) GetOobIp() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.OobIp != nil {
		r = e.OobIp.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *Device) GetCluster() *pb.Cluster {
	var r *pb.Cluster
	if e != nil && e.Cluster != nil {
		r = e.Cluster.ConvertToProtoMessage().(*pb.Cluster)
	}
	return r
}

func (e *Device) GetVirtualChassis() *pb.VirtualChassis {
	var r *pb.VirtualChassis
	if e != nil && e.VirtualChassis != nil {
		r = e.VirtualChassis.ConvertToProtoMessage().(*pb.VirtualChassis)
	}
	return r
}

func (e *Device) GetVcPosition() *int64 {
	var r *int64
	if e != nil && e.VcPosition != nil {
		r = e.VcPosition
	}
	return r
}

func (e *Device) GetVcPriority() *int64 {
	var r *int64
	if e != nil && e.VcPriority != nil {
		r = e.VcPriority
	}
	return r
}

func (e *Device) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Device) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Device) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Device) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Device) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Device) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *Device) GetConfig() *pb.DeviceConfig {
	var r *pb.DeviceConfig
	if e != nil && e.Config != nil {
		r = e.Config.ConvertToProtoMessage().(*pb.DeviceConfig)
	}
	return r
}

type DeviceBay struct {
	Device          *Device
	Name            *string
	Label           *string
	Description     *string
	InstalledDevice *Device
	Tags            []*Tag
	CustomFields    map[string]*CustomFieldValue
	Metadata        Metadata
	Owner           *Owner
	Enabled         *bool
}

func (e *DeviceBay) ConvertToProtoMessage() proto.Message {
	r := &pb.DeviceBay{
		Device:          e.GetDevice(),
		Name:            e.GetName(),
		Label:           e.GetLabel(),
		Description:     e.GetDescription(),
		InstalledDevice: e.GetInstalledDevice(),
		Tags:            e.GetTags(),
		CustomFields:    e.GetCustomFields(),
		Metadata:        e.GetMetadata(),
		Owner:           e.GetOwner(),
		Enabled:         e.GetEnabled(),
	}
	return r
}

func (e *DeviceBay) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_DeviceBay{
			DeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
		},
	}
}

func (e *DeviceBay) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *DeviceBay) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *DeviceBay) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *DeviceBay) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *DeviceBay) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *DeviceBay) GetInstalledDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.InstalledDevice != nil {
		r = e.InstalledDevice.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *DeviceBay) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *DeviceBay) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *DeviceBay) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *DeviceBay) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *DeviceBay) GetEnabled() *bool {
	var r *bool
	if e != nil && e.Enabled != nil {
		r = e.Enabled
	}
	return r
}

type DeviceRole struct {
	Name         *string
	Slug         *string
	Color        *string
	VmRole       *bool
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Parent       *DeviceRole
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *DeviceRole) ConvertToProtoMessage() proto.Message {
	r := &pb.DeviceRole{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Color:        e.GetColor(),
		VmRole:       e.GetVmRole(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Parent:       e.GetParent(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *DeviceRole) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_DeviceRole{
			DeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
		},
	}
}

func (e *DeviceRole) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *DeviceRole) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *DeviceRole) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *DeviceRole) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *DeviceRole) GetVmRole() *bool {
	var r *bool
	if e != nil && e.VmRole != nil {
		r = e.VmRole
	}
	return r
}

func (e *DeviceRole) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *DeviceRole) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *DeviceRole) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *DeviceRole) GetParent() *pb.DeviceRole {
	var r *pb.DeviceRole
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.DeviceRole)
	}
	return r
}

func (e *DeviceRole) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *DeviceRole) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *DeviceRole) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type DeviceType struct {
	Manufacturer           *Manufacturer
	DefaultPlatform        *Platform
	Model                  *string
	Slug                   *string
	PartNumber             *string
	UHeight                *float64
	ExcludeFromUtilization *bool
	IsFullDepth            *bool
	SubdeviceRole          *string
	Airflow                *string
	Weight                 *float64
	WeightUnit             *string
	Description            *string
	Comments               *string
	Tags                   []*Tag
	CustomFields           map[string]*CustomFieldValue
	Metadata               Metadata
	Owner                  *Owner
}

func (e *DeviceType) ConvertToProtoMessage() proto.Message {
	r := &pb.DeviceType{
		Manufacturer:           e.GetManufacturer(),
		DefaultPlatform:        e.GetDefaultPlatform(),
		Model:                  e.GetModel(),
		Slug:                   e.GetSlug(),
		PartNumber:             e.GetPartNumber(),
		UHeight:                e.GetUHeight(),
		ExcludeFromUtilization: e.GetExcludeFromUtilization(),
		IsFullDepth:            e.GetIsFullDepth(),
		SubdeviceRole:          e.GetSubdeviceRole(),
		Airflow:                e.GetAirflow(),
		Weight:                 e.GetWeight(),
		WeightUnit:             e.GetWeightUnit(),
		Description:            e.GetDescription(),
		Comments:               e.GetComments(),
		Tags:                   e.GetTags(),
		CustomFields:           e.GetCustomFields(),
		Metadata:               e.GetMetadata(),
		Owner:                  e.GetOwner(),
	}
	return r
}

func (e *DeviceType) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_DeviceType{
			DeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
		},
	}
}

func (e *DeviceType) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *DeviceType) GetManufacturer() *pb.Manufacturer {
	var r *pb.Manufacturer
	if e != nil && e.Manufacturer != nil {
		r = e.Manufacturer.ConvertToProtoMessage().(*pb.Manufacturer)
	}
	return r
}

func (e *DeviceType) GetDefaultPlatform() *pb.Platform {
	var r *pb.Platform
	if e != nil && e.DefaultPlatform != nil {
		r = e.DefaultPlatform.ConvertToProtoMessage().(*pb.Platform)
	}
	return r
}

func (e *DeviceType) GetModel() string {
	var r string
	if e != nil && e.Model != nil {
		r = *e.Model
	}
	return r
}

func (e *DeviceType) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *DeviceType) GetPartNumber() *string {
	var r *string
	if e != nil && e.PartNumber != nil {
		r = e.PartNumber
	}
	return r
}

func (e *DeviceType) GetUHeight() *float64 {
	var r *float64
	if e != nil && e.UHeight != nil {
		r = e.UHeight
	}
	return r
}

func (e *DeviceType) GetExcludeFromUtilization() *bool {
	var r *bool
	if e != nil && e.ExcludeFromUtilization != nil {
		r = e.ExcludeFromUtilization
	}
	return r
}

func (e *DeviceType) GetIsFullDepth() *bool {
	var r *bool
	if e != nil && e.IsFullDepth != nil {
		r = e.IsFullDepth
	}
	return r
}

func (e *DeviceType) GetSubdeviceRole() *string {
	var r *string
	if e != nil && e.SubdeviceRole != nil {
		r = e.SubdeviceRole
	}
	return r
}

func (e *DeviceType) GetAirflow() *string {
	var r *string
	if e != nil && e.Airflow != nil {
		r = e.Airflow
	}
	return r
}

func (e *DeviceType) GetWeight() *float64 {
	var r *float64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *DeviceType) GetWeightUnit() *string {
	var r *string
	if e != nil && e.WeightUnit != nil {
		r = e.WeightUnit
	}
	return r
}

func (e *DeviceType) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *DeviceType) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *DeviceType) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *DeviceType) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *DeviceType) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *DeviceType) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type FHRPGroup struct {
	Name         *string
	Protocol     *string
	GroupId      *int64
	AuthType     *string
	AuthKey      *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *FHRPGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.FHRPGroup{
		Name:         e.GetName(),
		Protocol:     e.GetProtocol(),
		GroupId:      e.GetGroupId(),
		AuthType:     e.GetAuthType(),
		AuthKey:      e.GetAuthKey(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *FHRPGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_FhrpGroup{
			FhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
		},
	}
}

func (e *FHRPGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *FHRPGroup) GetName() *string {
	var r *string
	if e != nil && e.Name != nil {
		r = e.Name
	}
	return r
}

func (e *FHRPGroup) GetProtocol() string {
	var r string
	if e != nil && e.Protocol != nil {
		r = *e.Protocol
	}
	return r
}

func (e *FHRPGroup) GetGroupId() int64 {
	var r int64
	if e != nil && e.GroupId != nil {
		r = *e.GroupId
	}
	return r
}

func (e *FHRPGroup) GetAuthType() *string {
	var r *string
	if e != nil && e.AuthType != nil {
		r = e.AuthType
	}
	return r
}

func (e *FHRPGroup) GetAuthKey() *string {
	var r *string
	if e != nil && e.AuthKey != nil {
		r = e.AuthKey
	}
	return r
}

func (e *FHRPGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *FHRPGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *FHRPGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *FHRPGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *FHRPGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *FHRPGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type FHRPGroupAssignment struct {
	Group *FHRPGroup
	// Interface can be:
	//  - ASN
	//  - ASNRange
	//  - Aggregate
	//  - Cable
	//  - CablePath
	//  - CableTermination
	//  - Circuit
	//  - CircuitGroup
	//  - CircuitGroupAssignment
	//  - CircuitTermination
	//  - CircuitType
	//  - Cluster
	//  - ClusterGroup
	//  - ClusterType
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - Contact
	//  - ContactAssignment
	//  - ContactGroup
	//  - ContactRole
	//  - Device
	//  - DeviceBay
	//  - DeviceRole
	//  - DeviceType
	//  - FHRPGroup
	//  - FHRPGroupAssignment
	//  - FrontPort
	//  - IKEPolicy
	//  - IKEProposal
	//  - IPAddress
	//  - IPRange
	//  - IPSecPolicy
	//  - IPSecProfile
	//  - IPSecProposal
	//  - Interface
	//  - InventoryItem
	//  - InventoryItemRole
	//  - L2VPN
	//  - L2VPNTermination
	//  - Location
	//  - MACAddress
	//  - Manufacturer
	//  - Module
	//  - ModuleBay
	//  - ModuleType
	//  - Platform
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPanel
	//  - PowerPort
	//  - Prefix
	//  - Provider
	//  - ProviderAccount
	//  - ProviderNetwork
	//  - RIR
	//  - Rack
	//  - RackReservation
	//  - RackRole
	//  - RackType
	//  - RearPort
	//  - Region
	//  - Role
	//  - RouteTarget
	//  - Service
	//  - Site
	//  - SiteGroup
	//  - Tag
	//  - Tenant
	//  - TenantGroup
	//  - Tunnel
	//  - TunnelGroup
	//  - TunnelTermination
	//  - VLAN
	//  - VLANGroup
	//  - VLANTranslationPolicy
	//  - VLANTranslationRule
	//  - VMInterface
	//  - VRF
	//  - VirtualChassis
	//  - VirtualCircuit
	//  - VirtualCircuitTermination
	//  - VirtualCircuitType
	//  - VirtualDeviceContext
	//  - VirtualDisk
	//  - VirtualMachine
	//  - WirelessLAN
	//  - WirelessLANGroup
	//  - WirelessLink
	//  - CustomField
	//  - CustomFieldChoiceSet
	//  - JournalEntry
	//  - ModuleTypeProfile
	//  - CustomLink
	//  - Owner
	//  - OwnerGroup
	//  - CableBundle
	//  - RackGroup
	//  - ScriptModule
	//  - VirtualMachineType
	Interface anyFHRPGroupAssignmentInterfaceValue
	Priority  *int64
	Metadata  Metadata
}

func (e *FHRPGroupAssignment) ConvertToProtoMessage() proto.Message {
	r := &pb.FHRPGroupAssignment{
		Group:    e.GetGroup(),
		Priority: e.GetPriority(),
		Metadata: e.GetMetadata(),
	}
	if e.Interface != nil {
		e.Interface.anyFHRPGroupAssignmentInterfaceValueApplyTo(r)
	}
	return r
}

func (e *FHRPGroupAssignment) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_FhrpGroupAssignment{
			FhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
		},
	}
}

func (e *FHRPGroupAssignment) GetGroup() *pb.FHRPGroup {
	var r *pb.FHRPGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.FHRPGroup)
	}
	return r
}

func (e *FHRPGroupAssignment) GetInterface() any {
	var r any
	if e != nil && e.Interface != nil {
		var tmp pb.FHRPGroupAssignment
		e.Interface.anyFHRPGroupAssignmentInterfaceValueApplyTo(&tmp)
		r = tmp.Interface
	}
	return r
}

func (e *FHRPGroupAssignment) GetPriority() int64 {
	var r int64
	if e != nil && e.Priority != nil {
		r = *e.Priority
	}
	return r
}

func (e *FHRPGroupAssignment) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type FrontPort struct {
	Device           *Device
	Module           *Module
	Name             *string
	Label            *string
	Type             *string
	Color            *string
	RearPort         *RearPort
	RearPortPosition *int64
	Description      *string
	MarkConnected    *bool
	Tags             []*Tag
	CustomFields     map[string]*CustomFieldValue
	Metadata         Metadata
	Positions        *int64
	Owner            *Owner
}

func (e *FrontPort) ConvertToProtoMessage() proto.Message {
	r := &pb.FrontPort{
		Device:           e.GetDevice(),
		Module:           e.GetModule(),
		Name:             e.GetName(),
		Label:            e.GetLabel(),
		Type:             e.GetType(),
		Color:            e.GetColor(),
		RearPort:         e.GetRearPort(),
		RearPortPosition: e.GetRearPortPosition(),
		Description:      e.GetDescription(),
		MarkConnected:    e.GetMarkConnected(),
		Tags:             e.GetTags(),
		CustomFields:     e.GetCustomFields(),
		Metadata:         e.GetMetadata(),
		Positions:        e.GetPositions(),
		Owner:            e.GetOwner(),
	}
	return r
}

func (e *FrontPort) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_FrontPort{
			FrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
		},
	}
}

func (e *FrontPort) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *FrontPort) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *FrontPort) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *FrontPort) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *FrontPort) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *FrontPort) GetType() string {
	var r string
	if e != nil && e.Type != nil {
		r = *e.Type
	}
	return r
}

func (e *FrontPort) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *FrontPort) GetRearPort() *pb.RearPort {
	var r *pb.RearPort
	if e != nil && e.RearPort != nil {
		r = e.RearPort.ConvertToProtoMessage().(*pb.RearPort)
	}
	return r
}

func (e *FrontPort) GetRearPortPosition() *int64 {
	var r *int64
	if e != nil && e.RearPortPosition != nil {
		r = e.RearPortPosition
	}
	return r
}

func (e *FrontPort) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *FrontPort) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *FrontPort) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *FrontPort) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *FrontPort) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *FrontPort) GetPositions() *int64 {
	var r *int64
	if e != nil && e.Positions != nil {
		r = e.Positions
	}
	return r
}

func (e *FrontPort) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type GenericObject struct {
	// Object can be:
	//  - ASN
	//  - ASNRange
	//  - Aggregate
	//  - Cable
	//  - CablePath
	//  - CableTermination
	//  - Circuit
	//  - CircuitGroup
	//  - CircuitGroupAssignment
	//  - CircuitTermination
	//  - CircuitType
	//  - Cluster
	//  - ClusterGroup
	//  - ClusterType
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - Contact
	//  - ContactAssignment
	//  - ContactGroup
	//  - ContactRole
	//  - Device
	//  - DeviceBay
	//  - DeviceRole
	//  - DeviceType
	//  - FHRPGroup
	//  - FHRPGroupAssignment
	//  - FrontPort
	//  - IKEPolicy
	//  - IKEProposal
	//  - IPAddress
	//  - IPRange
	//  - IPSecPolicy
	//  - IPSecProfile
	//  - IPSecProposal
	//  - Interface
	//  - InventoryItem
	//  - InventoryItemRole
	//  - L2VPN
	//  - L2VPNTermination
	//  - Location
	//  - MACAddress
	//  - Manufacturer
	//  - Module
	//  - ModuleBay
	//  - ModuleType
	//  - Platform
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPanel
	//  - PowerPort
	//  - Prefix
	//  - Provider
	//  - ProviderAccount
	//  - ProviderNetwork
	//  - RIR
	//  - Rack
	//  - RackReservation
	//  - RackRole
	//  - RackType
	//  - RearPort
	//  - Region
	//  - Role
	//  - RouteTarget
	//  - Service
	//  - Site
	//  - SiteGroup
	//  - Tag
	//  - Tenant
	//  - TenantGroup
	//  - Tunnel
	//  - TunnelGroup
	//  - TunnelTermination
	//  - VLAN
	//  - VLANGroup
	//  - VLANTranslationPolicy
	//  - VLANTranslationRule
	//  - VMInterface
	//  - VRF
	//  - VirtualChassis
	//  - VirtualCircuit
	//  - VirtualCircuitTermination
	//  - VirtualCircuitType
	//  - VirtualDeviceContext
	//  - VirtualDisk
	//  - VirtualMachine
	//  - WirelessLAN
	//  - WirelessLANGroup
	//  - WirelessLink
	//  - CustomField
	//  - CustomFieldChoiceSet
	//  - JournalEntry
	//  - ModuleTypeProfile
	//  - CustomLink
	//  - Owner
	//  - OwnerGroup
	//  - CableBundle
	//  - RackGroup
	//  - ScriptModule
	//  - VirtualMachineType
	Object anyGenericObjectObjectValue
}

func (e *GenericObject) ConvertToProtoMessage() proto.Message {
	r := &pb.GenericObject{}
	if e.Object != nil {
		e.Object.anyGenericObjectObjectValueApplyTo(r)
	}
	return r
}

func (e *GenericObject) GetObject() any {
	var r any
	if e != nil && e.Object != nil {
		var tmp pb.GenericObject
		e.Object.anyGenericObjectObjectValueApplyTo(&tmp)
		r = tmp.Object
	}
	return r
}

type IKEPolicy struct {
	Name         *string
	Description  *string
	Version      *int64
	Mode         *string
	PresharedKey *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Proposals    []*IKEProposal
	Metadata     Metadata
	Owner        *Owner
}

func (e *IKEPolicy) ConvertToProtoMessage() proto.Message {
	r := &pb.IKEPolicy{
		Name:         e.GetName(),
		Description:  e.GetDescription(),
		Version:      e.GetVersion(),
		Mode:         e.GetMode(),
		PresharedKey: e.GetPresharedKey(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Proposals:    e.GetProposals(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *IKEPolicy) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_IkePolicy{
			IkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
		},
	}
}

func (e *IKEPolicy) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *IKEPolicy) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *IKEPolicy) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *IKEPolicy) GetVersion() int64 {
	var r int64
	if e != nil && e.Version != nil {
		r = *e.Version
	}
	return r
}

func (e *IKEPolicy) GetMode() *string {
	var r *string
	if e != nil && e.Mode != nil {
		r = e.Mode
	}
	return r
}

func (e *IKEPolicy) GetPresharedKey() *string {
	var r *string
	if e != nil && e.PresharedKey != nil {
		r = e.PresharedKey
	}
	return r
}

func (e *IKEPolicy) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *IKEPolicy) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *IKEPolicy) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *IKEPolicy) GetProposals() []*pb.IKEProposal {
	var r []*pb.IKEProposal
	if e != nil && e.Proposals != nil {
		for _, v := range e.Proposals {
			r = append(r, v.ConvertToProtoMessage().(*pb.IKEProposal))
		}
	}
	return r
}

func (e *IKEPolicy) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *IKEPolicy) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type IKEProposal struct {
	Name                    *string
	Description             *string
	AuthenticationMethod    *string
	EncryptionAlgorithm     *string
	AuthenticationAlgorithm *string
	Group                   *int64
	SaLifetime              *int64
	Comments                *string
	Tags                    []*Tag
	CustomFields            map[string]*CustomFieldValue
	Metadata                Metadata
	Owner                   *Owner
}

func (e *IKEProposal) ConvertToProtoMessage() proto.Message {
	r := &pb.IKEProposal{
		Name:                    e.GetName(),
		Description:             e.GetDescription(),
		AuthenticationMethod:    e.GetAuthenticationMethod(),
		EncryptionAlgorithm:     e.GetEncryptionAlgorithm(),
		AuthenticationAlgorithm: e.GetAuthenticationAlgorithm(),
		Group:                   e.GetGroup(),
		SaLifetime:              e.GetSaLifetime(),
		Comments:                e.GetComments(),
		Tags:                    e.GetTags(),
		CustomFields:            e.GetCustomFields(),
		Metadata:                e.GetMetadata(),
		Owner:                   e.GetOwner(),
	}
	return r
}

func (e *IKEProposal) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_IkeProposal{
			IkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
		},
	}
}

func (e *IKEProposal) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *IKEProposal) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *IKEProposal) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *IKEProposal) GetAuthenticationMethod() string {
	var r string
	if e != nil && e.AuthenticationMethod != nil {
		r = *e.AuthenticationMethod
	}
	return r
}

func (e *IKEProposal) GetEncryptionAlgorithm() string {
	var r string
	if e != nil && e.EncryptionAlgorithm != nil {
		r = *e.EncryptionAlgorithm
	}
	return r
}

func (e *IKEProposal) GetAuthenticationAlgorithm() *string {
	var r *string
	if e != nil && e.AuthenticationAlgorithm != nil {
		r = e.AuthenticationAlgorithm
	}
	return r
}

func (e *IKEProposal) GetGroup() int64 {
	var r int64
	if e != nil && e.Group != nil {
		r = *e.Group
	}
	return r
}

func (e *IKEProposal) GetSaLifetime() *int64 {
	var r *int64
	if e != nil && e.SaLifetime != nil {
		r = e.SaLifetime
	}
	return r
}

func (e *IKEProposal) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *IKEProposal) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *IKEProposal) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *IKEProposal) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *IKEProposal) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type IPAddress struct {
	Address *string
	Vrf     *VRF
	Tenant  *Tenant
	Status  *string
	Role    *string
	// AssignedObject can be:
	//  - FHRPGroup
	//  - Interface
	//  - VMInterface
	AssignedObject anyIPAddressAssignedObjectValue
	NatInside      *IPAddress
	DnsName        *string
	Description    *string
	Comments       *string
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
	Owner          *Owner
}

func (e *IPAddress) ConvertToProtoMessage() proto.Message {
	r := &pb.IPAddress{
		Address:      e.GetAddress(),
		Vrf:          e.GetVrf(),
		Tenant:       e.GetTenant(),
		Status:       e.GetStatus(),
		Role:         e.GetRole(),
		NatInside:    e.GetNatInside(),
		DnsName:      e.GetDnsName(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	if e.AssignedObject != nil {
		e.AssignedObject.anyIPAddressAssignedObjectValueApplyTo(r)
	}
	return r
}

func (e *IPAddress) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_IpAddress{
			IpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
		},
	}
}

func (e *IPAddress) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *IPAddress) GetAddress() string {
	var r string
	if e != nil && e.Address != nil {
		r = *e.Address
	}
	return r
}

func (e *IPAddress) GetVrf() *pb.VRF {
	var r *pb.VRF
	if e != nil && e.Vrf != nil {
		r = e.Vrf.ConvertToProtoMessage().(*pb.VRF)
	}
	return r
}

func (e *IPAddress) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *IPAddress) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *IPAddress) GetRole() *string {
	var r *string
	if e != nil && e.Role != nil {
		r = e.Role
	}
	return r
}

func (e *IPAddress) GetAssignedObject() any {
	var r any
	if e != nil && e.AssignedObject != nil {
		var tmp pb.IPAddress
		e.AssignedObject.anyIPAddressAssignedObjectValueApplyTo(&tmp)
		r = tmp.AssignedObject
	}
	return r
}

func (e *IPAddress) GetNatInside() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.NatInside != nil {
		r = e.NatInside.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *IPAddress) GetDnsName() *string {
	var r *string
	if e != nil && e.DnsName != nil {
		r = e.DnsName
	}
	return r
}

func (e *IPAddress) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *IPAddress) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *IPAddress) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *IPAddress) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *IPAddress) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *IPAddress) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type IPRange struct {
	StartAddress  *string
	EndAddress    *string
	Vrf           *VRF
	Tenant        *Tenant
	Status        *string
	Role          *Role
	Description   *string
	Comments      *string
	Tags          []*Tag
	MarkUtilized  *bool
	CustomFields  map[string]*CustomFieldValue
	MarkPopulated *bool
	Metadata      Metadata
	Owner         *Owner
}

func (e *IPRange) ConvertToProtoMessage() proto.Message {
	r := &pb.IPRange{
		StartAddress:  e.GetStartAddress(),
		EndAddress:    e.GetEndAddress(),
		Vrf:           e.GetVrf(),
		Tenant:        e.GetTenant(),
		Status:        e.GetStatus(),
		Role:          e.GetRole(),
		Description:   e.GetDescription(),
		Comments:      e.GetComments(),
		Tags:          e.GetTags(),
		MarkUtilized:  e.GetMarkUtilized(),
		CustomFields:  e.GetCustomFields(),
		MarkPopulated: e.GetMarkPopulated(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *IPRange) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_IpRange{
			IpRange: e.ConvertToProtoMessage().(*pb.IPRange),
		},
	}
}

func (e *IPRange) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *IPRange) GetStartAddress() string {
	var r string
	if e != nil && e.StartAddress != nil {
		r = *e.StartAddress
	}
	return r
}

func (e *IPRange) GetEndAddress() string {
	var r string
	if e != nil && e.EndAddress != nil {
		r = *e.EndAddress
	}
	return r
}

func (e *IPRange) GetVrf() *pb.VRF {
	var r *pb.VRF
	if e != nil && e.Vrf != nil {
		r = e.Vrf.ConvertToProtoMessage().(*pb.VRF)
	}
	return r
}

func (e *IPRange) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *IPRange) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *IPRange) GetRole() *pb.Role {
	var r *pb.Role
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.Role)
	}
	return r
}

func (e *IPRange) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *IPRange) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *IPRange) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *IPRange) GetMarkUtilized() *bool {
	var r *bool
	if e != nil && e.MarkUtilized != nil {
		r = e.MarkUtilized
	}
	return r
}

func (e *IPRange) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *IPRange) GetMarkPopulated() *bool {
	var r *bool
	if e != nil && e.MarkPopulated != nil {
		r = e.MarkPopulated
	}
	return r
}

func (e *IPRange) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *IPRange) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type IPSecPolicy struct {
	Name         *string
	Description  *string
	PfsGroup     *int64
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Proposals    []*IPSecProposal
	Metadata     Metadata
	Owner        *Owner
}

func (e *IPSecPolicy) ConvertToProtoMessage() proto.Message {
	r := &pb.IPSecPolicy{
		Name:         e.GetName(),
		Description:  e.GetDescription(),
		PfsGroup:     e.GetPfsGroup(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Proposals:    e.GetProposals(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *IPSecPolicy) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_IpSecPolicy{
			IpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
		},
	}
}

func (e *IPSecPolicy) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *IPSecPolicy) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *IPSecPolicy) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *IPSecPolicy) GetPfsGroup() *int64 {
	var r *int64
	if e != nil && e.PfsGroup != nil {
		r = e.PfsGroup
	}
	return r
}

func (e *IPSecPolicy) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *IPSecPolicy) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *IPSecPolicy) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *IPSecPolicy) GetProposals() []*pb.IPSecProposal {
	var r []*pb.IPSecProposal
	if e != nil && e.Proposals != nil {
		for _, v := range e.Proposals {
			r = append(r, v.ConvertToProtoMessage().(*pb.IPSecProposal))
		}
	}
	return r
}

func (e *IPSecPolicy) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *IPSecPolicy) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type IPSecProfile struct {
	Name         *string
	Description  *string
	Mode         *string
	IkePolicy    *IKEPolicy
	IpsecPolicy  *IPSecPolicy
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *IPSecProfile) ConvertToProtoMessage() proto.Message {
	r := &pb.IPSecProfile{
		Name:         e.GetName(),
		Description:  e.GetDescription(),
		Mode:         e.GetMode(),
		IkePolicy:    e.GetIkePolicy(),
		IpsecPolicy:  e.GetIpsecPolicy(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *IPSecProfile) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_IpSecProfile{
			IpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
		},
	}
}

func (e *IPSecProfile) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *IPSecProfile) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *IPSecProfile) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *IPSecProfile) GetMode() string {
	var r string
	if e != nil && e.Mode != nil {
		r = *e.Mode
	}
	return r
}

func (e *IPSecProfile) GetIkePolicy() *pb.IKEPolicy {
	var r *pb.IKEPolicy
	if e != nil && e.IkePolicy != nil {
		r = e.IkePolicy.ConvertToProtoMessage().(*pb.IKEPolicy)
	}
	return r
}

func (e *IPSecProfile) GetIpsecPolicy() *pb.IPSecPolicy {
	var r *pb.IPSecPolicy
	if e != nil && e.IpsecPolicy != nil {
		r = e.IpsecPolicy.ConvertToProtoMessage().(*pb.IPSecPolicy)
	}
	return r
}

func (e *IPSecProfile) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *IPSecProfile) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *IPSecProfile) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *IPSecProfile) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *IPSecProfile) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type IPSecProposal struct {
	Name                    *string
	Description             *string
	EncryptionAlgorithm     *string
	AuthenticationAlgorithm *string
	SaLifetimeSeconds       *int64
	SaLifetimeData          *int64
	Comments                *string
	Tags                    []*Tag
	CustomFields            map[string]*CustomFieldValue
	Metadata                Metadata
	Owner                   *Owner
}

func (e *IPSecProposal) ConvertToProtoMessage() proto.Message {
	r := &pb.IPSecProposal{
		Name:                    e.GetName(),
		Description:             e.GetDescription(),
		EncryptionAlgorithm:     e.GetEncryptionAlgorithm(),
		AuthenticationAlgorithm: e.GetAuthenticationAlgorithm(),
		SaLifetimeSeconds:       e.GetSaLifetimeSeconds(),
		SaLifetimeData:          e.GetSaLifetimeData(),
		Comments:                e.GetComments(),
		Tags:                    e.GetTags(),
		CustomFields:            e.GetCustomFields(),
		Metadata:                e.GetMetadata(),
		Owner:                   e.GetOwner(),
	}
	return r
}

func (e *IPSecProposal) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_IpSecProposal{
			IpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
		},
	}
}

func (e *IPSecProposal) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *IPSecProposal) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *IPSecProposal) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *IPSecProposal) GetEncryptionAlgorithm() *string {
	var r *string
	if e != nil && e.EncryptionAlgorithm != nil {
		r = e.EncryptionAlgorithm
	}
	return r
}

func (e *IPSecProposal) GetAuthenticationAlgorithm() *string {
	var r *string
	if e != nil && e.AuthenticationAlgorithm != nil {
		r = e.AuthenticationAlgorithm
	}
	return r
}

func (e *IPSecProposal) GetSaLifetimeSeconds() *int64 {
	var r *int64
	if e != nil && e.SaLifetimeSeconds != nil {
		r = e.SaLifetimeSeconds
	}
	return r
}

func (e *IPSecProposal) GetSaLifetimeData() *int64 {
	var r *int64
	if e != nil && e.SaLifetimeData != nil {
		r = e.SaLifetimeData
	}
	return r
}

func (e *IPSecProposal) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *IPSecProposal) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *IPSecProposal) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *IPSecProposal) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *IPSecProposal) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Interface struct {
	Device                *Device
	Module                *Module
	Name                  *string
	Label                 *string
	Type                  *string
	Enabled               *bool
	Parent                *Interface
	Bridge                *Interface
	Lag                   *Interface
	Mtu                   *int64
	PrimaryMacAddress     *MACAddress
	Speed                 *int64
	Duplex                *string
	Wwn                   *string
	MgmtOnly              *bool
	Description           *string
	Mode                  *string
	RfRole                *string
	RfChannel             *string
	PoeMode               *string
	PoeType               *string
	RfChannelFrequency    *float64
	RfChannelWidth        *float64
	TxPower               *int64
	UntaggedVlan          *VLAN
	QinqSvlan             *VLAN
	VlanTranslationPolicy *VLANTranslationPolicy
	MarkConnected         *bool
	Vrf                   *VRF
	Tags                  []*Tag
	CustomFields          map[string]*CustomFieldValue
	Vdcs                  []*VirtualDeviceContext
	TaggedVlans           []*VLAN
	WirelessLans          []*WirelessLAN
	Metadata              Metadata
	Owner                 *Owner
}

func (e *Interface) ConvertToProtoMessage() proto.Message {
	r := &pb.Interface{
		Device:                e.GetDevice(),
		Module:                e.GetModule(),
		Name:                  e.GetName(),
		Label:                 e.GetLabel(),
		Type:                  e.GetType(),
		Enabled:               e.GetEnabled(),
		Parent:                e.GetParent(),
		Bridge:                e.GetBridge(),
		Lag:                   e.GetLag(),
		Mtu:                   e.GetMtu(),
		PrimaryMacAddress:     e.GetPrimaryMacAddress(),
		Speed:                 e.GetSpeed(),
		Duplex:                e.GetDuplex(),
		Wwn:                   e.GetWwn(),
		MgmtOnly:              e.GetMgmtOnly(),
		Description:           e.GetDescription(),
		Mode:                  e.GetMode(),
		RfRole:                e.GetRfRole(),
		RfChannel:             e.GetRfChannel(),
		PoeMode:               e.GetPoeMode(),
		PoeType:               e.GetPoeType(),
		RfChannelFrequency:    e.GetRfChannelFrequency(),
		RfChannelWidth:        e.GetRfChannelWidth(),
		TxPower:               e.GetTxPower(),
		UntaggedVlan:          e.GetUntaggedVlan(),
		QinqSvlan:             e.GetQinqSvlan(),
		VlanTranslationPolicy: e.GetVlanTranslationPolicy(),
		MarkConnected:         e.GetMarkConnected(),
		Vrf:                   e.GetVrf(),
		Tags:                  e.GetTags(),
		CustomFields:          e.GetCustomFields(),
		Vdcs:                  e.GetVdcs(),
		TaggedVlans:           e.GetTaggedVlans(),
		WirelessLans:          e.GetWirelessLans(),
		Metadata:              e.GetMetadata(),
		Owner:                 e.GetOwner(),
	}
	return r
}

func (e *Interface) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Interface{
			Interface: e.ConvertToProtoMessage().(*pb.Interface),
		},
	}
}

func (e *Interface) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Interface) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *Interface) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *Interface) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Interface) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *Interface) GetType() string {
	var r string
	if e != nil && e.Type != nil {
		r = *e.Type
	}
	return r
}

func (e *Interface) GetEnabled() *bool {
	var r *bool
	if e != nil && e.Enabled != nil {
		r = e.Enabled
	}
	return r
}

func (e *Interface) GetParent() *pb.Interface {
	var r *pb.Interface
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.Interface)
	}
	return r
}

func (e *Interface) GetBridge() *pb.Interface {
	var r *pb.Interface
	if e != nil && e.Bridge != nil {
		r = e.Bridge.ConvertToProtoMessage().(*pb.Interface)
	}
	return r
}

func (e *Interface) GetLag() *pb.Interface {
	var r *pb.Interface
	if e != nil && e.Lag != nil {
		r = e.Lag.ConvertToProtoMessage().(*pb.Interface)
	}
	return r
}

func (e *Interface) GetMtu() *int64 {
	var r *int64
	if e != nil && e.Mtu != nil {
		r = e.Mtu
	}
	return r
}

func (e *Interface) GetPrimaryMacAddress() *pb.MACAddress {
	var r *pb.MACAddress
	if e != nil && e.PrimaryMacAddress != nil {
		r = e.PrimaryMacAddress.ConvertToProtoMessage().(*pb.MACAddress)
	}
	return r
}

func (e *Interface) GetSpeed() *int64 {
	var r *int64
	if e != nil && e.Speed != nil {
		r = e.Speed
	}
	return r
}

func (e *Interface) GetDuplex() *string {
	var r *string
	if e != nil && e.Duplex != nil {
		r = e.Duplex
	}
	return r
}

func (e *Interface) GetWwn() *string {
	var r *string
	if e != nil && e.Wwn != nil {
		r = e.Wwn
	}
	return r
}

func (e *Interface) GetMgmtOnly() *bool {
	var r *bool
	if e != nil && e.MgmtOnly != nil {
		r = e.MgmtOnly
	}
	return r
}

func (e *Interface) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Interface) GetMode() *string {
	var r *string
	if e != nil && e.Mode != nil {
		r = e.Mode
	}
	return r
}

func (e *Interface) GetRfRole() *string {
	var r *string
	if e != nil && e.RfRole != nil {
		r = e.RfRole
	}
	return r
}

func (e *Interface) GetRfChannel() *string {
	var r *string
	if e != nil && e.RfChannel != nil {
		r = e.RfChannel
	}
	return r
}

func (e *Interface) GetPoeMode() *string {
	var r *string
	if e != nil && e.PoeMode != nil {
		r = e.PoeMode
	}
	return r
}

func (e *Interface) GetPoeType() *string {
	var r *string
	if e != nil && e.PoeType != nil {
		r = e.PoeType
	}
	return r
}

func (e *Interface) GetRfChannelFrequency() *float64 {
	var r *float64
	if e != nil && e.RfChannelFrequency != nil {
		r = e.RfChannelFrequency
	}
	return r
}

func (e *Interface) GetRfChannelWidth() *float64 {
	var r *float64
	if e != nil && e.RfChannelWidth != nil {
		r = e.RfChannelWidth
	}
	return r
}

func (e *Interface) GetTxPower() *int64 {
	var r *int64
	if e != nil && e.TxPower != nil {
		r = e.TxPower
	}
	return r
}

func (e *Interface) GetUntaggedVlan() *pb.VLAN {
	var r *pb.VLAN
	if e != nil && e.UntaggedVlan != nil {
		r = e.UntaggedVlan.ConvertToProtoMessage().(*pb.VLAN)
	}
	return r
}

func (e *Interface) GetQinqSvlan() *pb.VLAN {
	var r *pb.VLAN
	if e != nil && e.QinqSvlan != nil {
		r = e.QinqSvlan.ConvertToProtoMessage().(*pb.VLAN)
	}
	return r
}

func (e *Interface) GetVlanTranslationPolicy() *pb.VLANTranslationPolicy {
	var r *pb.VLANTranslationPolicy
	if e != nil && e.VlanTranslationPolicy != nil {
		r = e.VlanTranslationPolicy.ConvertToProtoMessage().(*pb.VLANTranslationPolicy)
	}
	return r
}

func (e *Interface) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *Interface) GetVrf() *pb.VRF {
	var r *pb.VRF
	if e != nil && e.Vrf != nil {
		r = e.Vrf.ConvertToProtoMessage().(*pb.VRF)
	}
	return r
}

func (e *Interface) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Interface) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Interface) GetVdcs() []*pb.VirtualDeviceContext {
	var r []*pb.VirtualDeviceContext
	if e != nil && e.Vdcs != nil {
		for _, v := range e.Vdcs {
			r = append(r, v.ConvertToProtoMessage().(*pb.VirtualDeviceContext))
		}
	}
	return r
}

func (e *Interface) GetTaggedVlans() []*pb.VLAN {
	var r []*pb.VLAN
	if e != nil && e.TaggedVlans != nil {
		for _, v := range e.TaggedVlans {
			r = append(r, v.ConvertToProtoMessage().(*pb.VLAN))
		}
	}
	return r
}

func (e *Interface) GetWirelessLans() []*pb.WirelessLAN {
	var r []*pb.WirelessLAN
	if e != nil && e.WirelessLans != nil {
		for _, v := range e.WirelessLans {
			r = append(r, v.ConvertToProtoMessage().(*pb.WirelessLAN))
		}
	}
	return r
}

func (e *Interface) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Interface) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type InventoryItem struct {
	Device       *Device
	Parent       *InventoryItem
	Name         *string
	Label        *string
	Status       *string
	Role         *InventoryItemRole
	Manufacturer *Manufacturer
	PartId       *string
	Serial       *string
	AssetTag     *string
	Discovered   *bool
	Description  *string
	// Component can be:
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - FrontPort
	//  - Interface
	//  - PowerOutlet
	//  - PowerPort
	//  - RearPort
	Component    anyInventoryItemComponentValue
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *InventoryItem) ConvertToProtoMessage() proto.Message {
	r := &pb.InventoryItem{
		Device:       e.GetDevice(),
		Parent:       e.GetParent(),
		Name:         e.GetName(),
		Label:        e.GetLabel(),
		Status:       e.GetStatus(),
		Role:         e.GetRole(),
		Manufacturer: e.GetManufacturer(),
		PartId:       e.GetPartId(),
		Serial:       e.GetSerial(),
		AssetTag:     e.GetAssetTag(),
		Discovered:   e.GetDiscovered(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	if e.Component != nil {
		e.Component.anyInventoryItemComponentValueApplyTo(r)
	}
	return r
}

func (e *InventoryItem) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_InventoryItem{
			InventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
		},
	}
}

func (e *InventoryItem) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *InventoryItem) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *InventoryItem) GetParent() *pb.InventoryItem {
	var r *pb.InventoryItem
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.InventoryItem)
	}
	return r
}

func (e *InventoryItem) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *InventoryItem) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *InventoryItem) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *InventoryItem) GetRole() *pb.InventoryItemRole {
	var r *pb.InventoryItemRole
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.InventoryItemRole)
	}
	return r
}

func (e *InventoryItem) GetManufacturer() *pb.Manufacturer {
	var r *pb.Manufacturer
	if e != nil && e.Manufacturer != nil {
		r = e.Manufacturer.ConvertToProtoMessage().(*pb.Manufacturer)
	}
	return r
}

func (e *InventoryItem) GetPartId() *string {
	var r *string
	if e != nil && e.PartId != nil {
		r = e.PartId
	}
	return r
}

func (e *InventoryItem) GetSerial() *string {
	var r *string
	if e != nil && e.Serial != nil {
		r = e.Serial
	}
	return r
}

func (e *InventoryItem) GetAssetTag() *string {
	var r *string
	if e != nil && e.AssetTag != nil {
		r = e.AssetTag
	}
	return r
}

func (e *InventoryItem) GetDiscovered() *bool {
	var r *bool
	if e != nil && e.Discovered != nil {
		r = e.Discovered
	}
	return r
}

func (e *InventoryItem) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *InventoryItem) GetComponent() any {
	var r any
	if e != nil && e.Component != nil {
		var tmp pb.InventoryItem
		e.Component.anyInventoryItemComponentValueApplyTo(&tmp)
		r = tmp.Component
	}
	return r
}

func (e *InventoryItem) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *InventoryItem) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *InventoryItem) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *InventoryItem) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type InventoryItemRole struct {
	Name         *string
	Slug         *string
	Color        *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *InventoryItemRole) ConvertToProtoMessage() proto.Message {
	r := &pb.InventoryItemRole{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Color:        e.GetColor(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *InventoryItemRole) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_InventoryItemRole{
			InventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
		},
	}
}

func (e *InventoryItemRole) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *InventoryItemRole) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *InventoryItemRole) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *InventoryItemRole) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *InventoryItemRole) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *InventoryItemRole) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *InventoryItemRole) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *InventoryItemRole) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *InventoryItemRole) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *InventoryItemRole) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type L2VPN struct {
	Identifier    *int64
	Name          *string
	Slug          *string
	Type          *string
	Description   *string
	Comments      *string
	Tenant        *Tenant
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	ImportTargets []*RouteTarget
	ExportTargets []*RouteTarget
	Status        *string
	Metadata      Metadata
	Owner         *Owner
}

func (e *L2VPN) ConvertToProtoMessage() proto.Message {
	r := &pb.L2VPN{
		Identifier:    e.GetIdentifier(),
		Name:          e.GetName(),
		Slug:          e.GetSlug(),
		Type:          e.GetType(),
		Description:   e.GetDescription(),
		Comments:      e.GetComments(),
		Tenant:        e.GetTenant(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		ImportTargets: e.GetImportTargets(),
		ExportTargets: e.GetExportTargets(),
		Status:        e.GetStatus(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *L2VPN) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_L2Vpn{
			L2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
		},
	}
}

func (e *L2VPN) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *L2VPN) GetIdentifier() *int64 {
	var r *int64
	if e != nil && e.Identifier != nil {
		r = e.Identifier
	}
	return r
}

func (e *L2VPN) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *L2VPN) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *L2VPN) GetType() *string {
	var r *string
	if e != nil && e.Type != nil {
		r = e.Type
	}
	return r
}

func (e *L2VPN) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *L2VPN) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *L2VPN) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *L2VPN) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *L2VPN) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *L2VPN) GetImportTargets() []*pb.RouteTarget {
	var r []*pb.RouteTarget
	if e != nil && e.ImportTargets != nil {
		for _, v := range e.ImportTargets {
			r = append(r, v.ConvertToProtoMessage().(*pb.RouteTarget))
		}
	}
	return r
}

func (e *L2VPN) GetExportTargets() []*pb.RouteTarget {
	var r []*pb.RouteTarget
	if e != nil && e.ExportTargets != nil {
		for _, v := range e.ExportTargets {
			r = append(r, v.ConvertToProtoMessage().(*pb.RouteTarget))
		}
	}
	return r
}

func (e *L2VPN) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *L2VPN) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *L2VPN) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type L2VPNTermination struct {
	L2Vpn *L2VPN
	// AssignedObject can be:
	//  - Interface
	//  - VLAN
	//  - VMInterface
	//  - ASN
	//  - ASNRange
	//  - Aggregate
	//  - Cable
	//  - CablePath
	//  - CableTermination
	//  - Circuit
	//  - CircuitGroup
	//  - CircuitGroupAssignment
	//  - CircuitTermination
	//  - CircuitType
	//  - Cluster
	//  - ClusterGroup
	//  - ClusterType
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - Contact
	//  - ContactAssignment
	//  - ContactGroup
	//  - ContactRole
	//  - CustomField
	//  - CustomFieldChoiceSet
	//  - Device
	//  - DeviceBay
	//  - DeviceRole
	//  - DeviceType
	//  - FHRPGroup
	//  - FHRPGroupAssignment
	//  - FrontPort
	//  - IKEPolicy
	//  - IKEProposal
	//  - IPAddress
	//  - IPRange
	//  - IPSecPolicy
	//  - IPSecProfile
	//  - IPSecProposal
	//  - InventoryItem
	//  - InventoryItemRole
	//  - JournalEntry
	//  - L2VPN
	//  - L2VPNTermination
	//  - Location
	//  - MACAddress
	//  - Manufacturer
	//  - Module
	//  - ModuleBay
	//  - ModuleType
	//  - ModuleTypeProfile
	//  - Platform
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPanel
	//  - PowerPort
	//  - Prefix
	//  - Provider
	//  - ProviderAccount
	//  - ProviderNetwork
	//  - RIR
	//  - Rack
	//  - RackReservation
	//  - RackRole
	//  - RackType
	//  - RearPort
	//  - Region
	//  - Role
	//  - RouteTarget
	//  - Service
	//  - Site
	//  - SiteGroup
	//  - Tag
	//  - Tenant
	//  - TenantGroup
	//  - Tunnel
	//  - TunnelGroup
	//  - TunnelTermination
	//  - VLANGroup
	//  - VLANTranslationPolicy
	//  - VLANTranslationRule
	//  - VRF
	//  - VirtualChassis
	//  - VirtualCircuit
	//  - VirtualCircuitTermination
	//  - VirtualCircuitType
	//  - VirtualDeviceContext
	//  - VirtualDisk
	//  - VirtualMachine
	//  - WirelessLAN
	//  - WirelessLANGroup
	//  - WirelessLink
	//  - CustomLink
	//  - Owner
	//  - OwnerGroup
	//  - CableBundle
	//  - RackGroup
	//  - ScriptModule
	//  - VirtualMachineType
	AssignedObject anyL2VPNTerminationAssignedObjectValue
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
}

func (e *L2VPNTermination) ConvertToProtoMessage() proto.Message {
	r := &pb.L2VPNTermination{
		L2Vpn:        e.GetL2Vpn(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
	}
	if e.AssignedObject != nil {
		e.AssignedObject.anyL2VPNTerminationAssignedObjectValueApplyTo(r)
	}
	return r
}

func (e *L2VPNTermination) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_L2VpnTermination{
			L2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
		},
	}
}

func (e *L2VPNTermination) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *L2VPNTermination) GetL2Vpn() *pb.L2VPN {
	var r *pb.L2VPN
	if e != nil && e.L2Vpn != nil {
		r = e.L2Vpn.ConvertToProtoMessage().(*pb.L2VPN)
	}
	return r
}

func (e *L2VPNTermination) GetAssignedObject() any {
	var r any
	if e != nil && e.AssignedObject != nil {
		var tmp pb.L2VPNTermination
		e.AssignedObject.anyL2VPNTerminationAssignedObjectValueApplyTo(&tmp)
		r = tmp.AssignedObject
	}
	return r
}

func (e *L2VPNTermination) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *L2VPNTermination) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *L2VPNTermination) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type Location struct {
	Name         *string
	Slug         *string
	Site         *Site
	Parent       *Location
	Status       *string
	Tenant       *Tenant
	Facility     *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *Location) ConvertToProtoMessage() proto.Message {
	r := &pb.Location{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Site:         e.GetSite(),
		Parent:       e.GetParent(),
		Status:       e.GetStatus(),
		Tenant:       e.GetTenant(),
		Facility:     e.GetFacility(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *Location) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Location{
			Location: e.ConvertToProtoMessage().(*pb.Location),
		},
	}
}

func (e *Location) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Location) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Location) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Location) GetSite() *pb.Site {
	var r *pb.Site
	if e != nil && e.Site != nil {
		r = e.Site.ConvertToProtoMessage().(*pb.Site)
	}
	return r
}

func (e *Location) GetParent() *pb.Location {
	var r *pb.Location
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.Location)
	}
	return r
}

func (e *Location) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Location) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Location) GetFacility() *string {
	var r *string
	if e != nil && e.Facility != nil {
		r = e.Facility
	}
	return r
}

func (e *Location) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Location) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Location) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Location) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Location) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Location) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type MACAddress struct {
	MacAddress *string
	// AssignedObject can be:
	//  - Interface
	//  - VMInterface
	AssignedObject anyMACAddressAssignedObjectValue
	Description    *string
	Comments       *string
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
	Owner          *Owner
}

func (e *MACAddress) ConvertToProtoMessage() proto.Message {
	r := &pb.MACAddress{
		MacAddress:   e.GetMacAddress(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	if e.AssignedObject != nil {
		e.AssignedObject.anyMACAddressAssignedObjectValueApplyTo(r)
	}
	return r
}

func (e *MACAddress) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_MacAddress{
			MacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
		},
	}
}

func (e *MACAddress) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *MACAddress) GetMacAddress() string {
	var r string
	if e != nil && e.MacAddress != nil {
		r = *e.MacAddress
	}
	return r
}

func (e *MACAddress) GetAssignedObject() any {
	var r any
	if e != nil && e.AssignedObject != nil {
		var tmp pb.MACAddress
		e.AssignedObject.anyMACAddressAssignedObjectValueApplyTo(&tmp)
		r = tmp.AssignedObject
	}
	return r
}

func (e *MACAddress) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *MACAddress) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *MACAddress) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *MACAddress) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *MACAddress) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *MACAddress) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Manufacturer struct {
	Name         *string
	Slug         *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *Manufacturer) ConvertToProtoMessage() proto.Message {
	r := &pb.Manufacturer{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *Manufacturer) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Manufacturer{
			Manufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
		},
	}
}

func (e *Manufacturer) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Manufacturer) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Manufacturer) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Manufacturer) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Manufacturer) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Manufacturer) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Manufacturer) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Manufacturer) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *Manufacturer) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type Module struct {
	Device              *Device
	ModuleBay           *ModuleBay
	ModuleType          *ModuleType
	Status              *string
	Serial              *string
	AssetTag            *string
	Description         *string
	Comments            *string
	Tags                []*Tag
	CustomFields        map[string]*CustomFieldValue
	Metadata            Metadata
	Owner               *Owner
	ReplicateComponents *bool
	AdoptComponents     *bool
}

func (e *Module) ConvertToProtoMessage() proto.Message {
	r := &pb.Module{
		Device:              e.GetDevice(),
		ModuleBay:           e.GetModuleBay(),
		ModuleType:          e.GetModuleType(),
		Status:              e.GetStatus(),
		Serial:              e.GetSerial(),
		AssetTag:            e.GetAssetTag(),
		Description:         e.GetDescription(),
		Comments:            e.GetComments(),
		Tags:                e.GetTags(),
		CustomFields:        e.GetCustomFields(),
		Metadata:            e.GetMetadata(),
		Owner:               e.GetOwner(),
		ReplicateComponents: e.GetReplicateComponents(),
		AdoptComponents:     e.GetAdoptComponents(),
	}
	return r
}

func (e *Module) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Module{
			Module: e.ConvertToProtoMessage().(*pb.Module),
		},
	}
}

func (e *Module) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Module) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *Module) GetModuleBay() *pb.ModuleBay {
	var r *pb.ModuleBay
	if e != nil && e.ModuleBay != nil {
		r = e.ModuleBay.ConvertToProtoMessage().(*pb.ModuleBay)
	}
	return r
}

func (e *Module) GetModuleType() *pb.ModuleType {
	var r *pb.ModuleType
	if e != nil && e.ModuleType != nil {
		r = e.ModuleType.ConvertToProtoMessage().(*pb.ModuleType)
	}
	return r
}

func (e *Module) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Module) GetSerial() *string {
	var r *string
	if e != nil && e.Serial != nil {
		r = e.Serial
	}
	return r
}

func (e *Module) GetAssetTag() *string {
	var r *string
	if e != nil && e.AssetTag != nil {
		r = e.AssetTag
	}
	return r
}

func (e *Module) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Module) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Module) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Module) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Module) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Module) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *Module) GetReplicateComponents() *bool {
	var r *bool
	if e != nil && e.ReplicateComponents != nil {
		r = e.ReplicateComponents
	}
	return r
}

func (e *Module) GetAdoptComponents() *bool {
	var r *bool
	if e != nil && e.AdoptComponents != nil {
		r = e.AdoptComponents
	}
	return r
}

type ModuleBay struct {
	Device          *Device
	Module          *Module
	Name            *string
	InstalledModule *Module
	Label           *string
	Position        *string
	Description     *string
	Tags            []*Tag
	CustomFields    map[string]*CustomFieldValue
	Metadata        Metadata
	Owner           *Owner
	Enabled         *bool
}

func (e *ModuleBay) ConvertToProtoMessage() proto.Message {
	r := &pb.ModuleBay{
		Device:          e.GetDevice(),
		Module:          e.GetModule(),
		Name:            e.GetName(),
		InstalledModule: e.GetInstalledModule(),
		Label:           e.GetLabel(),
		Position:        e.GetPosition(),
		Description:     e.GetDescription(),
		Tags:            e.GetTags(),
		CustomFields:    e.GetCustomFields(),
		Metadata:        e.GetMetadata(),
		Owner:           e.GetOwner(),
		Enabled:         e.GetEnabled(),
	}
	return r
}

func (e *ModuleBay) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ModuleBay{
			ModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
		},
	}
}

func (e *ModuleBay) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ModuleBay) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *ModuleBay) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *ModuleBay) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ModuleBay) GetInstalledModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.InstalledModule != nil {
		r = e.InstalledModule.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *ModuleBay) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *ModuleBay) GetPosition() *string {
	var r *string
	if e != nil && e.Position != nil {
		r = e.Position
	}
	return r
}

func (e *ModuleBay) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ModuleBay) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ModuleBay) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ModuleBay) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ModuleBay) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *ModuleBay) GetEnabled() *bool {
	var r *bool
	if e != nil && e.Enabled != nil {
		r = e.Enabled
	}
	return r
}

type ModuleType struct {
	Manufacturer *Manufacturer
	Model        *string
	PartNumber   *string
	Airflow      *string
	Weight       *float64
	WeightUnit   *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Profile      *ModuleTypeProfile
	Attributes   *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *ModuleType) ConvertToProtoMessage() proto.Message {
	r := &pb.ModuleType{
		Manufacturer: e.GetManufacturer(),
		Model:        e.GetModel(),
		PartNumber:   e.GetPartNumber(),
		Airflow:      e.GetAirflow(),
		Weight:       e.GetWeight(),
		WeightUnit:   e.GetWeightUnit(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Profile:      e.GetProfile(),
		Attributes:   e.GetAttributes(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *ModuleType) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ModuleType{
			ModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
		},
	}
}

func (e *ModuleType) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ModuleType) GetManufacturer() *pb.Manufacturer {
	var r *pb.Manufacturer
	if e != nil && e.Manufacturer != nil {
		r = e.Manufacturer.ConvertToProtoMessage().(*pb.Manufacturer)
	}
	return r
}

func (e *ModuleType) GetModel() string {
	var r string
	if e != nil && e.Model != nil {
		r = *e.Model
	}
	return r
}

func (e *ModuleType) GetPartNumber() *string {
	var r *string
	if e != nil && e.PartNumber != nil {
		r = e.PartNumber
	}
	return r
}

func (e *ModuleType) GetAirflow() *string {
	var r *string
	if e != nil && e.Airflow != nil {
		r = e.Airflow
	}
	return r
}

func (e *ModuleType) GetWeight() *float64 {
	var r *float64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *ModuleType) GetWeightUnit() *string {
	var r *string
	if e != nil && e.WeightUnit != nil {
		r = e.WeightUnit
	}
	return r
}

func (e *ModuleType) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ModuleType) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *ModuleType) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ModuleType) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ModuleType) GetProfile() *pb.ModuleTypeProfile {
	var r *pb.ModuleTypeProfile
	if e != nil && e.Profile != nil {
		r = e.Profile.ConvertToProtoMessage().(*pb.ModuleTypeProfile)
	}
	return r
}

func (e *ModuleType) GetAttributes() *string {
	var r *string
	if e != nil && e.Attributes != nil {
		r = e.Attributes
	}
	return r
}

func (e *ModuleType) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ModuleType) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Platform struct {
	Name         *string
	Slug         *string
	Manufacturer *Manufacturer
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Parent       *Platform
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *Platform) ConvertToProtoMessage() proto.Message {
	r := &pb.Platform{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Manufacturer: e.GetManufacturer(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Parent:       e.GetParent(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *Platform) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Platform{
			Platform: e.ConvertToProtoMessage().(*pb.Platform),
		},
	}
}

func (e *Platform) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Platform) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Platform) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Platform) GetManufacturer() *pb.Manufacturer {
	var r *pb.Manufacturer
	if e != nil && e.Manufacturer != nil {
		r = e.Manufacturer.ConvertToProtoMessage().(*pb.Manufacturer)
	}
	return r
}

func (e *Platform) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Platform) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Platform) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Platform) GetParent() *pb.Platform {
	var r *pb.Platform
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.Platform)
	}
	return r
}

func (e *Platform) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Platform) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Platform) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type PowerFeed struct {
	PowerPanel     *PowerPanel
	Rack           *Rack
	Name           *string
	Status         *string
	Type           *string
	Supply         *string
	Phase          *string
	Voltage        *int64
	Amperage       *int64
	MaxUtilization *int64
	MarkConnected  *bool
	Description    *string
	Tenant         *Tenant
	Comments       *string
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
	Owner          *Owner
}

func (e *PowerFeed) ConvertToProtoMessage() proto.Message {
	r := &pb.PowerFeed{
		PowerPanel:     e.GetPowerPanel(),
		Rack:           e.GetRack(),
		Name:           e.GetName(),
		Status:         e.GetStatus(),
		Type:           e.GetType(),
		Supply:         e.GetSupply(),
		Phase:          e.GetPhase(),
		Voltage:        e.GetVoltage(),
		Amperage:       e.GetAmperage(),
		MaxUtilization: e.GetMaxUtilization(),
		MarkConnected:  e.GetMarkConnected(),
		Description:    e.GetDescription(),
		Tenant:         e.GetTenant(),
		Comments:       e.GetComments(),
		Tags:           e.GetTags(),
		CustomFields:   e.GetCustomFields(),
		Metadata:       e.GetMetadata(),
		Owner:          e.GetOwner(),
	}
	return r
}

func (e *PowerFeed) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_PowerFeed{
			PowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
		},
	}
}

func (e *PowerFeed) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *PowerFeed) GetPowerPanel() *pb.PowerPanel {
	var r *pb.PowerPanel
	if e != nil && e.PowerPanel != nil {
		r = e.PowerPanel.ConvertToProtoMessage().(*pb.PowerPanel)
	}
	return r
}

func (e *PowerFeed) GetRack() *pb.Rack {
	var r *pb.Rack
	if e != nil && e.Rack != nil {
		r = e.Rack.ConvertToProtoMessage().(*pb.Rack)
	}
	return r
}

func (e *PowerFeed) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *PowerFeed) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *PowerFeed) GetType() *string {
	var r *string
	if e != nil && e.Type != nil {
		r = e.Type
	}
	return r
}

func (e *PowerFeed) GetSupply() *string {
	var r *string
	if e != nil && e.Supply != nil {
		r = e.Supply
	}
	return r
}

func (e *PowerFeed) GetPhase() *string {
	var r *string
	if e != nil && e.Phase != nil {
		r = e.Phase
	}
	return r
}

func (e *PowerFeed) GetVoltage() *int64 {
	var r *int64
	if e != nil && e.Voltage != nil {
		r = e.Voltage
	}
	return r
}

func (e *PowerFeed) GetAmperage() *int64 {
	var r *int64
	if e != nil && e.Amperage != nil {
		r = e.Amperage
	}
	return r
}

func (e *PowerFeed) GetMaxUtilization() *int64 {
	var r *int64
	if e != nil && e.MaxUtilization != nil {
		r = e.MaxUtilization
	}
	return r
}

func (e *PowerFeed) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *PowerFeed) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *PowerFeed) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *PowerFeed) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *PowerFeed) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *PowerFeed) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *PowerFeed) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *PowerFeed) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type PowerOutlet struct {
	Device        *Device
	Module        *Module
	Name          *string
	Label         *string
	Type          *string
	Color         *string
	PowerPort     *PowerPort
	FeedLeg       *string
	Description   *string
	MarkConnected *bool
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Status        *string
	Metadata      Metadata
	Owner         *Owner
}

func (e *PowerOutlet) ConvertToProtoMessage() proto.Message {
	r := &pb.PowerOutlet{
		Device:        e.GetDevice(),
		Module:        e.GetModule(),
		Name:          e.GetName(),
		Label:         e.GetLabel(),
		Type:          e.GetType(),
		Color:         e.GetColor(),
		PowerPort:     e.GetPowerPort(),
		FeedLeg:       e.GetFeedLeg(),
		Description:   e.GetDescription(),
		MarkConnected: e.GetMarkConnected(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Status:        e.GetStatus(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *PowerOutlet) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_PowerOutlet{
			PowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
		},
	}
}

func (e *PowerOutlet) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *PowerOutlet) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *PowerOutlet) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *PowerOutlet) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *PowerOutlet) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *PowerOutlet) GetType() *string {
	var r *string
	if e != nil && e.Type != nil {
		r = e.Type
	}
	return r
}

func (e *PowerOutlet) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *PowerOutlet) GetPowerPort() *pb.PowerPort {
	var r *pb.PowerPort
	if e != nil && e.PowerPort != nil {
		r = e.PowerPort.ConvertToProtoMessage().(*pb.PowerPort)
	}
	return r
}

func (e *PowerOutlet) GetFeedLeg() *string {
	var r *string
	if e != nil && e.FeedLeg != nil {
		r = e.FeedLeg
	}
	return r
}

func (e *PowerOutlet) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *PowerOutlet) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *PowerOutlet) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *PowerOutlet) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *PowerOutlet) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *PowerOutlet) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *PowerOutlet) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type PowerPanel struct {
	Site         *Site
	Location     *Location
	Name         *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *PowerPanel) ConvertToProtoMessage() proto.Message {
	r := &pb.PowerPanel{
		Site:         e.GetSite(),
		Location:     e.GetLocation(),
		Name:         e.GetName(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *PowerPanel) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_PowerPanel{
			PowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
		},
	}
}

func (e *PowerPanel) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *PowerPanel) GetSite() *pb.Site {
	var r *pb.Site
	if e != nil && e.Site != nil {
		r = e.Site.ConvertToProtoMessage().(*pb.Site)
	}
	return r
}

func (e *PowerPanel) GetLocation() *pb.Location {
	var r *pb.Location
	if e != nil && e.Location != nil {
		r = e.Location.ConvertToProtoMessage().(*pb.Location)
	}
	return r
}

func (e *PowerPanel) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *PowerPanel) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *PowerPanel) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *PowerPanel) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *PowerPanel) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *PowerPanel) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *PowerPanel) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type PowerPort struct {
	Device        *Device
	Module        *Module
	Name          *string
	Label         *string
	Type          *string
	MaximumDraw   *int64
	AllocatedDraw *int64
	Description   *string
	MarkConnected *bool
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Metadata      Metadata
	Owner         *Owner
}

func (e *PowerPort) ConvertToProtoMessage() proto.Message {
	r := &pb.PowerPort{
		Device:        e.GetDevice(),
		Module:        e.GetModule(),
		Name:          e.GetName(),
		Label:         e.GetLabel(),
		Type:          e.GetType(),
		MaximumDraw:   e.GetMaximumDraw(),
		AllocatedDraw: e.GetAllocatedDraw(),
		Description:   e.GetDescription(),
		MarkConnected: e.GetMarkConnected(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *PowerPort) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_PowerPort{
			PowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
		},
	}
}

func (e *PowerPort) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *PowerPort) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *PowerPort) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *PowerPort) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *PowerPort) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *PowerPort) GetType() *string {
	var r *string
	if e != nil && e.Type != nil {
		r = e.Type
	}
	return r
}

func (e *PowerPort) GetMaximumDraw() *int64 {
	var r *int64
	if e != nil && e.MaximumDraw != nil {
		r = e.MaximumDraw
	}
	return r
}

func (e *PowerPort) GetAllocatedDraw() *int64 {
	var r *int64
	if e != nil && e.AllocatedDraw != nil {
		r = e.AllocatedDraw
	}
	return r
}

func (e *PowerPort) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *PowerPort) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *PowerPort) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *PowerPort) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *PowerPort) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *PowerPort) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Prefix struct {
	Prefix *string
	Vrf    *VRF
	// Scope can be:
	//  - Location
	//  - Region
	//  - Site
	//  - SiteGroup
	Scope        anyPrefixScopeValue
	Tenant       *Tenant
	Vlan         *VLAN
	Status       *string
	Role         *Role
	IsPool       *bool
	MarkUtilized *bool
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *Prefix) ConvertToProtoMessage() proto.Message {
	r := &pb.Prefix{
		Prefix:       e.GetPrefix(),
		Vrf:          e.GetVrf(),
		Tenant:       e.GetTenant(),
		Vlan:         e.GetVlan(),
		Status:       e.GetStatus(),
		Role:         e.GetRole(),
		IsPool:       e.GetIsPool(),
		MarkUtilized: e.GetMarkUtilized(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	if e.Scope != nil {
		e.Scope.anyPrefixScopeValueApplyTo(r)
	}
	return r
}

func (e *Prefix) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Prefix{
			Prefix: e.ConvertToProtoMessage().(*pb.Prefix),
		},
	}
}

func (e *Prefix) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Prefix) GetPrefix() string {
	var r string
	if e != nil && e.Prefix != nil {
		r = *e.Prefix
	}
	return r
}

func (e *Prefix) GetVrf() *pb.VRF {
	var r *pb.VRF
	if e != nil && e.Vrf != nil {
		r = e.Vrf.ConvertToProtoMessage().(*pb.VRF)
	}
	return r
}

func (e *Prefix) GetScope() any {
	var r any
	if e != nil && e.Scope != nil {
		var tmp pb.Prefix
		e.Scope.anyPrefixScopeValueApplyTo(&tmp)
		r = tmp.Scope
	}
	return r
}

func (e *Prefix) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Prefix) GetVlan() *pb.VLAN {
	var r *pb.VLAN
	if e != nil && e.Vlan != nil {
		r = e.Vlan.ConvertToProtoMessage().(*pb.VLAN)
	}
	return r
}

func (e *Prefix) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Prefix) GetRole() *pb.Role {
	var r *pb.Role
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.Role)
	}
	return r
}

func (e *Prefix) GetIsPool() *bool {
	var r *bool
	if e != nil && e.IsPool != nil {
		r = e.IsPool
	}
	return r
}

func (e *Prefix) GetMarkUtilized() *bool {
	var r *bool
	if e != nil && e.MarkUtilized != nil {
		r = e.MarkUtilized
	}
	return r
}

func (e *Prefix) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Prefix) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Prefix) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Prefix) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Prefix) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Prefix) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Provider struct {
	Name         *string
	Slug         *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Accounts     []*ProviderAccount
	Asns         []*ASN
	Metadata     Metadata
	Owner        *Owner
}

func (e *Provider) ConvertToProtoMessage() proto.Message {
	r := &pb.Provider{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Accounts:     e.GetAccounts(),
		Asns:         e.GetAsns(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *Provider) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Provider{
			Provider: e.ConvertToProtoMessage().(*pb.Provider),
		},
	}
}

func (e *Provider) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Provider) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Provider) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Provider) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Provider) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Provider) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Provider) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Provider) GetAccounts() []*pb.ProviderAccount {
	var r []*pb.ProviderAccount
	if e != nil && e.Accounts != nil {
		for _, v := range e.Accounts {
			r = append(r, v.ConvertToProtoMessage().(*pb.ProviderAccount))
		}
	}
	return r
}

func (e *Provider) GetAsns() []*pb.ASN {
	var r []*pb.ASN
	if e != nil && e.Asns != nil {
		for _, v := range e.Asns {
			r = append(r, v.ConvertToProtoMessage().(*pb.ASN))
		}
	}
	return r
}

func (e *Provider) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Provider) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type ProviderAccount struct {
	Provider     *Provider
	Name         *string
	Account      *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *ProviderAccount) ConvertToProtoMessage() proto.Message {
	r := &pb.ProviderAccount{
		Provider:     e.GetProvider(),
		Name:         e.GetName(),
		Account:      e.GetAccount(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *ProviderAccount) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ProviderAccount{
			ProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
		},
	}
}

func (e *ProviderAccount) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ProviderAccount) GetProvider() *pb.Provider {
	var r *pb.Provider
	if e != nil && e.Provider != nil {
		r = e.Provider.ConvertToProtoMessage().(*pb.Provider)
	}
	return r
}

func (e *ProviderAccount) GetName() *string {
	var r *string
	if e != nil && e.Name != nil {
		r = e.Name
	}
	return r
}

func (e *ProviderAccount) GetAccount() string {
	var r string
	if e != nil && e.Account != nil {
		r = *e.Account
	}
	return r
}

func (e *ProviderAccount) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ProviderAccount) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *ProviderAccount) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ProviderAccount) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ProviderAccount) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ProviderAccount) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type ProviderNetwork struct {
	Provider     *Provider
	Name         *string
	ServiceId    *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *ProviderNetwork) ConvertToProtoMessage() proto.Message {
	r := &pb.ProviderNetwork{
		Provider:     e.GetProvider(),
		Name:         e.GetName(),
		ServiceId:    e.GetServiceId(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *ProviderNetwork) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ProviderNetwork{
			ProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
		},
	}
}

func (e *ProviderNetwork) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ProviderNetwork) GetProvider() *pb.Provider {
	var r *pb.Provider
	if e != nil && e.Provider != nil {
		r = e.Provider.ConvertToProtoMessage().(*pb.Provider)
	}
	return r
}

func (e *ProviderNetwork) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ProviderNetwork) GetServiceId() *string {
	var r *string
	if e != nil && e.ServiceId != nil {
		r = e.ServiceId
	}
	return r
}

func (e *ProviderNetwork) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ProviderNetwork) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *ProviderNetwork) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ProviderNetwork) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ProviderNetwork) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ProviderNetwork) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type RIR struct {
	Name         *string
	Slug         *string
	IsPrivate    *bool
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *RIR) ConvertToProtoMessage() proto.Message {
	r := &pb.RIR{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		IsPrivate:    e.GetIsPrivate(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *RIR) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Rir{
			Rir: e.ConvertToProtoMessage().(*pb.RIR),
		},
	}
}

func (e *RIR) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *RIR) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *RIR) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *RIR) GetIsPrivate() *bool {
	var r *bool
	if e != nil && e.IsPrivate != nil {
		r = e.IsPrivate
	}
	return r
}

func (e *RIR) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *RIR) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *RIR) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *RIR) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *RIR) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *RIR) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type Rack struct {
	Name          *string
	FacilityId    *string
	Site          *Site
	Location      *Location
	Tenant        *Tenant
	Status        *string
	Role          *RackRole
	Serial        *string
	AssetTag      *string
	RackType      *RackType
	FormFactor    *string
	Width         *int64
	UHeight       *int64
	StartingUnit  *int64
	Weight        *float64
	MaxWeight     *int64
	WeightUnit    *string
	DescUnits     *bool
	OuterWidth    *int64
	OuterDepth    *int64
	OuterUnit     *string
	MountingDepth *int64
	Airflow       *string
	Description   *string
	Comments      *string
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	OuterHeight   *int64
	Metadata      Metadata
	Owner         *Owner
	Group         *RackGroup
}

func (e *Rack) ConvertToProtoMessage() proto.Message {
	r := &pb.Rack{
		Name:          e.GetName(),
		FacilityId:    e.GetFacilityId(),
		Site:          e.GetSite(),
		Location:      e.GetLocation(),
		Tenant:        e.GetTenant(),
		Status:        e.GetStatus(),
		Role:          e.GetRole(),
		Serial:        e.GetSerial(),
		AssetTag:      e.GetAssetTag(),
		RackType:      e.GetRackType(),
		FormFactor:    e.GetFormFactor(),
		Width:         e.GetWidth(),
		UHeight:       e.GetUHeight(),
		StartingUnit:  e.GetStartingUnit(),
		Weight:        e.GetWeight(),
		MaxWeight:     e.GetMaxWeight(),
		WeightUnit:    e.GetWeightUnit(),
		DescUnits:     e.GetDescUnits(),
		OuterWidth:    e.GetOuterWidth(),
		OuterDepth:    e.GetOuterDepth(),
		OuterUnit:     e.GetOuterUnit(),
		MountingDepth: e.GetMountingDepth(),
		Airflow:       e.GetAirflow(),
		Description:   e.GetDescription(),
		Comments:      e.GetComments(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		OuterHeight:   e.GetOuterHeight(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
		Group:         e.GetGroup(),
	}
	return r
}

func (e *Rack) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Rack{
			Rack: e.ConvertToProtoMessage().(*pb.Rack),
		},
	}
}

func (e *Rack) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Rack) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Rack) GetFacilityId() *string {
	var r *string
	if e != nil && e.FacilityId != nil {
		r = e.FacilityId
	}
	return r
}

func (e *Rack) GetSite() *pb.Site {
	var r *pb.Site
	if e != nil && e.Site != nil {
		r = e.Site.ConvertToProtoMessage().(*pb.Site)
	}
	return r
}

func (e *Rack) GetLocation() *pb.Location {
	var r *pb.Location
	if e != nil && e.Location != nil {
		r = e.Location.ConvertToProtoMessage().(*pb.Location)
	}
	return r
}

func (e *Rack) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Rack) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Rack) GetRole() *pb.RackRole {
	var r *pb.RackRole
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.RackRole)
	}
	return r
}

func (e *Rack) GetSerial() *string {
	var r *string
	if e != nil && e.Serial != nil {
		r = e.Serial
	}
	return r
}

func (e *Rack) GetAssetTag() *string {
	var r *string
	if e != nil && e.AssetTag != nil {
		r = e.AssetTag
	}
	return r
}

func (e *Rack) GetRackType() *pb.RackType {
	var r *pb.RackType
	if e != nil && e.RackType != nil {
		r = e.RackType.ConvertToProtoMessage().(*pb.RackType)
	}
	return r
}

func (e *Rack) GetFormFactor() *string {
	var r *string
	if e != nil && e.FormFactor != nil {
		r = e.FormFactor
	}
	return r
}

func (e *Rack) GetWidth() *int64 {
	var r *int64
	if e != nil && e.Width != nil {
		r = e.Width
	}
	return r
}

func (e *Rack) GetUHeight() *int64 {
	var r *int64
	if e != nil && e.UHeight != nil {
		r = e.UHeight
	}
	return r
}

func (e *Rack) GetStartingUnit() *int64 {
	var r *int64
	if e != nil && e.StartingUnit != nil {
		r = e.StartingUnit
	}
	return r
}

func (e *Rack) GetWeight() *float64 {
	var r *float64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *Rack) GetMaxWeight() *int64 {
	var r *int64
	if e != nil && e.MaxWeight != nil {
		r = e.MaxWeight
	}
	return r
}

func (e *Rack) GetWeightUnit() *string {
	var r *string
	if e != nil && e.WeightUnit != nil {
		r = e.WeightUnit
	}
	return r
}

func (e *Rack) GetDescUnits() *bool {
	var r *bool
	if e != nil && e.DescUnits != nil {
		r = e.DescUnits
	}
	return r
}

func (e *Rack) GetOuterWidth() *int64 {
	var r *int64
	if e != nil && e.OuterWidth != nil {
		r = e.OuterWidth
	}
	return r
}

func (e *Rack) GetOuterDepth() *int64 {
	var r *int64
	if e != nil && e.OuterDepth != nil {
		r = e.OuterDepth
	}
	return r
}

func (e *Rack) GetOuterUnit() *string {
	var r *string
	if e != nil && e.OuterUnit != nil {
		r = e.OuterUnit
	}
	return r
}

func (e *Rack) GetMountingDepth() *int64 {
	var r *int64
	if e != nil && e.MountingDepth != nil {
		r = e.MountingDepth
	}
	return r
}

func (e *Rack) GetAirflow() *string {
	var r *string
	if e != nil && e.Airflow != nil {
		r = e.Airflow
	}
	return r
}

func (e *Rack) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Rack) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Rack) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Rack) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Rack) GetOuterHeight() *int64 {
	var r *int64
	if e != nil && e.OuterHeight != nil {
		r = e.OuterHeight
	}
	return r
}

func (e *Rack) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Rack) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *Rack) GetGroup() *pb.RackGroup {
	var r *pb.RackGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.RackGroup)
	}
	return r
}

type RackReservation struct {
	Rack         *Rack
	Units        []int64
	Tenant       *Tenant
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Status       *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *RackReservation) ConvertToProtoMessage() proto.Message {
	r := &pb.RackReservation{
		Rack:         e.GetRack(),
		Units:        e.GetUnits(),
		Tenant:       e.GetTenant(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Status:       e.GetStatus(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *RackReservation) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_RackReservation{
			RackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
		},
	}
}

func (e *RackReservation) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *RackReservation) GetRack() *pb.Rack {
	var r *pb.Rack
	if e != nil && e.Rack != nil {
		r = e.Rack.ConvertToProtoMessage().(*pb.Rack)
	}
	return r
}

func (e *RackReservation) GetUnits() []int64 {
	var r []int64
	if e != nil && e.Units != nil {
		for _, v := range e.Units {
			r = append(r, v)
		}
	}
	return r
}

func (e *RackReservation) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *RackReservation) GetDescription() string {
	var r string
	if e != nil && e.Description != nil {
		r = *e.Description
	}
	return r
}

func (e *RackReservation) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *RackReservation) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *RackReservation) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *RackReservation) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *RackReservation) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *RackReservation) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type RackRole struct {
	Name         *string
	Slug         *string
	Color        *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *RackRole) ConvertToProtoMessage() proto.Message {
	r := &pb.RackRole{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Color:        e.GetColor(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *RackRole) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_RackRole{
			RackRole: e.ConvertToProtoMessage().(*pb.RackRole),
		},
	}
}

func (e *RackRole) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *RackRole) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *RackRole) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *RackRole) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *RackRole) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *RackRole) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *RackRole) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *RackRole) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *RackRole) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *RackRole) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type RackType struct {
	Manufacturer  *Manufacturer
	Model         *string
	Slug          *string
	Description   *string
	FormFactor    *string
	Width         *int64
	UHeight       *int64
	StartingUnit  *int64
	DescUnits     *bool
	OuterWidth    *int64
	OuterDepth    *int64
	OuterUnit     *string
	Weight        *float64
	MaxWeight     *int64
	WeightUnit    *string
	MountingDepth *int64
	Comments      *string
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	OuterHeight   *int64
	Metadata      Metadata
	Owner         *Owner
}

func (e *RackType) ConvertToProtoMessage() proto.Message {
	r := &pb.RackType{
		Manufacturer:  e.GetManufacturer(),
		Model:         e.GetModel(),
		Slug:          e.GetSlug(),
		Description:   e.GetDescription(),
		FormFactor:    e.GetFormFactor(),
		Width:         e.GetWidth(),
		UHeight:       e.GetUHeight(),
		StartingUnit:  e.GetStartingUnit(),
		DescUnits:     e.GetDescUnits(),
		OuterWidth:    e.GetOuterWidth(),
		OuterDepth:    e.GetOuterDepth(),
		OuterUnit:     e.GetOuterUnit(),
		Weight:        e.GetWeight(),
		MaxWeight:     e.GetMaxWeight(),
		WeightUnit:    e.GetWeightUnit(),
		MountingDepth: e.GetMountingDepth(),
		Comments:      e.GetComments(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		OuterHeight:   e.GetOuterHeight(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *RackType) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_RackType{
			RackType: e.ConvertToProtoMessage().(*pb.RackType),
		},
	}
}

func (e *RackType) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *RackType) GetManufacturer() *pb.Manufacturer {
	var r *pb.Manufacturer
	if e != nil && e.Manufacturer != nil {
		r = e.Manufacturer.ConvertToProtoMessage().(*pb.Manufacturer)
	}
	return r
}

func (e *RackType) GetModel() string {
	var r string
	if e != nil && e.Model != nil {
		r = *e.Model
	}
	return r
}

func (e *RackType) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *RackType) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *RackType) GetFormFactor() *string {
	var r *string
	if e != nil && e.FormFactor != nil {
		r = e.FormFactor
	}
	return r
}

func (e *RackType) GetWidth() *int64 {
	var r *int64
	if e != nil && e.Width != nil {
		r = e.Width
	}
	return r
}

func (e *RackType) GetUHeight() *int64 {
	var r *int64
	if e != nil && e.UHeight != nil {
		r = e.UHeight
	}
	return r
}

func (e *RackType) GetStartingUnit() *int64 {
	var r *int64
	if e != nil && e.StartingUnit != nil {
		r = e.StartingUnit
	}
	return r
}

func (e *RackType) GetDescUnits() *bool {
	var r *bool
	if e != nil && e.DescUnits != nil {
		r = e.DescUnits
	}
	return r
}

func (e *RackType) GetOuterWidth() *int64 {
	var r *int64
	if e != nil && e.OuterWidth != nil {
		r = e.OuterWidth
	}
	return r
}

func (e *RackType) GetOuterDepth() *int64 {
	var r *int64
	if e != nil && e.OuterDepth != nil {
		r = e.OuterDepth
	}
	return r
}

func (e *RackType) GetOuterUnit() *string {
	var r *string
	if e != nil && e.OuterUnit != nil {
		r = e.OuterUnit
	}
	return r
}

func (e *RackType) GetWeight() *float64 {
	var r *float64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *RackType) GetMaxWeight() *int64 {
	var r *int64
	if e != nil && e.MaxWeight != nil {
		r = e.MaxWeight
	}
	return r
}

func (e *RackType) GetWeightUnit() *string {
	var r *string
	if e != nil && e.WeightUnit != nil {
		r = e.WeightUnit
	}
	return r
}

func (e *RackType) GetMountingDepth() *int64 {
	var r *int64
	if e != nil && e.MountingDepth != nil {
		r = e.MountingDepth
	}
	return r
}

func (e *RackType) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *RackType) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *RackType) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *RackType) GetOuterHeight() *int64 {
	var r *int64
	if e != nil && e.OuterHeight != nil {
		r = e.OuterHeight
	}
	return r
}

func (e *RackType) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *RackType) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type RearPort struct {
	Device        *Device
	Module        *Module
	Name          *string
	Label         *string
	Type          *string
	Color         *string
	Positions     *int64
	Description   *string
	MarkConnected *bool
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Metadata      Metadata
	Owner         *Owner
}

func (e *RearPort) ConvertToProtoMessage() proto.Message {
	r := &pb.RearPort{
		Device:        e.GetDevice(),
		Module:        e.GetModule(),
		Name:          e.GetName(),
		Label:         e.GetLabel(),
		Type:          e.GetType(),
		Color:         e.GetColor(),
		Positions:     e.GetPositions(),
		Description:   e.GetDescription(),
		MarkConnected: e.GetMarkConnected(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *RearPort) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_RearPort{
			RearPort: e.ConvertToProtoMessage().(*pb.RearPort),
		},
	}
}

func (e *RearPort) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *RearPort) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *RearPort) GetModule() *pb.Module {
	var r *pb.Module
	if e != nil && e.Module != nil {
		r = e.Module.ConvertToProtoMessage().(*pb.Module)
	}
	return r
}

func (e *RearPort) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *RearPort) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *RearPort) GetType() string {
	var r string
	if e != nil && e.Type != nil {
		r = *e.Type
	}
	return r
}

func (e *RearPort) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *RearPort) GetPositions() *int64 {
	var r *int64
	if e != nil && e.Positions != nil {
		r = e.Positions
	}
	return r
}

func (e *RearPort) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *RearPort) GetMarkConnected() *bool {
	var r *bool
	if e != nil && e.MarkConnected != nil {
		r = e.MarkConnected
	}
	return r
}

func (e *RearPort) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *RearPort) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *RearPort) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *RearPort) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Region struct {
	Name         *string
	Slug         *string
	Parent       *Region
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *Region) ConvertToProtoMessage() proto.Message {
	r := &pb.Region{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Parent:       e.GetParent(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *Region) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Region{
			Region: e.ConvertToProtoMessage().(*pb.Region),
		},
	}
}

func (e *Region) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Region) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Region) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Region) GetParent() *pb.Region {
	var r *pb.Region
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.Region)
	}
	return r
}

func (e *Region) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Region) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Region) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Region) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Region) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Region) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Role struct {
	Name         *string
	Slug         *string
	Weight       *int64
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *Role) ConvertToProtoMessage() proto.Message {
	r := &pb.Role{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Weight:       e.GetWeight(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *Role) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Role{
			Role: e.ConvertToProtoMessage().(*pb.Role),
		},
	}
}

func (e *Role) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Role) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Role) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Role) GetWeight() *int64 {
	var r *int64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *Role) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Role) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Role) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Role) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Role) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *Role) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type RouteTarget struct {
	Name         *string
	Tenant       *Tenant
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *RouteTarget) ConvertToProtoMessage() proto.Message {
	r := &pb.RouteTarget{
		Name:         e.GetName(),
		Tenant:       e.GetTenant(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *RouteTarget) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_RouteTarget{
			RouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
		},
	}
}

func (e *RouteTarget) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *RouteTarget) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *RouteTarget) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *RouteTarget) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *RouteTarget) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *RouteTarget) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *RouteTarget) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *RouteTarget) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *RouteTarget) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Service struct {
	Device         *Device
	VirtualMachine *VirtualMachine
	Name           *string
	Protocol       *string
	Ports          []int64
	Description    *string
	Comments       *string
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Ipaddresses    []*IPAddress
	// ParentObject can be:
	//  - Device
	//  - FHRPGroup
	//  - VirtualMachine
	ParentObject anyServiceParentObjectValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *Service) ConvertToProtoMessage() proto.Message {
	r := &pb.Service{
		Device:         e.GetDevice(),
		VirtualMachine: e.GetVirtualMachine(),
		Name:           e.GetName(),
		Protocol:       e.GetProtocol(),
		Ports:          e.GetPorts(),
		Description:    e.GetDescription(),
		Comments:       e.GetComments(),
		Tags:           e.GetTags(),
		CustomFields:   e.GetCustomFields(),
		Ipaddresses:    e.GetIpaddresses(),
		Metadata:       e.GetMetadata(),
		Owner:          e.GetOwner(),
	}
	if e.ParentObject != nil {
		e.ParentObject.anyServiceParentObjectValueApplyTo(r)
	}
	return r
}

func (e *Service) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Service{
			Service: e.ConvertToProtoMessage().(*pb.Service),
		},
	}
}

func (e *Service) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Service) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *Service) GetVirtualMachine() *pb.VirtualMachine {
	var r *pb.VirtualMachine
	if e != nil && e.VirtualMachine != nil {
		r = e.VirtualMachine.ConvertToProtoMessage().(*pb.VirtualMachine)
	}
	return r
}

func (e *Service) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Service) GetProtocol() *string {
	var r *string
	if e != nil && e.Protocol != nil {
		r = e.Protocol
	}
	return r
}

func (e *Service) GetPorts() []int64 {
	var r []int64
	if e != nil && e.Ports != nil {
		for _, v := range e.Ports {
			r = append(r, v)
		}
	}
	return r
}

func (e *Service) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Service) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Service) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Service) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Service) GetIpaddresses() []*pb.IPAddress {
	var r []*pb.IPAddress
	if e != nil && e.Ipaddresses != nil {
		for _, v := range e.Ipaddresses {
			r = append(r, v.ConvertToProtoMessage().(*pb.IPAddress))
		}
	}
	return r
}

func (e *Service) GetParentObject() any {
	var r any
	if e != nil && e.ParentObject != nil {
		var tmp pb.Service
		e.ParentObject.anyServiceParentObjectValueApplyTo(&tmp)
		r = tmp.ParentObject
	}
	return r
}

func (e *Service) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Service) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Site struct {
	Name            *string
	Slug            *string
	Status          *string
	Region          *Region
	Group           *SiteGroup
	Tenant          *Tenant
	Facility        *string
	TimeZone        *string
	Description     *string
	PhysicalAddress *string
	ShippingAddress *string
	Latitude        *float64
	Longitude       *float64
	Comments        *string
	Tags            []*Tag
	CustomFields    map[string]*CustomFieldValue
	Asns            []*ASN
	Metadata        Metadata
	Owner           *Owner
}

func (e *Site) ConvertToProtoMessage() proto.Message {
	r := &pb.Site{
		Name:            e.GetName(),
		Slug:            e.GetSlug(),
		Status:          e.GetStatus(),
		Region:          e.GetRegion(),
		Group:           e.GetGroup(),
		Tenant:          e.GetTenant(),
		Facility:        e.GetFacility(),
		TimeZone:        e.GetTimeZone(),
		Description:     e.GetDescription(),
		PhysicalAddress: e.GetPhysicalAddress(),
		ShippingAddress: e.GetShippingAddress(),
		Latitude:        e.GetLatitude(),
		Longitude:       e.GetLongitude(),
		Comments:        e.GetComments(),
		Tags:            e.GetTags(),
		CustomFields:    e.GetCustomFields(),
		Asns:            e.GetAsns(),
		Metadata:        e.GetMetadata(),
		Owner:           e.GetOwner(),
	}
	return r
}

func (e *Site) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Site{
			Site: e.ConvertToProtoMessage().(*pb.Site),
		},
	}
}

func (e *Site) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Site) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Site) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Site) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *Site) GetRegion() *pb.Region {
	var r *pb.Region
	if e != nil && e.Region != nil {
		r = e.Region.ConvertToProtoMessage().(*pb.Region)
	}
	return r
}

func (e *Site) GetGroup() *pb.SiteGroup {
	var r *pb.SiteGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.SiteGroup)
	}
	return r
}

func (e *Site) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Site) GetFacility() *string {
	var r *string
	if e != nil && e.Facility != nil {
		r = e.Facility
	}
	return r
}

func (e *Site) GetTimeZone() *string {
	var r *string
	if e != nil && e.TimeZone != nil {
		r = e.TimeZone
	}
	return r
}

func (e *Site) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Site) GetPhysicalAddress() *string {
	var r *string
	if e != nil && e.PhysicalAddress != nil {
		r = e.PhysicalAddress
	}
	return r
}

func (e *Site) GetShippingAddress() *string {
	var r *string
	if e != nil && e.ShippingAddress != nil {
		r = e.ShippingAddress
	}
	return r
}

func (e *Site) GetLatitude() *float64 {
	var r *float64
	if e != nil && e.Latitude != nil {
		r = e.Latitude
	}
	return r
}

func (e *Site) GetLongitude() *float64 {
	var r *float64
	if e != nil && e.Longitude != nil {
		r = e.Longitude
	}
	return r
}

func (e *Site) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Site) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Site) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Site) GetAsns() []*pb.ASN {
	var r []*pb.ASN
	if e != nil && e.Asns != nil {
		for _, v := range e.Asns {
			r = append(r, v.ConvertToProtoMessage().(*pb.ASN))
		}
	}
	return r
}

func (e *Site) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Site) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type SiteGroup struct {
	Name         *string
	Slug         *string
	Parent       *SiteGroup
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *SiteGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.SiteGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Parent:       e.GetParent(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *SiteGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_SiteGroup{
			SiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
		},
	}
}

func (e *SiteGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *SiteGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *SiteGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *SiteGroup) GetParent() *pb.SiteGroup {
	var r *pb.SiteGroup
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.SiteGroup)
	}
	return r
}

func (e *SiteGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *SiteGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *SiteGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *SiteGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *SiteGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *SiteGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Tag struct {
	Name        *string
	Slug        *string
	Color       *string
	Description *string
	Weight      *int64
	ObjectTypes []string
	Metadata    Metadata
}

func (e *Tag) ConvertToProtoMessage() proto.Message {
	r := &pb.Tag{
		Name:        e.GetName(),
		Slug:        e.GetSlug(),
		Color:       e.GetColor(),
		Description: e.GetDescription(),
		Weight:      e.GetWeight(),
		ObjectTypes: e.GetObjectTypes(),
		Metadata:    e.GetMetadata(),
	}
	return r
}

func (e *Tag) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Tag{
			Tag: e.ConvertToProtoMessage().(*pb.Tag),
		},
	}
}

func (e *Tag) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Tag) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Tag) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *Tag) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Tag) GetWeight() *int64 {
	var r *int64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *Tag) GetObjectTypes() []string {
	var r []string
	if e != nil && e.ObjectTypes != nil {
		for _, v := range e.ObjectTypes {
			r = append(r, v)
		}
	}
	return r
}

func (e *Tag) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type Tenant struct {
	Name         *string
	Slug         *string
	Group        *TenantGroup
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *Tenant) ConvertToProtoMessage() proto.Message {
	r := &pb.Tenant{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Group:        e.GetGroup(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *Tenant) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Tenant{
			Tenant: e.ConvertToProtoMessage().(*pb.Tenant),
		},
	}
}

func (e *Tenant) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Tenant) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Tenant) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *Tenant) GetGroup() *pb.TenantGroup {
	var r *pb.TenantGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.TenantGroup)
	}
	return r
}

func (e *Tenant) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Tenant) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Tenant) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Tenant) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Tenant) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Tenant) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type TenantGroup struct {
	Name         *string
	Slug         *string
	Parent       *TenantGroup
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *TenantGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.TenantGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Parent:       e.GetParent(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *TenantGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_TenantGroup{
			TenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
		},
	}
}

func (e *TenantGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *TenantGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *TenantGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *TenantGroup) GetParent() *pb.TenantGroup {
	var r *pb.TenantGroup
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.TenantGroup)
	}
	return r
}

func (e *TenantGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *TenantGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *TenantGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *TenantGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *TenantGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *TenantGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Tunnel struct {
	Name          *string
	Status        *string
	Group         *TunnelGroup
	Encapsulation *string
	IpsecProfile  *IPSecProfile
	Tenant        *Tenant
	TunnelId      *int64
	Description   *string
	Comments      *string
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	Metadata      Metadata
	Owner         *Owner
}

func (e *Tunnel) ConvertToProtoMessage() proto.Message {
	r := &pb.Tunnel{
		Name:          e.GetName(),
		Status:        e.GetStatus(),
		Group:         e.GetGroup(),
		Encapsulation: e.GetEncapsulation(),
		IpsecProfile:  e.GetIpsecProfile(),
		Tenant:        e.GetTenant(),
		TunnelId:      e.GetTunnelId(),
		Description:   e.GetDescription(),
		Comments:      e.GetComments(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *Tunnel) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Tunnel{
			Tunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
		},
	}
}

func (e *Tunnel) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *Tunnel) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Tunnel) GetStatus() string {
	var r string
	if e != nil && e.Status != nil {
		r = *e.Status
	}
	return r
}

func (e *Tunnel) GetGroup() *pb.TunnelGroup {
	var r *pb.TunnelGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.TunnelGroup)
	}
	return r
}

func (e *Tunnel) GetEncapsulation() string {
	var r string
	if e != nil && e.Encapsulation != nil {
		r = *e.Encapsulation
	}
	return r
}

func (e *Tunnel) GetIpsecProfile() *pb.IPSecProfile {
	var r *pb.IPSecProfile
	if e != nil && e.IpsecProfile != nil {
		r = e.IpsecProfile.ConvertToProtoMessage().(*pb.IPSecProfile)
	}
	return r
}

func (e *Tunnel) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *Tunnel) GetTunnelId() *int64 {
	var r *int64
	if e != nil && e.TunnelId != nil {
		r = e.TunnelId
	}
	return r
}

func (e *Tunnel) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Tunnel) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *Tunnel) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *Tunnel) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *Tunnel) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *Tunnel) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type TunnelGroup struct {
	Name         *string
	Slug         *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *TunnelGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.TunnelGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *TunnelGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_TunnelGroup{
			TunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
		},
	}
}

func (e *TunnelGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *TunnelGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *TunnelGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *TunnelGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *TunnelGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *TunnelGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *TunnelGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *TunnelGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *TunnelGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type TunnelTermination struct {
	Tunnel *Tunnel
	Role   *string
	// Termination can be:
	//  - ASN
	//  - ASNRange
	//  - Aggregate
	//  - Cable
	//  - CablePath
	//  - CableTermination
	//  - Circuit
	//  - CircuitGroup
	//  - CircuitGroupAssignment
	//  - CircuitTermination
	//  - CircuitType
	//  - Cluster
	//  - ClusterGroup
	//  - ClusterType
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - Contact
	//  - ContactAssignment
	//  - ContactGroup
	//  - ContactRole
	//  - Device
	//  - DeviceBay
	//  - DeviceRole
	//  - DeviceType
	//  - FHRPGroup
	//  - FHRPGroupAssignment
	//  - FrontPort
	//  - IKEPolicy
	//  - IKEProposal
	//  - IPAddress
	//  - IPRange
	//  - IPSecPolicy
	//  - IPSecProfile
	//  - IPSecProposal
	//  - Interface
	//  - InventoryItem
	//  - InventoryItemRole
	//  - L2VPN
	//  - L2VPNTermination
	//  - Location
	//  - MACAddress
	//  - Manufacturer
	//  - Module
	//  - ModuleBay
	//  - ModuleType
	//  - Platform
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPanel
	//  - PowerPort
	//  - Prefix
	//  - Provider
	//  - ProviderAccount
	//  - ProviderNetwork
	//  - RIR
	//  - Rack
	//  - RackReservation
	//  - RackRole
	//  - RackType
	//  - RearPort
	//  - Region
	//  - Role
	//  - RouteTarget
	//  - Service
	//  - Site
	//  - SiteGroup
	//  - Tag
	//  - Tenant
	//  - TenantGroup
	//  - Tunnel
	//  - TunnelGroup
	//  - TunnelTermination
	//  - VLAN
	//  - VLANGroup
	//  - VLANTranslationPolicy
	//  - VLANTranslationRule
	//  - VMInterface
	//  - VRF
	//  - VirtualChassis
	//  - VirtualCircuit
	//  - VirtualCircuitTermination
	//  - VirtualCircuitType
	//  - VirtualDeviceContext
	//  - VirtualDisk
	//  - VirtualMachine
	//  - WirelessLAN
	//  - WirelessLANGroup
	//  - WirelessLink
	//  - CustomField
	//  - CustomFieldChoiceSet
	//  - JournalEntry
	//  - ModuleTypeProfile
	//  - CustomLink
	//  - Owner
	//  - OwnerGroup
	//  - CableBundle
	//  - RackGroup
	//  - ScriptModule
	//  - VirtualMachineType
	Termination  anyTunnelTerminationTerminationValue
	OutsideIp    *IPAddress
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
}

func (e *TunnelTermination) ConvertToProtoMessage() proto.Message {
	r := &pb.TunnelTermination{
		Tunnel:       e.GetTunnel(),
		Role:         e.GetRole(),
		OutsideIp:    e.GetOutsideIp(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
	}
	if e.Termination != nil {
		e.Termination.anyTunnelTerminationTerminationValueApplyTo(r)
	}
	return r
}

func (e *TunnelTermination) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_TunnelTermination{
			TunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
		},
	}
}

func (e *TunnelTermination) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *TunnelTermination) GetTunnel() *pb.Tunnel {
	var r *pb.Tunnel
	if e != nil && e.Tunnel != nil {
		r = e.Tunnel.ConvertToProtoMessage().(*pb.Tunnel)
	}
	return r
}

func (e *TunnelTermination) GetRole() string {
	var r string
	if e != nil && e.Role != nil {
		r = *e.Role
	}
	return r
}

func (e *TunnelTermination) GetTermination() any {
	var r any
	if e != nil && e.Termination != nil {
		var tmp pb.TunnelTermination
		e.Termination.anyTunnelTerminationTerminationValueApplyTo(&tmp)
		r = tmp.Termination
	}
	return r
}

func (e *TunnelTermination) GetOutsideIp() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.OutsideIp != nil {
		r = e.OutsideIp.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *TunnelTermination) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *TunnelTermination) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *TunnelTermination) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type VLAN struct {
	Site         *Site
	Group        *VLANGroup
	Vid          *int64
	Name         *string
	Tenant       *Tenant
	Status       *string
	Role         *Role
	Description  *string
	QinqRole     *string
	QinqSvlan    *VLAN
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *VLAN) ConvertToProtoMessage() proto.Message {
	r := &pb.VLAN{
		Site:         e.GetSite(),
		Group:        e.GetGroup(),
		Vid:          e.GetVid(),
		Name:         e.GetName(),
		Tenant:       e.GetTenant(),
		Status:       e.GetStatus(),
		Role:         e.GetRole(),
		Description:  e.GetDescription(),
		QinqRole:     e.GetQinqRole(),
		QinqSvlan:    e.GetQinqSvlan(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *VLAN) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Vlan{
			Vlan: e.ConvertToProtoMessage().(*pb.VLAN),
		},
	}
}

func (e *VLAN) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VLAN) GetSite() *pb.Site {
	var r *pb.Site
	if e != nil && e.Site != nil {
		r = e.Site.ConvertToProtoMessage().(*pb.Site)
	}
	return r
}

func (e *VLAN) GetGroup() *pb.VLANGroup {
	var r *pb.VLANGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.VLANGroup)
	}
	return r
}

func (e *VLAN) GetVid() int64 {
	var r int64
	if e != nil && e.Vid != nil {
		r = *e.Vid
	}
	return r
}

func (e *VLAN) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VLAN) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *VLAN) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *VLAN) GetRole() *pb.Role {
	var r *pb.Role
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.Role)
	}
	return r
}

func (e *VLAN) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VLAN) GetQinqRole() *string {
	var r *string
	if e != nil && e.QinqRole != nil {
		r = e.QinqRole
	}
	return r
}

func (e *VLAN) GetQinqSvlan() *pb.VLAN {
	var r *pb.VLAN
	if e != nil && e.QinqSvlan != nil {
		r = e.QinqSvlan.ConvertToProtoMessage().(*pb.VLAN)
	}
	return r
}

func (e *VLAN) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *VLAN) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VLAN) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VLAN) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VLAN) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type VLANGroup struct {
	Name *string
	Slug *string
	// Scope can be:
	//  - Cluster
	//  - ClusterGroup
	//  - Location
	//  - Rack
	//  - Region
	//  - Site
	//  - SiteGroup
	//  - RackGroup
	Scope        anyVLANGroupScopeValue
	VidRanges    []int64
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Tenant       *Tenant
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *VLANGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.VLANGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		VidRanges:    e.GetVidRanges(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Tenant:       e.GetTenant(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	if e.Scope != nil {
		e.Scope.anyVLANGroupScopeValueApplyTo(r)
	}
	return r
}

func (e *VLANGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VlanGroup{
			VlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
		},
	}
}

func (e *VLANGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VLANGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VLANGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *VLANGroup) GetScope() any {
	var r any
	if e != nil && e.Scope != nil {
		var tmp pb.VLANGroup
		e.Scope.anyVLANGroupScopeValueApplyTo(&tmp)
		r = tmp.Scope
	}
	return r
}

func (e *VLANGroup) GetVidRanges() []int64 {
	var r []int64
	if e != nil && e.VidRanges != nil {
		for _, v := range e.VidRanges {
			r = append(r, v)
		}
	}
	return r
}

func (e *VLANGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VLANGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VLANGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VLANGroup) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *VLANGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VLANGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *VLANGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type VLANTranslationPolicy struct {
	Name        *string
	Description *string
	Metadata    Metadata
	Owner       *Owner
	Comments    *string
}

func (e *VLANTranslationPolicy) ConvertToProtoMessage() proto.Message {
	r := &pb.VLANTranslationPolicy{
		Name:        e.GetName(),
		Description: e.GetDescription(),
		Metadata:    e.GetMetadata(),
		Owner:       e.GetOwner(),
		Comments:    e.GetComments(),
	}
	return r
}

func (e *VLANTranslationPolicy) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VlanTranslationPolicy{
			VlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
		},
	}
}

func (e *VLANTranslationPolicy) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VLANTranslationPolicy) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VLANTranslationPolicy) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VLANTranslationPolicy) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *VLANTranslationPolicy) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type VLANTranslationRule struct {
	Policy      *VLANTranslationPolicy
	LocalVid    *int64
	RemoteVid   *int64
	Description *string
	Metadata    Metadata
}

func (e *VLANTranslationRule) ConvertToProtoMessage() proto.Message {
	r := &pb.VLANTranslationRule{
		Policy:      e.GetPolicy(),
		LocalVid:    e.GetLocalVid(),
		RemoteVid:   e.GetRemoteVid(),
		Description: e.GetDescription(),
		Metadata:    e.GetMetadata(),
	}
	return r
}

func (e *VLANTranslationRule) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VlanTranslationRule{
			VlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
		},
	}
}

func (e *VLANTranslationRule) GetPolicy() *pb.VLANTranslationPolicy {
	var r *pb.VLANTranslationPolicy
	if e != nil && e.Policy != nil {
		r = e.Policy.ConvertToProtoMessage().(*pb.VLANTranslationPolicy)
	}
	return r
}

func (e *VLANTranslationRule) GetLocalVid() int64 {
	var r int64
	if e != nil && e.LocalVid != nil {
		r = *e.LocalVid
	}
	return r
}

func (e *VLANTranslationRule) GetRemoteVid() int64 {
	var r int64
	if e != nil && e.RemoteVid != nil {
		r = *e.RemoteVid
	}
	return r
}

func (e *VLANTranslationRule) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VLANTranslationRule) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type VMInterface struct {
	VirtualMachine        *VirtualMachine
	Name                  *string
	Enabled               *bool
	Parent                *VMInterface
	Bridge                *VMInterface
	Mtu                   *int64
	PrimaryMacAddress     *MACAddress
	Description           *string
	Mode                  *string
	UntaggedVlan          *VLAN
	QinqSvlan             *VLAN
	VlanTranslationPolicy *VLANTranslationPolicy
	Vrf                   *VRF
	Tags                  []*Tag
	CustomFields          map[string]*CustomFieldValue
	TaggedVlans           []*VLAN
	Metadata              Metadata
	Owner                 *Owner
}

func (e *VMInterface) ConvertToProtoMessage() proto.Message {
	r := &pb.VMInterface{
		VirtualMachine:        e.GetVirtualMachine(),
		Name:                  e.GetName(),
		Enabled:               e.GetEnabled(),
		Parent:                e.GetParent(),
		Bridge:                e.GetBridge(),
		Mtu:                   e.GetMtu(),
		PrimaryMacAddress:     e.GetPrimaryMacAddress(),
		Description:           e.GetDescription(),
		Mode:                  e.GetMode(),
		UntaggedVlan:          e.GetUntaggedVlan(),
		QinqSvlan:             e.GetQinqSvlan(),
		VlanTranslationPolicy: e.GetVlanTranslationPolicy(),
		Vrf:                   e.GetVrf(),
		Tags:                  e.GetTags(),
		CustomFields:          e.GetCustomFields(),
		TaggedVlans:           e.GetTaggedVlans(),
		Metadata:              e.GetMetadata(),
		Owner:                 e.GetOwner(),
	}
	return r
}

func (e *VMInterface) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VmInterface{
			VmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
		},
	}
}

func (e *VMInterface) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VMInterface) GetVirtualMachine() *pb.VirtualMachine {
	var r *pb.VirtualMachine
	if e != nil && e.VirtualMachine != nil {
		r = e.VirtualMachine.ConvertToProtoMessage().(*pb.VirtualMachine)
	}
	return r
}

func (e *VMInterface) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VMInterface) GetEnabled() *bool {
	var r *bool
	if e != nil && e.Enabled != nil {
		r = e.Enabled
	}
	return r
}

func (e *VMInterface) GetParent() *pb.VMInterface {
	var r *pb.VMInterface
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.VMInterface)
	}
	return r
}

func (e *VMInterface) GetBridge() *pb.VMInterface {
	var r *pb.VMInterface
	if e != nil && e.Bridge != nil {
		r = e.Bridge.ConvertToProtoMessage().(*pb.VMInterface)
	}
	return r
}

func (e *VMInterface) GetMtu() *int64 {
	var r *int64
	if e != nil && e.Mtu != nil {
		r = e.Mtu
	}
	return r
}

func (e *VMInterface) GetPrimaryMacAddress() *pb.MACAddress {
	var r *pb.MACAddress
	if e != nil && e.PrimaryMacAddress != nil {
		r = e.PrimaryMacAddress.ConvertToProtoMessage().(*pb.MACAddress)
	}
	return r
}

func (e *VMInterface) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VMInterface) GetMode() *string {
	var r *string
	if e != nil && e.Mode != nil {
		r = e.Mode
	}
	return r
}

func (e *VMInterface) GetUntaggedVlan() *pb.VLAN {
	var r *pb.VLAN
	if e != nil && e.UntaggedVlan != nil {
		r = e.UntaggedVlan.ConvertToProtoMessage().(*pb.VLAN)
	}
	return r
}

func (e *VMInterface) GetQinqSvlan() *pb.VLAN {
	var r *pb.VLAN
	if e != nil && e.QinqSvlan != nil {
		r = e.QinqSvlan.ConvertToProtoMessage().(*pb.VLAN)
	}
	return r
}

func (e *VMInterface) GetVlanTranslationPolicy() *pb.VLANTranslationPolicy {
	var r *pb.VLANTranslationPolicy
	if e != nil && e.VlanTranslationPolicy != nil {
		r = e.VlanTranslationPolicy.ConvertToProtoMessage().(*pb.VLANTranslationPolicy)
	}
	return r
}

func (e *VMInterface) GetVrf() *pb.VRF {
	var r *pb.VRF
	if e != nil && e.Vrf != nil {
		r = e.Vrf.ConvertToProtoMessage().(*pb.VRF)
	}
	return r
}

func (e *VMInterface) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VMInterface) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VMInterface) GetTaggedVlans() []*pb.VLAN {
	var r []*pb.VLAN
	if e != nil && e.TaggedVlans != nil {
		for _, v := range e.TaggedVlans {
			r = append(r, v.ConvertToProtoMessage().(*pb.VLAN))
		}
	}
	return r
}

func (e *VMInterface) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VMInterface) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type VRF struct {
	Name          *string
	Rd            *string
	Tenant        *Tenant
	EnforceUnique *bool
	Description   *string
	Comments      *string
	Tags          []*Tag
	CustomFields  map[string]*CustomFieldValue
	ImportTargets []*RouteTarget
	ExportTargets []*RouteTarget
	Metadata      Metadata
	Owner         *Owner
}

func (e *VRF) ConvertToProtoMessage() proto.Message {
	r := &pb.VRF{
		Name:          e.GetName(),
		Rd:            e.GetRd(),
		Tenant:        e.GetTenant(),
		EnforceUnique: e.GetEnforceUnique(),
		Description:   e.GetDescription(),
		Comments:      e.GetComments(),
		Tags:          e.GetTags(),
		CustomFields:  e.GetCustomFields(),
		ImportTargets: e.GetImportTargets(),
		ExportTargets: e.GetExportTargets(),
		Metadata:      e.GetMetadata(),
		Owner:         e.GetOwner(),
	}
	return r
}

func (e *VRF) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Vrf{
			Vrf: e.ConvertToProtoMessage().(*pb.VRF),
		},
	}
}

func (e *VRF) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VRF) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VRF) GetRd() *string {
	var r *string
	if e != nil && e.Rd != nil {
		r = e.Rd
	}
	return r
}

func (e *VRF) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *VRF) GetEnforceUnique() *bool {
	var r *bool
	if e != nil && e.EnforceUnique != nil {
		r = e.EnforceUnique
	}
	return r
}

func (e *VRF) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VRF) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *VRF) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VRF) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VRF) GetImportTargets() []*pb.RouteTarget {
	var r []*pb.RouteTarget
	if e != nil && e.ImportTargets != nil {
		for _, v := range e.ImportTargets {
			r = append(r, v.ConvertToProtoMessage().(*pb.RouteTarget))
		}
	}
	return r
}

func (e *VRF) GetExportTargets() []*pb.RouteTarget {
	var r []*pb.RouteTarget
	if e != nil && e.ExportTargets != nil {
		for _, v := range e.ExportTargets {
			r = append(r, v.ConvertToProtoMessage().(*pb.RouteTarget))
		}
	}
	return r
}

func (e *VRF) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VRF) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type VirtualChassis struct {
	Name         *string
	Domain       *string
	Master       *Device
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *VirtualChassis) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualChassis{
		Name:         e.GetName(),
		Domain:       e.GetDomain(),
		Master:       e.GetMaster(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *VirtualChassis) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualChassis{
			VirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
		},
	}
}

func (e *VirtualChassis) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualChassis) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VirtualChassis) GetDomain() *string {
	var r *string
	if e != nil && e.Domain != nil {
		r = e.Domain
	}
	return r
}

func (e *VirtualChassis) GetMaster() *pb.Device {
	var r *pb.Device
	if e != nil && e.Master != nil {
		r = e.Master.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *VirtualChassis) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualChassis) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *VirtualChassis) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualChassis) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualChassis) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VirtualChassis) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type VirtualCircuit struct {
	Cid             *string
	ProviderNetwork *ProviderNetwork
	ProviderAccount *ProviderAccount
	Type            *VirtualCircuitType
	Status          *string
	Tenant          *Tenant
	Description     *string
	Comments        *string
	Tags            []*Tag
	CustomFields    map[string]*CustomFieldValue
	Metadata        Metadata
	Owner           *Owner
}

func (e *VirtualCircuit) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualCircuit{
		Cid:             e.GetCid(),
		ProviderNetwork: e.GetProviderNetwork(),
		ProviderAccount: e.GetProviderAccount(),
		Type:            e.GetType(),
		Status:          e.GetStatus(),
		Tenant:          e.GetTenant(),
		Description:     e.GetDescription(),
		Comments:        e.GetComments(),
		Tags:            e.GetTags(),
		CustomFields:    e.GetCustomFields(),
		Metadata:        e.GetMetadata(),
		Owner:           e.GetOwner(),
	}
	return r
}

func (e *VirtualCircuit) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualCircuit{
			VirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
		},
	}
}

func (e *VirtualCircuit) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualCircuit) GetCid() string {
	var r string
	if e != nil && e.Cid != nil {
		r = *e.Cid
	}
	return r
}

func (e *VirtualCircuit) GetProviderNetwork() *pb.ProviderNetwork {
	var r *pb.ProviderNetwork
	if e != nil && e.ProviderNetwork != nil {
		r = e.ProviderNetwork.ConvertToProtoMessage().(*pb.ProviderNetwork)
	}
	return r
}

func (e *VirtualCircuit) GetProviderAccount() *pb.ProviderAccount {
	var r *pb.ProviderAccount
	if e != nil && e.ProviderAccount != nil {
		r = e.ProviderAccount.ConvertToProtoMessage().(*pb.ProviderAccount)
	}
	return r
}

func (e *VirtualCircuit) GetType() *pb.VirtualCircuitType {
	var r *pb.VirtualCircuitType
	if e != nil && e.Type != nil {
		r = e.Type.ConvertToProtoMessage().(*pb.VirtualCircuitType)
	}
	return r
}

func (e *VirtualCircuit) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *VirtualCircuit) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *VirtualCircuit) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualCircuit) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *VirtualCircuit) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualCircuit) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualCircuit) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VirtualCircuit) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type VirtualCircuitTermination struct {
	VirtualCircuit *VirtualCircuit
	Role           *string
	Interface      *Interface
	Description    *string
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
}

func (e *VirtualCircuitTermination) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualCircuitTermination{
		VirtualCircuit: e.GetVirtualCircuit(),
		Role:           e.GetRole(),
		Interface:      e.GetInterface(),
		Description:    e.GetDescription(),
		Tags:           e.GetTags(),
		CustomFields:   e.GetCustomFields(),
		Metadata:       e.GetMetadata(),
	}
	return r
}

func (e *VirtualCircuitTermination) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualCircuitTermination{
			VirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
		},
	}
}

func (e *VirtualCircuitTermination) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualCircuitTermination) GetVirtualCircuit() *pb.VirtualCircuit {
	var r *pb.VirtualCircuit
	if e != nil && e.VirtualCircuit != nil {
		r = e.VirtualCircuit.ConvertToProtoMessage().(*pb.VirtualCircuit)
	}
	return r
}

func (e *VirtualCircuitTermination) GetRole() *string {
	var r *string
	if e != nil && e.Role != nil {
		r = e.Role
	}
	return r
}

func (e *VirtualCircuitTermination) GetInterface() *pb.Interface {
	var r *pb.Interface
	if e != nil && e.Interface != nil {
		r = e.Interface.ConvertToProtoMessage().(*pb.Interface)
	}
	return r
}

func (e *VirtualCircuitTermination) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualCircuitTermination) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualCircuitTermination) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualCircuitTermination) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type VirtualCircuitType struct {
	Name         *string
	Slug         *string
	Color        *string
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
	Comments     *string
}

func (e *VirtualCircuitType) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualCircuitType{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Color:        e.GetColor(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
	}
	return r
}

func (e *VirtualCircuitType) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualCircuitType{
			VirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
		},
	}
}

func (e *VirtualCircuitType) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualCircuitType) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VirtualCircuitType) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *VirtualCircuitType) GetColor() *string {
	var r *string
	if e != nil && e.Color != nil {
		r = e.Color
	}
	return r
}

func (e *VirtualCircuitType) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualCircuitType) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualCircuitType) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualCircuitType) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VirtualCircuitType) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *VirtualCircuitType) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

type VirtualDeviceContext struct {
	Name         *string
	Device       *Device
	Identifier   *int64
	Tenant       *Tenant
	PrimaryIp4   *IPAddress
	PrimaryIp6   *IPAddress
	Status       *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *VirtualDeviceContext) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualDeviceContext{
		Name:         e.GetName(),
		Device:       e.GetDevice(),
		Identifier:   e.GetIdentifier(),
		Tenant:       e.GetTenant(),
		PrimaryIp4:   e.GetPrimaryIp4(),
		PrimaryIp6:   e.GetPrimaryIp6(),
		Status:       e.GetStatus(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *VirtualDeviceContext) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualDeviceContext{
			VirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
		},
	}
}

func (e *VirtualDeviceContext) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualDeviceContext) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VirtualDeviceContext) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *VirtualDeviceContext) GetIdentifier() *int64 {
	var r *int64
	if e != nil && e.Identifier != nil {
		r = e.Identifier
	}
	return r
}

func (e *VirtualDeviceContext) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *VirtualDeviceContext) GetPrimaryIp4() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.PrimaryIp4 != nil {
		r = e.PrimaryIp4.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *VirtualDeviceContext) GetPrimaryIp6() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.PrimaryIp6 != nil {
		r = e.PrimaryIp6.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *VirtualDeviceContext) GetStatus() string {
	var r string
	if e != nil && e.Status != nil {
		r = *e.Status
	}
	return r
}

func (e *VirtualDeviceContext) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualDeviceContext) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *VirtualDeviceContext) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualDeviceContext) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualDeviceContext) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VirtualDeviceContext) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type VirtualDisk struct {
	VirtualMachine *VirtualMachine
	Name           *string
	Description    *string
	Size           *int64
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
	Owner          *Owner
}

func (e *VirtualDisk) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualDisk{
		VirtualMachine: e.GetVirtualMachine(),
		Name:           e.GetName(),
		Description:    e.GetDescription(),
		Size:           e.GetSize(),
		Tags:           e.GetTags(),
		CustomFields:   e.GetCustomFields(),
		Metadata:       e.GetMetadata(),
		Owner:          e.GetOwner(),
	}
	return r
}

func (e *VirtualDisk) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualDisk{
			VirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
		},
	}
}

func (e *VirtualDisk) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualDisk) GetVirtualMachine() *pb.VirtualMachine {
	var r *pb.VirtualMachine
	if e != nil && e.VirtualMachine != nil {
		r = e.VirtualMachine.ConvertToProtoMessage().(*pb.VirtualMachine)
	}
	return r
}

func (e *VirtualDisk) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VirtualDisk) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualDisk) GetSize() int64 {
	var r int64
	if e != nil && e.Size != nil {
		r = *e.Size
	}
	return r
}

func (e *VirtualDisk) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualDisk) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualDisk) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VirtualDisk) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type VirtualMachine struct {
	Name               *string
	Status             *string
	Site               *Site
	Cluster            *Cluster
	Device             *Device
	Serial             *string
	Role               *DeviceRole
	Tenant             *Tenant
	Platform           *Platform
	PrimaryIp4         *IPAddress
	PrimaryIp6         *IPAddress
	Vcpus              *float64
	Memory             *int64
	Disk               *int64
	Description        *string
	Comments           *string
	Tags               []*Tag
	CustomFields       map[string]*CustomFieldValue
	Metadata           Metadata
	StartOnBoot        *string
	Owner              *Owner
	VirtualMachineType *VirtualMachineType
}

func (e *VirtualMachine) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualMachine{
		Name:               e.GetName(),
		Status:             e.GetStatus(),
		Site:               e.GetSite(),
		Cluster:            e.GetCluster(),
		Device:             e.GetDevice(),
		Serial:             e.GetSerial(),
		Role:               e.GetRole(),
		Tenant:             e.GetTenant(),
		Platform:           e.GetPlatform(),
		PrimaryIp4:         e.GetPrimaryIp4(),
		PrimaryIp6:         e.GetPrimaryIp6(),
		Vcpus:              e.GetVcpus(),
		Memory:             e.GetMemory(),
		Disk:               e.GetDisk(),
		Description:        e.GetDescription(),
		Comments:           e.GetComments(),
		Tags:               e.GetTags(),
		CustomFields:       e.GetCustomFields(),
		Metadata:           e.GetMetadata(),
		StartOnBoot:        e.GetStartOnBoot(),
		Owner:              e.GetOwner(),
		VirtualMachineType: e.GetVirtualMachineType(),
	}
	return r
}

func (e *VirtualMachine) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualMachine{
			VirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
		},
	}
}

func (e *VirtualMachine) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualMachine) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VirtualMachine) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *VirtualMachine) GetSite() *pb.Site {
	var r *pb.Site
	if e != nil && e.Site != nil {
		r = e.Site.ConvertToProtoMessage().(*pb.Site)
	}
	return r
}

func (e *VirtualMachine) GetCluster() *pb.Cluster {
	var r *pb.Cluster
	if e != nil && e.Cluster != nil {
		r = e.Cluster.ConvertToProtoMessage().(*pb.Cluster)
	}
	return r
}

func (e *VirtualMachine) GetDevice() *pb.Device {
	var r *pb.Device
	if e != nil && e.Device != nil {
		r = e.Device.ConvertToProtoMessage().(*pb.Device)
	}
	return r
}

func (e *VirtualMachine) GetSerial() *string {
	var r *string
	if e != nil && e.Serial != nil {
		r = e.Serial
	}
	return r
}

func (e *VirtualMachine) GetRole() *pb.DeviceRole {
	var r *pb.DeviceRole
	if e != nil && e.Role != nil {
		r = e.Role.ConvertToProtoMessage().(*pb.DeviceRole)
	}
	return r
}

func (e *VirtualMachine) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *VirtualMachine) GetPlatform() *pb.Platform {
	var r *pb.Platform
	if e != nil && e.Platform != nil {
		r = e.Platform.ConvertToProtoMessage().(*pb.Platform)
	}
	return r
}

func (e *VirtualMachine) GetPrimaryIp4() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.PrimaryIp4 != nil {
		r = e.PrimaryIp4.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *VirtualMachine) GetPrimaryIp6() *pb.IPAddress {
	var r *pb.IPAddress
	if e != nil && e.PrimaryIp6 != nil {
		r = e.PrimaryIp6.ConvertToProtoMessage().(*pb.IPAddress)
	}
	return r
}

func (e *VirtualMachine) GetVcpus() *float64 {
	var r *float64
	if e != nil && e.Vcpus != nil {
		r = e.Vcpus
	}
	return r
}

func (e *VirtualMachine) GetMemory() *int64 {
	var r *int64
	if e != nil && e.Memory != nil {
		r = e.Memory
	}
	return r
}

func (e *VirtualMachine) GetDisk() *int64 {
	var r *int64
	if e != nil && e.Disk != nil {
		r = e.Disk
	}
	return r
}

func (e *VirtualMachine) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualMachine) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *VirtualMachine) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualMachine) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualMachine) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *VirtualMachine) GetStartOnBoot() *string {
	var r *string
	if e != nil && e.StartOnBoot != nil {
		r = e.StartOnBoot
	}
	return r
}

func (e *VirtualMachine) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *VirtualMachine) GetVirtualMachineType() *pb.VirtualMachineType {
	var r *pb.VirtualMachineType
	if e != nil && e.VirtualMachineType != nil {
		r = e.VirtualMachineType.ConvertToProtoMessage().(*pb.VirtualMachineType)
	}
	return r
}

type WirelessLAN struct {
	Ssid        *string
	Description *string
	Group       *WirelessLANGroup
	Status      *string
	Vlan        *VLAN
	// Scope can be:
	//  - Location
	//  - Region
	//  - Site
	//  - SiteGroup
	Scope        anyWirelessLANScopeValue
	Tenant       *Tenant
	AuthType     *string
	AuthCipher   *string
	AuthPsk      *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *WirelessLAN) ConvertToProtoMessage() proto.Message {
	r := &pb.WirelessLAN{
		Ssid:         e.GetSsid(),
		Description:  e.GetDescription(),
		Group:        e.GetGroup(),
		Status:       e.GetStatus(),
		Vlan:         e.GetVlan(),
		Tenant:       e.GetTenant(),
		AuthType:     e.GetAuthType(),
		AuthCipher:   e.GetAuthCipher(),
		AuthPsk:      e.GetAuthPsk(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	if e.Scope != nil {
		e.Scope.anyWirelessLANScopeValueApplyTo(r)
	}
	return r
}

func (e *WirelessLAN) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_WirelessLan{
			WirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
		},
	}
}

func (e *WirelessLAN) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *WirelessLAN) GetSsid() string {
	var r string
	if e != nil && e.Ssid != nil {
		r = *e.Ssid
	}
	return r
}

func (e *WirelessLAN) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *WirelessLAN) GetGroup() *pb.WirelessLANGroup {
	var r *pb.WirelessLANGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.WirelessLANGroup)
	}
	return r
}

func (e *WirelessLAN) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *WirelessLAN) GetVlan() *pb.VLAN {
	var r *pb.VLAN
	if e != nil && e.Vlan != nil {
		r = e.Vlan.ConvertToProtoMessage().(*pb.VLAN)
	}
	return r
}

func (e *WirelessLAN) GetScope() any {
	var r any
	if e != nil && e.Scope != nil {
		var tmp pb.WirelessLAN
		e.Scope.anyWirelessLANScopeValueApplyTo(&tmp)
		r = tmp.Scope
	}
	return r
}

func (e *WirelessLAN) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *WirelessLAN) GetAuthType() *string {
	var r *string
	if e != nil && e.AuthType != nil {
		r = e.AuthType
	}
	return r
}

func (e *WirelessLAN) GetAuthCipher() *string {
	var r *string
	if e != nil && e.AuthCipher != nil {
		r = e.AuthCipher
	}
	return r
}

func (e *WirelessLAN) GetAuthPsk() *string {
	var r *string
	if e != nil && e.AuthPsk != nil {
		r = e.AuthPsk
	}
	return r
}

func (e *WirelessLAN) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *WirelessLAN) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *WirelessLAN) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *WirelessLAN) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *WirelessLAN) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type WirelessLANGroup struct {
	Name         *string
	Slug         *string
	Parent       *WirelessLANGroup
	Description  *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Comments     *string
	Metadata     Metadata
	Owner        *Owner
}

func (e *WirelessLANGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.WirelessLANGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Parent:       e.GetParent(),
		Description:  e.GetDescription(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Comments:     e.GetComments(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *WirelessLANGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_WirelessLanGroup{
			WirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
		},
	}
}

func (e *WirelessLANGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *WirelessLANGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *WirelessLANGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *WirelessLANGroup) GetParent() *pb.WirelessLANGroup {
	var r *pb.WirelessLANGroup
	if e != nil && e.Parent != nil {
		r = e.Parent.ConvertToProtoMessage().(*pb.WirelessLANGroup)
	}
	return r
}

func (e *WirelessLANGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *WirelessLANGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *WirelessLANGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *WirelessLANGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *WirelessLANGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *WirelessLANGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type WirelessLink struct {
	InterfaceA   *Interface
	InterfaceB   *Interface
	Ssid         *string
	Status       *string
	Tenant       *Tenant
	AuthType     *string
	AuthCipher   *string
	AuthPsk      *string
	Distance     *float64
	DistanceUnit *string
	Description  *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *WirelessLink) ConvertToProtoMessage() proto.Message {
	r := &pb.WirelessLink{
		InterfaceA:   e.GetInterfaceA(),
		InterfaceB:   e.GetInterfaceB(),
		Ssid:         e.GetSsid(),
		Status:       e.GetStatus(),
		Tenant:       e.GetTenant(),
		AuthType:     e.GetAuthType(),
		AuthCipher:   e.GetAuthCipher(),
		AuthPsk:      e.GetAuthPsk(),
		Distance:     e.GetDistance(),
		DistanceUnit: e.GetDistanceUnit(),
		Description:  e.GetDescription(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *WirelessLink) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_WirelessLink{
			WirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
		},
	}
}

func (e *WirelessLink) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *WirelessLink) GetInterfaceA() *pb.Interface {
	var r *pb.Interface
	if e != nil && e.InterfaceA != nil {
		r = e.InterfaceA.ConvertToProtoMessage().(*pb.Interface)
	}
	return r
}

func (e *WirelessLink) GetInterfaceB() *pb.Interface {
	var r *pb.Interface
	if e != nil && e.InterfaceB != nil {
		r = e.InterfaceB.ConvertToProtoMessage().(*pb.Interface)
	}
	return r
}

func (e *WirelessLink) GetSsid() *string {
	var r *string
	if e != nil && e.Ssid != nil {
		r = e.Ssid
	}
	return r
}

func (e *WirelessLink) GetStatus() *string {
	var r *string
	if e != nil && e.Status != nil {
		r = e.Status
	}
	return r
}

func (e *WirelessLink) GetTenant() *pb.Tenant {
	var r *pb.Tenant
	if e != nil && e.Tenant != nil {
		r = e.Tenant.ConvertToProtoMessage().(*pb.Tenant)
	}
	return r
}

func (e *WirelessLink) GetAuthType() *string {
	var r *string
	if e != nil && e.AuthType != nil {
		r = e.AuthType
	}
	return r
}

func (e *WirelessLink) GetAuthCipher() *string {
	var r *string
	if e != nil && e.AuthCipher != nil {
		r = e.AuthCipher
	}
	return r
}

func (e *WirelessLink) GetAuthPsk() *string {
	var r *string
	if e != nil && e.AuthPsk != nil {
		r = e.AuthPsk
	}
	return r
}

func (e *WirelessLink) GetDistance() *float64 {
	var r *float64
	if e != nil && e.Distance != nil {
		r = e.Distance
	}
	return r
}

func (e *WirelessLink) GetDistanceUnit() *string {
	var r *string
	if e != nil && e.DistanceUnit != nil {
		r = e.DistanceUnit
	}
	return r
}

func (e *WirelessLink) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *WirelessLink) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *WirelessLink) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *WirelessLink) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *WirelessLink) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *WirelessLink) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type CustomField struct {
	Type                *string
	RelatedObjectType   *string
	Name                *string
	Label               *string
	GroupName           *string
	Description         *string
	Required            *bool
	Unique              *bool
	SearchWeight        *int64
	FilterLogic         *string
	UiVisible           *string
	UiEditable          *string
	IsCloneable         *bool
	Default             *string
	RelatedObjectFilter *string
	Weight              *int64
	ValidationMinimum   *float64
	ValidationMaximum   *float64
	ValidationRegex     *string
	ChoiceSet           *CustomFieldChoiceSet
	Comments            *string
	ObjectTypes         []string
	Metadata            Metadata
	Owner               *Owner
	ValidationSchema    *string
}

func (e *CustomField) ConvertToProtoMessage() proto.Message {
	r := &pb.CustomField{
		Type:                e.GetType(),
		RelatedObjectType:   e.GetRelatedObjectType(),
		Name:                e.GetName(),
		Label:               e.GetLabel(),
		GroupName:           e.GetGroupName(),
		Description:         e.GetDescription(),
		Required:            e.GetRequired(),
		Unique:              e.GetUnique(),
		SearchWeight:        e.GetSearchWeight(),
		FilterLogic:         e.GetFilterLogic(),
		UiVisible:           e.GetUiVisible(),
		UiEditable:          e.GetUiEditable(),
		IsCloneable:         e.GetIsCloneable(),
		Default:             e.GetDefault(),
		RelatedObjectFilter: e.GetRelatedObjectFilter(),
		Weight:              e.GetWeight(),
		ValidationMinimum:   e.GetValidationMinimum(),
		ValidationMaximum:   e.GetValidationMaximum(),
		ValidationRegex:     e.GetValidationRegex(),
		ChoiceSet:           e.GetChoiceSet(),
		Comments:            e.GetComments(),
		ObjectTypes:         e.GetObjectTypes(),
		Metadata:            e.GetMetadata(),
		Owner:               e.GetOwner(),
		ValidationSchema:    e.GetValidationSchema(),
	}
	return r
}

func (e *CustomField) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CustomField{
			CustomField: e.ConvertToProtoMessage().(*pb.CustomField),
		},
	}
}

func (e *CustomField) GetType() string {
	var r string
	if e != nil && e.Type != nil {
		r = *e.Type
	}
	return r
}

func (e *CustomField) GetRelatedObjectType() *string {
	var r *string
	if e != nil && e.RelatedObjectType != nil {
		r = e.RelatedObjectType
	}
	return r
}

func (e *CustomField) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *CustomField) GetLabel() *string {
	var r *string
	if e != nil && e.Label != nil {
		r = e.Label
	}
	return r
}

func (e *CustomField) GetGroupName() *string {
	var r *string
	if e != nil && e.GroupName != nil {
		r = e.GroupName
	}
	return r
}

func (e *CustomField) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *CustomField) GetRequired() *bool {
	var r *bool
	if e != nil && e.Required != nil {
		r = e.Required
	}
	return r
}

func (e *CustomField) GetUnique() *bool {
	var r *bool
	if e != nil && e.Unique != nil {
		r = e.Unique
	}
	return r
}

func (e *CustomField) GetSearchWeight() *int64 {
	var r *int64
	if e != nil && e.SearchWeight != nil {
		r = e.SearchWeight
	}
	return r
}

func (e *CustomField) GetFilterLogic() *string {
	var r *string
	if e != nil && e.FilterLogic != nil {
		r = e.FilterLogic
	}
	return r
}

func (e *CustomField) GetUiVisible() *string {
	var r *string
	if e != nil && e.UiVisible != nil {
		r = e.UiVisible
	}
	return r
}

func (e *CustomField) GetUiEditable() *string {
	var r *string
	if e != nil && e.UiEditable != nil {
		r = e.UiEditable
	}
	return r
}

func (e *CustomField) GetIsCloneable() *bool {
	var r *bool
	if e != nil && e.IsCloneable != nil {
		r = e.IsCloneable
	}
	return r
}

func (e *CustomField) GetDefault() *string {
	var r *string
	if e != nil && e.Default != nil {
		r = e.Default
	}
	return r
}

func (e *CustomField) GetRelatedObjectFilter() *string {
	var r *string
	if e != nil && e.RelatedObjectFilter != nil {
		r = e.RelatedObjectFilter
	}
	return r
}

func (e *CustomField) GetWeight() *int64 {
	var r *int64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *CustomField) GetValidationMinimum() *float64 {
	var r *float64
	if e != nil && e.ValidationMinimum != nil {
		r = e.ValidationMinimum
	}
	return r
}

func (e *CustomField) GetValidationMaximum() *float64 {
	var r *float64
	if e != nil && e.ValidationMaximum != nil {
		r = e.ValidationMaximum
	}
	return r
}

func (e *CustomField) GetValidationRegex() *string {
	var r *string
	if e != nil && e.ValidationRegex != nil {
		r = e.ValidationRegex
	}
	return r
}

func (e *CustomField) GetChoiceSet() *pb.CustomFieldChoiceSet {
	var r *pb.CustomFieldChoiceSet
	if e != nil && e.ChoiceSet != nil {
		r = e.ChoiceSet.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet)
	}
	return r
}

func (e *CustomField) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *CustomField) GetObjectTypes() []string {
	var r []string
	if e != nil && e.ObjectTypes != nil {
		for _, v := range e.ObjectTypes {
			r = append(r, v)
		}
	}
	return r
}

func (e *CustomField) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *CustomField) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *CustomField) GetValidationSchema() *string {
	var r *string
	if e != nil && e.ValidationSchema != nil {
		r = e.ValidationSchema
	}
	return r
}

type CustomFieldChoiceSet struct {
	Name                *string
	Description         *string
	BaseChoices         *string
	OrderAlphabetically *bool
	ExtraChoices        []string
	Metadata            Metadata
	Owner               *Owner
	ChoiceColors        *string
}

func (e *CustomFieldChoiceSet) ConvertToProtoMessage() proto.Message {
	r := &pb.CustomFieldChoiceSet{
		Name:                e.GetName(),
		Description:         e.GetDescription(),
		BaseChoices:         e.GetBaseChoices(),
		OrderAlphabetically: e.GetOrderAlphabetically(),
		ExtraChoices:        e.GetExtraChoices(),
		Metadata:            e.GetMetadata(),
		Owner:               e.GetOwner(),
		ChoiceColors:        e.GetChoiceColors(),
	}
	return r
}

func (e *CustomFieldChoiceSet) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CustomFieldChoiceSet{
			CustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
		},
	}
}

func (e *CustomFieldChoiceSet) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *CustomFieldChoiceSet) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *CustomFieldChoiceSet) GetBaseChoices() *string {
	var r *string
	if e != nil && e.BaseChoices != nil {
		r = e.BaseChoices
	}
	return r
}

func (e *CustomFieldChoiceSet) GetOrderAlphabetically() *bool {
	var r *bool
	if e != nil && e.OrderAlphabetically != nil {
		r = e.OrderAlphabetically
	}
	return r
}

func (e *CustomFieldChoiceSet) GetExtraChoices() []string {
	var r []string
	if e != nil && e.ExtraChoices != nil {
		for _, v := range e.ExtraChoices {
			r = append(r, v)
		}
	}
	return r
}

func (e *CustomFieldChoiceSet) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *CustomFieldChoiceSet) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *CustomFieldChoiceSet) GetChoiceColors() *string {
	var r *string
	if e != nil && e.ChoiceColors != nil {
		r = e.ChoiceColors
	}
	return r
}

type JournalEntry struct {
	// AssignedObject can be:
	//  - ASN
	//  - ASNRange
	//  - Aggregate
	//  - Cable
	//  - CablePath
	//  - CableTermination
	//  - Circuit
	//  - CircuitGroup
	//  - CircuitGroupAssignment
	//  - CircuitTermination
	//  - CircuitType
	//  - Cluster
	//  - ClusterGroup
	//  - ClusterType
	//  - ConsolePort
	//  - ConsoleServerPort
	//  - Contact
	//  - ContactAssignment
	//  - ContactGroup
	//  - ContactRole
	//  - CustomField
	//  - CustomFieldChoiceSet
	//  - Device
	//  - DeviceBay
	//  - DeviceRole
	//  - DeviceType
	//  - FHRPGroup
	//  - FHRPGroupAssignment
	//  - FrontPort
	//  - IKEPolicy
	//  - IKEProposal
	//  - IPAddress
	//  - IPRange
	//  - IPSecPolicy
	//  - IPSecProfile
	//  - IPSecProposal
	//  - Interface
	//  - InventoryItem
	//  - InventoryItemRole
	//  - JournalEntry
	//  - L2VPN
	//  - L2VPNTermination
	//  - Location
	//  - MACAddress
	//  - Manufacturer
	//  - Module
	//  - ModuleBay
	//  - ModuleType
	//  - ModuleTypeProfile
	//  - Platform
	//  - PowerFeed
	//  - PowerOutlet
	//  - PowerPanel
	//  - PowerPort
	//  - Prefix
	//  - Provider
	//  - ProviderAccount
	//  - ProviderNetwork
	//  - RIR
	//  - Rack
	//  - RackReservation
	//  - RackRole
	//  - RackType
	//  - RearPort
	//  - Region
	//  - Role
	//  - RouteTarget
	//  - Service
	//  - Site
	//  - SiteGroup
	//  - Tag
	//  - Tenant
	//  - TenantGroup
	//  - Tunnel
	//  - TunnelGroup
	//  - TunnelTermination
	//  - VLAN
	//  - VLANGroup
	//  - VLANTranslationPolicy
	//  - VLANTranslationRule
	//  - VMInterface
	//  - VRF
	//  - VirtualChassis
	//  - VirtualCircuit
	//  - VirtualCircuitTermination
	//  - VirtualCircuitType
	//  - VirtualDeviceContext
	//  - VirtualDisk
	//  - VirtualMachine
	//  - WirelessLAN
	//  - WirelessLANGroup
	//  - WirelessLink
	//  - CustomLink
	//  - Owner
	//  - OwnerGroup
	//  - CableBundle
	//  - RackGroup
	//  - ScriptModule
	//  - VirtualMachineType
	AssignedObject anyJournalEntryAssignedObjectValue
	Kind           *string
	Comments       *string
	Tags           []*Tag
	CustomFields   map[string]*CustomFieldValue
	Metadata       Metadata
}

func (e *JournalEntry) ConvertToProtoMessage() proto.Message {
	r := &pb.JournalEntry{
		Kind:         e.GetKind(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
	}
	if e.AssignedObject != nil {
		e.AssignedObject.anyJournalEntryAssignedObjectValueApplyTo(r)
	}
	return r
}

func (e *JournalEntry) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_JournalEntry{
			JournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
		},
	}
}

func (e *JournalEntry) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *JournalEntry) GetAssignedObject() any {
	var r any
	if e != nil && e.AssignedObject != nil {
		var tmp pb.JournalEntry
		e.AssignedObject.anyJournalEntryAssignedObjectValueApplyTo(&tmp)
		r = tmp.AssignedObject
	}
	return r
}

func (e *JournalEntry) GetKind() *string {
	var r *string
	if e != nil && e.Kind != nil {
		r = e.Kind
	}
	return r
}

func (e *JournalEntry) GetComments() string {
	var r string
	if e != nil && e.Comments != nil {
		r = *e.Comments
	}
	return r
}

func (e *JournalEntry) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *JournalEntry) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *JournalEntry) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type ModuleTypeProfile struct {
	Name         *string
	Description  *string
	Schema       *string
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
	Owner        *Owner
}

func (e *ModuleTypeProfile) ConvertToProtoMessage() proto.Message {
	r := &pb.ModuleTypeProfile{
		Name:         e.GetName(),
		Description:  e.GetDescription(),
		Schema:       e.GetSchema(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
		Owner:        e.GetOwner(),
	}
	return r
}

func (e *ModuleTypeProfile) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ModuleTypeProfile{
			ModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
		},
	}
}

func (e *ModuleTypeProfile) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *ModuleTypeProfile) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *ModuleTypeProfile) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *ModuleTypeProfile) GetSchema() *string {
	var r *string
	if e != nil && e.Schema != nil {
		r = e.Schema
	}
	return r
}

func (e *ModuleTypeProfile) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *ModuleTypeProfile) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *ModuleTypeProfile) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *ModuleTypeProfile) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *ModuleTypeProfile) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type CustomLink struct {
	Name        *string
	Enabled     *bool
	LinkText    *string
	LinkUrl     *string
	Weight      *int64
	GroupName   *string
	ButtonClass *string
	NewWindow   *bool
	ObjectTypes []string
	Metadata    Metadata
	Owner       *Owner
}

func (e *CustomLink) ConvertToProtoMessage() proto.Message {
	r := &pb.CustomLink{
		Name:        e.GetName(),
		Enabled:     e.GetEnabled(),
		LinkText:    e.GetLinkText(),
		LinkUrl:     e.GetLinkUrl(),
		Weight:      e.GetWeight(),
		GroupName:   e.GetGroupName(),
		ButtonClass: e.GetButtonClass(),
		NewWindow:   e.GetNewWindow(),
		ObjectTypes: e.GetObjectTypes(),
		Metadata:    e.GetMetadata(),
		Owner:       e.GetOwner(),
	}
	return r
}

func (e *CustomLink) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CustomLink{
			CustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
		},
	}
}

func (e *CustomLink) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *CustomLink) GetEnabled() *bool {
	var r *bool
	if e != nil && e.Enabled != nil {
		r = e.Enabled
	}
	return r
}

func (e *CustomLink) GetLinkText() string {
	var r string
	if e != nil && e.LinkText != nil {
		r = *e.LinkText
	}
	return r
}

func (e *CustomLink) GetLinkUrl() string {
	var r string
	if e != nil && e.LinkUrl != nil {
		r = *e.LinkUrl
	}
	return r
}

func (e *CustomLink) GetWeight() *int64 {
	var r *int64
	if e != nil && e.Weight != nil {
		r = e.Weight
	}
	return r
}

func (e *CustomLink) GetGroupName() *string {
	var r *string
	if e != nil && e.GroupName != nil {
		r = e.GroupName
	}
	return r
}

func (e *CustomLink) GetButtonClass() *string {
	var r *string
	if e != nil && e.ButtonClass != nil {
		r = e.ButtonClass
	}
	return r
}

func (e *CustomLink) GetNewWindow() *bool {
	var r *bool
	if e != nil && e.NewWindow != nil {
		r = e.NewWindow
	}
	return r
}

func (e *CustomLink) GetObjectTypes() []string {
	var r []string
	if e != nil && e.ObjectTypes != nil {
		for _, v := range e.ObjectTypes {
			r = append(r, v)
		}
	}
	return r
}

func (e *CustomLink) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func (e *CustomLink) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

type Owner struct {
	Name        *string
	Group       *OwnerGroup
	Description *string
	Metadata    Metadata
}

func (e *Owner) ConvertToProtoMessage() proto.Message {
	r := &pb.Owner{
		Name:        e.GetName(),
		Group:       e.GetGroup(),
		Description: e.GetDescription(),
		Metadata:    e.GetMetadata(),
	}
	return r
}

func (e *Owner) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_Owner{
			Owner: e.ConvertToProtoMessage().(*pb.Owner),
		},
	}
}

func (e *Owner) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *Owner) GetGroup() *pb.OwnerGroup {
	var r *pb.OwnerGroup
	if e != nil && e.Group != nil {
		r = e.Group.ConvertToProtoMessage().(*pb.OwnerGroup)
	}
	return r
}

func (e *Owner) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *Owner) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type OwnerGroup struct {
	Name        *string
	Description *string
	Metadata    Metadata
}

func (e *OwnerGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.OwnerGroup{
		Name:        e.GetName(),
		Description: e.GetDescription(),
		Metadata:    e.GetMetadata(),
	}
	return r
}

func (e *OwnerGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_OwnerGroup{
			OwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
		},
	}
}

func (e *OwnerGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *OwnerGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *OwnerGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type DeviceConfig struct {
	Startup   []byte
	Running   []byte
	Candidate []byte
	Metadata  Metadata
}

func (e *DeviceConfig) ConvertToProtoMessage() proto.Message {
	r := &pb.DeviceConfig{
		Startup:   e.GetStartup(),
		Running:   e.GetRunning(),
		Candidate: e.GetCandidate(),
		Metadata:  e.GetMetadata(),
	}
	return r
}

func (e *DeviceConfig) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_DeviceConfig{
			DeviceConfig: e.ConvertToProtoMessage().(*pb.DeviceConfig),
		},
	}
}

func (e *DeviceConfig) GetStartup() []byte {
	if e != nil {
		return e.Startup
	}
	return nil
}

func (e *DeviceConfig) GetRunning() []byte {
	if e != nil {
		return e.Running
	}
	return nil
}

func (e *DeviceConfig) GetCandidate() []byte {
	if e != nil {
		return e.Candidate
	}
	return nil
}

func (e *DeviceConfig) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type CableBundle struct {
	Name         *string
	Description  *string
	Owner        *Owner
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
}

func (e *CableBundle) ConvertToProtoMessage() proto.Message {
	r := &pb.CableBundle{
		Name:         e.GetName(),
		Description:  e.GetDescription(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
	}
	return r
}

func (e *CableBundle) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_CableBundle{
			CableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
		},
	}
}

func (e *CableBundle) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *CableBundle) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *CableBundle) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *CableBundle) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *CableBundle) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *CableBundle) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *CableBundle) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *CableBundle) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type RackGroup struct {
	Name         *string
	Slug         *string
	Description  *string
	Owner        *Owner
	Comments     *string
	Tags         []*Tag
	CustomFields map[string]*CustomFieldValue
	Metadata     Metadata
}

func (e *RackGroup) ConvertToProtoMessage() proto.Message {
	r := &pb.RackGroup{
		Name:         e.GetName(),
		Slug:         e.GetSlug(),
		Description:  e.GetDescription(),
		Owner:        e.GetOwner(),
		Comments:     e.GetComments(),
		Tags:         e.GetTags(),
		CustomFields: e.GetCustomFields(),
		Metadata:     e.GetMetadata(),
	}
	return r
}

func (e *RackGroup) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_RackGroup{
			RackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
		},
	}
}

func (e *RackGroup) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *RackGroup) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *RackGroup) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *RackGroup) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *RackGroup) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *RackGroup) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *RackGroup) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *RackGroup) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *RackGroup) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type ScriptModule struct {
	File     *string
	Metadata Metadata
}

func (e *ScriptModule) ConvertToProtoMessage() proto.Message {
	r := &pb.ScriptModule{
		File:     e.GetFile(),
		Metadata: e.GetMetadata(),
	}
	return r
}

func (e *ScriptModule) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_ScriptModule{
			ScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
		},
	}
}

func (e *ScriptModule) GetFile() string {
	var r string
	if e != nil && e.File != nil {
		r = *e.File
	}
	return r
}

func (e *ScriptModule) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

type VirtualMachineType struct {
	Name            *string
	Slug            *string
	DefaultPlatform *Platform
	DefaultVcpus    *float64
	DefaultMemory   *int64
	Description     *string
	Owner           *Owner
	Comments        *string
	Tags            []*Tag
	CustomFields    map[string]*CustomFieldValue
	Metadata        Metadata
}

func (e *VirtualMachineType) ConvertToProtoMessage() proto.Message {
	r := &pb.VirtualMachineType{
		Name:            e.GetName(),
		Slug:            e.GetSlug(),
		DefaultPlatform: e.GetDefaultPlatform(),
		DefaultVcpus:    e.GetDefaultVcpus(),
		DefaultMemory:   e.GetDefaultMemory(),
		Description:     e.GetDescription(),
		Owner:           e.GetOwner(),
		Comments:        e.GetComments(),
		Tags:            e.GetTags(),
		CustomFields:    e.GetCustomFields(),
		Metadata:        e.GetMetadata(),
	}
	return r
}

func (e *VirtualMachineType) ConvertToProtoEntity() *pb.Entity {
	return &pb.Entity{
		Entity: &pb.Entity_VirtualMachineType{
			VirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
		},
	}
}

func (e *VirtualMachineType) SetCustomField(key string, value any) error {
	if e.CustomFields == nil {
		e.CustomFields = make(map[string]*CustomFieldValue)
	}
	return setCustomField(e.CustomFields, key, value)
}

func (e *VirtualMachineType) GetName() string {
	var r string
	if e != nil && e.Name != nil {
		r = *e.Name
	}
	return r
}

func (e *VirtualMachineType) GetSlug() string {
	var r string
	if e != nil && e.Slug != nil {
		r = *e.Slug
	}
	return r
}

func (e *VirtualMachineType) GetDefaultPlatform() *pb.Platform {
	var r *pb.Platform
	if e != nil && e.DefaultPlatform != nil {
		r = e.DefaultPlatform.ConvertToProtoMessage().(*pb.Platform)
	}
	return r
}

func (e *VirtualMachineType) GetDefaultVcpus() *float64 {
	var r *float64
	if e != nil && e.DefaultVcpus != nil {
		r = e.DefaultVcpus
	}
	return r
}

func (e *VirtualMachineType) GetDefaultMemory() *int64 {
	var r *int64
	if e != nil && e.DefaultMemory != nil {
		r = e.DefaultMemory
	}
	return r
}

func (e *VirtualMachineType) GetDescription() *string {
	var r *string
	if e != nil && e.Description != nil {
		r = e.Description
	}
	return r
}

func (e *VirtualMachineType) GetOwner() *pb.Owner {
	var r *pb.Owner
	if e != nil && e.Owner != nil {
		r = e.Owner.ConvertToProtoMessage().(*pb.Owner)
	}
	return r
}

func (e *VirtualMachineType) GetComments() *string {
	var r *string
	if e != nil && e.Comments != nil {
		r = e.Comments
	}
	return r
}

func (e *VirtualMachineType) GetTags() []*pb.Tag {
	var r []*pb.Tag
	if e != nil && e.Tags != nil {
		for _, v := range e.Tags {
			r = append(r, v.ConvertToProtoMessage().(*pb.Tag))
		}
	}
	return r
}

func (e *VirtualMachineType) GetCustomFields() map[string]*pb.CustomFieldValue {
	var r map[string]*pb.CustomFieldValue
	if e != nil && e.CustomFields != nil {
		r = make(map[string]*pb.CustomFieldValue)
		for k, v := range e.CustomFields {
			r[k] = v.ConvertToProtoMessage().(*pb.CustomFieldValue)
		}
	}
	return r
}

func (e *VirtualMachineType) GetMetadata() *structpb.Struct {
	var r *structpb.Struct
	if e != nil && e.Metadata != nil {
		r, _ = structpb.NewStruct(convertMetadata(e.Metadata))
	}
	return r
}

func setCustomField(cf map[string]*CustomFieldValue, key string, value any) error {
	switch v := value.(type) {
	case json.RawMessage:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueJson(string(v)),
		}
	case string:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueText(v),
		}
	case int32:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueInteger(v),
		}
	case int64:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueInteger(v),
		}
	case bool:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueBoolean(v),
		}
	case float32:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueDecimal(v),
		}
	case float64:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueDecimal(v),
		}
	case time.Time:
		cf[key] = &CustomFieldValue{
			Value: CustomFieldValueDatetime(v),
		}
	case *CustomFieldValue:
		cf[key] = v
	case CustomFieldValue:
		cf[key] = &v
	case anyCustomFieldObjectReferenceObjectValue:
		cf[key] = &CustomFieldValue{
			Value: &CustomFieldObjectReference{
				Object: v,
			},
		}
	case []string:
		cf[key] = &CustomFieldValue{
			MultipleSelection: value.([]string),
		}
	case []Entity:
		vals := make([]*CustomFieldObjectReference, 0)
		for _, e := range v {
			if ref, ok := e.(anyCustomFieldObjectReferenceObjectValue); ok {
				vals = append(vals, &CustomFieldObjectReference{
					Object: ref,
				})
			} else {
				return fmt.Errorf("Unsupported custom field value for reference in list: %T", e)
			}
		}
		cf[key] = &CustomFieldValue{
			MultipleObjects: vals,
		}
	default:
		return fmt.Errorf("Unsupported custom field value type: %T", v)
	}
	return nil
}

// implementation of oneof interfaces for ASN.

func (e *ASN) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectAsn{
		ObjectAsn: e.ConvertToProtoMessage().(*pb.ASN),
	}
}

func (e *ASN) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Asn{
		Asn: e.ConvertToProtoMessage().(*pb.ASN),
	}
}

func (e *ASN) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceAsn{
		InterfaceAsn: e.ConvertToProtoMessage().(*pb.ASN),
	}
}

func (e *ASN) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectAsn{
		ObjectAsn: e.ConvertToProtoMessage().(*pb.ASN),
	}
}

func (e *ASN) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectAsn{
		AssignedObjectAsn: e.ConvertToProtoMessage().(*pb.ASN),
	}
}

func (e *ASN) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationAsn{
		TerminationAsn: e.ConvertToProtoMessage().(*pb.ASN),
	}
}

func (e *ASN) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectAsn{
		AssignedObjectAsn: e.ConvertToProtoMessage().(*pb.ASN),
	}
}

// implementation of oneof interfaces for ASNRange.

func (e *ASNRange) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectAsnRange{
		ObjectAsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
	}
}

func (e *ASNRange) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_AsnRange{
		AsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
	}
}

func (e *ASNRange) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceAsnRange{
		InterfaceAsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
	}
}

func (e *ASNRange) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectAsnRange{
		ObjectAsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
	}
}

func (e *ASNRange) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectAsnRange{
		AssignedObjectAsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
	}
}

func (e *ASNRange) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationAsnRange{
		TerminationAsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
	}
}

func (e *ASNRange) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectAsnRange{
		AssignedObjectAsnRange: e.ConvertToProtoMessage().(*pb.ASNRange),
	}
}

// implementation of oneof interfaces for Aggregate.

func (e *Aggregate) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectAggregate{
		ObjectAggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
	}
}

func (e *Aggregate) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Aggregate{
		Aggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
	}
}

func (e *Aggregate) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceAggregate{
		InterfaceAggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
	}
}

func (e *Aggregate) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectAggregate{
		ObjectAggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
	}
}

func (e *Aggregate) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectAggregate{
		AssignedObjectAggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
	}
}

func (e *Aggregate) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationAggregate{
		TerminationAggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
	}
}

func (e *Aggregate) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectAggregate{
		AssignedObjectAggregate: e.ConvertToProtoMessage().(*pb.Aggregate),
	}
}

// implementation of oneof interfaces for Cable.

func (e *Cable) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCable{
		ObjectCable: e.ConvertToProtoMessage().(*pb.Cable),
	}
}

func (e *Cable) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Cable{
		Cable: e.ConvertToProtoMessage().(*pb.Cable),
	}
}

func (e *Cable) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCable{
		InterfaceCable: e.ConvertToProtoMessage().(*pb.Cable),
	}
}

func (e *Cable) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCable{
		ObjectCable: e.ConvertToProtoMessage().(*pb.Cable),
	}
}

func (e *Cable) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCable{
		AssignedObjectCable: e.ConvertToProtoMessage().(*pb.Cable),
	}
}

func (e *Cable) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCable{
		TerminationCable: e.ConvertToProtoMessage().(*pb.Cable),
	}
}

func (e *Cable) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCable{
		AssignedObjectCable: e.ConvertToProtoMessage().(*pb.Cable),
	}
}

// implementation of oneof interfaces for CablePath.

func (e *CablePath) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCablePath{
		ObjectCablePath: e.ConvertToProtoMessage().(*pb.CablePath),
	}
}

func (e *CablePath) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CablePath{
		CablePath: e.ConvertToProtoMessage().(*pb.CablePath),
	}
}

func (e *CablePath) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCablePath{
		InterfaceCablePath: e.ConvertToProtoMessage().(*pb.CablePath),
	}
}

func (e *CablePath) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCablePath{
		ObjectCablePath: e.ConvertToProtoMessage().(*pb.CablePath),
	}
}

func (e *CablePath) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCablePath{
		AssignedObjectCablePath: e.ConvertToProtoMessage().(*pb.CablePath),
	}
}

func (e *CablePath) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCablePath{
		TerminationCablePath: e.ConvertToProtoMessage().(*pb.CablePath),
	}
}

func (e *CablePath) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCablePath{
		AssignedObjectCablePath: e.ConvertToProtoMessage().(*pb.CablePath),
	}
}

// implementation of oneof interfaces for CableTermination.

func (e *CableTermination) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCableTermination{
		ObjectCableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
	}
}

func (e *CableTermination) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CableTermination{
		CableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
	}
}

func (e *CableTermination) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCableTermination{
		InterfaceCableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
	}
}

func (e *CableTermination) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCableTermination{
		ObjectCableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
	}
}

func (e *CableTermination) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCableTermination{
		AssignedObjectCableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
	}
}

func (e *CableTermination) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCableTermination{
		TerminationCableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
	}
}

func (e *CableTermination) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCableTermination{
		AssignedObjectCableTermination: e.ConvertToProtoMessage().(*pb.CableTermination),
	}
}

// implementation of oneof interfaces for Circuit.

func (e *Circuit) anyCircuitGroupAssignmentMemberValueApplyTo(p *pb.CircuitGroupAssignment) {
	p.Member = &pb.CircuitGroupAssignment_MemberCircuit{
		MemberCircuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

func (e *Circuit) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCircuit{
		ObjectCircuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

func (e *Circuit) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Circuit{
		Circuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

func (e *Circuit) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCircuit{
		InterfaceCircuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

func (e *Circuit) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCircuit{
		ObjectCircuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

func (e *Circuit) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCircuit{
		AssignedObjectCircuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

func (e *Circuit) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCircuit{
		TerminationCircuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

func (e *Circuit) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCircuit{
		AssignedObjectCircuit: e.ConvertToProtoMessage().(*pb.Circuit),
	}
}

// implementation of oneof interfaces for CircuitGroup.

func (e *CircuitGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCircuitGroup{
		ObjectCircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
	}
}

func (e *CircuitGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CircuitGroup{
		CircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
	}
}

func (e *CircuitGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCircuitGroup{
		InterfaceCircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
	}
}

func (e *CircuitGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCircuitGroup{
		ObjectCircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
	}
}

func (e *CircuitGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCircuitGroup{
		AssignedObjectCircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
	}
}

func (e *CircuitGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCircuitGroup{
		TerminationCircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
	}
}

func (e *CircuitGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCircuitGroup{
		AssignedObjectCircuitGroup: e.ConvertToProtoMessage().(*pb.CircuitGroup),
	}
}

// implementation of oneof interfaces for CircuitGroupAssignment.

func (e *CircuitGroupAssignment) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCircuitGroupAssignment{
		ObjectCircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
	}
}

func (e *CircuitGroupAssignment) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CircuitGroupAssignment{
		CircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
	}
}

func (e *CircuitGroupAssignment) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCircuitGroupAssignment{
		InterfaceCircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
	}
}

func (e *CircuitGroupAssignment) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCircuitGroupAssignment{
		ObjectCircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
	}
}

func (e *CircuitGroupAssignment) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCircuitGroupAssignment{
		AssignedObjectCircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
	}
}

func (e *CircuitGroupAssignment) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCircuitGroupAssignment{
		TerminationCircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
	}
}

func (e *CircuitGroupAssignment) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCircuitGroupAssignment{
		AssignedObjectCircuitGroupAssignment: e.ConvertToProtoMessage().(*pb.CircuitGroupAssignment),
	}
}

// implementation of oneof interfaces for CircuitTermination.

func (e *CircuitTermination) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationCircuitTermination{
		TerminationCircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

func (e *CircuitTermination) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCircuitTermination{
		ObjectCircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

func (e *CircuitTermination) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CircuitTermination{
		CircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

func (e *CircuitTermination) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCircuitTermination{
		InterfaceCircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

func (e *CircuitTermination) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCircuitTermination{
		ObjectCircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

func (e *CircuitTermination) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCircuitTermination{
		AssignedObjectCircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

func (e *CircuitTermination) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCircuitTermination{
		TerminationCircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

func (e *CircuitTermination) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCircuitTermination{
		AssignedObjectCircuitTermination: e.ConvertToProtoMessage().(*pb.CircuitTermination),
	}
}

// implementation of oneof interfaces for CircuitType.

func (e *CircuitType) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCircuitType{
		ObjectCircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
	}
}

func (e *CircuitType) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CircuitType{
		CircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
	}
}

func (e *CircuitType) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCircuitType{
		InterfaceCircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
	}
}

func (e *CircuitType) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCircuitType{
		ObjectCircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
	}
}

func (e *CircuitType) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCircuitType{
		AssignedObjectCircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
	}
}

func (e *CircuitType) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCircuitType{
		TerminationCircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
	}
}

func (e *CircuitType) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCircuitType{
		AssignedObjectCircuitType: e.ConvertToProtoMessage().(*pb.CircuitType),
	}
}

// implementation of oneof interfaces for Cluster.

func (e *Cluster) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCluster{
		ObjectCluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

func (e *Cluster) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Cluster{
		Cluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

func (e *Cluster) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCluster{
		InterfaceCluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

func (e *Cluster) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCluster{
		ObjectCluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

func (e *Cluster) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCluster{
		AssignedObjectCluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

func (e *Cluster) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCluster{
		TerminationCluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

func (e *Cluster) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeCluster{
		ScopeCluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

func (e *Cluster) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCluster{
		AssignedObjectCluster: e.ConvertToProtoMessage().(*pb.Cluster),
	}
}

// implementation of oneof interfaces for ClusterGroup.

func (e *ClusterGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectClusterGroup{
		ObjectClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

func (e *ClusterGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ClusterGroup{
		ClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

func (e *ClusterGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceClusterGroup{
		InterfaceClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

func (e *ClusterGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectClusterGroup{
		ObjectClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

func (e *ClusterGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectClusterGroup{
		AssignedObjectClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

func (e *ClusterGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationClusterGroup{
		TerminationClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

func (e *ClusterGroup) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeClusterGroup{
		ScopeClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

func (e *ClusterGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectClusterGroup{
		AssignedObjectClusterGroup: e.ConvertToProtoMessage().(*pb.ClusterGroup),
	}
}

// implementation of oneof interfaces for ClusterType.

func (e *ClusterType) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectClusterType{
		ObjectClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
	}
}

func (e *ClusterType) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ClusterType{
		ClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
	}
}

func (e *ClusterType) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceClusterType{
		InterfaceClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
	}
}

func (e *ClusterType) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectClusterType{
		ObjectClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
	}
}

func (e *ClusterType) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectClusterType{
		AssignedObjectClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
	}
}

func (e *ClusterType) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationClusterType{
		TerminationClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
	}
}

func (e *ClusterType) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectClusterType{
		AssignedObjectClusterType: e.ConvertToProtoMessage().(*pb.ClusterType),
	}
}

// implementation of oneof interfaces for ConsolePort.

func (e *ConsolePort) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationConsolePort{
		TerminationConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectConsolePort{
		ObjectConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ConsolePort{
		ConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceConsolePort{
		InterfaceConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectConsolePort{
		ObjectConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyInventoryItemComponentValueApplyTo(p *pb.InventoryItem) {
	p.Component = &pb.InventoryItem_ComponentConsolePort{
		ComponentConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectConsolePort{
		AssignedObjectConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationConsolePort{
		TerminationConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

func (e *ConsolePort) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectConsolePort{
		AssignedObjectConsolePort: e.ConvertToProtoMessage().(*pb.ConsolePort),
	}
}

// implementation of oneof interfaces for ConsoleServerPort.

func (e *ConsoleServerPort) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationConsoleServerPort{
		TerminationConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectConsoleServerPort{
		ObjectConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ConsoleServerPort{
		ConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceConsoleServerPort{
		InterfaceConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectConsoleServerPort{
		ObjectConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyInventoryItemComponentValueApplyTo(p *pb.InventoryItem) {
	p.Component = &pb.InventoryItem_ComponentConsoleServerPort{
		ComponentConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectConsoleServerPort{
		AssignedObjectConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationConsoleServerPort{
		TerminationConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

func (e *ConsoleServerPort) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectConsoleServerPort{
		AssignedObjectConsoleServerPort: e.ConvertToProtoMessage().(*pb.ConsoleServerPort),
	}
}

// implementation of oneof interfaces for Contact.

func (e *Contact) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectContact{
		ObjectContact: e.ConvertToProtoMessage().(*pb.Contact),
	}
}

func (e *Contact) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Contact{
		Contact: e.ConvertToProtoMessage().(*pb.Contact),
	}
}

func (e *Contact) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceContact{
		InterfaceContact: e.ConvertToProtoMessage().(*pb.Contact),
	}
}

func (e *Contact) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectContact{
		ObjectContact: e.ConvertToProtoMessage().(*pb.Contact),
	}
}

func (e *Contact) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectContact{
		AssignedObjectContact: e.ConvertToProtoMessage().(*pb.Contact),
	}
}

func (e *Contact) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationContact{
		TerminationContact: e.ConvertToProtoMessage().(*pb.Contact),
	}
}

func (e *Contact) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectContact{
		AssignedObjectContact: e.ConvertToProtoMessage().(*pb.Contact),
	}
}

// implementation of oneof interfaces for ContactAssignment.

func (e *ContactAssignment) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectContactAssignment{
		ObjectContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
	}
}

func (e *ContactAssignment) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ContactAssignment{
		ContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
	}
}

func (e *ContactAssignment) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceContactAssignment{
		InterfaceContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
	}
}

func (e *ContactAssignment) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectContactAssignment{
		ObjectContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
	}
}

func (e *ContactAssignment) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectContactAssignment{
		AssignedObjectContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
	}
}

func (e *ContactAssignment) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationContactAssignment{
		TerminationContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
	}
}

func (e *ContactAssignment) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectContactAssignment{
		AssignedObjectContactAssignment: e.ConvertToProtoMessage().(*pb.ContactAssignment),
	}
}

// implementation of oneof interfaces for ContactGroup.

func (e *ContactGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectContactGroup{
		ObjectContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
	}
}

func (e *ContactGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ContactGroup{
		ContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
	}
}

func (e *ContactGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceContactGroup{
		InterfaceContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
	}
}

func (e *ContactGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectContactGroup{
		ObjectContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
	}
}

func (e *ContactGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectContactGroup{
		AssignedObjectContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
	}
}

func (e *ContactGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationContactGroup{
		TerminationContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
	}
}

func (e *ContactGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectContactGroup{
		AssignedObjectContactGroup: e.ConvertToProtoMessage().(*pb.ContactGroup),
	}
}

// implementation of oneof interfaces for ContactRole.

func (e *ContactRole) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectContactRole{
		ObjectContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
	}
}

func (e *ContactRole) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ContactRole{
		ContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
	}
}

func (e *ContactRole) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceContactRole{
		InterfaceContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
	}
}

func (e *ContactRole) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectContactRole{
		ObjectContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
	}
}

func (e *ContactRole) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectContactRole{
		AssignedObjectContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
	}
}

func (e *ContactRole) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationContactRole{
		TerminationContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
	}
}

func (e *ContactRole) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectContactRole{
		AssignedObjectContactRole: e.ConvertToProtoMessage().(*pb.ContactRole),
	}
}

// implementation of oneof interfaces for CustomFieldObjectReference.

func (e *CustomFieldObjectReference) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Object{
		Object: e.ConvertToProtoMessage().(*pb.CustomFieldObjectReference),
	}
}

// implementation of oneof interfaces for Device.

func (e *Device) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectDevice{
		ObjectDevice: e.ConvertToProtoMessage().(*pb.Device),
	}
}

func (e *Device) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Device{
		Device: e.ConvertToProtoMessage().(*pb.Device),
	}
}

func (e *Device) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceDevice{
		InterfaceDevice: e.ConvertToProtoMessage().(*pb.Device),
	}
}

func (e *Device) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectDevice{
		ObjectDevice: e.ConvertToProtoMessage().(*pb.Device),
	}
}

func (e *Device) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectDevice{
		AssignedObjectDevice: e.ConvertToProtoMessage().(*pb.Device),
	}
}

func (e *Device) anyServiceParentObjectValueApplyTo(p *pb.Service) {
	p.ParentObject = &pb.Service_ParentObjectDevice{
		ParentObjectDevice: e.ConvertToProtoMessage().(*pb.Device),
	}
}

func (e *Device) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationDevice{
		TerminationDevice: e.ConvertToProtoMessage().(*pb.Device),
	}
}

func (e *Device) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectDevice{
		AssignedObjectDevice: e.ConvertToProtoMessage().(*pb.Device),
	}
}

// implementation of oneof interfaces for DeviceBay.

func (e *DeviceBay) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectDeviceBay{
		ObjectDeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
	}
}

func (e *DeviceBay) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_DeviceBay{
		DeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
	}
}

func (e *DeviceBay) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceDeviceBay{
		InterfaceDeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
	}
}

func (e *DeviceBay) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectDeviceBay{
		ObjectDeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
	}
}

func (e *DeviceBay) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectDeviceBay{
		AssignedObjectDeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
	}
}

func (e *DeviceBay) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationDeviceBay{
		TerminationDeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
	}
}

func (e *DeviceBay) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectDeviceBay{
		AssignedObjectDeviceBay: e.ConvertToProtoMessage().(*pb.DeviceBay),
	}
}

// implementation of oneof interfaces for DeviceRole.

func (e *DeviceRole) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectDeviceRole{
		ObjectDeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
	}
}

func (e *DeviceRole) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_DeviceRole{
		DeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
	}
}

func (e *DeviceRole) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceDeviceRole{
		InterfaceDeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
	}
}

func (e *DeviceRole) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectDeviceRole{
		ObjectDeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
	}
}

func (e *DeviceRole) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectDeviceRole{
		AssignedObjectDeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
	}
}

func (e *DeviceRole) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationDeviceRole{
		TerminationDeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
	}
}

func (e *DeviceRole) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectDeviceRole{
		AssignedObjectDeviceRole: e.ConvertToProtoMessage().(*pb.DeviceRole),
	}
}

// implementation of oneof interfaces for DeviceType.

func (e *DeviceType) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectDeviceType{
		ObjectDeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
	}
}

func (e *DeviceType) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_DeviceType{
		DeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
	}
}

func (e *DeviceType) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceDeviceType{
		InterfaceDeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
	}
}

func (e *DeviceType) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectDeviceType{
		ObjectDeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
	}
}

func (e *DeviceType) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectDeviceType{
		AssignedObjectDeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
	}
}

func (e *DeviceType) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationDeviceType{
		TerminationDeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
	}
}

func (e *DeviceType) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectDeviceType{
		AssignedObjectDeviceType: e.ConvertToProtoMessage().(*pb.DeviceType),
	}
}

// implementation of oneof interfaces for FHRPGroup.

func (e *FHRPGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectFhrpGroup{
		ObjectFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_FhrpGroup{
		FhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceFhrpGroup{
		InterfaceFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectFhrpGroup{
		ObjectFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyIPAddressAssignedObjectValueApplyTo(p *pb.IPAddress) {
	p.AssignedObject = &pb.IPAddress_AssignedObjectFhrpGroup{
		AssignedObjectFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectFhrpGroup{
		AssignedObjectFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyServiceParentObjectValueApplyTo(p *pb.Service) {
	p.ParentObject = &pb.Service_ParentObjectFhrpGroup{
		ParentObjectFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationFhrpGroup{
		TerminationFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

func (e *FHRPGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectFhrpGroup{
		AssignedObjectFhrpGroup: e.ConvertToProtoMessage().(*pb.FHRPGroup),
	}
}

// implementation of oneof interfaces for FHRPGroupAssignment.

func (e *FHRPGroupAssignment) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectFhrpGroupAssignment{
		ObjectFhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
	}
}

func (e *FHRPGroupAssignment) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_FhrpGroupAssignment{
		FhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
	}
}

func (e *FHRPGroupAssignment) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceFhrpGroupAssignment{
		InterfaceFhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
	}
}

func (e *FHRPGroupAssignment) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectFhrpGroupAssignment{
		ObjectFhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
	}
}

func (e *FHRPGroupAssignment) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectFhrpGroupAssignment{
		AssignedObjectFhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
	}
}

func (e *FHRPGroupAssignment) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationFhrpGroupAssignment{
		TerminationFhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
	}
}

func (e *FHRPGroupAssignment) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectFhrpGroupAssignment{
		AssignedObjectFhrpGroupAssignment: e.ConvertToProtoMessage().(*pb.FHRPGroupAssignment),
	}
}

// implementation of oneof interfaces for FrontPort.

func (e *FrontPort) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationFrontPort{
		TerminationFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectFrontPort{
		ObjectFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_FrontPort{
		FrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceFrontPort{
		InterfaceFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectFrontPort{
		ObjectFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyInventoryItemComponentValueApplyTo(p *pb.InventoryItem) {
	p.Component = &pb.InventoryItem_ComponentFrontPort{
		ComponentFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectFrontPort{
		AssignedObjectFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationFrontPort{
		TerminationFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

func (e *FrontPort) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectFrontPort{
		AssignedObjectFrontPort: e.ConvertToProtoMessage().(*pb.FrontPort),
	}
}

// implementation of oneof interfaces for IKEPolicy.

func (e *IKEPolicy) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectIkePolicy{
		ObjectIkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
	}
}

func (e *IKEPolicy) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_IkePolicy{
		IkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
	}
}

func (e *IKEPolicy) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceIkePolicy{
		InterfaceIkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
	}
}

func (e *IKEPolicy) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectIkePolicy{
		ObjectIkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
	}
}

func (e *IKEPolicy) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectIkePolicy{
		AssignedObjectIkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
	}
}

func (e *IKEPolicy) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationIkePolicy{
		TerminationIkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
	}
}

func (e *IKEPolicy) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectIkePolicy{
		AssignedObjectIkePolicy: e.ConvertToProtoMessage().(*pb.IKEPolicy),
	}
}

// implementation of oneof interfaces for IKEProposal.

func (e *IKEProposal) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectIkeProposal{
		ObjectIkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
	}
}

func (e *IKEProposal) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_IkeProposal{
		IkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
	}
}

func (e *IKEProposal) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceIkeProposal{
		InterfaceIkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
	}
}

func (e *IKEProposal) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectIkeProposal{
		ObjectIkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
	}
}

func (e *IKEProposal) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectIkeProposal{
		AssignedObjectIkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
	}
}

func (e *IKEProposal) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationIkeProposal{
		TerminationIkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
	}
}

func (e *IKEProposal) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectIkeProposal{
		AssignedObjectIkeProposal: e.ConvertToProtoMessage().(*pb.IKEProposal),
	}
}

// implementation of oneof interfaces for IPAddress.

func (e *IPAddress) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectIpAddress{
		ObjectIpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
	}
}

func (e *IPAddress) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_IpAddress{
		IpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
	}
}

func (e *IPAddress) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceIpAddress{
		InterfaceIpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
	}
}

func (e *IPAddress) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectIpAddress{
		ObjectIpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
	}
}

func (e *IPAddress) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectIpAddress{
		AssignedObjectIpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
	}
}

func (e *IPAddress) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationIpAddress{
		TerminationIpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
	}
}

func (e *IPAddress) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectIpAddress{
		AssignedObjectIpAddress: e.ConvertToProtoMessage().(*pb.IPAddress),
	}
}

// implementation of oneof interfaces for IPRange.

func (e *IPRange) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectIpRange{
		ObjectIpRange: e.ConvertToProtoMessage().(*pb.IPRange),
	}
}

func (e *IPRange) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_IpRange{
		IpRange: e.ConvertToProtoMessage().(*pb.IPRange),
	}
}

func (e *IPRange) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceIpRange{
		InterfaceIpRange: e.ConvertToProtoMessage().(*pb.IPRange),
	}
}

func (e *IPRange) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectIpRange{
		ObjectIpRange: e.ConvertToProtoMessage().(*pb.IPRange),
	}
}

func (e *IPRange) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectIpRange{
		AssignedObjectIpRange: e.ConvertToProtoMessage().(*pb.IPRange),
	}
}

func (e *IPRange) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationIpRange{
		TerminationIpRange: e.ConvertToProtoMessage().(*pb.IPRange),
	}
}

func (e *IPRange) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectIpRange{
		AssignedObjectIpRange: e.ConvertToProtoMessage().(*pb.IPRange),
	}
}

// implementation of oneof interfaces for IPSecPolicy.

func (e *IPSecPolicy) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectIpSecPolicy{
		ObjectIpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
	}
}

func (e *IPSecPolicy) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_IpSecPolicy{
		IpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
	}
}

func (e *IPSecPolicy) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceIpSecPolicy{
		InterfaceIpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
	}
}

func (e *IPSecPolicy) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectIpSecPolicy{
		ObjectIpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
	}
}

func (e *IPSecPolicy) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectIpSecPolicy{
		AssignedObjectIpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
	}
}

func (e *IPSecPolicy) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationIpSecPolicy{
		TerminationIpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
	}
}

func (e *IPSecPolicy) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectIpSecPolicy{
		AssignedObjectIpSecPolicy: e.ConvertToProtoMessage().(*pb.IPSecPolicy),
	}
}

// implementation of oneof interfaces for IPSecProfile.

func (e *IPSecProfile) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectIpSecProfile{
		ObjectIpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
	}
}

func (e *IPSecProfile) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_IpSecProfile{
		IpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
	}
}

func (e *IPSecProfile) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceIpSecProfile{
		InterfaceIpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
	}
}

func (e *IPSecProfile) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectIpSecProfile{
		ObjectIpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
	}
}

func (e *IPSecProfile) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectIpSecProfile{
		AssignedObjectIpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
	}
}

func (e *IPSecProfile) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationIpSecProfile{
		TerminationIpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
	}
}

func (e *IPSecProfile) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectIpSecProfile{
		AssignedObjectIpSecProfile: e.ConvertToProtoMessage().(*pb.IPSecProfile),
	}
}

// implementation of oneof interfaces for IPSecProposal.

func (e *IPSecProposal) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectIpSecProposal{
		ObjectIpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
	}
}

func (e *IPSecProposal) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_IpSecProposal{
		IpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
	}
}

func (e *IPSecProposal) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceIpSecProposal{
		InterfaceIpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
	}
}

func (e *IPSecProposal) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectIpSecProposal{
		ObjectIpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
	}
}

func (e *IPSecProposal) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectIpSecProposal{
		AssignedObjectIpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
	}
}

func (e *IPSecProposal) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationIpSecProposal{
		TerminationIpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
	}
}

func (e *IPSecProposal) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectIpSecProposal{
		AssignedObjectIpSecProposal: e.ConvertToProtoMessage().(*pb.IPSecProposal),
	}
}

// implementation of oneof interfaces for Interface.

func (e *Interface) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationInterface{
		TerminationInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectInterface{
		ObjectInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Interface{
		Interface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceInterface{
		InterfaceInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectInterface{
		ObjectInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyIPAddressAssignedObjectValueApplyTo(p *pb.IPAddress) {
	p.AssignedObject = &pb.IPAddress_AssignedObjectInterface{
		AssignedObjectInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyInventoryItemComponentValueApplyTo(p *pb.InventoryItem) {
	p.Component = &pb.InventoryItem_ComponentInterface{
		ComponentInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectInterface{
		AssignedObjectInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyMACAddressAssignedObjectValueApplyTo(p *pb.MACAddress) {
	p.AssignedObject = &pb.MACAddress_AssignedObjectInterface{
		AssignedObjectInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationInterface{
		TerminationInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

func (e *Interface) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectInterface{
		AssignedObjectInterface: e.ConvertToProtoMessage().(*pb.Interface),
	}
}

// implementation of oneof interfaces for InventoryItem.

func (e *InventoryItem) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectInventoryItem{
		ObjectInventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
	}
}

func (e *InventoryItem) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_InventoryItem{
		InventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
	}
}

func (e *InventoryItem) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceInventoryItem{
		InterfaceInventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
	}
}

func (e *InventoryItem) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectInventoryItem{
		ObjectInventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
	}
}

func (e *InventoryItem) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectInventoryItem{
		AssignedObjectInventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
	}
}

func (e *InventoryItem) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationInventoryItem{
		TerminationInventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
	}
}

func (e *InventoryItem) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectInventoryItem{
		AssignedObjectInventoryItem: e.ConvertToProtoMessage().(*pb.InventoryItem),
	}
}

// implementation of oneof interfaces for InventoryItemRole.

func (e *InventoryItemRole) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectInventoryItemRole{
		ObjectInventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
	}
}

func (e *InventoryItemRole) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_InventoryItemRole{
		InventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
	}
}

func (e *InventoryItemRole) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceInventoryItemRole{
		InterfaceInventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
	}
}

func (e *InventoryItemRole) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectInventoryItemRole{
		ObjectInventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
	}
}

func (e *InventoryItemRole) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectInventoryItemRole{
		AssignedObjectInventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
	}
}

func (e *InventoryItemRole) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationInventoryItemRole{
		TerminationInventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
	}
}

func (e *InventoryItemRole) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectInventoryItemRole{
		AssignedObjectInventoryItemRole: e.ConvertToProtoMessage().(*pb.InventoryItemRole),
	}
}

// implementation of oneof interfaces for L2VPN.

func (e *L2VPN) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectL2Vpn{
		ObjectL2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
	}
}

func (e *L2VPN) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_L2Vpn{
		L2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
	}
}

func (e *L2VPN) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceL2Vpn{
		InterfaceL2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
	}
}

func (e *L2VPN) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectL2Vpn{
		ObjectL2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
	}
}

func (e *L2VPN) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectL2Vpn{
		AssignedObjectL2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
	}
}

func (e *L2VPN) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationL2Vpn{
		TerminationL2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
	}
}

func (e *L2VPN) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectL2Vpn{
		AssignedObjectL2Vpn: e.ConvertToProtoMessage().(*pb.L2VPN),
	}
}

// implementation of oneof interfaces for L2VPNTermination.

func (e *L2VPNTermination) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectL2VpnTermination{
		ObjectL2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
	}
}

func (e *L2VPNTermination) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_L2VpnTermination{
		L2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
	}
}

func (e *L2VPNTermination) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceL2VpnTermination{
		InterfaceL2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
	}
}

func (e *L2VPNTermination) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectL2VpnTermination{
		ObjectL2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
	}
}

func (e *L2VPNTermination) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectL2VpnTermination{
		AssignedObjectL2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
	}
}

func (e *L2VPNTermination) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationL2VpnTermination{
		TerminationL2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
	}
}

func (e *L2VPNTermination) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectL2VpnTermination{
		AssignedObjectL2VpnTermination: e.ConvertToProtoMessage().(*pb.L2VPNTermination),
	}
}

// implementation of oneof interfaces for Location.

func (e *Location) anyCircuitTerminationTerminationValueApplyTo(p *pb.CircuitTermination) {
	p.Termination = &pb.CircuitTermination_TerminationLocation{
		TerminationLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyClusterScopeValueApplyTo(p *pb.Cluster) {
	p.Scope = &pb.Cluster_ScopeLocation{
		ScopeLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectLocation{
		ObjectLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Location{
		Location: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceLocation{
		InterfaceLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectLocation{
		ObjectLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectLocation{
		AssignedObjectLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyPrefixScopeValueApplyTo(p *pb.Prefix) {
	p.Scope = &pb.Prefix_ScopeLocation{
		ScopeLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationLocation{
		TerminationLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeLocation{
		ScopeLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyWirelessLANScopeValueApplyTo(p *pb.WirelessLAN) {
	p.Scope = &pb.WirelessLAN_ScopeLocation{
		ScopeLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

func (e *Location) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectLocation{
		AssignedObjectLocation: e.ConvertToProtoMessage().(*pb.Location),
	}
}

// implementation of oneof interfaces for MACAddress.

func (e *MACAddress) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectMacAddress{
		ObjectMacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
	}
}

func (e *MACAddress) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_MacAddress{
		MacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
	}
}

func (e *MACAddress) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceMacAddress{
		InterfaceMacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
	}
}

func (e *MACAddress) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectMacAddress{
		ObjectMacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
	}
}

func (e *MACAddress) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectMacAddress{
		AssignedObjectMacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
	}
}

func (e *MACAddress) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationMacAddress{
		TerminationMacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
	}
}

func (e *MACAddress) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectMacAddress{
		AssignedObjectMacAddress: e.ConvertToProtoMessage().(*pb.MACAddress),
	}
}

// implementation of oneof interfaces for Manufacturer.

func (e *Manufacturer) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectManufacturer{
		ObjectManufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
	}
}

func (e *Manufacturer) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Manufacturer{
		Manufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
	}
}

func (e *Manufacturer) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceManufacturer{
		InterfaceManufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
	}
}

func (e *Manufacturer) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectManufacturer{
		ObjectManufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
	}
}

func (e *Manufacturer) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectManufacturer{
		AssignedObjectManufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
	}
}

func (e *Manufacturer) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationManufacturer{
		TerminationManufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
	}
}

func (e *Manufacturer) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectManufacturer{
		AssignedObjectManufacturer: e.ConvertToProtoMessage().(*pb.Manufacturer),
	}
}

// implementation of oneof interfaces for Module.

func (e *Module) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectModule{
		ObjectModule: e.ConvertToProtoMessage().(*pb.Module),
	}
}

func (e *Module) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Module{
		Module: e.ConvertToProtoMessage().(*pb.Module),
	}
}

func (e *Module) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceModule{
		InterfaceModule: e.ConvertToProtoMessage().(*pb.Module),
	}
}

func (e *Module) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectModule{
		ObjectModule: e.ConvertToProtoMessage().(*pb.Module),
	}
}

func (e *Module) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectModule{
		AssignedObjectModule: e.ConvertToProtoMessage().(*pb.Module),
	}
}

func (e *Module) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationModule{
		TerminationModule: e.ConvertToProtoMessage().(*pb.Module),
	}
}

func (e *Module) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectModule{
		AssignedObjectModule: e.ConvertToProtoMessage().(*pb.Module),
	}
}

// implementation of oneof interfaces for ModuleBay.

func (e *ModuleBay) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectModuleBay{
		ObjectModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
	}
}

func (e *ModuleBay) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ModuleBay{
		ModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
	}
}

func (e *ModuleBay) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceModuleBay{
		InterfaceModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
	}
}

func (e *ModuleBay) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectModuleBay{
		ObjectModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
	}
}

func (e *ModuleBay) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectModuleBay{
		AssignedObjectModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
	}
}

func (e *ModuleBay) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationModuleBay{
		TerminationModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
	}
}

func (e *ModuleBay) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectModuleBay{
		AssignedObjectModuleBay: e.ConvertToProtoMessage().(*pb.ModuleBay),
	}
}

// implementation of oneof interfaces for ModuleType.

func (e *ModuleType) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectModuleType{
		ObjectModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
	}
}

func (e *ModuleType) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ModuleType{
		ModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
	}
}

func (e *ModuleType) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceModuleType{
		InterfaceModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
	}
}

func (e *ModuleType) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectModuleType{
		ObjectModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
	}
}

func (e *ModuleType) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectModuleType{
		AssignedObjectModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
	}
}

func (e *ModuleType) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationModuleType{
		TerminationModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
	}
}

func (e *ModuleType) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectModuleType{
		AssignedObjectModuleType: e.ConvertToProtoMessage().(*pb.ModuleType),
	}
}

// implementation of oneof interfaces for Platform.

func (e *Platform) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectPlatform{
		ObjectPlatform: e.ConvertToProtoMessage().(*pb.Platform),
	}
}

func (e *Platform) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Platform{
		Platform: e.ConvertToProtoMessage().(*pb.Platform),
	}
}

func (e *Platform) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfacePlatform{
		InterfacePlatform: e.ConvertToProtoMessage().(*pb.Platform),
	}
}

func (e *Platform) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectPlatform{
		ObjectPlatform: e.ConvertToProtoMessage().(*pb.Platform),
	}
}

func (e *Platform) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectPlatform{
		AssignedObjectPlatform: e.ConvertToProtoMessage().(*pb.Platform),
	}
}

func (e *Platform) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationPlatform{
		TerminationPlatform: e.ConvertToProtoMessage().(*pb.Platform),
	}
}

func (e *Platform) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectPlatform{
		AssignedObjectPlatform: e.ConvertToProtoMessage().(*pb.Platform),
	}
}

// implementation of oneof interfaces for PowerFeed.

func (e *PowerFeed) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationPowerFeed{
		TerminationPowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

func (e *PowerFeed) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectPowerFeed{
		ObjectPowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

func (e *PowerFeed) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_PowerFeed{
		PowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

func (e *PowerFeed) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfacePowerFeed{
		InterfacePowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

func (e *PowerFeed) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectPowerFeed{
		ObjectPowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

func (e *PowerFeed) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectPowerFeed{
		AssignedObjectPowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

func (e *PowerFeed) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationPowerFeed{
		TerminationPowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

func (e *PowerFeed) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectPowerFeed{
		AssignedObjectPowerFeed: e.ConvertToProtoMessage().(*pb.PowerFeed),
	}
}

// implementation of oneof interfaces for PowerOutlet.

func (e *PowerOutlet) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationPowerOutlet{
		TerminationPowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectPowerOutlet{
		ObjectPowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_PowerOutlet{
		PowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfacePowerOutlet{
		InterfacePowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectPowerOutlet{
		ObjectPowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyInventoryItemComponentValueApplyTo(p *pb.InventoryItem) {
	p.Component = &pb.InventoryItem_ComponentPowerOutlet{
		ComponentPowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectPowerOutlet{
		AssignedObjectPowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationPowerOutlet{
		TerminationPowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

func (e *PowerOutlet) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectPowerOutlet{
		AssignedObjectPowerOutlet: e.ConvertToProtoMessage().(*pb.PowerOutlet),
	}
}

// implementation of oneof interfaces for PowerPanel.

func (e *PowerPanel) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectPowerPanel{
		ObjectPowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
	}
}

func (e *PowerPanel) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_PowerPanel{
		PowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
	}
}

func (e *PowerPanel) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfacePowerPanel{
		InterfacePowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
	}
}

func (e *PowerPanel) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectPowerPanel{
		ObjectPowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
	}
}

func (e *PowerPanel) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectPowerPanel{
		AssignedObjectPowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
	}
}

func (e *PowerPanel) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationPowerPanel{
		TerminationPowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
	}
}

func (e *PowerPanel) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectPowerPanel{
		AssignedObjectPowerPanel: e.ConvertToProtoMessage().(*pb.PowerPanel),
	}
}

// implementation of oneof interfaces for PowerPort.

func (e *PowerPort) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationPowerPort{
		TerminationPowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectPowerPort{
		ObjectPowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_PowerPort{
		PowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfacePowerPort{
		InterfacePowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectPowerPort{
		ObjectPowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyInventoryItemComponentValueApplyTo(p *pb.InventoryItem) {
	p.Component = &pb.InventoryItem_ComponentPowerPort{
		ComponentPowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectPowerPort{
		AssignedObjectPowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationPowerPort{
		TerminationPowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

func (e *PowerPort) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectPowerPort{
		AssignedObjectPowerPort: e.ConvertToProtoMessage().(*pb.PowerPort),
	}
}

// implementation of oneof interfaces for Prefix.

func (e *Prefix) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectPrefix{
		ObjectPrefix: e.ConvertToProtoMessage().(*pb.Prefix),
	}
}

func (e *Prefix) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Prefix{
		Prefix: e.ConvertToProtoMessage().(*pb.Prefix),
	}
}

func (e *Prefix) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfacePrefix{
		InterfacePrefix: e.ConvertToProtoMessage().(*pb.Prefix),
	}
}

func (e *Prefix) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectPrefix{
		ObjectPrefix: e.ConvertToProtoMessage().(*pb.Prefix),
	}
}

func (e *Prefix) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectPrefix{
		AssignedObjectPrefix: e.ConvertToProtoMessage().(*pb.Prefix),
	}
}

func (e *Prefix) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationPrefix{
		TerminationPrefix: e.ConvertToProtoMessage().(*pb.Prefix),
	}
}

func (e *Prefix) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectPrefix{
		AssignedObjectPrefix: e.ConvertToProtoMessage().(*pb.Prefix),
	}
}

// implementation of oneof interfaces for Provider.

func (e *Provider) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectProvider{
		ObjectProvider: e.ConvertToProtoMessage().(*pb.Provider),
	}
}

func (e *Provider) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Provider{
		Provider: e.ConvertToProtoMessage().(*pb.Provider),
	}
}

func (e *Provider) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceProvider{
		InterfaceProvider: e.ConvertToProtoMessage().(*pb.Provider),
	}
}

func (e *Provider) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectProvider{
		ObjectProvider: e.ConvertToProtoMessage().(*pb.Provider),
	}
}

func (e *Provider) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectProvider{
		AssignedObjectProvider: e.ConvertToProtoMessage().(*pb.Provider),
	}
}

func (e *Provider) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationProvider{
		TerminationProvider: e.ConvertToProtoMessage().(*pb.Provider),
	}
}

func (e *Provider) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectProvider{
		AssignedObjectProvider: e.ConvertToProtoMessage().(*pb.Provider),
	}
}

// implementation of oneof interfaces for ProviderAccount.

func (e *ProviderAccount) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectProviderAccount{
		ObjectProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
	}
}

func (e *ProviderAccount) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ProviderAccount{
		ProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
	}
}

func (e *ProviderAccount) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceProviderAccount{
		InterfaceProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
	}
}

func (e *ProviderAccount) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectProviderAccount{
		ObjectProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
	}
}

func (e *ProviderAccount) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectProviderAccount{
		AssignedObjectProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
	}
}

func (e *ProviderAccount) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationProviderAccount{
		TerminationProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
	}
}

func (e *ProviderAccount) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectProviderAccount{
		AssignedObjectProviderAccount: e.ConvertToProtoMessage().(*pb.ProviderAccount),
	}
}

// implementation of oneof interfaces for ProviderNetwork.

func (e *ProviderNetwork) anyCircuitTerminationTerminationValueApplyTo(p *pb.CircuitTermination) {
	p.Termination = &pb.CircuitTermination_TerminationProviderNetwork{
		TerminationProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

func (e *ProviderNetwork) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectProviderNetwork{
		ObjectProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

func (e *ProviderNetwork) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ProviderNetwork{
		ProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

func (e *ProviderNetwork) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceProviderNetwork{
		InterfaceProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

func (e *ProviderNetwork) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectProviderNetwork{
		ObjectProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

func (e *ProviderNetwork) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectProviderNetwork{
		AssignedObjectProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

func (e *ProviderNetwork) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationProviderNetwork{
		TerminationProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

func (e *ProviderNetwork) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectProviderNetwork{
		AssignedObjectProviderNetwork: e.ConvertToProtoMessage().(*pb.ProviderNetwork),
	}
}

// implementation of oneof interfaces for RIR.

func (e *RIR) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRir{
		ObjectRir: e.ConvertToProtoMessage().(*pb.RIR),
	}
}

func (e *RIR) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Rir{
		Rir: e.ConvertToProtoMessage().(*pb.RIR),
	}
}

func (e *RIR) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRir{
		InterfaceRir: e.ConvertToProtoMessage().(*pb.RIR),
	}
}

func (e *RIR) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRir{
		ObjectRir: e.ConvertToProtoMessage().(*pb.RIR),
	}
}

func (e *RIR) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRir{
		AssignedObjectRir: e.ConvertToProtoMessage().(*pb.RIR),
	}
}

func (e *RIR) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRir{
		TerminationRir: e.ConvertToProtoMessage().(*pb.RIR),
	}
}

func (e *RIR) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRir{
		AssignedObjectRir: e.ConvertToProtoMessage().(*pb.RIR),
	}
}

// implementation of oneof interfaces for Rack.

func (e *Rack) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRack{
		ObjectRack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

func (e *Rack) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Rack{
		Rack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

func (e *Rack) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRack{
		InterfaceRack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

func (e *Rack) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRack{
		ObjectRack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

func (e *Rack) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRack{
		AssignedObjectRack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

func (e *Rack) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRack{
		TerminationRack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

func (e *Rack) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeRack{
		ScopeRack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

func (e *Rack) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRack{
		AssignedObjectRack: e.ConvertToProtoMessage().(*pb.Rack),
	}
}

// implementation of oneof interfaces for RackReservation.

func (e *RackReservation) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRackReservation{
		ObjectRackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
	}
}

func (e *RackReservation) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_RackReservation{
		RackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
	}
}

func (e *RackReservation) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRackReservation{
		InterfaceRackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
	}
}

func (e *RackReservation) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRackReservation{
		ObjectRackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
	}
}

func (e *RackReservation) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRackReservation{
		AssignedObjectRackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
	}
}

func (e *RackReservation) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRackReservation{
		TerminationRackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
	}
}

func (e *RackReservation) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRackReservation{
		AssignedObjectRackReservation: e.ConvertToProtoMessage().(*pb.RackReservation),
	}
}

// implementation of oneof interfaces for RackRole.

func (e *RackRole) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRackRole{
		ObjectRackRole: e.ConvertToProtoMessage().(*pb.RackRole),
	}
}

func (e *RackRole) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_RackRole{
		RackRole: e.ConvertToProtoMessage().(*pb.RackRole),
	}
}

func (e *RackRole) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRackRole{
		InterfaceRackRole: e.ConvertToProtoMessage().(*pb.RackRole),
	}
}

func (e *RackRole) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRackRole{
		ObjectRackRole: e.ConvertToProtoMessage().(*pb.RackRole),
	}
}

func (e *RackRole) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRackRole{
		AssignedObjectRackRole: e.ConvertToProtoMessage().(*pb.RackRole),
	}
}

func (e *RackRole) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRackRole{
		TerminationRackRole: e.ConvertToProtoMessage().(*pb.RackRole),
	}
}

func (e *RackRole) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRackRole{
		AssignedObjectRackRole: e.ConvertToProtoMessage().(*pb.RackRole),
	}
}

// implementation of oneof interfaces for RackType.

func (e *RackType) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRackType{
		ObjectRackType: e.ConvertToProtoMessage().(*pb.RackType),
	}
}

func (e *RackType) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_RackType{
		RackType: e.ConvertToProtoMessage().(*pb.RackType),
	}
}

func (e *RackType) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRackType{
		InterfaceRackType: e.ConvertToProtoMessage().(*pb.RackType),
	}
}

func (e *RackType) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRackType{
		ObjectRackType: e.ConvertToProtoMessage().(*pb.RackType),
	}
}

func (e *RackType) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRackType{
		AssignedObjectRackType: e.ConvertToProtoMessage().(*pb.RackType),
	}
}

func (e *RackType) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRackType{
		TerminationRackType: e.ConvertToProtoMessage().(*pb.RackType),
	}
}

func (e *RackType) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRackType{
		AssignedObjectRackType: e.ConvertToProtoMessage().(*pb.RackType),
	}
}

// implementation of oneof interfaces for RearPort.

func (e *RearPort) anyCableTerminationTerminationValueApplyTo(p *pb.CableTermination) {
	p.Termination = &pb.CableTermination_TerminationRearPort{
		TerminationRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRearPort{
		ObjectRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_RearPort{
		RearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRearPort{
		InterfaceRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRearPort{
		ObjectRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyInventoryItemComponentValueApplyTo(p *pb.InventoryItem) {
	p.Component = &pb.InventoryItem_ComponentRearPort{
		ComponentRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRearPort{
		AssignedObjectRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRearPort{
		TerminationRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

func (e *RearPort) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRearPort{
		AssignedObjectRearPort: e.ConvertToProtoMessage().(*pb.RearPort),
	}
}

// implementation of oneof interfaces for Region.

func (e *Region) anyCircuitTerminationTerminationValueApplyTo(p *pb.CircuitTermination) {
	p.Termination = &pb.CircuitTermination_TerminationRegion{
		TerminationRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyClusterScopeValueApplyTo(p *pb.Cluster) {
	p.Scope = &pb.Cluster_ScopeRegion{
		ScopeRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRegion{
		ObjectRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Region{
		Region: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRegion{
		InterfaceRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRegion{
		ObjectRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRegion{
		AssignedObjectRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyPrefixScopeValueApplyTo(p *pb.Prefix) {
	p.Scope = &pb.Prefix_ScopeRegion{
		ScopeRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRegion{
		TerminationRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeRegion{
		ScopeRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyWirelessLANScopeValueApplyTo(p *pb.WirelessLAN) {
	p.Scope = &pb.WirelessLAN_ScopeRegion{
		ScopeRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

func (e *Region) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRegion{
		AssignedObjectRegion: e.ConvertToProtoMessage().(*pb.Region),
	}
}

// implementation of oneof interfaces for Role.

func (e *Role) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRole{
		ObjectRole: e.ConvertToProtoMessage().(*pb.Role),
	}
}

func (e *Role) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Role{
		Role: e.ConvertToProtoMessage().(*pb.Role),
	}
}

func (e *Role) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRole{
		InterfaceRole: e.ConvertToProtoMessage().(*pb.Role),
	}
}

func (e *Role) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRole{
		ObjectRole: e.ConvertToProtoMessage().(*pb.Role),
	}
}

func (e *Role) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRole{
		AssignedObjectRole: e.ConvertToProtoMessage().(*pb.Role),
	}
}

func (e *Role) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRole{
		TerminationRole: e.ConvertToProtoMessage().(*pb.Role),
	}
}

func (e *Role) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRole{
		AssignedObjectRole: e.ConvertToProtoMessage().(*pb.Role),
	}
}

// implementation of oneof interfaces for RouteTarget.

func (e *RouteTarget) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRouteTarget{
		ObjectRouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
	}
}

func (e *RouteTarget) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_RouteTarget{
		RouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
	}
}

func (e *RouteTarget) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRouteTarget{
		InterfaceRouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
	}
}

func (e *RouteTarget) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRouteTarget{
		ObjectRouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
	}
}

func (e *RouteTarget) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRouteTarget{
		AssignedObjectRouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
	}
}

func (e *RouteTarget) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRouteTarget{
		TerminationRouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
	}
}

func (e *RouteTarget) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRouteTarget{
		AssignedObjectRouteTarget: e.ConvertToProtoMessage().(*pb.RouteTarget),
	}
}

// implementation of oneof interfaces for Service.

func (e *Service) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectService{
		ObjectService: e.ConvertToProtoMessage().(*pb.Service),
	}
}

func (e *Service) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Service{
		Service: e.ConvertToProtoMessage().(*pb.Service),
	}
}

func (e *Service) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceService{
		InterfaceService: e.ConvertToProtoMessage().(*pb.Service),
	}
}

func (e *Service) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectService{
		ObjectService: e.ConvertToProtoMessage().(*pb.Service),
	}
}

func (e *Service) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectService{
		AssignedObjectService: e.ConvertToProtoMessage().(*pb.Service),
	}
}

func (e *Service) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationService{
		TerminationService: e.ConvertToProtoMessage().(*pb.Service),
	}
}

func (e *Service) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectService{
		AssignedObjectService: e.ConvertToProtoMessage().(*pb.Service),
	}
}

// implementation of oneof interfaces for Site.

func (e *Site) anyCircuitTerminationTerminationValueApplyTo(p *pb.CircuitTermination) {
	p.Termination = &pb.CircuitTermination_TerminationSite{
		TerminationSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyClusterScopeValueApplyTo(p *pb.Cluster) {
	p.Scope = &pb.Cluster_ScopeSite{
		ScopeSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectSite{
		ObjectSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Site{
		Site: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceSite{
		InterfaceSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectSite{
		ObjectSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectSite{
		AssignedObjectSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyPrefixScopeValueApplyTo(p *pb.Prefix) {
	p.Scope = &pb.Prefix_ScopeSite{
		ScopeSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationSite{
		TerminationSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeSite{
		ScopeSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyWirelessLANScopeValueApplyTo(p *pb.WirelessLAN) {
	p.Scope = &pb.WirelessLAN_ScopeSite{
		ScopeSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

func (e *Site) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectSite{
		AssignedObjectSite: e.ConvertToProtoMessage().(*pb.Site),
	}
}

// implementation of oneof interfaces for SiteGroup.

func (e *SiteGroup) anyCircuitTerminationTerminationValueApplyTo(p *pb.CircuitTermination) {
	p.Termination = &pb.CircuitTermination_TerminationSiteGroup{
		TerminationSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyClusterScopeValueApplyTo(p *pb.Cluster) {
	p.Scope = &pb.Cluster_ScopeSiteGroup{
		ScopeSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectSiteGroup{
		ObjectSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_SiteGroup{
		SiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceSiteGroup{
		InterfaceSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectSiteGroup{
		ObjectSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectSiteGroup{
		AssignedObjectSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyPrefixScopeValueApplyTo(p *pb.Prefix) {
	p.Scope = &pb.Prefix_ScopeSiteGroup{
		ScopeSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationSiteGroup{
		TerminationSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeSiteGroup{
		ScopeSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyWirelessLANScopeValueApplyTo(p *pb.WirelessLAN) {
	p.Scope = &pb.WirelessLAN_ScopeSiteGroup{
		ScopeSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

func (e *SiteGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectSiteGroup{
		AssignedObjectSiteGroup: e.ConvertToProtoMessage().(*pb.SiteGroup),
	}
}

// implementation of oneof interfaces for Tag.

func (e *Tag) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectTag{
		ObjectTag: e.ConvertToProtoMessage().(*pb.Tag),
	}
}

func (e *Tag) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Tag{
		Tag: e.ConvertToProtoMessage().(*pb.Tag),
	}
}

func (e *Tag) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceTag{
		InterfaceTag: e.ConvertToProtoMessage().(*pb.Tag),
	}
}

func (e *Tag) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectTag{
		ObjectTag: e.ConvertToProtoMessage().(*pb.Tag),
	}
}

func (e *Tag) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectTag{
		AssignedObjectTag: e.ConvertToProtoMessage().(*pb.Tag),
	}
}

func (e *Tag) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationTag{
		TerminationTag: e.ConvertToProtoMessage().(*pb.Tag),
	}
}

func (e *Tag) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectTag{
		AssignedObjectTag: e.ConvertToProtoMessage().(*pb.Tag),
	}
}

// implementation of oneof interfaces for Tenant.

func (e *Tenant) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectTenant{
		ObjectTenant: e.ConvertToProtoMessage().(*pb.Tenant),
	}
}

func (e *Tenant) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Tenant{
		Tenant: e.ConvertToProtoMessage().(*pb.Tenant),
	}
}

func (e *Tenant) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceTenant{
		InterfaceTenant: e.ConvertToProtoMessage().(*pb.Tenant),
	}
}

func (e *Tenant) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectTenant{
		ObjectTenant: e.ConvertToProtoMessage().(*pb.Tenant),
	}
}

func (e *Tenant) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectTenant{
		AssignedObjectTenant: e.ConvertToProtoMessage().(*pb.Tenant),
	}
}

func (e *Tenant) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationTenant{
		TerminationTenant: e.ConvertToProtoMessage().(*pb.Tenant),
	}
}

func (e *Tenant) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectTenant{
		AssignedObjectTenant: e.ConvertToProtoMessage().(*pb.Tenant),
	}
}

// implementation of oneof interfaces for TenantGroup.

func (e *TenantGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectTenantGroup{
		ObjectTenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
	}
}

func (e *TenantGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_TenantGroup{
		TenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
	}
}

func (e *TenantGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceTenantGroup{
		InterfaceTenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
	}
}

func (e *TenantGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectTenantGroup{
		ObjectTenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
	}
}

func (e *TenantGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectTenantGroup{
		AssignedObjectTenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
	}
}

func (e *TenantGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationTenantGroup{
		TerminationTenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
	}
}

func (e *TenantGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectTenantGroup{
		AssignedObjectTenantGroup: e.ConvertToProtoMessage().(*pb.TenantGroup),
	}
}

// implementation of oneof interfaces for Tunnel.

func (e *Tunnel) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectTunnel{
		ObjectTunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
	}
}

func (e *Tunnel) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Tunnel{
		Tunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
	}
}

func (e *Tunnel) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceTunnel{
		InterfaceTunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
	}
}

func (e *Tunnel) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectTunnel{
		ObjectTunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
	}
}

func (e *Tunnel) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectTunnel{
		AssignedObjectTunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
	}
}

func (e *Tunnel) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationTunnel{
		TerminationTunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
	}
}

func (e *Tunnel) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectTunnel{
		AssignedObjectTunnel: e.ConvertToProtoMessage().(*pb.Tunnel),
	}
}

// implementation of oneof interfaces for TunnelGroup.

func (e *TunnelGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectTunnelGroup{
		ObjectTunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
	}
}

func (e *TunnelGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_TunnelGroup{
		TunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
	}
}

func (e *TunnelGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceTunnelGroup{
		InterfaceTunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
	}
}

func (e *TunnelGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectTunnelGroup{
		ObjectTunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
	}
}

func (e *TunnelGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectTunnelGroup{
		AssignedObjectTunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
	}
}

func (e *TunnelGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationTunnelGroup{
		TerminationTunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
	}
}

func (e *TunnelGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectTunnelGroup{
		AssignedObjectTunnelGroup: e.ConvertToProtoMessage().(*pb.TunnelGroup),
	}
}

// implementation of oneof interfaces for TunnelTermination.

func (e *TunnelTermination) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectTunnelTermination{
		ObjectTunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
	}
}

func (e *TunnelTermination) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_TunnelTermination{
		TunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
	}
}

func (e *TunnelTermination) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceTunnelTermination{
		InterfaceTunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
	}
}

func (e *TunnelTermination) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectTunnelTermination{
		ObjectTunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
	}
}

func (e *TunnelTermination) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectTunnelTermination{
		AssignedObjectTunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
	}
}

func (e *TunnelTermination) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationTunnelTermination{
		TerminationTunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
	}
}

func (e *TunnelTermination) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectTunnelTermination{
		AssignedObjectTunnelTermination: e.ConvertToProtoMessage().(*pb.TunnelTermination),
	}
}

// implementation of oneof interfaces for VLAN.

func (e *VLAN) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVlan{
		ObjectVlan: e.ConvertToProtoMessage().(*pb.VLAN),
	}
}

func (e *VLAN) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Vlan{
		Vlan: e.ConvertToProtoMessage().(*pb.VLAN),
	}
}

func (e *VLAN) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVlan{
		InterfaceVlan: e.ConvertToProtoMessage().(*pb.VLAN),
	}
}

func (e *VLAN) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVlan{
		ObjectVlan: e.ConvertToProtoMessage().(*pb.VLAN),
	}
}

func (e *VLAN) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVlan{
		AssignedObjectVlan: e.ConvertToProtoMessage().(*pb.VLAN),
	}
}

func (e *VLAN) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVlan{
		TerminationVlan: e.ConvertToProtoMessage().(*pb.VLAN),
	}
}

func (e *VLAN) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVlan{
		AssignedObjectVlan: e.ConvertToProtoMessage().(*pb.VLAN),
	}
}

// implementation of oneof interfaces for VLANGroup.

func (e *VLANGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVlanGroup{
		ObjectVlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
	}
}

func (e *VLANGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VlanGroup{
		VlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
	}
}

func (e *VLANGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVlanGroup{
		InterfaceVlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
	}
}

func (e *VLANGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVlanGroup{
		ObjectVlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
	}
}

func (e *VLANGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVlanGroup{
		AssignedObjectVlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
	}
}

func (e *VLANGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVlanGroup{
		TerminationVlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
	}
}

func (e *VLANGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVlanGroup{
		AssignedObjectVlanGroup: e.ConvertToProtoMessage().(*pb.VLANGroup),
	}
}

// implementation of oneof interfaces for VLANTranslationPolicy.

func (e *VLANTranslationPolicy) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVlanTranslationPolicy{
		ObjectVlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
	}
}

func (e *VLANTranslationPolicy) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VlanTranslationPolicy{
		VlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
	}
}

func (e *VLANTranslationPolicy) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVlanTranslationPolicy{
		InterfaceVlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
	}
}

func (e *VLANTranslationPolicy) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVlanTranslationPolicy{
		ObjectVlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
	}
}

func (e *VLANTranslationPolicy) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVlanTranslationPolicy{
		AssignedObjectVlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
	}
}

func (e *VLANTranslationPolicy) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVlanTranslationPolicy{
		TerminationVlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
	}
}

func (e *VLANTranslationPolicy) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVlanTranslationPolicy{
		AssignedObjectVlanTranslationPolicy: e.ConvertToProtoMessage().(*pb.VLANTranslationPolicy),
	}
}

// implementation of oneof interfaces for VLANTranslationRule.

func (e *VLANTranslationRule) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVlanTranslationRule{
		ObjectVlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
	}
}

func (e *VLANTranslationRule) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VlanTranslationRule{
		VlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
	}
}

func (e *VLANTranslationRule) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVlanTranslationRule{
		InterfaceVlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
	}
}

func (e *VLANTranslationRule) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVlanTranslationRule{
		ObjectVlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
	}
}

func (e *VLANTranslationRule) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVlanTranslationRule{
		AssignedObjectVlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
	}
}

func (e *VLANTranslationRule) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVlanTranslationRule{
		TerminationVlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
	}
}

func (e *VLANTranslationRule) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVlanTranslationRule{
		AssignedObjectVlanTranslationRule: e.ConvertToProtoMessage().(*pb.VLANTranslationRule),
	}
}

// implementation of oneof interfaces for VMInterface.

func (e *VMInterface) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVmInterface{
		ObjectVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VmInterface{
		VmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVmInterface{
		InterfaceVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVmInterface{
		ObjectVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyIPAddressAssignedObjectValueApplyTo(p *pb.IPAddress) {
	p.AssignedObject = &pb.IPAddress_AssignedObjectVmInterface{
		AssignedObjectVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVmInterface{
		AssignedObjectVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyMACAddressAssignedObjectValueApplyTo(p *pb.MACAddress) {
	p.AssignedObject = &pb.MACAddress_AssignedObjectVmInterface{
		AssignedObjectVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVmInterface{
		TerminationVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

func (e *VMInterface) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVmInterface{
		AssignedObjectVmInterface: e.ConvertToProtoMessage().(*pb.VMInterface),
	}
}

// implementation of oneof interfaces for VRF.

func (e *VRF) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVrf{
		ObjectVrf: e.ConvertToProtoMessage().(*pb.VRF),
	}
}

func (e *VRF) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Vrf{
		Vrf: e.ConvertToProtoMessage().(*pb.VRF),
	}
}

func (e *VRF) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVrf{
		InterfaceVrf: e.ConvertToProtoMessage().(*pb.VRF),
	}
}

func (e *VRF) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVrf{
		ObjectVrf: e.ConvertToProtoMessage().(*pb.VRF),
	}
}

func (e *VRF) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVrf{
		AssignedObjectVrf: e.ConvertToProtoMessage().(*pb.VRF),
	}
}

func (e *VRF) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVrf{
		TerminationVrf: e.ConvertToProtoMessage().(*pb.VRF),
	}
}

func (e *VRF) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVrf{
		AssignedObjectVrf: e.ConvertToProtoMessage().(*pb.VRF),
	}
}

// implementation of oneof interfaces for VirtualChassis.

func (e *VirtualChassis) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualChassis{
		ObjectVirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
	}
}

func (e *VirtualChassis) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualChassis{
		VirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
	}
}

func (e *VirtualChassis) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualChassis{
		InterfaceVirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
	}
}

func (e *VirtualChassis) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualChassis{
		ObjectVirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
	}
}

func (e *VirtualChassis) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualChassis{
		AssignedObjectVirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
	}
}

func (e *VirtualChassis) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualChassis{
		TerminationVirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
	}
}

func (e *VirtualChassis) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualChassis{
		AssignedObjectVirtualChassis: e.ConvertToProtoMessage().(*pb.VirtualChassis),
	}
}

// implementation of oneof interfaces for VirtualCircuit.

func (e *VirtualCircuit) anyCircuitGroupAssignmentMemberValueApplyTo(p *pb.CircuitGroupAssignment) {
	p.Member = &pb.CircuitGroupAssignment_MemberVirtualCircuit{
		MemberVirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

func (e *VirtualCircuit) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualCircuit{
		ObjectVirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

func (e *VirtualCircuit) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualCircuit{
		VirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

func (e *VirtualCircuit) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualCircuit{
		InterfaceVirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

func (e *VirtualCircuit) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualCircuit{
		ObjectVirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

func (e *VirtualCircuit) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualCircuit{
		AssignedObjectVirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

func (e *VirtualCircuit) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualCircuit{
		TerminationVirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

func (e *VirtualCircuit) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualCircuit{
		AssignedObjectVirtualCircuit: e.ConvertToProtoMessage().(*pb.VirtualCircuit),
	}
}

// implementation of oneof interfaces for VirtualCircuitTermination.

func (e *VirtualCircuitTermination) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualCircuitTermination{
		ObjectVirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
	}
}

func (e *VirtualCircuitTermination) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualCircuitTermination{
		VirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
	}
}

func (e *VirtualCircuitTermination) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualCircuitTermination{
		InterfaceVirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
	}
}

func (e *VirtualCircuitTermination) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualCircuitTermination{
		ObjectVirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
	}
}

func (e *VirtualCircuitTermination) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualCircuitTermination{
		AssignedObjectVirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
	}
}

func (e *VirtualCircuitTermination) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualCircuitTermination{
		TerminationVirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
	}
}

func (e *VirtualCircuitTermination) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualCircuitTermination{
		AssignedObjectVirtualCircuitTermination: e.ConvertToProtoMessage().(*pb.VirtualCircuitTermination),
	}
}

// implementation of oneof interfaces for VirtualCircuitType.

func (e *VirtualCircuitType) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualCircuitType{
		ObjectVirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
	}
}

func (e *VirtualCircuitType) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualCircuitType{
		VirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
	}
}

func (e *VirtualCircuitType) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualCircuitType{
		InterfaceVirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
	}
}

func (e *VirtualCircuitType) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualCircuitType{
		ObjectVirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
	}
}

func (e *VirtualCircuitType) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualCircuitType{
		AssignedObjectVirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
	}
}

func (e *VirtualCircuitType) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualCircuitType{
		TerminationVirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
	}
}

func (e *VirtualCircuitType) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualCircuitType{
		AssignedObjectVirtualCircuitType: e.ConvertToProtoMessage().(*pb.VirtualCircuitType),
	}
}

// implementation of oneof interfaces for VirtualDeviceContext.

func (e *VirtualDeviceContext) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualDeviceContext{
		ObjectVirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
	}
}

func (e *VirtualDeviceContext) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualDeviceContext{
		VirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
	}
}

func (e *VirtualDeviceContext) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualDeviceContext{
		InterfaceVirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
	}
}

func (e *VirtualDeviceContext) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualDeviceContext{
		ObjectVirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
	}
}

func (e *VirtualDeviceContext) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualDeviceContext{
		AssignedObjectVirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
	}
}

func (e *VirtualDeviceContext) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualDeviceContext{
		TerminationVirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
	}
}

func (e *VirtualDeviceContext) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualDeviceContext{
		AssignedObjectVirtualDeviceContext: e.ConvertToProtoMessage().(*pb.VirtualDeviceContext),
	}
}

// implementation of oneof interfaces for VirtualDisk.

func (e *VirtualDisk) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualDisk{
		ObjectVirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
	}
}

func (e *VirtualDisk) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualDisk{
		VirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
	}
}

func (e *VirtualDisk) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualDisk{
		InterfaceVirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
	}
}

func (e *VirtualDisk) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualDisk{
		ObjectVirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
	}
}

func (e *VirtualDisk) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualDisk{
		AssignedObjectVirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
	}
}

func (e *VirtualDisk) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualDisk{
		TerminationVirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
	}
}

func (e *VirtualDisk) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualDisk{
		AssignedObjectVirtualDisk: e.ConvertToProtoMessage().(*pb.VirtualDisk),
	}
}

// implementation of oneof interfaces for VirtualMachine.

func (e *VirtualMachine) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualMachine{
		ObjectVirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

func (e *VirtualMachine) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualMachine{
		VirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

func (e *VirtualMachine) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualMachine{
		InterfaceVirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

func (e *VirtualMachine) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualMachine{
		ObjectVirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

func (e *VirtualMachine) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualMachine{
		AssignedObjectVirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

func (e *VirtualMachine) anyServiceParentObjectValueApplyTo(p *pb.Service) {
	p.ParentObject = &pb.Service_ParentObjectVirtualMachine{
		ParentObjectVirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

func (e *VirtualMachine) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualMachine{
		TerminationVirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

func (e *VirtualMachine) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualMachine{
		AssignedObjectVirtualMachine: e.ConvertToProtoMessage().(*pb.VirtualMachine),
	}
}

// implementation of oneof interfaces for WirelessLAN.

func (e *WirelessLAN) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectWirelessLan{
		ObjectWirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
	}
}

func (e *WirelessLAN) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_WirelessLan{
		WirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
	}
}

func (e *WirelessLAN) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceWirelessLan{
		InterfaceWirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
	}
}

func (e *WirelessLAN) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectWirelessLan{
		ObjectWirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
	}
}

func (e *WirelessLAN) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectWirelessLan{
		AssignedObjectWirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
	}
}

func (e *WirelessLAN) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationWirelessLan{
		TerminationWirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
	}
}

func (e *WirelessLAN) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectWirelessLan{
		AssignedObjectWirelessLan: e.ConvertToProtoMessage().(*pb.WirelessLAN),
	}
}

// implementation of oneof interfaces for WirelessLANGroup.

func (e *WirelessLANGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectWirelessLanGroup{
		ObjectWirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
	}
}

func (e *WirelessLANGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_WirelessLanGroup{
		WirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
	}
}

func (e *WirelessLANGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceWirelessLanGroup{
		InterfaceWirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
	}
}

func (e *WirelessLANGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectWirelessLanGroup{
		ObjectWirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
	}
}

func (e *WirelessLANGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectWirelessLanGroup{
		AssignedObjectWirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
	}
}

func (e *WirelessLANGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationWirelessLanGroup{
		TerminationWirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
	}
}

func (e *WirelessLANGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectWirelessLanGroup{
		AssignedObjectWirelessLanGroup: e.ConvertToProtoMessage().(*pb.WirelessLANGroup),
	}
}

// implementation of oneof interfaces for WirelessLink.

func (e *WirelessLink) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectWirelessLink{
		ObjectWirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
	}
}

func (e *WirelessLink) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_WirelessLink{
		WirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
	}
}

func (e *WirelessLink) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceWirelessLink{
		InterfaceWirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
	}
}

func (e *WirelessLink) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectWirelessLink{
		ObjectWirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
	}
}

func (e *WirelessLink) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectWirelessLink{
		AssignedObjectWirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
	}
}

func (e *WirelessLink) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationWirelessLink{
		TerminationWirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
	}
}

func (e *WirelessLink) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectWirelessLink{
		AssignedObjectWirelessLink: e.ConvertToProtoMessage().(*pb.WirelessLink),
	}
}

// implementation of oneof interfaces for CustomField.

func (e *CustomField) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCustomField{
		ObjectCustomField: e.ConvertToProtoMessage().(*pb.CustomField),
	}
}

func (e *CustomField) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CustomField{
		CustomField: e.ConvertToProtoMessage().(*pb.CustomField),
	}
}

func (e *CustomField) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCustomField{
		InterfaceCustomField: e.ConvertToProtoMessage().(*pb.CustomField),
	}
}

func (e *CustomField) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCustomField{
		ObjectCustomField: e.ConvertToProtoMessage().(*pb.CustomField),
	}
}

func (e *CustomField) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCustomField{
		AssignedObjectCustomField: e.ConvertToProtoMessage().(*pb.CustomField),
	}
}

func (e *CustomField) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCustomField{
		TerminationCustomField: e.ConvertToProtoMessage().(*pb.CustomField),
	}
}

func (e *CustomField) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCustomField{
		AssignedObjectCustomField: e.ConvertToProtoMessage().(*pb.CustomField),
	}
}

// implementation of oneof interfaces for CustomFieldChoiceSet.

func (e *CustomFieldChoiceSet) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCustomFieldChoiceSet{
		ObjectCustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
	}
}

func (e *CustomFieldChoiceSet) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CustomFieldChoiceSet{
		CustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
	}
}

func (e *CustomFieldChoiceSet) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCustomFieldChoiceSet{
		InterfaceCustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
	}
}

func (e *CustomFieldChoiceSet) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCustomFieldChoiceSet{
		ObjectCustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
	}
}

func (e *CustomFieldChoiceSet) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCustomFieldChoiceSet{
		AssignedObjectCustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
	}
}

func (e *CustomFieldChoiceSet) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCustomFieldChoiceSet{
		TerminationCustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
	}
}

func (e *CustomFieldChoiceSet) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCustomFieldChoiceSet{
		AssignedObjectCustomFieldChoiceSet: e.ConvertToProtoMessage().(*pb.CustomFieldChoiceSet),
	}
}

// implementation of oneof interfaces for JournalEntry.

func (e *JournalEntry) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectJournalEntry{
		ObjectJournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
	}
}

func (e *JournalEntry) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_JournalEntry{
		JournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
	}
}

func (e *JournalEntry) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceJournalEntry{
		InterfaceJournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
	}
}

func (e *JournalEntry) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectJournalEntry{
		ObjectJournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
	}
}

func (e *JournalEntry) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectJournalEntry{
		AssignedObjectJournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
	}
}

func (e *JournalEntry) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationJournalEntry{
		TerminationJournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
	}
}

func (e *JournalEntry) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectJournalEntry{
		AssignedObjectJournalEntry: e.ConvertToProtoMessage().(*pb.JournalEntry),
	}
}

// implementation of oneof interfaces for ModuleTypeProfile.

func (e *ModuleTypeProfile) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectModuleTypeProfile{
		ObjectModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
	}
}

func (e *ModuleTypeProfile) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ModuleTypeProfile{
		ModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
	}
}

func (e *ModuleTypeProfile) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceModuleTypeProfile{
		InterfaceModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
	}
}

func (e *ModuleTypeProfile) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectModuleTypeProfile{
		ObjectModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
	}
}

func (e *ModuleTypeProfile) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectModuleTypeProfile{
		AssignedObjectModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
	}
}

func (e *ModuleTypeProfile) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationModuleTypeProfile{
		TerminationModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
	}
}

func (e *ModuleTypeProfile) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectModuleTypeProfile{
		AssignedObjectModuleTypeProfile: e.ConvertToProtoMessage().(*pb.ModuleTypeProfile),
	}
}

// implementation of oneof interfaces for CustomLink.

func (e *CustomLink) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCustomLink{
		ObjectCustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
	}
}

func (e *CustomLink) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CustomLink{
		CustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
	}
}

func (e *CustomLink) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCustomLink{
		InterfaceCustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
	}
}

func (e *CustomLink) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCustomLink{
		ObjectCustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
	}
}

func (e *CustomLink) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCustomLink{
		AssignedObjectCustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
	}
}

func (e *CustomLink) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCustomLink{
		TerminationCustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
	}
}

func (e *CustomLink) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCustomLink{
		AssignedObjectCustomLink: e.ConvertToProtoMessage().(*pb.CustomLink),
	}
}

// implementation of oneof interfaces for Owner.

func (e *Owner) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectOwner{
		ObjectOwner: e.ConvertToProtoMessage().(*pb.Owner),
	}
}

func (e *Owner) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_Owner{
		Owner: e.ConvertToProtoMessage().(*pb.Owner),
	}
}

func (e *Owner) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceOwner{
		InterfaceOwner: e.ConvertToProtoMessage().(*pb.Owner),
	}
}

func (e *Owner) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectOwner{
		ObjectOwner: e.ConvertToProtoMessage().(*pb.Owner),
	}
}

func (e *Owner) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectOwner{
		AssignedObjectOwner: e.ConvertToProtoMessage().(*pb.Owner),
	}
}

func (e *Owner) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationOwner{
		TerminationOwner: e.ConvertToProtoMessage().(*pb.Owner),
	}
}

func (e *Owner) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectOwner{
		AssignedObjectOwner: e.ConvertToProtoMessage().(*pb.Owner),
	}
}

// implementation of oneof interfaces for OwnerGroup.

func (e *OwnerGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectOwnerGroup{
		ObjectOwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
	}
}

func (e *OwnerGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_OwnerGroup{
		OwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
	}
}

func (e *OwnerGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceOwnerGroup{
		InterfaceOwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
	}
}

func (e *OwnerGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectOwnerGroup{
		ObjectOwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
	}
}

func (e *OwnerGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectOwnerGroup{
		AssignedObjectOwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
	}
}

func (e *OwnerGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationOwnerGroup{
		TerminationOwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
	}
}

func (e *OwnerGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectOwnerGroup{
		AssignedObjectOwnerGroup: e.ConvertToProtoMessage().(*pb.OwnerGroup),
	}
}

// implementation of oneof interfaces for CableBundle.

func (e *CableBundle) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectCableBundle{
		ObjectCableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
	}
}

func (e *CableBundle) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_CableBundle{
		CableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
	}
}

func (e *CableBundle) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceCableBundle{
		InterfaceCableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
	}
}

func (e *CableBundle) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectCableBundle{
		ObjectCableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
	}
}

func (e *CableBundle) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectCableBundle{
		AssignedObjectCableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
	}
}

func (e *CableBundle) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationCableBundle{
		TerminationCableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
	}
}

func (e *CableBundle) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectCableBundle{
		AssignedObjectCableBundle: e.ConvertToProtoMessage().(*pb.CableBundle),
	}
}

// implementation of oneof interfaces for RackGroup.

func (e *RackGroup) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectRackGroup{
		ObjectRackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

func (e *RackGroup) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_RackGroup{
		RackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

func (e *RackGroup) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceRackGroup{
		InterfaceRackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

func (e *RackGroup) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectRackGroup{
		ObjectRackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

func (e *RackGroup) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectRackGroup{
		AssignedObjectRackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

func (e *RackGroup) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationRackGroup{
		TerminationRackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

func (e *RackGroup) anyVLANGroupScopeValueApplyTo(p *pb.VLANGroup) {
	p.Scope = &pb.VLANGroup_ScopeRackGroup{
		ScopeRackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

func (e *RackGroup) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectRackGroup{
		AssignedObjectRackGroup: e.ConvertToProtoMessage().(*pb.RackGroup),
	}
}

// implementation of oneof interfaces for ScriptModule.

func (e *ScriptModule) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectScriptModule{
		ObjectScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
	}
}

func (e *ScriptModule) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_ScriptModule{
		ScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
	}
}

func (e *ScriptModule) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceScriptModule{
		InterfaceScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
	}
}

func (e *ScriptModule) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectScriptModule{
		ObjectScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
	}
}

func (e *ScriptModule) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectScriptModule{
		AssignedObjectScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
	}
}

func (e *ScriptModule) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationScriptModule{
		TerminationScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
	}
}

func (e *ScriptModule) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectScriptModule{
		AssignedObjectScriptModule: e.ConvertToProtoMessage().(*pb.ScriptModule),
	}
}

// implementation of oneof interfaces for VirtualMachineType.

func (e *VirtualMachineType) anyContactAssignmentObjectValueApplyTo(p *pb.ContactAssignment) {
	p.Object = &pb.ContactAssignment_ObjectVirtualMachineType{
		ObjectVirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
	}
}

func (e *VirtualMachineType) anyCustomFieldObjectReferenceObjectValueApplyTo(p *pb.CustomFieldObjectReference) {
	p.Object = &pb.CustomFieldObjectReference_VirtualMachineType{
		VirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
	}
}

func (e *VirtualMachineType) anyFHRPGroupAssignmentInterfaceValueApplyTo(p *pb.FHRPGroupAssignment) {
	p.Interface = &pb.FHRPGroupAssignment_InterfaceVirtualMachineType{
		InterfaceVirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
	}
}

func (e *VirtualMachineType) anyGenericObjectObjectValueApplyTo(p *pb.GenericObject) {
	p.Object = &pb.GenericObject_ObjectVirtualMachineType{
		ObjectVirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
	}
}

func (e *VirtualMachineType) anyL2VPNTerminationAssignedObjectValueApplyTo(p *pb.L2VPNTermination) {
	p.AssignedObject = &pb.L2VPNTermination_AssignedObjectVirtualMachineType{
		AssignedObjectVirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
	}
}

func (e *VirtualMachineType) anyTunnelTerminationTerminationValueApplyTo(p *pb.TunnelTermination) {
	p.Termination = &pb.TunnelTermination_TerminationVirtualMachineType{
		TerminationVirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
	}
}

func (e *VirtualMachineType) anyJournalEntryAssignedObjectValueApplyTo(p *pb.JournalEntry) {
	p.AssignedObject = &pb.JournalEntry_AssignedObjectVirtualMachineType{
		AssignedObjectVirtualMachineType: e.ConvertToProtoMessage().(*pb.VirtualMachineType),
	}
}

// interface for values of CableTermination.termination
type anyCableTerminationTerminationValue interface {
	anyCableTerminationTerminationValueApplyTo(e *pb.CableTermination)
}

// interface for values of CircuitGroupAssignment.member
type anyCircuitGroupAssignmentMemberValue interface {
	anyCircuitGroupAssignmentMemberValueApplyTo(e *pb.CircuitGroupAssignment)
}

// interface for values of CircuitTermination.termination
type anyCircuitTerminationTerminationValue interface {
	anyCircuitTerminationTerminationValueApplyTo(e *pb.CircuitTermination)
}

// interface for values of Cluster.scope
type anyClusterScopeValue interface {
	anyClusterScopeValueApplyTo(e *pb.Cluster)
}

// interface for values of ContactAssignment.object
type anyContactAssignmentObjectValue interface {
	anyContactAssignmentObjectValueApplyTo(e *pb.ContactAssignment)
}

// interface for values of CustomFieldObjectReference.object
type anyCustomFieldObjectReferenceObjectValue interface {
	anyCustomFieldObjectReferenceObjectValueApplyTo(e *pb.CustomFieldObjectReference)
}

// interface for values of CustomFieldValue.value
type anyCustomFieldValueValueValue interface {
	anyCustomFieldValueValueValueApplyTo(e *pb.CustomFieldValue)
}
type CustomFieldValueText string

func (e CustomFieldValueText) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Text{
		Text: string(e),
	}
}

type CustomFieldValueLongText string

func (e CustomFieldValueLongText) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_LongText{
		LongText: string(e),
	}
}

type CustomFieldValueInteger int64

func (e CustomFieldValueInteger) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Integer{
		Integer: int64(e),
	}
}

type CustomFieldValueDecimal float64

func (e CustomFieldValueDecimal) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Decimal{
		Decimal: float64(e),
	}
}

type CustomFieldValueBoolean bool

func (e CustomFieldValueBoolean) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Boolean{
		Boolean: bool(e),
	}
}

type CustomFieldValueDate time.Time

func (e CustomFieldValueDate) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Date{
		Date: timestamppb.New(time.Time(e)),
	}
}

type CustomFieldValueDatetime time.Time

func (e CustomFieldValueDatetime) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Datetime{
		Datetime: timestamppb.New(time.Time(e)),
	}
}

type CustomFieldValueUrl string

func (e CustomFieldValueUrl) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Url{
		Url: string(e),
	}
}

type CustomFieldValueJson string

func (e CustomFieldValueJson) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Json{
		Json: string(e),
	}
}

type CustomFieldValueSelection string

func (e CustomFieldValueSelection) anyCustomFieldValueValueValueApplyTo(p *pb.CustomFieldValue) {
	p.Value = &pb.CustomFieldValue_Selection{
		Selection: string(e),
	}
}

// interface for values of FHRPGroupAssignment.interface
type anyFHRPGroupAssignmentInterfaceValue interface {
	anyFHRPGroupAssignmentInterfaceValueApplyTo(e *pb.FHRPGroupAssignment)
}

// interface for values of GenericObject.object
type anyGenericObjectObjectValue interface {
	anyGenericObjectObjectValueApplyTo(e *pb.GenericObject)
}

// interface for values of IPAddress.assigned_object
type anyIPAddressAssignedObjectValue interface {
	anyIPAddressAssignedObjectValueApplyTo(e *pb.IPAddress)
}

// interface for values of InventoryItem.component
type anyInventoryItemComponentValue interface {
	anyInventoryItemComponentValueApplyTo(e *pb.InventoryItem)
}

// interface for values of L2VPNTermination.assigned_object
type anyL2VPNTerminationAssignedObjectValue interface {
	anyL2VPNTerminationAssignedObjectValueApplyTo(e *pb.L2VPNTermination)
}

// interface for values of MACAddress.assigned_object
type anyMACAddressAssignedObjectValue interface {
	anyMACAddressAssignedObjectValueApplyTo(e *pb.MACAddress)
}

// interface for values of Prefix.scope
type anyPrefixScopeValue interface {
	anyPrefixScopeValueApplyTo(e *pb.Prefix)
}

// interface for values of Service.parent_object
type anyServiceParentObjectValue interface {
	anyServiceParentObjectValueApplyTo(e *pb.Service)
}

// interface for values of TunnelTermination.termination
type anyTunnelTerminationTerminationValue interface {
	anyTunnelTerminationTerminationValueApplyTo(e *pb.TunnelTermination)
}

// interface for values of VLANGroup.scope
type anyVLANGroupScopeValue interface {
	anyVLANGroupScopeValueApplyTo(e *pb.VLANGroup)
}

// interface for values of WirelessLAN.scope
type anyWirelessLANScopeValue interface {
	anyWirelessLANScopeValueApplyTo(e *pb.WirelessLAN)
}

// interface for values of JournalEntry.assigned_object
type anyJournalEntryAssignedObjectValue interface {
	anyJournalEntryAssignedObjectValueApplyTo(e *pb.JournalEntry)
}
