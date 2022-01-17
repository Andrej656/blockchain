// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: motor/v1/response.proto

// Package Motor is used for defining a Motor node and its properties.

package motor

import (
	v1 "github.com/sonr-io/sonr/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// (Client) ShareResponse is response to ShareRequest
type ShareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // True if Supply is Active
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`      // Error Message if Supply is not Active
}

func (x *ShareResponse) Reset() {
	*x = ShareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareResponse) ProtoMessage() {}

func (x *ShareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareResponse.ProtoReflect.Descriptor instead.
func (*ShareResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *ShareResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ShareResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// (Client) RespondResponse is response to RespondRequest
type DecideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // True if Supply is Active
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`      // Error Message if Supply is not Active
}

func (x *DecideResponse) Reset() {
	*x = DecideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecideResponse) ProtoMessage() {}

func (x *DecideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecideResponse.ProtoReflect.Descriptor instead.
func (*DecideResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{1}
}

func (x *DecideResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DecideResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// (Client) SearchResponse is Message for Searching for Peer
type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`            // Success
	Error   string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`                 // Error Message
	Peer    *v1.Peer `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`                   // Peer Data
	PeerId  string   `protobuf:"bytes,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"` // Peer ID
	SName   string   `protobuf:"bytes,5,opt,name=s_name,json=sName,proto3" json:"s_name,omitempty"`    // SName
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SearchResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SearchResponse) GetPeer() *v1.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *SearchResponse) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *SearchResponse) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

// DecisionEvent is emitted when a decision is made by Peer.
type OnTransmitDecisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision bool     `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"` // true = accept, false = reject
	From     *v1.Peer `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`          // Peer that made decision
	Received int64    `protobuf:"varint,3,opt,name=received,proto3" json:"received,omitempty"` // Timestamp
}

func (x *OnTransmitDecisionResponse) Reset() {
	*x = OnTransmitDecisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnTransmitDecisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnTransmitDecisionResponse) ProtoMessage() {}

func (x *OnTransmitDecisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnTransmitDecisionResponse.ProtoReflect.Descriptor instead.
func (*OnTransmitDecisionResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{3}
}

func (x *OnTransmitDecisionResponse) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *OnTransmitDecisionResponse) GetFrom() *v1.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *OnTransmitDecisionResponse) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

// Message Sent when peer messages Lobby
type OnLobbyRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Olc      string     `protobuf:"bytes,1,opt,name=olc,proto3" json:"olc,omitempty"`            // OLC Code of Topic
	Peers    []*v1.Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`        // User Information
	Received int64      `protobuf:"varint,3,opt,name=received,proto3" json:"received,omitempty"` // Invite received Timestamp
}

func (x *OnLobbyRefreshResponse) Reset() {
	*x = OnLobbyRefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnLobbyRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnLobbyRefreshResponse) ProtoMessage() {}

func (x *OnLobbyRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnLobbyRefreshResponse.ProtoReflect.Descriptor instead.
func (*OnLobbyRefreshResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{4}
}

func (x *OnLobbyRefreshResponse) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *OnLobbyRefreshResponse) GetPeers() []*v1.Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *OnLobbyRefreshResponse) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

// InviteEvent notifies Peer that an Invite has been received
type OnTransmitInviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Received int64       `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"` // Invite received Timestamp
	From     *v1.Peer    `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`          // Peer that sent the Invite
	Payload  *v1.Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`    // Attached Data
}

func (x *OnTransmitInviteResponse) Reset() {
	*x = OnTransmitInviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnTransmitInviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnTransmitInviteResponse) ProtoMessage() {}

func (x *OnTransmitInviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnTransmitInviteResponse.ProtoReflect.Descriptor instead.
func (*OnTransmitInviteResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{5}
}

func (x *OnTransmitInviteResponse) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *OnTransmitInviteResponse) GetFrom() *v1.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *OnTransmitInviteResponse) GetPayload() *v1.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Received Mail Event
type OnMailboxMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`             // ID is the Message ID
	Buffer   []byte       `protobuf:"bytes,2,opt,name=buffer,proto3" json:"buffer,omitempty"`     // Buffer is the message data
	From     *v1.Profile  `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`         // Users Peer Data
	To       *v1.Profile  `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`             // Receivers Peer Data
	Metadata *v1.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"` // Metadata
}

func (x *OnMailboxMessageResponse) Reset() {
	*x = OnMailboxMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnMailboxMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnMailboxMessageResponse) ProtoMessage() {}

func (x *OnMailboxMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnMailboxMessageResponse.ProtoReflect.Descriptor instead.
func (*OnMailboxMessageResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{6}
}

func (x *OnMailboxMessageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OnMailboxMessageResponse) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *OnMailboxMessageResponse) GetFrom() *v1.Profile {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *OnMailboxMessageResponse) GetTo() *v1.Profile {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *OnMailboxMessageResponse) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Transfer Progress Event
type OnTransmitProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress  float64      `protobuf:"fixed64,1,opt,name=progress,proto3" json:"progress,omitempty"`                           // Current Transfer Progress
	Received  int64        `protobuf:"varint,2,opt,name=received,proto3" json:"received,omitempty"`                            // Timestamp
	Current   int32        `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`                              // Current position of item in list
	Total     int32        `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`                                  // Total number of items in list
	Direction v1.Direction `protobuf:"varint,5,opt,name=direction,proto3,enum=common.v1.Direction" json:"direction,omitempty"` // Direction of Transfer
}

func (x *OnTransmitProgressResponse) Reset() {
	*x = OnTransmitProgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnTransmitProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnTransmitProgressResponse) ProtoMessage() {}

func (x *OnTransmitProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnTransmitProgressResponse.ProtoReflect.Descriptor instead.
func (*OnTransmitProgressResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{7}
}

func (x *OnTransmitProgressResponse) GetProgress() float64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *OnTransmitProgressResponse) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *OnTransmitProgressResponse) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *OnTransmitProgressResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OnTransmitProgressResponse) GetDirection() v1.Direction {
	if x != nil {
		return x.Direction
	}
	return v1.Direction(0)
}

// Message Sent after Completed Transfer
type OnTransmitCompleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction  v1.Direction   `protobuf:"varint,1,opt,name=direction,proto3,enum=common.v1.Direction" json:"direction,omitempty"`                                                             // Direction of Transfer
	Payload    *v1.Payload    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                                                                                           // Transfer Data
	From       *v1.Peer       `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`                                                                                                 // Peer that sent the Complete Event
	To         *v1.Peer       `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`                                                                                                     // Peer that received the Complete Event
	CreatedAt  int64          `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                                                     // Transfer Created Timestamp
	ReceivedAt int64          `protobuf:"varint,6,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`                                                                  // Transfer Received Timestamp
	Results    map[int32]bool `protobuf:"bytes,7,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // Transfer Success
}

func (x *OnTransmitCompleteResponse) Reset() {
	*x = OnTransmitCompleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motor_v1_response_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnTransmitCompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnTransmitCompleteResponse) ProtoMessage() {}

func (x *OnTransmitCompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motor_v1_response_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnTransmitCompleteResponse.ProtoReflect.Descriptor instead.
func (*OnTransmitCompleteResponse) Descriptor() ([]byte, []int) {
	return file_motor_v1_response_proto_rawDescGZIP(), []int{8}
}

func (x *OnTransmitCompleteResponse) GetDirection() v1.Direction {
	if x != nil {
		return x.Direction
	}
	return v1.Direction(0)
}

func (x *OnTransmitCompleteResponse) GetPayload() *v1.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *OnTransmitCompleteResponse) GetFrom() *v1.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *OnTransmitCompleteResponse) GetTo() *v1.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *OnTransmitCompleteResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *OnTransmitCompleteResponse) GetReceivedAt() int64 {
	if x != nil {
		return x.ReceivedAt
	}
	return 0
}

func (x *OnTransmitCompleteResponse) GetResults() map[int32]bool {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_motor_v1_response_proto protoreflect.FileDescriptor

var file_motor_v1_response_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x6f, 0x74, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3f, 0x0a, 0x0d, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x40, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x1a, 0x4f, 0x6e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x22, 0x6d, 0x0a, 0x16, 0x4f, 0x6e, 0x4c, 0x6f, 0x62, 0x62, 0x79,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x6c,
	0x63, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x18, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0xbf, 0x01, 0x0a, 0x18, 0x4f, 0x6e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xb8, 0x01, 0x0a, 0x1a, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8d, 0x03,
	0x0a, 0x1a, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1f, 0x5a,
	0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_motor_v1_response_proto_rawDescOnce sync.Once
	file_motor_v1_response_proto_rawDescData = file_motor_v1_response_proto_rawDesc
)

func file_motor_v1_response_proto_rawDescGZIP() []byte {
	file_motor_v1_response_proto_rawDescOnce.Do(func() {
		file_motor_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_motor_v1_response_proto_rawDescData)
	})
	return file_motor_v1_response_proto_rawDescData
}

var file_motor_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_motor_v1_response_proto_goTypes = []interface{}{
	(*ShareResponse)(nil),              // 0: motor.v1.ShareResponse
	(*DecideResponse)(nil),             // 1: motor.v1.DecideResponse
	(*SearchResponse)(nil),             // 2: motor.v1.SearchResponse
	(*OnTransmitDecisionResponse)(nil), // 3: motor.v1.OnTransmitDecisionResponse
	(*OnLobbyRefreshResponse)(nil),     // 4: motor.v1.OnLobbyRefreshResponse
	(*OnTransmitInviteResponse)(nil),   // 5: motor.v1.OnTransmitInviteResponse
	(*OnMailboxMessageResponse)(nil),   // 6: motor.v1.OnMailboxMessageResponse
	(*OnTransmitProgressResponse)(nil), // 7: motor.v1.OnTransmitProgressResponse
	(*OnTransmitCompleteResponse)(nil), // 8: motor.v1.OnTransmitCompleteResponse
	nil,                                // 9: motor.v1.OnTransmitCompleteResponse.ResultsEntry
	(*v1.Peer)(nil),                    // 10: common.v1.Peer
	(*v1.Payload)(nil),                 // 11: common.v1.Payload
	(*v1.Profile)(nil),                 // 12: common.v1.Profile
	(*v1.Metadata)(nil),                // 13: common.v1.Metadata
	(v1.Direction)(0),                  // 14: common.v1.Direction
}
var file_motor_v1_response_proto_depIdxs = []int32{
	10, // 0: motor.v1.SearchResponse.peer:type_name -> common.v1.Peer
	10, // 1: motor.v1.OnTransmitDecisionResponse.from:type_name -> common.v1.Peer
	10, // 2: motor.v1.OnLobbyRefreshResponse.peers:type_name -> common.v1.Peer
	10, // 3: motor.v1.OnTransmitInviteResponse.from:type_name -> common.v1.Peer
	11, // 4: motor.v1.OnTransmitInviteResponse.payload:type_name -> common.v1.Payload
	12, // 5: motor.v1.OnMailboxMessageResponse.from:type_name -> common.v1.Profile
	12, // 6: motor.v1.OnMailboxMessageResponse.to:type_name -> common.v1.Profile
	13, // 7: motor.v1.OnMailboxMessageResponse.metadata:type_name -> common.v1.Metadata
	14, // 8: motor.v1.OnTransmitProgressResponse.direction:type_name -> common.v1.Direction
	14, // 9: motor.v1.OnTransmitCompleteResponse.direction:type_name -> common.v1.Direction
	11, // 10: motor.v1.OnTransmitCompleteResponse.payload:type_name -> common.v1.Payload
	10, // 11: motor.v1.OnTransmitCompleteResponse.from:type_name -> common.v1.Peer
	10, // 12: motor.v1.OnTransmitCompleteResponse.to:type_name -> common.v1.Peer
	9,  // 13: motor.v1.OnTransmitCompleteResponse.results:type_name -> motor.v1.OnTransmitCompleteResponse.ResultsEntry
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_motor_v1_response_proto_init() }
func file_motor_v1_response_proto_init() {
	if File_motor_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_motor_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecideResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnTransmitDecisionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnLobbyRefreshResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnTransmitInviteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnMailboxMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnTransmitProgressResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_motor_v1_response_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnTransmitCompleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_motor_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_motor_v1_response_proto_goTypes,
		DependencyIndexes: file_motor_v1_response_proto_depIdxs,
		MessageInfos:      file_motor_v1_response_proto_msgTypes,
	}.Build()
	File_motor_v1_response_proto = out.File
	file_motor_v1_response_proto_rawDesc = nil
	file_motor_v1_response_proto_goTypes = nil
	file_motor_v1_response_proto_depIdxs = nil
}
