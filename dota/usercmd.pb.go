// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.5
// source: usercmd.proto

package dota

import (
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

type CInButtonStatePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buttonstate1 *uint64 `protobuf:"varint,1,opt,name=buttonstate1" json:"buttonstate1,omitempty"`
	Buttonstate2 *uint64 `protobuf:"varint,2,opt,name=buttonstate2" json:"buttonstate2,omitempty"`
	Buttonstate3 *uint64 `protobuf:"varint,3,opt,name=buttonstate3" json:"buttonstate3,omitempty"`
}

func (x *CInButtonStatePB) Reset() {
	*x = CInButtonStatePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CInButtonStatePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CInButtonStatePB) ProtoMessage() {}

func (x *CInButtonStatePB) ProtoReflect() protoreflect.Message {
	mi := &file_usercmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CInButtonStatePB.ProtoReflect.Descriptor instead.
func (*CInButtonStatePB) Descriptor() ([]byte, []int) {
	return file_usercmd_proto_rawDescGZIP(), []int{0}
}

func (x *CInButtonStatePB) GetButtonstate1() uint64 {
	if x != nil && x.Buttonstate1 != nil {
		return *x.Buttonstate1
	}
	return 0
}

func (x *CInButtonStatePB) GetButtonstate2() uint64 {
	if x != nil && x.Buttonstate2 != nil {
		return *x.Buttonstate2
	}
	return 0
}

func (x *CInButtonStatePB) GetButtonstate3() uint64 {
	if x != nil && x.Buttonstate3 != nil {
		return *x.Buttonstate3
	}
	return 0
}

type CSubtickMoveStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Button             *uint64  `protobuf:"varint,1,opt,name=button" json:"button,omitempty"`
	Pressed            *bool    `protobuf:"varint,2,opt,name=pressed" json:"pressed,omitempty"`
	When               *float32 `protobuf:"fixed32,3,opt,name=when" json:"when,omitempty"`
	AnalogForwardDelta *float32 `protobuf:"fixed32,4,opt,name=analog_forward_delta,json=analogForwardDelta" json:"analog_forward_delta,omitempty"`
	AnalogLeftDelta    *float32 `protobuf:"fixed32,5,opt,name=analog_left_delta,json=analogLeftDelta" json:"analog_left_delta,omitempty"`
}

func (x *CSubtickMoveStep) Reset() {
	*x = CSubtickMoveStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSubtickMoveStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSubtickMoveStep) ProtoMessage() {}

func (x *CSubtickMoveStep) ProtoReflect() protoreflect.Message {
	mi := &file_usercmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSubtickMoveStep.ProtoReflect.Descriptor instead.
func (*CSubtickMoveStep) Descriptor() ([]byte, []int) {
	return file_usercmd_proto_rawDescGZIP(), []int{1}
}

func (x *CSubtickMoveStep) GetButton() uint64 {
	if x != nil && x.Button != nil {
		return *x.Button
	}
	return 0
}

func (x *CSubtickMoveStep) GetPressed() bool {
	if x != nil && x.Pressed != nil {
		return *x.Pressed
	}
	return false
}

func (x *CSubtickMoveStep) GetWhen() float32 {
	if x != nil && x.When != nil {
		return *x.When
	}
	return 0
}

func (x *CSubtickMoveStep) GetAnalogForwardDelta() float32 {
	if x != nil && x.AnalogForwardDelta != nil {
		return *x.AnalogForwardDelta
	}
	return 0
}

func (x *CSubtickMoveStep) GetAnalogLeftDelta() float32 {
	if x != nil && x.AnalogLeftDelta != nil {
		return *x.AnalogLeftDelta
	}
	return 0
}

type CBaseUserCmdPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LegacyCommandNumber        *int32              `protobuf:"varint,1,opt,name=legacy_command_number,json=legacyCommandNumber" json:"legacy_command_number,omitempty"`
	ClientTick                 *int32              `protobuf:"varint,2,opt,name=client_tick,json=clientTick" json:"client_tick,omitempty"`
	ButtonsPb                  *CInButtonStatePB   `protobuf:"bytes,3,opt,name=buttons_pb,json=buttonsPb" json:"buttons_pb,omitempty"`
	Viewangles                 *CMsgQAngle         `protobuf:"bytes,4,opt,name=viewangles" json:"viewangles,omitempty"`
	Forwardmove                *float32            `protobuf:"fixed32,5,opt,name=forwardmove" json:"forwardmove,omitempty"`
	Leftmove                   *float32            `protobuf:"fixed32,6,opt,name=leftmove" json:"leftmove,omitempty"`
	Upmove                     *float32            `protobuf:"fixed32,7,opt,name=upmove" json:"upmove,omitempty"`
	Impulse                    *int32              `protobuf:"varint,8,opt,name=impulse" json:"impulse,omitempty"`
	Weaponselect               *int32              `protobuf:"varint,9,opt,name=weaponselect" json:"weaponselect,omitempty"`
	RandomSeed                 *int32              `protobuf:"varint,10,opt,name=random_seed,json=randomSeed" json:"random_seed,omitempty"`
	Mousedx                    *int32              `protobuf:"varint,11,opt,name=mousedx" json:"mousedx,omitempty"`
	Mousedy                    *int32              `protobuf:"varint,12,opt,name=mousedy" json:"mousedy,omitempty"`
	PawnEntityHandle           *uint32             `protobuf:"varint,14,opt,name=pawn_entity_handle,json=pawnEntityHandle" json:"pawn_entity_handle,omitempty"`
	SubtickMoves               []*CSubtickMoveStep `protobuf:"bytes,18,rep,name=subtick_moves,json=subtickMoves" json:"subtick_moves,omitempty"`
	MoveCrc                    []byte              `protobuf:"bytes,19,opt,name=move_crc,json=moveCrc" json:"move_crc,omitempty"`
	ConsumedServerAngleChanges *uint32             `protobuf:"varint,20,opt,name=consumed_server_angle_changes,json=consumedServerAngleChanges" json:"consumed_server_angle_changes,omitempty"`
	CmdFlags                   *int32              `protobuf:"varint,21,opt,name=cmd_flags,json=cmdFlags" json:"cmd_flags,omitempty"`
}

func (x *CBaseUserCmdPB) Reset() {
	*x = CBaseUserCmdPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBaseUserCmdPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBaseUserCmdPB) ProtoMessage() {}

func (x *CBaseUserCmdPB) ProtoReflect() protoreflect.Message {
	mi := &file_usercmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBaseUserCmdPB.ProtoReflect.Descriptor instead.
func (*CBaseUserCmdPB) Descriptor() ([]byte, []int) {
	return file_usercmd_proto_rawDescGZIP(), []int{2}
}

func (x *CBaseUserCmdPB) GetLegacyCommandNumber() int32 {
	if x != nil && x.LegacyCommandNumber != nil {
		return *x.LegacyCommandNumber
	}
	return 0
}

func (x *CBaseUserCmdPB) GetClientTick() int32 {
	if x != nil && x.ClientTick != nil {
		return *x.ClientTick
	}
	return 0
}

func (x *CBaseUserCmdPB) GetButtonsPb() *CInButtonStatePB {
	if x != nil {
		return x.ButtonsPb
	}
	return nil
}

func (x *CBaseUserCmdPB) GetViewangles() *CMsgQAngle {
	if x != nil {
		return x.Viewangles
	}
	return nil
}

func (x *CBaseUserCmdPB) GetForwardmove() float32 {
	if x != nil && x.Forwardmove != nil {
		return *x.Forwardmove
	}
	return 0
}

func (x *CBaseUserCmdPB) GetLeftmove() float32 {
	if x != nil && x.Leftmove != nil {
		return *x.Leftmove
	}
	return 0
}

func (x *CBaseUserCmdPB) GetUpmove() float32 {
	if x != nil && x.Upmove != nil {
		return *x.Upmove
	}
	return 0
}

func (x *CBaseUserCmdPB) GetImpulse() int32 {
	if x != nil && x.Impulse != nil {
		return *x.Impulse
	}
	return 0
}

func (x *CBaseUserCmdPB) GetWeaponselect() int32 {
	if x != nil && x.Weaponselect != nil {
		return *x.Weaponselect
	}
	return 0
}

func (x *CBaseUserCmdPB) GetRandomSeed() int32 {
	if x != nil && x.RandomSeed != nil {
		return *x.RandomSeed
	}
	return 0
}

func (x *CBaseUserCmdPB) GetMousedx() int32 {
	if x != nil && x.Mousedx != nil {
		return *x.Mousedx
	}
	return 0
}

func (x *CBaseUserCmdPB) GetMousedy() int32 {
	if x != nil && x.Mousedy != nil {
		return *x.Mousedy
	}
	return 0
}

func (x *CBaseUserCmdPB) GetPawnEntityHandle() uint32 {
	if x != nil && x.PawnEntityHandle != nil {
		return *x.PawnEntityHandle
	}
	return 0
}

func (x *CBaseUserCmdPB) GetSubtickMoves() []*CSubtickMoveStep {
	if x != nil {
		return x.SubtickMoves
	}
	return nil
}

func (x *CBaseUserCmdPB) GetMoveCrc() []byte {
	if x != nil {
		return x.MoveCrc
	}
	return nil
}

func (x *CBaseUserCmdPB) GetConsumedServerAngleChanges() uint32 {
	if x != nil && x.ConsumedServerAngleChanges != nil {
		return *x.ConsumedServerAngleChanges
	}
	return 0
}

func (x *CBaseUserCmdPB) GetCmdFlags() int32 {
	if x != nil && x.CmdFlags != nil {
		return *x.CmdFlags
	}
	return 0
}

type CUserCmdBasePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *CBaseUserCmdPB `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
}

func (x *CUserCmdBasePB) Reset() {
	*x = CUserCmdBasePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercmd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserCmdBasePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserCmdBasePB) ProtoMessage() {}

func (x *CUserCmdBasePB) ProtoReflect() protoreflect.Message {
	mi := &file_usercmd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserCmdBasePB.ProtoReflect.Descriptor instead.
func (*CUserCmdBasePB) Descriptor() ([]byte, []int) {
	return file_usercmd_proto_rawDescGZIP(), []int{3}
}

func (x *CUserCmdBasePB) GetBase() *CBaseUserCmdPB {
	if x != nil {
		return x.Base
	}
	return nil
}

var File_usercmd_proto protoreflect.FileDescriptor

var file_usercmd_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x64, 0x6f, 0x74, 0x61, 0x1a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a,
	0x10, 0x43, 0x49, 0x6e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50,
	0x42, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x31, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x62, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x33, 0x22, 0xb6, 0x01,
	0x0a, 0x10, 0x43, 0x53, 0x75, 0x62, 0x74, 0x69, 0x63, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x6e, 0x61, 0x6c,
	0x6f, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x61, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6e,
	0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x61, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x66,
	0x74, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x9d, 0x05, 0x0a, 0x0e, 0x43, 0x42, 0x61, 0x73, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6d, 0x64, 0x50, 0x42, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x35,
	0x0a, 0x0a, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x49, 0x6e, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x42, 0x52, 0x09, 0x62, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x73, 0x50, 0x62, 0x12, 0x30, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x74, 0x61,
	0x2e, 0x43, 0x4d, 0x73, 0x67, 0x51, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x0a, 0x76, 0x69, 0x65,
	0x77, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x66,
	0x74, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x65, 0x66,
	0x74, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x6d, 0x6f, 0x76, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x75, 0x70, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x6d, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x69, 0x6d, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x77,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x6f, 0x75, 0x73, 0x65, 0x64, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x6f, 0x75, 0x73, 0x65, 0x64, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x64,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x64, 0x79,
	0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x70, 0x61,
	0x77, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x3b,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x18,
	0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x63, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x74, 0x69, 0x63, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x76, 0x65, 0x5f, 0x63, 0x72, 0x63, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x6f, 0x76, 0x65, 0x43, 0x72, 0x63, 0x12, 0x41, 0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x6e, 0x67,
	0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6d, 0x64,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6d,
	0x64, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x3a, 0x0a, 0x0e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6d, 0x64, 0x42, 0x61, 0x73, 0x65, 0x50, 0x42, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x42,
	0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6d, 0x64, 0x50, 0x42, 0x52, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6f, 0x74, 0x61, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x61, 0x2f,
	0x64, 0x6f, 0x74, 0x61, 0x3b, 0x64, 0x6f, 0x74, 0x61,
}

var (
	file_usercmd_proto_rawDescOnce sync.Once
	file_usercmd_proto_rawDescData = file_usercmd_proto_rawDesc
)

func file_usercmd_proto_rawDescGZIP() []byte {
	file_usercmd_proto_rawDescOnce.Do(func() {
		file_usercmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_usercmd_proto_rawDescData)
	})
	return file_usercmd_proto_rawDescData
}

var file_usercmd_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_usercmd_proto_goTypes = []interface{}{
	(*CInButtonStatePB)(nil), // 0: dota.CInButtonStatePB
	(*CSubtickMoveStep)(nil), // 1: dota.CSubtickMoveStep
	(*CBaseUserCmdPB)(nil),   // 2: dota.CBaseUserCmdPB
	(*CUserCmdBasePB)(nil),   // 3: dota.CUserCmdBasePB
	(*CMsgQAngle)(nil),       // 4: dota.CMsgQAngle
}
var file_usercmd_proto_depIdxs = []int32{
	0, // 0: dota.CBaseUserCmdPB.buttons_pb:type_name -> dota.CInButtonStatePB
	4, // 1: dota.CBaseUserCmdPB.viewangles:type_name -> dota.CMsgQAngle
	1, // 2: dota.CBaseUserCmdPB.subtick_moves:type_name -> dota.CSubtickMoveStep
	2, // 3: dota.CUserCmdBasePB.base:type_name -> dota.CBaseUserCmdPB
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_usercmd_proto_init() }
func file_usercmd_proto_init() {
	if File_usercmd_proto != nil {
		return
	}
	file_networkbasetypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_usercmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CInButtonStatePB); i {
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
		file_usercmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSubtickMoveStep); i {
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
		file_usercmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBaseUserCmdPB); i {
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
		file_usercmd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CUserCmdBasePB); i {
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
			RawDescriptor: file_usercmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_usercmd_proto_goTypes,
		DependencyIndexes: file_usercmd_proto_depIdxs,
		MessageInfos:      file_usercmd_proto_msgTypes,
	}.Build()
	File_usercmd_proto = out.File
	file_usercmd_proto_rawDesc = nil
	file_usercmd_proto_goTypes = nil
	file_usercmd_proto_depIdxs = nil
}
