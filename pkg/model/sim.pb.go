// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pb/model/sim.proto

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

type TerminationReason int32

const (
	TerminationReason_INVALID_TERMINATION TerminationReason = 0
	TerminationReason_BATTLE_LOSS         TerminationReason = 1
	TerminationReason_BATTLE_WIN          TerminationReason = 2
	TerminationReason_TIMEOUT             TerminationReason = 3
)

// Enum value maps for TerminationReason.
var (
	TerminationReason_name = map[int32]string{
		0: "INVALID_TERMINATION",
		1: "BATTLE_LOSS",
		2: "BATTLE_WIN",
		3: "TIMEOUT",
	}
	TerminationReason_value = map[string]int32{
		"INVALID_TERMINATION": 0,
		"BATTLE_LOSS":         1,
		"BATTLE_WIN":          2,
		"TIMEOUT":             3,
	}
)

func (x TerminationReason) Enum() *TerminationReason {
	p := new(TerminationReason)
	*p = x
	return p
}

func (x TerminationReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TerminationReason) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_model_sim_proto_enumTypes[0].Descriptor()
}

func (TerminationReason) Type() protoreflect.EnumType {
	return &file_pb_model_sim_proto_enumTypes[0]
}

func (x TerminationReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TerminationReason.Descriptor instead.
func (TerminationReason) EnumDescriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{0}
}

type SimConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iterations  int32              `protobuf:"varint,1,opt,name=iterations,proto3" json:"iterations,omitempty"`
	WorkerCount int32              `protobuf:"varint,2,opt,name=worker_count,proto3" json:"worker_count,omitempty"`
	Settings    *SimulatorSettings `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	Characters  []*Character       `protobuf:"bytes,4,rep,name=characters,proto3" json:"characters,omitempty"`
	Enemies     []*Enemy           `protobuf:"bytes,5,rep,name=enemies,proto3" json:"enemies,omitempty"`
}

func (x *SimConfig) Reset() {
	*x = SimConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimConfig) ProtoMessage() {}

func (x *SimConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimConfig.ProtoReflect.Descriptor instead.
func (*SimConfig) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{0}
}

func (x *SimConfig) GetIterations() int32 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

func (x *SimConfig) GetWorkerCount() int32 {
	if x != nil {
		return x.WorkerCount
	}
	return 0
}

func (x *SimConfig) GetSettings() *SimulatorSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *SimConfig) GetCharacters() []*Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

func (x *SimConfig) GetEnemies() []*Enemy {
	if x != nil {
		return x.Enemies
	}
	return nil
}

type SimulatorSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CycleLimit int32 `protobuf:"varint,1,opt,name=cycle_limit,proto3" json:"cycle_limit,omitempty"`
	TtkMode    bool  `protobuf:"varint,2,opt,name=ttk_mode,proto3" json:"ttk_mode,omitempty"`
}

func (x *SimulatorSettings) Reset() {
	*x = SimulatorSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulatorSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorSettings) ProtoMessage() {}

func (x *SimulatorSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorSettings.ProtoReflect.Descriptor instead.
func (*SimulatorSettings) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{1}
}

func (x *SimulatorSettings) GetCycleLimit() int32 {
	if x != nil {
		return x.CycleLimit
	}
	return 0
}

func (x *SimulatorSettings) GetTtkMode() bool {
	if x != nil {
		return x.TtkMode
	}
	return false
}

type Relic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MainStat *RelicStat   `protobuf:"bytes,2,opt,name=main_stat,json=sub_stats,proto3" json:"main_stat,omitempty"`
	SubStats []*RelicStat `protobuf:"bytes,3,rep,name=sub_stats,proto3" json:"sub_stats,omitempty"`
}

func (x *Relic) Reset() {
	*x = Relic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relic) ProtoMessage() {}

func (x *Relic) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pb_model_sim_proto_rawDescGZIP(), []int{2}
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
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelicStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelicStat) ProtoMessage() {}

func (x *RelicStat) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pb_model_sim_proto_rawDescGZIP(), []int{3}
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

type LightCone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Level      uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	MaxLevel   uint32 `protobuf:"varint,3,opt,name=max_level,proto3" json:"max_level,omitempty"`
	Imposition uint32 `protobuf:"varint,4,opt,name=imposition,proto3" json:"imposition,omitempty"`
}

func (x *LightCone) Reset() {
	*x = LightCone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightCone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightCone) ProtoMessage() {}

func (x *LightCone) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pb_model_sim_proto_rawDescGZIP(), []int{4}
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

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Level       uint32     `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	MaxLevel    uint32     `protobuf:"varint,3,opt,name=max_level,proto3" json:"max_level,omitempty"`
	Eidols      uint32     `protobuf:"varint,4,opt,name=eidols,proto3" json:"eidols,omitempty"`
	Traces      []string   `protobuf:"bytes,5,rep,name=traces,proto3" json:"traces,omitempty"`
	Talents     []uint32   `protobuf:"varint,6,rep,packed,name=talents,proto3" json:"talents,omitempty"` // [attack, skill, ultimate, talent]
	Cone        *LightCone `protobuf:"bytes,7,opt,name=cone,proto3" json:"cone,omitempty"`
	Relics      []*Relic   `protobuf:"bytes,8,rep,name=relics,proto3" json:"relics,omitempty"`
	StartEnergy float64    `protobuf:"fixed64,9,opt,name=start_energy,proto3" json:"start_energy,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pb_model_sim_proto_rawDescGZIP(), []int{5}
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

func (x *Character) GetTalents() []uint32 {
	if x != nil {
		return x.Talents
	}
	return nil
}

func (x *Character) GetCone() *LightCone {
	if x != nil {
		return x.Cone
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

type Enemy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      uint32       `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Hp         float64      `protobuf:"fixed64,3,opt,name=hp,proto3" json:"hp,omitempty"`
	Toughness  float64      `protobuf:"fixed64,4,opt,name=toughness,proto3" json:"toughness,omitempty"`
	Weaknesses []DamageType `protobuf:"varint,5,rep,packed,name=weaknesses,proto3,enum=model.DamageType" json:"weaknesses,omitempty"`
	DebuffRes  []*DebuffRES `protobuf:"bytes,6,rep,name=debuff_res,proto3" json:"debuff_res,omitempty"`
}

func (x *Enemy) Reset() {
	*x = Enemy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enemy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enemy) ProtoMessage() {}

func (x *Enemy) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enemy.ProtoReflect.Descriptor instead.
func (*Enemy) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{6}
}

func (x *Enemy) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Enemy) GetHp() float64 {
	if x != nil {
		return x.Hp
	}
	return 0
}

func (x *Enemy) GetToughness() float64 {
	if x != nil {
		return x.Toughness
	}
	return 0
}

func (x *Enemy) GetWeaknesses() []DamageType {
	if x != nil {
		return x.Weaknesses
	}
	return nil
}

func (x *Enemy) GetDebuffRes() []*DebuffRES {
	if x != nil {
		return x.DebuffRes
	}
	return nil
}

type DebuffRES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat   BehaviorFlag `protobuf:"varint,1,opt,name=stat,proto3,enum=model.BehaviorFlag" json:"stat,omitempty"`
	Amount float64      `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DebuffRES) Reset() {
	*x = DebuffRES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebuffRES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebuffRES) ProtoMessage() {}

func (x *DebuffRES) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebuffRES.ProtoReflect.Descriptor instead.
func (*DebuffRES) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{7}
}

func (x *DebuffRES) GetStat() BehaviorFlag {
	if x != nil {
		return x.Stat
	}
	return BehaviorFlag_INVALID_FLAG
}

func (x *DebuffRES) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_pb_model_sim_proto protoreflect.FileDescriptor

var file_pb_model_sim_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x69, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x13, 0x70, 0x62, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdf, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0a,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x65, 0x6e,
	0x65, 0x6d, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69,
	0x65, 0x73, 0x22, 0x51, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x74, 0x6b,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x74, 0x6b,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x79, 0x0a, 0x05, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2e, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x2e, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x22, 0x48, 0x0a, 0x09, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x12, 0x23, 0x0a,
	0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x04, 0x73, 0x74,
	0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x71, 0x0a, 0x09, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x43, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8b, 0x02,
	0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x69, 0x64, 0x6f, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x65, 0x69, 0x64, 0x6f, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x07, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x63,
	0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x6e,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52,
	0x06, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x22, 0xb0, 0x01, 0x0a, 0x05,
	0x45, 0x6e, 0x65, 0x6d, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x68,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x68, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x6f, 0x75, 0x67, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x74, 0x6f, 0x75, 0x67, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x77, 0x65, 0x61,
	0x6b, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x77, 0x65, 0x61, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0a,
	0x64, 0x65, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x66, 0x66, 0x52,
	0x45, 0x53, 0x52, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x72, 0x65, 0x73, 0x22, 0x4c,
	0x0a, 0x09, 0x44, 0x65, 0x62, 0x75, 0x66, 0x66, 0x52, 0x45, 0x53, 0x12, 0x27, 0x0a, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x04,
	0x73, 0x74, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x5a, 0x0a, 0x11,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x45, 0x52,
	0x4d, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41,
	0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x42,
	0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54,
	0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x2f, 0x73, 0x72, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_model_sim_proto_rawDescOnce sync.Once
	file_pb_model_sim_proto_rawDescData = file_pb_model_sim_proto_rawDesc
)

func file_pb_model_sim_proto_rawDescGZIP() []byte {
	file_pb_model_sim_proto_rawDescOnce.Do(func() {
		file_pb_model_sim_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_model_sim_proto_rawDescData)
	})
	return file_pb_model_sim_proto_rawDescData
}

var file_pb_model_sim_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_model_sim_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_model_sim_proto_goTypes = []interface{}{
	(TerminationReason)(0),    // 0: model.TerminationReason
	(*SimConfig)(nil),         // 1: model.SimConfig
	(*SimulatorSettings)(nil), // 2: model.SimulatorSettings
	(*Relic)(nil),             // 3: model.Relic
	(*RelicStat)(nil),         // 4: model.RelicStat
	(*LightCone)(nil),         // 5: model.LightCone
	(*Character)(nil),         // 6: model.Character
	(*Enemy)(nil),             // 7: model.Enemy
	(*DebuffRES)(nil),         // 8: model.DebuffRES
	(Property)(0),             // 9: model.Property
	(DamageType)(0),           // 10: model.DamageType
	(BehaviorFlag)(0),         // 11: model.BehaviorFlag
}
var file_pb_model_sim_proto_depIdxs = []int32{
	2,  // 0: model.SimConfig.settings:type_name -> model.SimulatorSettings
	6,  // 1: model.SimConfig.characters:type_name -> model.Character
	7,  // 2: model.SimConfig.enemies:type_name -> model.Enemy
	4,  // 3: model.Relic.main_stat:type_name -> model.RelicStat
	4,  // 4: model.Relic.sub_stats:type_name -> model.RelicStat
	9,  // 5: model.RelicStat.stat:type_name -> model.Property
	5,  // 6: model.Character.cone:type_name -> model.LightCone
	3,  // 7: model.Character.relics:type_name -> model.Relic
	10, // 8: model.Enemy.weaknesses:type_name -> model.DamageType
	8,  // 9: model.Enemy.debuff_res:type_name -> model.DebuffRES
	11, // 10: model.DebuffRES.stat:type_name -> model.BehaviorFlag
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pb_model_sim_proto_init() }
func file_pb_model_sim_proto_init() {
	if File_pb_model_sim_proto != nil {
		return
	}
	file_pb_model_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_model_sim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimConfig); i {
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
		file_pb_model_sim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulatorSettings); i {
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
		file_pb_model_sim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relic); i {
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
		file_pb_model_sim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelicStat); i {
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
		file_pb_model_sim_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightCone); i {
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
		file_pb_model_sim_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_pb_model_sim_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enemy); i {
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
		file_pb_model_sim_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebuffRES); i {
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
			RawDescriptor: file_pb_model_sim_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_model_sim_proto_goTypes,
		DependencyIndexes: file_pb_model_sim_proto_depIdxs,
		EnumInfos:         file_pb_model_sim_proto_enumTypes,
		MessageInfos:      file_pb_model_sim_proto_msgTypes,
	}.Build()
	File_pb_model_sim_proto = out.File
	file_pb_model_sim_proto_rawDesc = nil
	file_pb_model_sim_proto_goTypes = nil
	file_pb_model_sim_proto_depIdxs = nil
}
