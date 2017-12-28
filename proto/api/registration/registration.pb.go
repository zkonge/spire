// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registration.proto

/*
Package registration is a generated protocol buffer package.

It is generated from these files:
	registration.proto

It has these top-level messages:
	RegistrationEntryID
	ParentID
	SpiffeID
	UpdateEntryRequest
	FederatedBundle
	CreateFederatedBundleRequest
	ListFederatedBundlesReply
	FederatedSpiffeID
	JoinToken
*/
package registration

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api1 "google.golang.org/genproto/googleapis/api/annotations"
import spire_common "github.com/spiffe/spire/proto/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// http from public import google/api/annotations.proto
var E_Http = google_api1.E_Http

// Empty from public import github.com/spiffe/spire/proto/common/common.proto
type Empty spire_common.Empty

func (m *Empty) Reset()         { (*spire_common.Empty)(m).Reset() }
func (m *Empty) String() string { return (*spire_common.Empty)(m).String() }
func (*Empty) ProtoMessage()    {}

// AttestedData from public import github.com/spiffe/spire/proto/common/common.proto
type AttestedData spire_common.AttestedData

func (m *AttestedData) Reset()          { (*spire_common.AttestedData)(m).Reset() }
func (m *AttestedData) String() string  { return (*spire_common.AttestedData)(m).String() }
func (*AttestedData) ProtoMessage()     {}
func (m *AttestedData) GetType() string { return (*spire_common.AttestedData)(m).GetType() }
func (m *AttestedData) GetData() []byte { return (*spire_common.AttestedData)(m).GetData() }

// Selector from public import github.com/spiffe/spire/proto/common/common.proto
type Selector spire_common.Selector

func (m *Selector) Reset()           { (*spire_common.Selector)(m).Reset() }
func (m *Selector) String() string   { return (*spire_common.Selector)(m).String() }
func (*Selector) ProtoMessage()      {}
func (m *Selector) GetType() string  { return (*spire_common.Selector)(m).GetType() }
func (m *Selector) GetValue() string { return (*spire_common.Selector)(m).GetValue() }

// Selectors from public import github.com/spiffe/spire/proto/common/common.proto
type Selectors spire_common.Selectors

func (m *Selectors) Reset()         { (*spire_common.Selectors)(m).Reset() }
func (m *Selectors) String() string { return (*spire_common.Selectors)(m).String() }
func (*Selectors) ProtoMessage()    {}
func (m *Selectors) GetEntries() []*Selector {
	o := (*spire_common.Selectors)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}

// RegistrationEntry from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntry spire_common.RegistrationEntry

func (m *RegistrationEntry) Reset()         { (*spire_common.RegistrationEntry)(m).Reset() }
func (m *RegistrationEntry) String() string { return (*spire_common.RegistrationEntry)(m).String() }
func (*RegistrationEntry) ProtoMessage()    {}
func (m *RegistrationEntry) GetSelectors() []*Selector {
	o := (*spire_common.RegistrationEntry)(m).GetSelectors()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}
func (m *RegistrationEntry) GetParentId() string {
	return (*spire_common.RegistrationEntry)(m).GetParentId()
}
func (m *RegistrationEntry) GetSpiffeId() string {
	return (*spire_common.RegistrationEntry)(m).GetSpiffeId()
}
func (m *RegistrationEntry) GetTtl() int32 { return (*spire_common.RegistrationEntry)(m).GetTtl() }
func (m *RegistrationEntry) GetFbSpiffeIds() []string {
	return (*spire_common.RegistrationEntry)(m).GetFbSpiffeIds()
}

// RegistrationEntries from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntries spire_common.RegistrationEntries

func (m *RegistrationEntries) Reset()         { (*spire_common.RegistrationEntries)(m).Reset() }
func (m *RegistrationEntries) String() string { return (*spire_common.RegistrationEntries)(m).String() }
func (*RegistrationEntries) ProtoMessage()    {}
func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	o := (*spire_common.RegistrationEntries)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*RegistrationEntry, len(o))
	for i, x := range o {
		s[i] = (*RegistrationEntry)(x)
	}
	return s
}

// A type that represents the id of an entry.
type RegistrationEntryID struct {
	// RegistrationEntryID.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RegistrationEntryID) Reset()                    { *m = RegistrationEntryID{} }
func (m *RegistrationEntryID) String() string            { return proto.CompactTextString(m) }
func (*RegistrationEntryID) ProtoMessage()               {}
func (*RegistrationEntryID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegistrationEntryID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A type that represents a parent Id.
type ParentID struct {
	// ParentId.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ParentID) Reset()                    { *m = ParentID{} }
func (m *ParentID) String() string            { return proto.CompactTextString(m) }
func (*ParentID) ProtoMessage()               {}
func (*ParentID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ParentID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A type that represents a SPIFFE Id.
type SpiffeID struct {
	// SpiffeId.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SpiffeID) Reset()                    { *m = SpiffeID{} }
func (m *SpiffeID) String() string            { return proto.CompactTextString(m) }
func (*SpiffeID) ProtoMessage()               {}
func (*SpiffeID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SpiffeID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A type with the id with want to update plus values to modify.
type UpdateEntryRequest struct {
	// Id of the entry to update.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Values in the RegistrationEntry to update.
	Entry *spire_common.RegistrationEntry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateEntryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateEntryRequest) GetEntry() *spire_common.RegistrationEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// A CA bundle for a different Trust Domain than the one used and managed by the Control Plane.
type FederatedBundle struct {
	// A SPIFFE ID that has a Federated Bundle
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId" json:"spiffe_id,omitempty"`
	// A trusted cert bundle that is not part of Control Planes trust domain but belongs to a different Trust Domain
	FederatedBundle []byte `protobuf:"bytes,2,opt,name=federated_bundle,json=federatedBundle,proto3" json:"federated_bundle,omitempty"`
	// Time to live.
	Ttl int32 `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *FederatedBundle) Reset()                    { *m = FederatedBundle{} }
func (m *FederatedBundle) String() string            { return proto.CompactTextString(m) }
func (*FederatedBundle) ProtoMessage()               {}
func (*FederatedBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FederatedBundle) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *FederatedBundle) GetFederatedBundle() []byte {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

func (m *FederatedBundle) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

// It represents a request with a FederatedBundle to create.
type CreateFederatedBundleRequest struct {
	// A trusted cert bundle that is not part of Control Planes trust domain but belongs to a different Trust Domain.
	FederatedBundle *FederatedBundle `protobuf:"bytes,1,opt,name=federated_bundle,json=federatedBundle" json:"federated_bundle,omitempty"`
}

func (m *CreateFederatedBundleRequest) Reset()                    { *m = CreateFederatedBundleRequest{} }
func (m *CreateFederatedBundleRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFederatedBundleRequest) ProtoMessage()               {}
func (*CreateFederatedBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateFederatedBundleRequest) GetFederatedBundle() *FederatedBundle {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

// It represents a reply with a list of FederatedBundle.
type ListFederatedBundlesReply struct {
	// A list of FederatedBundle.
	Bundles []*FederatedBundle `protobuf:"bytes,1,rep,name=bundles" json:"bundles,omitempty"`
}

func (m *ListFederatedBundlesReply) Reset()                    { *m = ListFederatedBundlesReply{} }
func (m *ListFederatedBundlesReply) String() string            { return proto.CompactTextString(m) }
func (*ListFederatedBundlesReply) ProtoMessage()               {}
func (*ListFederatedBundlesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListFederatedBundlesReply) GetBundles() []*FederatedBundle {
	if m != nil {
		return m.Bundles
	}
	return nil
}

// A type that represents a Federated SPIFFE Id.
type FederatedSpiffeID struct {
	// FederatedSpiffeID
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *FederatedSpiffeID) Reset()                    { *m = FederatedSpiffeID{} }
func (m *FederatedSpiffeID) String() string            { return proto.CompactTextString(m) }
func (*FederatedSpiffeID) ProtoMessage()               {}
func (*FederatedSpiffeID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FederatedSpiffeID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// JoinToken message is used for registering a new token
type JoinToken struct {
	// The join token. If not set, one will be generated
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// TTL in seconds
	Ttl int32 `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *JoinToken) Reset()                    { *m = JoinToken{} }
func (m *JoinToken) String() string            { return proto.CompactTextString(m) }
func (*JoinToken) ProtoMessage()               {}
func (*JoinToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *JoinToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *JoinToken) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func init() {
	proto.RegisterType((*RegistrationEntryID)(nil), "spire.api.registration.RegistrationEntryID")
	proto.RegisterType((*ParentID)(nil), "spire.api.registration.ParentID")
	proto.RegisterType((*SpiffeID)(nil), "spire.api.registration.SpiffeID")
	proto.RegisterType((*UpdateEntryRequest)(nil), "spire.api.registration.UpdateEntryRequest")
	proto.RegisterType((*FederatedBundle)(nil), "spire.api.registration.FederatedBundle")
	proto.RegisterType((*CreateFederatedBundleRequest)(nil), "spire.api.registration.CreateFederatedBundleRequest")
	proto.RegisterType((*ListFederatedBundlesReply)(nil), "spire.api.registration.ListFederatedBundlesReply")
	proto.RegisterType((*FederatedSpiffeID)(nil), "spire.api.registration.FederatedSpiffeID")
	proto.RegisterType((*JoinToken)(nil), "spire.api.registration.JoinToken")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Registration service

type RegistrationClient interface {
	// Creates an entry in the Registration table, used to assign SPIFFE IDs to nodes and workloads.
	CreateEntry(ctx context.Context, in *spire_common.RegistrationEntry, opts ...grpc.CallOption) (*RegistrationEntryID, error)
	// Deletes an entry and returns the deleted entry.
	DeleteEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error)
	// Retrieve a specific registered entry.
	FetchEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error)
	// Retrieve all registered entries.
	FetchEntries(ctx context.Context, in *spire_common.Empty, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error)
	// Updates a specific registered entry.
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error)
	// Returns all the Entries associated with the ParentID value.
	ListByParentID(ctx context.Context, in *ParentID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error)
	// Returns all the entries associated with a selector value.
	ListBySelector(ctx context.Context, in *spire_common.Selector, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error)
	// Return all registration entries for which SPIFFE ID matches.
	ListBySpiffeID(ctx context.Context, in *SpiffeID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error)
	// Creates an entry in the Federated bundle table to store the mappings of Federated SPIFFE IDs and their associated CA bundle.
	CreateFederatedBundle(ctx context.Context, in *CreateFederatedBundleRequest, opts ...grpc.CallOption) (*spire_common.Empty, error)
	// Retrieves Federated bundles for all the Federated SPIFFE IDs.
	ListFederatedBundles(ctx context.Context, in *spire_common.Empty, opts ...grpc.CallOption) (*ListFederatedBundlesReply, error)
	// Updates a particular Federated Bundle. Useful for rotation.
	UpdateFederatedBundle(ctx context.Context, in *FederatedBundle, opts ...grpc.CallOption) (*spire_common.Empty, error)
	// Delete a particular Federated Bundle. Used to destroy inter-domain trust.
	DeleteFederatedBundle(ctx context.Context, in *FederatedSpiffeID, opts ...grpc.CallOption) (*spire_common.Empty, error)
	// Create a new join token
	CreateJoinToken(ctx context.Context, in *JoinToken, opts ...grpc.CallOption) (*JoinToken, error)
}

type registrationClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationClient(cc *grpc.ClientConn) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) CreateEntry(ctx context.Context, in *spire_common.RegistrationEntry, opts ...grpc.CallOption) (*RegistrationEntryID, error) {
	out := new(RegistrationEntryID)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/CreateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) DeleteEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error) {
	out := new(spire_common.RegistrationEntry)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/DeleteEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) FetchEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error) {
	out := new(spire_common.RegistrationEntry)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/FetchEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) FetchEntries(ctx context.Context, in *spire_common.Empty, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error) {
	out := new(spire_common.RegistrationEntries)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/FetchEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error) {
	out := new(spire_common.RegistrationEntry)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListByParentID(ctx context.Context, in *ParentID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error) {
	out := new(spire_common.RegistrationEntries)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListByParentID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListBySelector(ctx context.Context, in *spire_common.Selector, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error) {
	out := new(spire_common.RegistrationEntries)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListBySelector", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListBySpiffeID(ctx context.Context, in *SpiffeID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error) {
	out := new(spire_common.RegistrationEntries)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListBySpiffeID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) CreateFederatedBundle(ctx context.Context, in *CreateFederatedBundleRequest, opts ...grpc.CallOption) (*spire_common.Empty, error) {
	out := new(spire_common.Empty)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/CreateFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListFederatedBundles(ctx context.Context, in *spire_common.Empty, opts ...grpc.CallOption) (*ListFederatedBundlesReply, error) {
	out := new(ListFederatedBundlesReply)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListFederatedBundles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) UpdateFederatedBundle(ctx context.Context, in *FederatedBundle, opts ...grpc.CallOption) (*spire_common.Empty, error) {
	out := new(spire_common.Empty)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/UpdateFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) DeleteFederatedBundle(ctx context.Context, in *FederatedSpiffeID, opts ...grpc.CallOption) (*spire_common.Empty, error) {
	out := new(spire_common.Empty)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/DeleteFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) CreateJoinToken(ctx context.Context, in *JoinToken, opts ...grpc.CallOption) (*JoinToken, error) {
	out := new(JoinToken)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/CreateJoinToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registration service

type RegistrationServer interface {
	// Creates an entry in the Registration table, used to assign SPIFFE IDs to nodes and workloads.
	CreateEntry(context.Context, *spire_common.RegistrationEntry) (*RegistrationEntryID, error)
	// Deletes an entry and returns the deleted entry.
	DeleteEntry(context.Context, *RegistrationEntryID) (*spire_common.RegistrationEntry, error)
	// Retrieve a specific registered entry.
	FetchEntry(context.Context, *RegistrationEntryID) (*spire_common.RegistrationEntry, error)
	// Retrieve all registered entries.
	FetchEntries(context.Context, *spire_common.Empty) (*spire_common.RegistrationEntries, error)
	// Updates a specific registered entry.
	UpdateEntry(context.Context, *UpdateEntryRequest) (*spire_common.RegistrationEntry, error)
	// Returns all the Entries associated with the ParentID value.
	ListByParentID(context.Context, *ParentID) (*spire_common.RegistrationEntries, error)
	// Returns all the entries associated with a selector value.
	ListBySelector(context.Context, *spire_common.Selector) (*spire_common.RegistrationEntries, error)
	// Return all registration entries for which SPIFFE ID matches.
	ListBySpiffeID(context.Context, *SpiffeID) (*spire_common.RegistrationEntries, error)
	// Creates an entry in the Federated bundle table to store the mappings of Federated SPIFFE IDs and their associated CA bundle.
	CreateFederatedBundle(context.Context, *CreateFederatedBundleRequest) (*spire_common.Empty, error)
	// Retrieves Federated bundles for all the Federated SPIFFE IDs.
	ListFederatedBundles(context.Context, *spire_common.Empty) (*ListFederatedBundlesReply, error)
	// Updates a particular Federated Bundle. Useful for rotation.
	UpdateFederatedBundle(context.Context, *FederatedBundle) (*spire_common.Empty, error)
	// Delete a particular Federated Bundle. Used to destroy inter-domain trust.
	DeleteFederatedBundle(context.Context, *FederatedSpiffeID) (*spire_common.Empty, error)
	// Create a new join token
	CreateJoinToken(context.Context, *JoinToken) (*JoinToken, error)
}

func RegisterRegistrationServer(s *grpc.Server, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common.RegistrationEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).CreateEntry(ctx, req.(*spire_common.RegistrationEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationEntryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).DeleteEntry(ctx, req.(*RegistrationEntryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_FetchEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationEntryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).FetchEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/FetchEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).FetchEntry(ctx, req.(*RegistrationEntryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_FetchEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).FetchEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/FetchEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).FetchEntries(ctx, req.(*spire_common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListByParentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListByParentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListByParentID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListByParentID(ctx, req.(*ParentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListBySelector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common.Selector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListBySelector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListBySelector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListBySelector(ctx, req.(*spire_common.Selector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListBySpiffeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiffeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListBySpiffeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListBySpiffeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListBySpiffeID(ctx, req.(*SpiffeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_CreateFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederatedBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).CreateFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/CreateFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).CreateFederatedBundle(ctx, req.(*CreateFederatedBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListFederatedBundles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListFederatedBundles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListFederatedBundles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListFederatedBundles(ctx, req.(*spire_common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_UpdateFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FederatedBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).UpdateFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/UpdateFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).UpdateFederatedBundle(ctx, req.(*FederatedBundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_DeleteFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FederatedSpiffeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).DeleteFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/DeleteFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).DeleteFederatedBundle(ctx, req.(*FederatedSpiffeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_CreateJoinToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).CreateJoinToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/CreateJoinToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).CreateJoinToken(ctx, req.(*JoinToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.registration.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntry",
			Handler:    _Registration_CreateEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _Registration_DeleteEntry_Handler,
		},
		{
			MethodName: "FetchEntry",
			Handler:    _Registration_FetchEntry_Handler,
		},
		{
			MethodName: "FetchEntries",
			Handler:    _Registration_FetchEntries_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _Registration_UpdateEntry_Handler,
		},
		{
			MethodName: "ListByParentID",
			Handler:    _Registration_ListByParentID_Handler,
		},
		{
			MethodName: "ListBySelector",
			Handler:    _Registration_ListBySelector_Handler,
		},
		{
			MethodName: "ListBySpiffeID",
			Handler:    _Registration_ListBySpiffeID_Handler,
		},
		{
			MethodName: "CreateFederatedBundle",
			Handler:    _Registration_CreateFederatedBundle_Handler,
		},
		{
			MethodName: "ListFederatedBundles",
			Handler:    _Registration_ListFederatedBundles_Handler,
		},
		{
			MethodName: "UpdateFederatedBundle",
			Handler:    _Registration_UpdateFederatedBundle_Handler,
		},
		{
			MethodName: "DeleteFederatedBundle",
			Handler:    _Registration_DeleteFederatedBundle_Handler,
		},
		{
			MethodName: "CreateJoinToken",
			Handler:    _Registration_CreateJoinToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registration.proto",
}

func init() { proto.RegisterFile("registration.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x9d, 0xaa, 0xfd, 0x9b, 0x71, 0x48, 0xdb, 0x4d, 0x53, 0x05, 0x53, 0x89, 0xd4, 0x08,
	0x91, 0x06, 0xc9, 0x56, 0x5b, 0xb8, 0x70, 0x23, 0xb4, 0x95, 0x0a, 0x1c, 0x22, 0x97, 0x80, 0x04,
	0x52, 0x2b, 0x27, 0x9e, 0xa4, 0x4b, 0x1c, 0xaf, 0xb1, 0x37, 0x87, 0x08, 0x71, 0xe1, 0x15, 0x78,
	0x34, 0x5e, 0x81, 0x03, 0x8f, 0x81, 0xb2, 0x9b, 0x35, 0xa9, 0x6b, 0x37, 0xe9, 0x81, 0x53, 0x36,
	0x3b, 0x33, 0xdf, 0x37, 0xdf, 0xcc, 0xec, 0x18, 0x48, 0x84, 0x03, 0x1a, 0xf3, 0xc8, 0xe5, 0x94,
	0x05, 0x56, 0x18, 0x31, 0xce, 0xc8, 0x4e, 0x1c, 0xd2, 0x08, 0x2d, 0x37, 0xa4, 0xd6, 0xbc, 0xd5,
	0xd8, 0x1d, 0x30, 0x36, 0xf0, 0xd1, 0x76, 0x43, 0x6a, 0xbb, 0x41, 0xc0, 0xb8, 0xb8, 0x8e, 0x65,
	0x94, 0x71, 0x30, 0xa0, 0xfc, 0x6a, 0xdc, 0xb5, 0x7a, 0x6c, 0x64, 0xc7, 0x21, 0xed, 0xf7, 0xd1,
	0x16, 0x38, 0xb6, 0x30, 0xdb, 0x3d, 0x36, 0x1a, 0xb1, 0x60, 0xf6, 0x23, 0x43, 0xcc, 0xc7, 0x50,
	0x71, 0xe6, 0x08, 0x4e, 0x02, 0x1e, 0x4d, 0xce, 0x8e, 0x49, 0x19, 0x0a, 0xd4, 0xab, 0x69, 0x75,
	0xad, 0x51, 0x74, 0x0a, 0xd4, 0x33, 0x0d, 0x58, 0x6f, 0xbb, 0x11, 0x06, 0x3c, 0xdb, 0x76, 0x2e,
	0xc8, 0x32, 0x6c, 0x9f, 0x80, 0x74, 0x42, 0xcf, 0xe5, 0x28, 0x80, 0x1d, 0xfc, 0x32, 0xc6, 0x98,
	0xa7, 0xbd, 0xc8, 0x73, 0x58, 0xc5, 0xa9, 0xbd, 0x56, 0xa8, 0x6b, 0x0d, 0xfd, 0xf0, 0xa1, 0x25,
	0xd5, 0xcf, 0x12, 0xbd, 0x91, 0x9f, 0x23, 0xbd, 0xcd, 0x21, 0x6c, 0x9c, 0xa2, 0x87, 0x91, 0xcb,
	0xd1, 0x6b, 0x8d, 0x03, 0xcf, 0x47, 0xf2, 0x00, 0x8a, 0x52, 0xf8, 0x65, 0x42, 0xb0, 0x2e, 0x2f,
	0xce, 0x3c, 0xb2, 0x0f, 0x9b, 0x7d, 0xe5, 0x7f, 0xd9, 0x15, 0x01, 0x82, 0xb1, 0xe4, 0x6c, 0xf4,
	0x53, 0x38, 0x9b, 0xb0, 0xc2, 0xb9, 0x5f, 0x5b, 0xa9, 0x6b, 0x8d, 0x55, 0x67, 0x7a, 0x34, 0x23,
	0xd8, 0x7d, 0x15, 0xa1, 0xcb, 0x31, 0x45, 0xa9, 0x34, 0x39, 0x19, 0xe0, 0x9a, 0x90, 0xf3, 0xc4,
	0xca, 0x6e, 0xa6, 0x95, 0x46, 0x4a, 0x67, 0x61, 0x5e, 0xc0, 0xfd, 0xb7, 0x34, 0xe6, 0x29, 0xbf,
	0xd8, 0xc1, 0xd0, 0x9f, 0x90, 0x97, 0xf0, 0xbf, 0xa4, 0x89, 0x6b, 0x5a, 0x7d, 0xe5, 0x2e, 0x3c,
	0x2a, 0xce, 0x7c, 0x04, 0x5b, 0x89, 0x2d, 0xb7, 0x85, 0x47, 0x50, 0x7c, 0xcd, 0x68, 0xf0, 0x8e,
	0x0d, 0x31, 0x20, 0xdb, 0xb0, 0xca, 0xa7, 0x87, 0x99, 0x5d, 0xfe, 0x51, 0xd5, 0x2a, 0x24, 0xd5,
	0x3a, 0xfc, 0x5d, 0x84, 0xd2, 0x7c, 0xdf, 0x48, 0x00, 0xba, 0x2c, 0x9f, 0xe8, 0x20, 0x59, 0xd4,
	0x62, 0xe3, 0x69, 0x9e, 0x98, 0x8c, 0x69, 0x35, 0xb7, 0xbe, 0xff, 0xfc, 0xf5, 0xa3, 0xa0, 0x9b,
	0x6b, 0xb6, 0x18, 0x8c, 0x17, 0x5a, 0x93, 0x0c, 0x41, 0x3f, 0x46, 0x1f, 0x15, 0xdf, 0x5d, 0xe0,
	0x8c, 0x45, 0xc9, 0x99, 0x65, 0xc1, 0xb7, 0xde, 0x9c, 0xf1, 0x11, 0x06, 0x70, 0x8a, 0xbc, 0x77,
	0xf5, 0x2f, 0xb8, 0x2a, 0x82, 0xeb, 0x1e, 0xd1, 0x25, 0x97, 0xfd, 0x95, 0x7a, 0xdf, 0xc8, 0x7b,
	0x28, 0x25, 0x84, 0x14, 0x63, 0x52, 0xb9, 0x8e, 0x72, 0x32, 0x0a, 0xf9, 0xc4, 0xd8, 0xbb, 0x1d,
	0x9a, 0x62, 0xac, 0x84, 0x10, 0x25, 0xe4, 0x33, 0xe8, 0x73, 0xcf, 0x95, 0x34, 0xf3, 0x94, 0xdc,
	0x7c, 0xd3, 0x4b, 0x17, 0xcd, 0x50, 0x5c, 0x1d, 0x28, 0x4f, 0x87, 0xbb, 0x35, 0x49, 0x16, 0x4b,
	0x3d, 0x8f, 0x4e, 0x79, 0x2c, 0x21, 0x89, 0xbc, 0x51, 0xb0, 0xe7, 0xe8, 0x63, 0x8f, 0xb3, 0x88,
	0xec, 0x5c, 0x0f, 0x52, 0xf7, 0xcb, 0x80, 0x25, 0x39, 0x26, 0xaf, 0x23, 0x37, 0x47, 0xe5, 0xb1,
	0x0c, 0x6c, 0x17, 0xaa, 0x99, 0xbb, 0x84, 0x3c, 0xcb, 0x43, 0xbf, 0x6d, 0xf5, 0x18, 0x59, 0xdd,
	0x27, 0x17, 0xb0, 0x9d, 0xb5, 0x3b, 0xb2, 0x47, 0xe5, 0x20, 0x8f, 0x37, 0x7f, 0xfd, 0x74, 0xa0,
	0x2a, 0xa7, 0x20, 0xad, 0x61, 0xd9, 0x35, 0x94, 0x9d, 0xf6, 0x07, 0xa8, 0xca, 0x77, 0x9b, 0x86,
	0xdd, 0x5f, 0x08, 0x9b, 0x74, 0x20, 0x07, 0x78, 0x43, 0x16, 0xf1, 0xef, 0x32, 0xdb, 0xcb, 0x83,
	0x4c, 0x5c, 0x8c, 0xc5, 0x2e, 0xad, 0xf2, 0xc7, 0xd2, 0xbc, 0xa5, 0xfd, 0x5f, 0x5b, 0xeb, 0xae,
	0x89, 0x8f, 0xeb, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x0c, 0xd4, 0xb9, 0xdb, 0x07,
	0x00, 0x00,
}
