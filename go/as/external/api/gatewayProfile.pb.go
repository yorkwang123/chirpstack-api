// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/gatewayProfile.proto

package api

import (
	context "context"
	fmt "fmt"
	common "github.com/brocaar/chirpstack-api/go/v3/common"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GatewayProfile struct {
	// Gateway-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the gateway-profile.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server ID of the gateway-profile.
	NetworkServerId int64 `protobuf:"varint,3,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Default channels (channels specified by the LoRaWAN Regional Parameters
	// specification) enabled for this configuration.
	Channels []uint32 `protobuf:"varint,4,rep,packed,name=channels,proto3" json:"channels,omitempty"`
	// Extra channels added to the channel-configuration (in case the LoRaWAN
	// region supports adding custom channels).
	ExtraChannels        []*GatewayProfileExtraChannel `protobuf:"bytes,5,rep,name=extra_channels,json=extraChannels,proto3" json:"extra_channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GatewayProfile) Reset()         { *m = GatewayProfile{} }
func (m *GatewayProfile) String() string { return proto.CompactTextString(m) }
func (*GatewayProfile) ProtoMessage()    {}
func (*GatewayProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{0}
}

func (m *GatewayProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfile.Unmarshal(m, b)
}
func (m *GatewayProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfile.Marshal(b, m, deterministic)
}
func (m *GatewayProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfile.Merge(m, src)
}
func (m *GatewayProfile) XXX_Size() int {
	return xxx_messageInfo_GatewayProfile.Size(m)
}
func (m *GatewayProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfile.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfile proto.InternalMessageInfo

func (m *GatewayProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GatewayProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewayProfile) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *GatewayProfile) GetChannels() []uint32 {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *GatewayProfile) GetExtraChannels() []*GatewayProfileExtraChannel {
	if m != nil {
		return m.ExtraChannels
	}
	return nil
}

type GatewayProfileListItem struct {
	// Gateway-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Gateway-profile name,
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server ID on which the gateway-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,3,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Network-server name.
	NetworkServerName string `protobuf:"bytes,7,opt,name=network_server_name,json=networkServerName,proto3" json:"network_server_name,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GatewayProfileListItem) Reset()         { *m = GatewayProfileListItem{} }
func (m *GatewayProfileListItem) String() string { return proto.CompactTextString(m) }
func (*GatewayProfileListItem) ProtoMessage()    {}
func (*GatewayProfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{1}
}

func (m *GatewayProfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfileListItem.Unmarshal(m, b)
}
func (m *GatewayProfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfileListItem.Marshal(b, m, deterministic)
}
func (m *GatewayProfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfileListItem.Merge(m, src)
}
func (m *GatewayProfileListItem) XXX_Size() int {
	return xxx_messageInfo_GatewayProfileListItem.Size(m)
}
func (m *GatewayProfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfileListItem proto.InternalMessageInfo

func (m *GatewayProfileListItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GatewayProfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewayProfileListItem) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *GatewayProfileListItem) GetNetworkServerName() string {
	if m != nil {
		return m.NetworkServerName
	}
	return ""
}

func (m *GatewayProfileListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GatewayProfileListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GatewayProfileExtraChannel struct {
	// Modulation.
	Modulation common.Modulation `protobuf:"varint,1,opt,name=modulation,proto3,enum=common.Modulation" json:"modulation,omitempty"`
	// Frequency.
	Frequency uint32 `protobuf:"varint,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Bandwidth.
	Bandwidth uint32 `protobuf:"varint,3,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Bitrate (in case of FSK modulation).
	Bitrate uint32 `protobuf:"varint,4,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	// Spreading factors (in case of LoRa modulation).
	SpreadingFactors     []uint32 `protobuf:"varint,5,rep,packed,name=spreading_factors,json=spreadingFactors,proto3" json:"spreading_factors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayProfileExtraChannel) Reset()         { *m = GatewayProfileExtraChannel{} }
func (m *GatewayProfileExtraChannel) String() string { return proto.CompactTextString(m) }
func (*GatewayProfileExtraChannel) ProtoMessage()    {}
func (*GatewayProfileExtraChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{2}
}

func (m *GatewayProfileExtraChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfileExtraChannel.Unmarshal(m, b)
}
func (m *GatewayProfileExtraChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfileExtraChannel.Marshal(b, m, deterministic)
}
func (m *GatewayProfileExtraChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfileExtraChannel.Merge(m, src)
}
func (m *GatewayProfileExtraChannel) XXX_Size() int {
	return xxx_messageInfo_GatewayProfileExtraChannel.Size(m)
}
func (m *GatewayProfileExtraChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfileExtraChannel.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfileExtraChannel proto.InternalMessageInfo

func (m *GatewayProfileExtraChannel) GetModulation() common.Modulation {
	if m != nil {
		return m.Modulation
	}
	return common.Modulation_LORA
}

func (m *GatewayProfileExtraChannel) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *GatewayProfileExtraChannel) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *GatewayProfileExtraChannel) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

func (m *GatewayProfileExtraChannel) GetSpreadingFactors() []uint32 {
	if m != nil {
		return m.SpreadingFactors
	}
	return nil
}

type CreateGatewayProfileRequest struct {
	// Gateway-profile object to create.
	GatewayProfile       *GatewayProfile `protobuf:"bytes,1,opt,name=gateway_profile,json=gatewayProfile,proto3" json:"gateway_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateGatewayProfileRequest) Reset()         { *m = CreateGatewayProfileRequest{} }
func (m *CreateGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGatewayProfileRequest) ProtoMessage()    {}
func (*CreateGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{3}
}

func (m *CreateGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGatewayProfileRequest.Unmarshal(m, b)
}
func (m *CreateGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGatewayProfileRequest.Merge(m, src)
}
func (m *CreateGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGatewayProfileRequest.Size(m)
}
func (m *CreateGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGatewayProfileRequest proto.InternalMessageInfo

func (m *CreateGatewayProfileRequest) GetGatewayProfile() *GatewayProfile {
	if m != nil {
		return m.GatewayProfile
	}
	return nil
}

type CreateGatewayProfileResponse struct {
	// Gateway-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGatewayProfileResponse) Reset()         { *m = CreateGatewayProfileResponse{} }
func (m *CreateGatewayProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGatewayProfileResponse) ProtoMessage()    {}
func (*CreateGatewayProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{4}
}

func (m *CreateGatewayProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGatewayProfileResponse.Unmarshal(m, b)
}
func (m *CreateGatewayProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGatewayProfileResponse.Marshal(b, m, deterministic)
}
func (m *CreateGatewayProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGatewayProfileResponse.Merge(m, src)
}
func (m *CreateGatewayProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGatewayProfileResponse.Size(m)
}
func (m *CreateGatewayProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGatewayProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGatewayProfileResponse proto.InternalMessageInfo

func (m *CreateGatewayProfileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetGatewayProfileRequest struct {
	// Gateway-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayProfileRequest) Reset()         { *m = GetGatewayProfileRequest{} }
func (m *GetGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileRequest) ProtoMessage()    {}
func (*GetGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{5}
}

func (m *GetGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileRequest.Unmarshal(m, b)
}
func (m *GetGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileRequest.Merge(m, src)
}
func (m *GetGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileRequest.Size(m)
}
func (m *GetGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileRequest proto.InternalMessageInfo

func (m *GetGatewayProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetGatewayProfileResponse struct {
	// Gateway-profile object.
	GatewayProfile *GatewayProfile `protobuf:"bytes,1,opt,name=gateway_profile,json=gatewayProfile,proto3" json:"gateway_profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetGatewayProfileResponse) Reset()         { *m = GetGatewayProfileResponse{} }
func (m *GetGatewayProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileResponse) ProtoMessage()    {}
func (*GetGatewayProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{6}
}

func (m *GetGatewayProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileResponse.Unmarshal(m, b)
}
func (m *GetGatewayProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileResponse.Merge(m, src)
}
func (m *GetGatewayProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileResponse.Size(m)
}
func (m *GetGatewayProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileResponse proto.InternalMessageInfo

func (m *GetGatewayProfileResponse) GetGatewayProfile() *GatewayProfile {
	if m != nil {
		return m.GatewayProfile
	}
	return nil
}

func (m *GetGatewayProfileResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetGatewayProfileResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateGatewayProfileRequest struct {
	// Gateway-profile object to update.
	GatewayProfile       *GatewayProfile `protobuf:"bytes,1,opt,name=gateway_profile,json=gatewayProfile,proto3" json:"gateway_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateGatewayProfileRequest) Reset()         { *m = UpdateGatewayProfileRequest{} }
func (m *UpdateGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGatewayProfileRequest) ProtoMessage()    {}
func (*UpdateGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{7}
}

func (m *UpdateGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGatewayProfileRequest.Unmarshal(m, b)
}
func (m *UpdateGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGatewayProfileRequest.Merge(m, src)
}
func (m *UpdateGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGatewayProfileRequest.Size(m)
}
func (m *UpdateGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGatewayProfileRequest proto.InternalMessageInfo

func (m *UpdateGatewayProfileRequest) GetGatewayProfile() *GatewayProfile {
	if m != nil {
		return m.GatewayProfile
	}
	return nil
}

type DeleteGatewayProfileRequest struct {
	// Gateway-profile id (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGatewayProfileRequest) Reset()         { *m = DeleteGatewayProfileRequest{} }
func (m *DeleteGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGatewayProfileRequest) ProtoMessage()    {}
func (*DeleteGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{8}
}

func (m *DeleteGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGatewayProfileRequest.Unmarshal(m, b)
}
func (m *DeleteGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGatewayProfileRequest.Merge(m, src)
}
func (m *DeleteGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGatewayProfileRequest.Size(m)
}
func (m *DeleteGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGatewayProfileRequest proto.InternalMessageInfo

func (m *DeleteGatewayProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListGatewayProfilesRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Network-server ID to filter on (optional).
	NetworkServerId      int64    `protobuf:"varint,3,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGatewayProfilesRequest) Reset()         { *m = ListGatewayProfilesRequest{} }
func (m *ListGatewayProfilesRequest) String() string { return proto.CompactTextString(m) }
func (*ListGatewayProfilesRequest) ProtoMessage()    {}
func (*ListGatewayProfilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{9}
}

func (m *ListGatewayProfilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGatewayProfilesRequest.Unmarshal(m, b)
}
func (m *ListGatewayProfilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGatewayProfilesRequest.Marshal(b, m, deterministic)
}
func (m *ListGatewayProfilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGatewayProfilesRequest.Merge(m, src)
}
func (m *ListGatewayProfilesRequest) XXX_Size() int {
	return xxx_messageInfo_ListGatewayProfilesRequest.Size(m)
}
func (m *ListGatewayProfilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGatewayProfilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGatewayProfilesRequest proto.InternalMessageInfo

func (m *ListGatewayProfilesRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListGatewayProfilesRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListGatewayProfilesRequest) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

type ListGatewayProfilesResponse struct {
	// Total number of gateway-profiles.
	TotalCount           int64                     `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Result               []*GatewayProfileListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListGatewayProfilesResponse) Reset()         { *m = ListGatewayProfilesResponse{} }
func (m *ListGatewayProfilesResponse) String() string { return proto.CompactTextString(m) }
func (*ListGatewayProfilesResponse) ProtoMessage()    {}
func (*ListGatewayProfilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{10}
}

func (m *ListGatewayProfilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGatewayProfilesResponse.Unmarshal(m, b)
}
func (m *ListGatewayProfilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGatewayProfilesResponse.Marshal(b, m, deterministic)
}
func (m *ListGatewayProfilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGatewayProfilesResponse.Merge(m, src)
}
func (m *ListGatewayProfilesResponse) XXX_Size() int {
	return xxx_messageInfo_ListGatewayProfilesResponse.Size(m)
}
func (m *ListGatewayProfilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGatewayProfilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGatewayProfilesResponse proto.InternalMessageInfo

func (m *ListGatewayProfilesResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListGatewayProfilesResponse) GetResult() []*GatewayProfileListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*GatewayProfile)(nil), "api.GatewayProfile")
	proto.RegisterType((*GatewayProfileListItem)(nil), "api.GatewayProfileListItem")
	proto.RegisterType((*GatewayProfileExtraChannel)(nil), "api.GatewayProfileExtraChannel")
	proto.RegisterType((*CreateGatewayProfileRequest)(nil), "api.CreateGatewayProfileRequest")
	proto.RegisterType((*CreateGatewayProfileResponse)(nil), "api.CreateGatewayProfileResponse")
	proto.RegisterType((*GetGatewayProfileRequest)(nil), "api.GetGatewayProfileRequest")
	proto.RegisterType((*GetGatewayProfileResponse)(nil), "api.GetGatewayProfileResponse")
	proto.RegisterType((*UpdateGatewayProfileRequest)(nil), "api.UpdateGatewayProfileRequest")
	proto.RegisterType((*DeleteGatewayProfileRequest)(nil), "api.DeleteGatewayProfileRequest")
	proto.RegisterType((*ListGatewayProfilesRequest)(nil), "api.ListGatewayProfilesRequest")
	proto.RegisterType((*ListGatewayProfilesResponse)(nil), "api.ListGatewayProfilesResponse")
}

func init() {
	proto.RegisterFile("as/external/api/gatewayProfile.proto", fileDescriptor_a425cd68ff0462c4)
}

var fileDescriptor_a425cd68ff0462c4 = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0x97, 0xe3, 0x6c, 0xca, 0xbe, 0x28, 0x29, 0x3b, 0x4b, 0x17, 0xe3, 0xa4, 0xac, 0xb1, 0x38,
	0x44, 0x41, 0x6b, 0x4b, 0x59, 0x21, 0x54, 0xc4, 0xa5, 0x6c, 0xdb, 0x55, 0x25, 0x40, 0xc8, 0xc0,
	0x85, 0x1e, 0xa2, 0x89, 0x3d, 0x49, 0x46, 0xb5, 0x3d, 0x66, 0x66, 0xb2, 0xdb, 0x08, 0xf5, 0xc2,
	0x07, 0xe0, 0xc2, 0x07, 0xe2, 0xc2, 0x9d, 0x43, 0xbf, 0x02, 0x1f, 0x04, 0x79, 0x66, 0x92, 0xae,
	0x53, 0x3b, 0x68, 0x81, 0x9e, 0x92, 0xf7, 0xde, 0x6f, 0xde, 0xdf, 0xdf, 0x9b, 0x31, 0x7c, 0x8c,
	0x45, 0x48, 0x5e, 0x48, 0xc2, 0x73, 0x9c, 0x86, 0xb8, 0xa0, 0xe1, 0x02, 0x4b, 0x72, 0x8d, 0xd7,
	0xdf, 0x72, 0x36, 0xa7, 0x29, 0x09, 0x0a, 0xce, 0x24, 0x43, 0x36, 0x2e, 0xa8, 0x3b, 0x5c, 0x30,
	0xb6, 0x48, 0x89, 0x42, 0xe1, 0x3c, 0x67, 0x12, 0x4b, 0xca, 0x72, 0xa1, 0x21, 0xee, 0xa9, 0xb1,
	0x2a, 0x69, 0xb6, 0x9a, 0x87, 0x92, 0x66, 0x44, 0x48, 0x9c, 0x15, 0x06, 0x30, 0xd8, 0x05, 0x90,
	0xac, 0x90, 0x6b, 0x63, 0x3c, 0x8e, 0x59, 0x96, 0xb1, 0x3c, 0xd4, 0x3f, 0x5a, 0xe9, 0xff, 0x61,
	0x41, 0xff, 0xb2, 0x92, 0x0e, 0xea, 0x43, 0x8b, 0x26, 0x8e, 0xe5, 0x59, 0xa3, 0xc3, 0xa8, 0x45,
	0x13, 0x84, 0xa0, 0x9d, 0xe3, 0x8c, 0x38, 0x2d, 0xa5, 0x51, 0xff, 0xd1, 0x18, 0x8e, 0x72, 0x22,
	0xaf, 0x19, 0x7f, 0x3e, 0x15, 0x84, 0x5f, 0x11, 0x3e, 0xa5, 0x89, 0x63, 0x7b, 0xd6, 0xc8, 0x8e,
	0xee, 0x1a, 0xc3, 0x77, 0x4a, 0xff, 0xf4, 0x11, 0x72, 0xe1, 0x9d, 0x78, 0x89, 0xf3, 0x9c, 0xa4,
	0xc2, 0x69, 0x7b, 0xf6, 0xa8, 0x17, 0x6d, 0x65, 0xf4, 0x04, 0xfa, 0xe4, 0x85, 0xe4, 0x78, 0xba,
	0x45, 0x1c, 0x78, 0xf6, 0xa8, 0x3b, 0x39, 0x0d, 0x70, 0x41, 0x83, 0x6a, 0x62, 0x8f, 0x4b, 0xe0,
	0x85, 0xc6, 0x45, 0x3d, 0x72, 0x43, 0x12, 0xfe, 0xaf, 0x2d, 0x38, 0xa9, 0xa2, 0xbf, 0xa2, 0x42,
	0x3e, 0x95, 0x24, 0xfb, 0xdf, 0xcb, 0x09, 0xe0, 0x78, 0x07, 0xab, 0xdc, 0xdd, 0x51, 0xee, 0x8e,
	0x2a, 0xe8, 0x6f, 0x4a, 0xdf, 0x0f, 0x00, 0x62, 0x4e, 0xb0, 0x24, 0xc9, 0x14, 0x4b, 0xe7, 0xc0,
	0xb3, 0x46, 0xdd, 0x89, 0x1b, 0xe8, 0x41, 0x05, 0x9b, 0x41, 0x05, 0xdf, 0x6f, 0x26, 0x19, 0x1d,
	0x1a, 0xf4, 0x43, 0x59, 0x1e, 0x5d, 0x15, 0xc9, 0xe6, 0x68, 0xe7, 0x9f, 0x8f, 0x1a, 0xf4, 0x43,
	0xe9, 0xbf, 0xb2, 0xc0, 0x6d, 0x6e, 0x1f, 0x9a, 0x00, 0x64, 0x2c, 0x59, 0xa5, 0x8a, 0x5e, 0xaa,
	0x39, 0xfd, 0x09, 0x0a, 0x0c, 0x33, 0xbe, 0xde, 0x5a, 0xa2, 0x1b, 0x28, 0x34, 0x84, 0xc3, 0x39,
	0x27, 0x3f, 0xad, 0x48, 0x1e, 0xaf, 0x55, 0xf7, 0x7a, 0xd1, 0x6b, 0x45, 0x69, 0x9d, 0xe1, 0x3c,
	0xb9, 0xa6, 0x89, 0x5c, 0xaa, 0xd6, 0xf5, 0xa2, 0xd7, 0x0a, 0xe4, 0xc0, 0x9d, 0x19, 0x95, 0x1c,
	0x4b, 0xe2, 0xb4, 0x95, 0x6d, 0x23, 0xa2, 0x4f, 0xe0, 0x48, 0x14, 0x9c, 0xe0, 0x84, 0xe6, 0x8b,
	0xe9, 0x1c, 0xc7, 0x92, 0x71, 0x4d, 0x82, 0x5e, 0xf4, 0xee, 0xd6, 0xf0, 0x44, 0xeb, 0xfd, 0x67,
	0x30, 0xb8, 0x50, 0xdd, 0xa9, 0x96, 0x16, 0x95, 0x49, 0x08, 0x89, 0xbe, 0x80, 0xbb, 0x66, 0xb5,
	0xa6, 0x85, 0xb6, 0xa8, 0xd2, 0xba, 0x93, 0xe3, 0x1a, 0x3a, 0x45, 0xfd, 0xea, 0x1a, 0xfa, 0x01,
	0x0c, 0xeb, 0x9d, 0x8b, 0x82, 0xe5, 0xe2, 0x8d, 0xbd, 0xf0, 0xc7, 0xe0, 0x5c, 0x12, 0x59, 0x9f,
	0xc9, 0x2e, 0xf6, 0x4f, 0x0b, 0x3e, 0xa8, 0x01, 0x1b, 0xcf, 0xff, 0x29, 0xef, 0x1d, 0x82, 0xb5,
	0xfe, 0x3d, 0xc1, 0xec, 0xdb, 0x10, 0xec, 0x19, 0x0c, 0x7e, 0x50, 0xc2, 0xdb, 0x18, 0xc5, 0x19,
	0x0c, 0x1e, 0x91, 0x94, 0x34, 0x39, 0xdf, 0xed, 0xee, 0x15, 0xb8, 0xe5, 0xba, 0x57, 0xc1, 0x62,
	0x83, 0x7e, 0x0f, 0x0e, 0x52, 0x9a, 0x51, 0xa9, 0x0e, 0xd8, 0x91, 0x16, 0xd0, 0x09, 0x74, 0xd8,
	0x7c, 0x2e, 0x88, 0xee, 0x98, 0x1d, 0x19, 0xe9, 0x36, 0x57, 0x81, 0x2f, 0x60, 0x50, 0x1b, 0xd7,
	0x8c, 0xf5, 0x14, 0xba, 0x92, 0x49, 0x9c, 0x4e, 0x63, 0xb6, 0xca, 0x37, 0xe1, 0x41, 0xa9, 0x2e,
	0x4a, 0x0d, 0x3a, 0x87, 0x0e, 0x27, 0x62, 0x95, 0x96, 0x39, 0x94, 0xb7, 0xde, 0xa0, 0xa6, 0x37,
	0x9b, 0x7b, 0x2c, 0x32, 0xd0, 0xc9, 0xef, 0x6d, 0xb8, 0x57, 0x85, 0x94, 0xf9, 0xd0, 0x98, 0x20,
	0x06, 0x1d, 0x4d, 0x60, 0xe4, 0x29, 0x47, 0x7b, 0x56, 0xc5, 0xfd, 0x68, 0x0f, 0x42, 0xa7, 0xef,
	0x7b, 0xbf, 0xbc, 0xfa, 0xeb, 0xb7, 0x96, 0xeb, 0xdf, 0xbb, 0xf9, 0x66, 0x9d, 0x99, 0x69, 0x8a,
	0xcf, 0xad, 0x31, 0x5a, 0x82, 0x7d, 0x49, 0x24, 0xba, 0xaf, 0xd3, 0x6e, 0xd8, 0x05, 0xf7, 0xc3,
	0x26, 0xb3, 0x89, 0xe3, 0xab, 0x38, 0x43, 0xe4, 0xd6, 0xc6, 0x09, 0x7f, 0xa6, 0xc9, 0x4b, 0xb4,
	0x86, 0x8e, 0x66, 0x9b, 0x29, 0x6d, 0x0f, 0xf5, 0xdc, 0x93, 0x37, 0x08, 0xfc, 0xb8, 0x7c, 0x05,
	0xfd, 0x4f, 0x55, 0x9c, 0xd0, 0x1d, 0x37, 0xc4, 0xd9, 0xe1, 0x6b, 0x40, 0x93, 0x97, 0x65, 0x91,
	0x73, 0xe8, 0x68, 0x2e, 0x9a, 0xd0, 0x7b, 0x88, 0xd9, 0x18, 0xda, 0x94, 0x38, 0xde, 0x57, 0xe2,
	0x12, 0xda, 0xe5, 0xac, 0x91, 0x7e, 0xfa, 0x9a, 0xf9, 0xec, 0x7a, 0xcd, 0x00, 0xd3, 0xd1, 0xfb,
	0x2a, 0xdc, 0xfb, 0xa8, 0x7e, 0x72, 0x5f, 0x3e, 0xf8, 0xf1, 0xb3, 0x05, 0x95, 0xcb, 0xd5, 0xac,
	0xbc, 0xf0, 0xc3, 0x19, 0x67, 0x31, 0xc6, 0x3c, 0x8c, 0x97, 0x94, 0x17, 0x42, 0xe2, 0xf8, 0xf9,
	0x99, 0x3a, 0xc5, 0xc2, 0xab, 0xf3, 0x70, 0xe7, 0xcb, 0x65, 0xd6, 0x51, 0x85, 0x9d, 0xff, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xd3, 0xc5, 0xe3, 0xd4, 0xd3, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayProfileServiceClient is the client API for GatewayProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayProfileServiceClient interface {
	// Create creates the given gateway-profile.
	Create(ctx context.Context, in *CreateGatewayProfileRequest, opts ...grpc.CallOption) (*CreateGatewayProfileResponse, error)
	// Get returns the gateway-profile matching the given id.
	Get(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error)
	// Update updates the given gateway-profile.
	Update(ctx context.Context, in *UpdateGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the gateway-profile matching the given id.
	Delete(ctx context.Context, in *DeleteGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List returns the existing gateway-profiles.
	List(ctx context.Context, in *ListGatewayProfilesRequest, opts ...grpc.CallOption) (*ListGatewayProfilesResponse, error)
}

type gatewayProfileServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayProfileServiceClient(cc *grpc.ClientConn) GatewayProfileServiceClient {
	return &gatewayProfileServiceClient{cc}
}

func (c *gatewayProfileServiceClient) Create(ctx context.Context, in *CreateGatewayProfileRequest, opts ...grpc.CallOption) (*CreateGatewayProfileResponse, error) {
	out := new(CreateGatewayProfileResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) Get(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error) {
	out := new(GetGatewayProfileResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) Update(ctx context.Context, in *UpdateGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) Delete(ctx context.Context, in *DeleteGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) List(ctx context.Context, in *ListGatewayProfilesRequest, opts ...grpc.CallOption) (*ListGatewayProfilesResponse, error) {
	out := new(ListGatewayProfilesResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayProfileServiceServer is the server API for GatewayProfileService service.
type GatewayProfileServiceServer interface {
	// Create creates the given gateway-profile.
	Create(context.Context, *CreateGatewayProfileRequest) (*CreateGatewayProfileResponse, error)
	// Get returns the gateway-profile matching the given id.
	Get(context.Context, *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error)
	// Update updates the given gateway-profile.
	Update(context.Context, *UpdateGatewayProfileRequest) (*empty.Empty, error)
	// Delete deletes the gateway-profile matching the given id.
	Delete(context.Context, *DeleteGatewayProfileRequest) (*empty.Empty, error)
	// List returns the existing gateway-profiles.
	List(context.Context, *ListGatewayProfilesRequest) (*ListGatewayProfilesResponse, error)
}

// UnimplementedGatewayProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayProfileServiceServer struct {
}

func (*UnimplementedGatewayProfileServiceServer) Create(ctx context.Context, req *CreateGatewayProfileRequest) (*CreateGatewayProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) Get(ctx context.Context, req *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) Update(ctx context.Context, req *UpdateGatewayProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) Delete(ctx context.Context, req *DeleteGatewayProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) List(ctx context.Context, req *ListGatewayProfilesRequest) (*ListGatewayProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterGatewayProfileServiceServer(s *grpc.Server, srv GatewayProfileServiceServer) {
	s.RegisterService(&_GatewayProfileService_serviceDesc, srv)
}

func _GatewayProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Create(ctx, req.(*CreateGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Get(ctx, req.(*GetGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Update(ctx, req.(*UpdateGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Delete(ctx, req.(*DeleteGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewayProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).List(ctx, req.(*ListGatewayProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GatewayProfileService",
	HandlerType: (*GatewayProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GatewayProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GatewayProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GatewayProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GatewayProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GatewayProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/gatewayProfile.proto",
}
