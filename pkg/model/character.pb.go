// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.3
// source: pb/model/character.proto

package model

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

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Level       uint32     `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	MaxLevel    uint32     `protobuf:"varint,3,opt,name=max_level,json=maxLevel,proto3" json:"max_level,omitempty"`
	Eidols      uint32     `protobuf:"varint,4,opt,name=eidols,proto3" json:"eidols,omitempty"`
	Traces      []string   `protobuf:"bytes,5,rep,name=traces,proto3" json:"traces,omitempty"`
	Abilities   *Abilities `protobuf:"bytes,6,opt,name=abilities,proto3" json:"abilities,omitempty"`
	LightCone   *LightCone `protobuf:"bytes,7,opt,name=light_cone,json=lightCone,proto3" json:"light_cone,omitempty"`
	Relics      []*Relic   `protobuf:"bytes,8,rep,name=relics,proto3" json:"relics,omitempty"` // TODO: oneof for alternative options
	StartEnergy float64    `protobuf:"fixed64,9,opt,name=start_energy,json=startEnergy,proto3" json:"start_energy,omitempty"`
	StartHp     float64    `protobuf:"fixed64,10,opt,name=start_hp,json=startHp,proto3" json:"start_hp,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	mi := &file_pb_model_character_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_character_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_pb_model_character_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Character) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Character) GetMaxLevel() uint32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *Character) GetEidols() uint32 {
	if x != nil {
		return x.Eidols
	}
	return 0
}

func (x *Character) GetTraces() []string {
	if x != nil {
		return x.Traces
	}
	return nil
}

func (x *Character) GetAbilities() *Abilities {
	if x != nil {
		return x.Abilities
	}
	return nil
}

func (x *Character) GetLightCone() *LightCone {
	if x != nil {
		return x.LightCone
	}
	return nil
}

func (x *Character) GetRelics() []*Relic {
	if x != nil {
		return x.Relics
	}
	return nil
}

func (x *Character) GetStartEnergy() float64 {
	if x != nil {
		return x.StartEnergy
	}
	return 0
}

func (x *Character) GetStartHp() float64 {
	if x != nil {
		return x.StartHp
	}
	return 0
}

type Abilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attack uint32 `protobuf:"varint,1,opt,name=attack,proto3" json:"attack,omitempty"`
	Skill  uint32 `protobuf:"varint,2,opt,name=skill,proto3" json:"skill,omitempty"`
	Ult    uint32 `protobuf:"varint,3,opt,name=ult,proto3" json:"ult,omitempty"`
	Talent uint32 `protobuf:"varint,4,opt,name=talent,proto3" json:"talent,omitempty"`
}

func (x *Abilities) Reset() {
	*x = Abilities{}
	mi := &file_pb_model_character_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Abilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Abilities) ProtoMessage() {}

func (x *Abilities) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_character_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Abilities.ProtoReflect.Descriptor instead.
func (*Abilities) Descriptor() ([]byte, []int) {
	return file_pb_model_character_proto_rawDescGZIP(), []int{1}
}

func (x *Abilities) GetAttack() uint32 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *Abilities) GetSkill() uint32 {
	if x != nil {
		return x.Skill
	}
	return 0
}

func (x *Abilities) GetUlt() uint32 {
	if x != nil {
		return x.Ult
	}
	return 0
}

func (x *Abilities) GetTalent() uint32 {
	if x != nil {
		return x.Talent
	}
	return 0
}

type LightCone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Level      uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	MaxLevel   uint32 `protobuf:"varint,3,opt,name=max_level,json=maxLevel,proto3" json:"max_level,omitempty"`
	Imposition uint32 `protobuf:"varint,4,opt,name=imposition,proto3" json:"imposition,omitempty"`
}

func (x *LightCone) Reset() {
	*x = LightCone{}
	mi := &file_pb_model_character_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LightCone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightCone) ProtoMessage() {}

func (x *LightCone) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_character_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightCone.ProtoReflect.Descriptor instead.
func (*LightCone) Descriptor() ([]byte, []int) {
	return file_pb_model_character_proto_rawDescGZIP(), []int{2}
}

func (x *LightCone) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LightCone) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LightCone) GetMaxLevel() uint32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *LightCone) GetImposition() uint32 {
	if x != nil {
		return x.Imposition
	}
	return 0
}

type Relic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MainStat *RelicStat   `protobuf:"bytes,2,opt,name=main_stat,json=mainStat,proto3" json:"main_stat,omitempty"`
	SubStats []*RelicStat `protobuf:"bytes,3,rep,name=sub_stats,json=subStats,proto3" json:"sub_stats,omitempty"`
}

func (x *Relic) Reset() {
	*x = Relic{}
	mi := &file_pb_model_character_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Relic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relic) ProtoMessage() {}

func (x *Relic) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_character_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relic.ProtoReflect.Descriptor instead.
func (*Relic) Descriptor() ([]byte, []int) {
	return file_pb_model_character_proto_rawDescGZIP(), []int{3}
}

func (x *Relic) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Relic) GetMainStat() *RelicStat {
	if x != nil {
		return x.MainStat
	}
	return nil
}

func (x *Relic) GetSubStats() []*RelicStat {
	if x != nil {
		return x.SubStats
	}
	return nil
}

type RelicStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat   Property `protobuf:"varint,1,opt,name=stat,proto3,enum=model.Property" json:"stat,omitempty"`
	Amount float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *RelicStat) Reset() {
	*x = RelicStat{}
	mi := &file_pb_model_character_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelicStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelicStat) ProtoMessage() {}

func (x *RelicStat) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_character_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelicStat.ProtoReflect.Descriptor instead.
func (*RelicStat) Descriptor() ([]byte, []int) {
	return file_pb_model_character_proto_rawDescGZIP(), []int{4}
}

func (x *RelicStat) GetStat() Property {
	if x != nil {
		return x.Stat
	}
	return Property_INVALID_PROP
}

func (x *RelicStat) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_pb_model_character_proto protoreflect.FileDescriptor

var file_pb_model_character_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x13, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x69, 0x64,
	0x6f, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x69, 0x64, 0x6f, 0x6c,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x09,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6e, 0x65, 0x52,
	0x09, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65,
	0x6c, 0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x70, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x70, 0x22, 0x63,
	0x0a, 0x09, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x6c, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x6c,
	0x65, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x09, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6e, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x05, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2d, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x2d, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x69, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x08, 0x73, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x48,
	0x0a, 0x09, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x2f, 0x73, 0x72, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_model_character_proto_rawDescOnce sync.Once
	file_pb_model_character_proto_rawDescData = file_pb_model_character_proto_rawDesc
)

func file_pb_model_character_proto_rawDescGZIP() []byte {
	file_pb_model_character_proto_rawDescOnce.Do(func() {
		file_pb_model_character_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_model_character_proto_rawDescData)
	})
	return file_pb_model_character_proto_rawDescData
}

var file_pb_model_character_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_model_character_proto_goTypes = []any{
	(*Character)(nil), // 0: model.Character
	(*Abilities)(nil), // 1: model.Abilities
	(*LightCone)(nil), // 2: model.LightCone
	(*Relic)(nil),     // 3: model.Relic
	(*RelicStat)(nil), // 4: model.RelicStat
	(Property)(0),     // 5: model.Property
}
var file_pb_model_character_proto_depIdxs = []int32{
	1, // 0: model.Character.abilities:type_name -> model.Abilities
	2, // 1: model.Character.light_cone:type_name -> model.LightCone
	3, // 2: model.Character.relics:type_name -> model.Relic
	4, // 3: model.Relic.main_stat:type_name -> model.RelicStat
	4, // 4: model.Relic.sub_stats:type_name -> model.RelicStat
	5, // 5: model.RelicStat.stat:type_name -> model.Property
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_model_character_proto_init() }
func file_pb_model_character_proto_init() {
	if File_pb_model_character_proto != nil {
		return
	}
	file_pb_model_enum_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_model_character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_model_character_proto_goTypes,
		DependencyIndexes: file_pb_model_character_proto_depIdxs,
		MessageInfos:      file_pb_model_character_proto_msgTypes,
	}.Build()
	File_pb_model_character_proto = out.File
	file_pb_model_character_proto_rawDesc = nil
	file_pb_model_character_proto_goTypes = nil
	file_pb_model_character_proto_depIdxs = nil
}
