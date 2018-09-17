// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Signed
	Packet
	Coins
	Signatures
	Message
	Address
	Registration
	VerificationKey
	EncryptionKey
	DecryptionKey
	Hash
	Signature
	Transaction
	Blame
	Invalid
	Inputs
	Packets
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Phase int32

const (
	Phase_NONE                        Phase = 0
	Phase_ANNOUNCEMENT                Phase = 1
	Phase_SHUFFLE                     Phase = 2
	Phase_BROADCAST                   Phase = 3
	Phase_EQUIVOCATION_CHECK          Phase = 4
	Phase_SIGNING                     Phase = 5
	Phase_VERIFICATION_AND_SUBMISSION Phase = 6
	Phase_BLAME                       Phase = 7
)

var Phase_name = map[int32]string{
	0: "NONE",
	1: "ANNOUNCEMENT",
	2: "SHUFFLE",
	3: "BROADCAST",
	4: "EQUIVOCATION_CHECK",
	5: "SIGNING",
	6: "VERIFICATION_AND_SUBMISSION",
	7: "BLAME",
}
var Phase_value = map[string]int32{
	"NONE":                        0,
	"ANNOUNCEMENT":                1,
	"SHUFFLE":                     2,
	"BROADCAST":                   3,
	"EQUIVOCATION_CHECK":          4,
	"SIGNING":                     5,
	"VERIFICATION_AND_SUBMISSION": 6,
	"BLAME": 7,
}

func (x Phase) String() string {
	return proto.EnumName(Phase_name, int32(x))
}
func (Phase) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Reason int32

const (
	Reason_INSUFFICIENTFUNDS             Reason = 0
	Reason_DOUBLESPEND                   Reason = 1
	Reason_EQUIVOCATIONFAILURE           Reason = 2
	Reason_SHUFFLEFAILURE                Reason = 3
	Reason_SHUFFLEANDEQUIVOCATIONFAILURE Reason = 4
	Reason_INVALIDSIGNATURE              Reason = 5
	Reason_MISSINGOUTPUT                 Reason = 6
	Reason_LIAR                          Reason = 7
	Reason_INVALIDFORMAT                 Reason = 8
)

var Reason_name = map[int32]string{
	0: "INSUFFICIENTFUNDS",
	1: "DOUBLESPEND",
	2: "EQUIVOCATIONFAILURE",
	3: "SHUFFLEFAILURE",
	4: "SHUFFLEANDEQUIVOCATIONFAILURE",
	5: "INVALIDSIGNATURE",
	6: "MISSINGOUTPUT",
	7: "LIAR",
	8: "INVALIDFORMAT",
}
var Reason_value = map[string]int32{
	"INSUFFICIENTFUNDS":             0,
	"DOUBLESPEND":                   1,
	"EQUIVOCATIONFAILURE":           2,
	"SHUFFLEFAILURE":                3,
	"SHUFFLEANDEQUIVOCATIONFAILURE": 4,
	"INVALIDSIGNATURE":              5,
	"MISSINGOUTPUT":                 6,
	"LIAR":                          7,
	"INVALIDFORMAT":                 8,
}

func (x Reason) String() string {
	return proto.EnumName(Reason_name, int32(x))
}
func (Reason) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Signed struct {
	Packet    *Packet    `protobuf:"bytes,1,opt,name=packet" json:"packet,omitempty"`
	Signature *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (m *Signed) Reset()                    { *m = Signed{} }
func (m *Signed) String() string            { return proto.CompactTextString(m) }
func (*Signed) ProtoMessage()               {}
func (*Signed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Signed) GetPacket() *Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *Signed) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Packet struct {
	Session      []byte           `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Number       uint32           `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	FromKey      *VerificationKey `protobuf:"bytes,3,opt,name=from_key,json=fromKey" json:"from_key,omitempty"`
	ToKey        *VerificationKey `protobuf:"bytes,4,opt,name=to_key,json=toKey" json:"to_key,omitempty"`
	Phase        Phase            `protobuf:"varint,5,opt,name=phase,enum=Phase" json:"phase,omitempty"`
	Message      *Message         `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	Registration *Registration    `protobuf:"bytes,7,opt,name=registration" json:"registration,omitempty"`
}

func (m *Packet) Reset()                    { *m = Packet{} }
func (m *Packet) String() string            { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()               {}
func (*Packet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Packet) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Packet) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Packet) GetFromKey() *VerificationKey {
	if m != nil {
		return m.FromKey
	}
	return nil
}

func (m *Packet) GetToKey() *VerificationKey {
	if m != nil {
		return m.ToKey
	}
	return nil
}

func (m *Packet) GetPhase() Phase {
	if m != nil {
		return m.Phase
	}
	return Phase_NONE
}

func (m *Packet) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Packet) GetRegistration() *Registration {
	if m != nil {
		return m.Registration
	}
	return nil
}

type Coins struct {
	Coins []string `protobuf:"bytes,1,rep,name=coins" json:"coins,omitempty"`
}

func (m *Coins) Reset()                    { *m = Coins{} }
func (m *Coins) String() string            { return proto.CompactTextString(m) }
func (*Coins) ProtoMessage()               {}
func (*Coins) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Coins) GetCoins() []string {
	if m != nil {
		return m.Coins
	}
	return nil
}

type Signatures struct {
	Utxo      string     `protobuf:"bytes,1,opt,name=utxo" json:"utxo,omitempty"`
	Signature *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (m *Signatures) Reset()                    { *m = Signatures{} }
func (m *Signatures) String() string            { return proto.CompactTextString(m) }
func (*Signatures) ProtoMessage()               {}
func (*Signatures) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Signatures) GetUtxo() string {
	if m != nil {
		return m.Utxo
	}
	return ""
}

func (m *Signatures) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Message struct {
	Address *Address       `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Key     *EncryptionKey `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Hash    *Hash          `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	// map<string, Signature> signatures = 4;
	Signatures []*Signatures     `protobuf:"bytes,4,rep,name=signatures" json:"signatures,omitempty"`
	Str        string            `protobuf:"bytes,5,opt,name=str" json:"str,omitempty"`
	Blame      *Blame            `protobuf:"bytes,6,opt,name=blame" json:"blame,omitempty"`
	Inputs     map[string]*Coins `protobuf:"bytes,7,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Message) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Message) GetKey() *EncryptionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetHash() *Hash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Message) GetSignatures() []*Signatures {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Message) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Message) GetBlame() *Blame {
	if m != nil {
		return m.Blame
	}
	return nil
}

func (m *Message) GetInputs() map[string]*Coins {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type Address struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Registration struct {
	Amount uint64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *Registration) Reset()                    { *m = Registration{} }
func (m *Registration) String() string            { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()               {}
func (*Registration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Registration) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VerificationKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *VerificationKey) Reset()                    { *m = VerificationKey{} }
func (m *VerificationKey) String() string            { return proto.CompactTextString(m) }
func (*VerificationKey) ProtoMessage()               {}
func (*VerificationKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *VerificationKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type EncryptionKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *EncryptionKey) Reset()                    { *m = EncryptionKey{} }
func (m *EncryptionKey) String() string            { return proto.CompactTextString(m) }
func (*EncryptionKey) ProtoMessage()               {}
func (*EncryptionKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *EncryptionKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DecryptionKey struct {
	Key    string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Public string `protobuf:"bytes,2,opt,name=public" json:"public,omitempty"`
}

func (m *DecryptionKey) Reset()                    { *m = DecryptionKey{} }
func (m *DecryptionKey) String() string            { return proto.CompactTextString(m) }
func (*DecryptionKey) ProtoMessage()               {}
func (*DecryptionKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DecryptionKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DecryptionKey) GetPublic() string {
	if m != nil {
		return m.Public
	}
	return ""
}

type Hash struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *Hash) Reset()                    { *m = Hash{} }
func (m *Hash) String() string            { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()               {}
func (*Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Hash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Signature struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Transaction struct {
	Transaction []byte `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Transaction) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type Blame struct {
	Reason      Reason           `protobuf:"varint,1,opt,name=reason,enum=Reason" json:"reason,omitempty"`
	Accused     *VerificationKey `protobuf:"bytes,2,opt,name=accused" json:"accused,omitempty"`
	Key         *DecryptionKey   `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Transaction *Transaction     `protobuf:"bytes,4,opt,name=transaction" json:"transaction,omitempty"`
	Invalid     *Invalid         `protobuf:"bytes,5,opt,name=invalid" json:"invalid,omitempty"`
	Packets     *Packets         `protobuf:"bytes,6,opt,name=packets" json:"packets,omitempty"`
}

func (m *Blame) Reset()                    { *m = Blame{} }
func (m *Blame) String() string            { return proto.CompactTextString(m) }
func (*Blame) ProtoMessage()               {}
func (*Blame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Blame) GetReason() Reason {
	if m != nil {
		return m.Reason
	}
	return Reason_INSUFFICIENTFUNDS
}

func (m *Blame) GetAccused() *VerificationKey {
	if m != nil {
		return m.Accused
	}
	return nil
}

func (m *Blame) GetKey() *DecryptionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Blame) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *Blame) GetInvalid() *Invalid {
	if m != nil {
		return m.Invalid
	}
	return nil
}

func (m *Blame) GetPackets() *Packets {
	if m != nil {
		return m.Packets
	}
	return nil
}

type Invalid struct {
	Invalid []byte `protobuf:"bytes,1,opt,name=invalid,proto3" json:"invalid,omitempty"`
}

func (m *Invalid) Reset()                    { *m = Invalid{} }
func (m *Invalid) String() string            { return proto.CompactTextString(m) }
func (*Invalid) ProtoMessage()               {}
func (*Invalid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Invalid) GetInvalid() []byte {
	if m != nil {
		return m.Invalid
	}
	return nil
}

type Inputs struct {
	Address string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Coins   []string `protobuf:"bytes,2,rep,name=coins" json:"coins,omitempty"`
}

func (m *Inputs) Reset()                    { *m = Inputs{} }
func (m *Inputs) String() string            { return proto.CompactTextString(m) }
func (*Inputs) ProtoMessage()               {}
func (*Inputs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Inputs) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Inputs) GetCoins() []string {
	if m != nil {
		return m.Coins
	}
	return nil
}

type Packets struct {
	Packet []*Signed `protobuf:"bytes,1,rep,name=packet" json:"packet,omitempty"`
}

func (m *Packets) Reset()                    { *m = Packets{} }
func (m *Packets) String() string            { return proto.CompactTextString(m) }
func (*Packets) ProtoMessage()               {}
func (*Packets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Packets) GetPacket() []*Signed {
	if m != nil {
		return m.Packet
	}
	return nil
}

func init() {
	proto.RegisterType((*Signed)(nil), "Signed")
	proto.RegisterType((*Packet)(nil), "Packet")
	proto.RegisterType((*Coins)(nil), "Coins")
	proto.RegisterType((*Signatures)(nil), "Signatures")
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Address)(nil), "Address")
	proto.RegisterType((*Registration)(nil), "Registration")
	proto.RegisterType((*VerificationKey)(nil), "VerificationKey")
	proto.RegisterType((*EncryptionKey)(nil), "EncryptionKey")
	proto.RegisterType((*DecryptionKey)(nil), "DecryptionKey")
	proto.RegisterType((*Hash)(nil), "Hash")
	proto.RegisterType((*Signature)(nil), "Signature")
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*Blame)(nil), "Blame")
	proto.RegisterType((*Invalid)(nil), "Invalid")
	proto.RegisterType((*Inputs)(nil), "Inputs")
	proto.RegisterType((*Packets)(nil), "Packets")
	proto.RegisterEnum("Phase", Phase_name, Phase_value)
	proto.RegisterEnum("Reason", Reason_name, Reason_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xae, 0x6c, 0x4b, 0xb2, 0x8f, 0xed, 0x54, 0x5d, 0x42, 0x11, 0x25, 0x9d, 0xba, 0xea, 0x0c,
	0x98, 0x94, 0x11, 0x43, 0xb8, 0x29, 0xdc, 0xc9, 0xb6, 0x9c, 0x88, 0xd8, 0xab, 0xb0, 0x92, 0x72,
	0x9b, 0x51, 0x6c, 0x35, 0xd1, 0x34, 0x96, 0x3c, 0x5a, 0xb9, 0x83, 0x1f, 0x80, 0x7b, 0x1e, 0x87,
	0x27, 0xe0, 0x71, 0x78, 0x06, 0x66, 0x7f, 0xe4, 0xc8, 0x90, 0x30, 0xd3, 0x3b, 0x7d, 0xdf, 0xf9,
	0x76, 0xf7, 0xec, 0x77, 0xce, 0x1e, 0x41, 0x7f, 0x95, 0x50, 0x1a, 0xdf, 0x24, 0xf6, 0xba, 0xc8,
	0xcb, 0xdc, 0x0a, 0x40, 0x0b, 0xd2, 0x9b, 0x2c, 0x59, 0xa2, 0x57, 0xa0, 0xad, 0xe3, 0xc5, 0x87,
	0xa4, 0x34, 0x95, 0x81, 0x32, 0xec, 0x9e, 0xe8, 0xf6, 0x05, 0x87, 0x44, 0xd2, 0x68, 0x08, 0x1d,
	0x9a, 0xde, 0x64, 0x71, 0xb9, 0x29, 0x12, 0xb3, 0xc1, 0x35, 0x60, 0x07, 0x15, 0x43, 0xee, 0x83,
	0xd6, 0xef, 0x0d, 0xd0, 0xc4, 0x62, 0x64, 0x82, 0x4e, 0x13, 0x4a, 0xd3, 0x3c, 0xe3, 0xdb, 0xf6,
	0x48, 0x05, 0xd1, 0x73, 0xd0, 0xb2, 0xcd, 0xea, 0x3a, 0x29, 0xf8, 0x5e, 0x7d, 0x22, 0x11, 0x7a,
	0x0b, 0xed, 0xf7, 0x45, 0xbe, 0xba, 0xfa, 0x90, 0x6c, 0xcd, 0x26, 0x3f, 0xc5, 0xb0, 0x2f, 0x93,
	0x22, 0x7d, 0x9f, 0x2e, 0xe2, 0x32, 0xcd, 0xb3, 0xf3, 0x64, 0x4b, 0x74, 0xa6, 0x38, 0x4f, 0xb6,
	0xe8, 0x1b, 0xd0, 0xca, 0x9c, 0x4b, 0x5b, 0x8f, 0x48, 0xd5, 0x32, 0x67, 0xc2, 0x23, 0x50, 0xd7,
	0xb7, 0x31, 0x4d, 0x4c, 0x75, 0xa0, 0x0c, 0x0f, 0x4e, 0x34, 0xfb, 0x82, 0x21, 0x22, 0x48, 0x64,
	0x81, 0x2e, 0x6d, 0x31, 0x35, 0xbe, 0x4f, 0xdb, 0x9e, 0x0b, 0x4c, 0xaa, 0x00, 0xfa, 0x01, 0x7a,
	0x45, 0x72, 0x93, 0xd2, 0xb2, 0xe0, 0x7b, 0x9b, 0x3a, 0x17, 0xf6, 0x6d, 0x52, 0x23, 0xc9, 0x9e,
	0xc4, 0x7a, 0x09, 0xea, 0x38, 0x4f, 0x33, 0x8a, 0x0e, 0x41, 0x5d, 0xb0, 0x0f, 0x53, 0x19, 0x34,
	0x87, 0x1d, 0x22, 0x80, 0xf5, 0x0b, 0xc0, 0xce, 0x3e, 0x8a, 0x10, 0xb4, 0x36, 0xe5, 0x6f, 0x39,
	0xb7, 0xa9, 0x43, 0xf8, 0xf7, 0x27, 0x58, 0xfe, 0x67, 0x03, 0x74, 0x99, 0x32, 0xbb, 0x4d, 0xbc,
	0x5c, 0x16, 0x09, 0xa5, 0xb2, 0x94, 0x6d, 0xdb, 0x11, 0x98, 0x54, 0x01, 0x34, 0x80, 0x26, 0x73,
	0x4d, 0xec, 0x79, 0x60, 0xbb, 0xd9, 0xa2, 0xd8, 0xae, 0x2b, 0xcf, 0x58, 0x08, 0x7d, 0x09, 0xad,
	0xdb, 0x98, 0xde, 0xca, 0x1a, 0xa8, 0xf6, 0x59, 0x4c, 0x6f, 0x09, 0xa7, 0xd0, 0x5b, 0x80, 0xdd,
	0xc9, 0xd4, 0x6c, 0x0d, 0x9a, 0xc3, 0xee, 0x49, 0xf7, 0x3e, 0x2f, 0x4a, 0x6a, 0x61, 0x64, 0x40,
	0x93, 0x96, 0x05, 0xf7, 0xbd, 0x43, 0xd8, 0x27, 0xab, 0xc5, 0xf5, 0x5d, 0xbc, 0xaa, 0xbc, 0xd6,
	0xec, 0x11, 0x43, 0x44, 0x90, 0xe8, 0x3b, 0xd0, 0xd2, 0x6c, 0xbd, 0x29, 0xa9, 0xa9, 0xf3, 0x8d,
	0x0f, 0xab, 0x52, 0xd8, 0x1e, 0xa7, 0xdd, 0xac, 0x2c, 0xb6, 0x44, 0x6a, 0x5e, 0x38, 0xd0, 0xad,
	0xd1, 0xec, 0x30, 0x76, 0x2d, 0xe1, 0x21, 0xbf, 0xc6, 0x11, 0xa8, 0x1f, 0xe3, 0xbb, 0x4d, 0x65,
	0x9f, 0x66, 0xf3, 0x8a, 0x10, 0x41, 0xfe, 0xdc, 0x78, 0xa7, 0x58, 0x6f, 0x40, 0x97, 0xf6, 0xb0,
	0x6e, 0xad, 0x3b, 0xd7, 0xd9, 0xf9, 0x65, 0x7d, 0x0d, 0xbd, 0x7a, 0xa1, 0x59, 0xf7, 0xc6, 0xab,
	0x7c, 0x93, 0x89, 0xd7, 0xd2, 0x22, 0x12, 0x59, 0x6f, 0xe0, 0xe9, 0xbf, 0x3a, 0xf0, 0xbf, 0x39,
	0x59, 0xaf, 0xa1, 0xbf, 0x67, 0xf8, 0x03, 0x92, 0x9f, 0xa0, 0x3f, 0x49, 0xfe, 0x57, 0xc2, 0x52,
	0x58, 0x6f, 0xae, 0xef, 0xd2, 0x05, 0xbf, 0x5a, 0x87, 0x48, 0x64, 0xbd, 0x80, 0x16, 0xab, 0x15,
	0x6b, 0x28, 0x5e, 0x40, 0xf1, 0xee, 0xf8, 0xb7, 0xf5, 0x2d, 0x74, 0x76, 0x65, 0x42, 0x47, 0xf5,
	0xee, 0x12, 0xaa, 0x5a, 0x47, 0x7d, 0x0f, 0xdd, 0xb0, 0x88, 0x33, 0x1a, 0x2f, 0xf8, 0x85, 0x07,
	0xd0, 0x2d, 0xef, 0xa1, 0x94, 0xd7, 0x29, 0xeb, 0x6f, 0x05, 0x54, 0x5e, 0x49, 0x36, 0x4a, 0x8a,
	0x24, 0xa6, 0x52, 0x76, 0x70, 0xa2, 0xdb, 0x84, 0x43, 0x22, 0x69, 0x74, 0x0c, 0x7a, 0xbc, 0x58,
	0x6c, 0x68, 0xb2, 0x94, 0x65, 0x79, 0xe0, 0x89, 0x4b, 0x41, 0xd5, 0xa9, 0x4d, 0xd9, 0xa9, 0x7b,
	0xae, 0x08, 0x23, 0xec, 0xfd, 0xd4, 0xc4, 0x24, 0xe8, 0xd9, 0xb5, 0xec, 0xf7, 0x12, 0x65, 0xef,
	0x23, 0xcd, 0x3e, 0xc6, 0x77, 0xe9, 0x92, 0x77, 0x25, 0x7b, 0x1f, 0x9e, 0xc0, 0xa4, 0x0a, 0x30,
	0x8d, 0x18, 0x7b, 0x74, 0x37, 0x11, 0xc4, 0x44, 0xa3, 0xa4, 0x0a, 0xb0, 0xc6, 0x91, 0xeb, 0x58,
	0xe3, 0x54, 0x5b, 0xca, 0x31, 0x27, 0xa1, 0xf5, 0x0e, 0x34, 0xd1, 0xa0, 0x8f, 0x37, 0xd7, 0xfd,
	0x78, 0x68, 0xd4, 0xc7, 0xc3, 0x31, 0xe8, 0xf2, 0xc8, 0xbd, 0xd9, 0xdc, 0xe4, 0xb3, 0x59, 0x0c,
	0xed, 0x6a, 0x36, 0x1f, 0xff, 0xa1, 0x80, 0xca, 0x27, 0x1a, 0x6a, 0x43, 0x0b, 0xfb, 0xd8, 0x35,
	0x9e, 0x20, 0x03, 0x7a, 0x0e, 0xc6, 0x7e, 0x84, 0xc7, 0xee, 0xdc, 0xc5, 0xa1, 0xa1, 0xa0, 0x2e,
	0xe8, 0xc1, 0x59, 0x34, 0x9d, 0xce, 0x5c, 0xa3, 0x81, 0xfa, 0xd0, 0x19, 0x11, 0xdf, 0x99, 0x8c,
	0x9d, 0x20, 0x34, 0x9a, 0xe8, 0x39, 0x20, 0xf7, 0xd7, 0xc8, 0xbb, 0xf4, 0xc7, 0x4e, 0xe8, 0xf9,
	0xf8, 0x6a, 0x7c, 0xe6, 0x8e, 0xcf, 0x8d, 0x16, 0x5f, 0xe3, 0x9d, 0x62, 0x0f, 0x9f, 0x1a, 0x2a,
	0x7a, 0x05, 0x5f, 0x5d, 0xba, 0xc4, 0x9b, 0x7a, 0x52, 0xe4, 0xe0, 0xc9, 0x55, 0x10, 0x8d, 0xe6,
	0x5e, 0x10, 0x78, 0x3e, 0x36, 0x34, 0xd4, 0x01, 0x75, 0x34, 0x73, 0xe6, 0xae, 0xa1, 0x1f, 0xff,
	0xa5, 0x80, 0x26, 0xca, 0x8e, 0x3e, 0x87, 0x67, 0x1e, 0x0e, 0xa2, 0xe9, 0xd4, 0x1b, 0x7b, 0x2e,
	0x0e, 0xa7, 0x11, 0x9e, 0x04, 0xc6, 0x13, 0xf4, 0x14, 0xba, 0x13, 0x3f, 0x1a, 0xcd, 0xdc, 0xe0,
	0xc2, 0xc5, 0x13, 0x43, 0x41, 0x5f, 0xc0, 0x67, 0xf5, 0x1c, 0xa6, 0x8e, 0x37, 0x8b, 0x08, 0xcb,
	0x15, 0xc1, 0x81, 0x4c, 0xbc, 0xe2, 0x9a, 0xe8, 0x35, 0xbc, 0x94, 0x9c, 0x83, 0x27, 0x0f, 0x2d,
	0x6b, 0xa1, 0x43, 0x30, 0x3c, 0x7c, 0xe9, 0xcc, 0xbc, 0x09, 0xbb, 0x82, 0x13, 0x32, 0x56, 0x45,
	0xcf, 0xa0, 0xcf, 0x13, 0xc6, 0xa7, 0x7e, 0x14, 0x5e, 0x44, 0xa1, 0xa1, 0x31, 0xd3, 0x66, 0x9e,
	0x43, 0x0c, 0x9d, 0x05, 0xe5, 0x92, 0xa9, 0x4f, 0xe6, 0x4e, 0x68, 0xb4, 0xaf, 0x35, 0xfe, 0xa7,
	0xfc, 0xf1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xf4, 0x8f, 0xf3, 0x3a, 0x07, 0x00, 0x00,
}
